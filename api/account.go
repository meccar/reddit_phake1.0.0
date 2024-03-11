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
	
	_, claim, err := jwtauth.FromContext(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if claim["Role"].(string) == "admin" {
		// web.Render(c.Writer, "admin", nil)
		c.HTML(http.StatusOK, "admin", nil)
	} else if claim["Role"].(string) == "user" {
		// web.Render(c.Writer, "user", nil)
		c.HTML(http.StatusOK, "user", nil)
	} else {
		c.JSON(http.StatusForbidden, errorResponse(fmt.Errorf("Access Forbidden")))
	}
}
