package api

import (
	"net/http"

	jwtauth "token"

	"github.com/gin-gonic/gin"

)

func (server *Server) SetupRoutes() {
	// Load HTML templates
	server.Router.LoadHTMLGlob("./templates/*.html")

	// Serve static files
	server.Router.Static("/public", "./public")

	// Public routes
	server.Router.GET("/lienhe", GetHandler("lienhe"))
	server.Router.GET("/login", GetHandler("login"))
	server.Router.GET("/register", GetHandler("register"))
	server.Router.GET("/", GetHandler("home"))
	server.Router.GET("/thankyou", GetHandler("thankyou"))
	server.Router.GET("/tintuc", server.handleNews)

	// Route grouping for authenticated routes
	authRoutes := server.Router.Group("/")
	authRoutes.Use(
		jwtauth.Verifier(server.TokenAuthRS256),
		jwtauth.Authenticator(server.TokenAuthRS256),
		// jwtauth.VerifyPaseto(server.Pv4),
		// jwtauth.VerifyPaseto(*http.Request),
		roleMiddleware(),
	)
	{
		authRoutes.GET("/:role/:token", UserHandler)
	}

	// API routes
	apiRoutes := server.Router.Group("/api/v1")
	apiRoutes.POST("/form", server.handlerWrapper(server.formHandler))
	// apiRoutes.Use(
	// 	jwtauth.Verifier(server.TokenAuthRS256),
	// 	jwtauth.Authenticator(server.TokenAuthRS256),
	// )
	{
		apiRoutes.POST("/register", server.handlerWrapper(server.registerHandler))
		apiRoutes.GET("/verify_email", server.handlerWrapper(server.VerifyEmailHandler))
		apiRoutes.POST("/login", server.handlerWrapper(server.loginHandler))
		apiRoutes.POST("/post", server.handlerWrapper(server.postHandler))
		apiRoutes.POST("/logout", server.handlerWrapper(server.logoutHandler))
	}

	// No route handler
	server.Router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})
}





