package api

import (
	"net/http"
	"strings"
	"fmt"
	"time"
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
  	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/requestid"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/timeout"

  	
	"github.com/rs/zerolog/log"
	"github.com/fatih/color"
	"go.uber.org/ratelimit"

	token "token"
)

var (
	limit ratelimit.Limiter
	rps   = flag.Int("rps", 100, "request per second")
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func (server *Server) MountMiddleware() {
	// Session
	store := cookie.NewStore([]byte("secret"))
	server.Router.Use(sessions.Sessions("session", store))

	// Use RequestID middleware to generate request IDs
	server.Router.Use(requestid.New())
	
	// Use httprate middleware to limit requests by IP
	limit = ratelimit.New(*rps)
	server.Router.Use(leakBucket())

	// Use CORS middleware to handle Cross-Origin Resource Sharing
	server.Router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token","Accept-Encoding","Cache-Control","X-Requested-With","Origin"},
		ExposeHeaders:    []string{"Content-Length", "Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// server.Router.Use(timeoutMiddleware())
}

func roleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		tokenClaims, err := token.GetClaims(c.Request)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		// fmt.Println("<<< After tokenClaims ", tokenClaims)
		// fmt.Println("<<< tokenClaims['Role'].(string) ", tokenClaims["Role"].(string))
		// fmt.Println("<<< URL admin? ", strings.HasPrefix(c.Request.URL.Path, "/admin/"))
		// fmt.Println("<<< URL user? ", strings.HasPrefix(c.Request.URL.Path, "/user/"))


		// Check allowed paths based on role
		if tokenClaims["Role"].(string) == "admin" && !strings.HasPrefix(c.Request.URL.Path, "/admin/") {
			// fmt.Println("<<< admin")
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(fmt.Errorf("Access Forbidden")))
			return

		} else if tokenClaims["Role"].(string) == "user" && !strings.HasPrefix(c.Request.URL.Path, "/user/") {
			// fmt.Println("<<< user")
			c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(fmt.Errorf("Access Forbidden")))
			return
		}
		// fmt.Println("<<< After Check allowed paths")

		// Continue processing the request
		c.Next()
	}
}

func CookieTool() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get cookie
		if cookie, err := c.Cookie("label"); err == nil {
			if cookie == "ok" {
				c.Next()
				return
			}
		}

		// Cookie verification failed
		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden with no cookie"})
		c.Abort()
	}
}

func timeoutMiddleware() gin.HandlerFunc {
	return timeout.New(
		timeout.WithTimeout(500*time.Millisecond),
		timeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		timeout.WithResponse(timeoutResponse),
	)
}

func timeoutResponse(c *gin.Context) {
	c.JSON(http.StatusRequestTimeout, gin.H{
		"error": "Request Timeout",
	})
}

func leakBucket() gin.HandlerFunc {
	prev := time.Now()
	return func(c *gin.Context) {
		now := limit.Take()
		log.Print(color.CyanString("%v", now.Sub(prev)))
		prev = now
	}
}
