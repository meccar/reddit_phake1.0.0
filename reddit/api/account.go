package api

import (
	"net/http"
	"fmt"
	
	jwtauth "token"

	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	// Extract the role and token from the URL parameters
	_ = c.Param("role")
	_ = c.Param("token")

	tokenClaims, err := jwtauth.GetClaims(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if tokenClaims["role"].(string) == "admin" {
		// web.Render(c.Writer, "admin", nil)
		c.HTML(http.StatusOK, "admin", nil)
	} else if tokenClaims["role"].(string) == "user" {
		// web.Render(c.Writer, "user", nil)
		c.HTML(http.StatusOK, "user", nil)
	} else {
		c.AbortWithStatusJSON(http.StatusForbidden, errorResponse(fmt.Errorf("Access Forbidden")))
	}
}
