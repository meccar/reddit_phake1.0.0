package api

import (
	"net/http"
	"fmt"
	
	jwtauth "token"

	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	var requestBody map[string]interface{}
	// Extract the role and token from the URL parameters
	_ = c.Param("role")
	_ = c.Param("token")
	
	token, claim, err := jwtauth.FromContext(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	tokenTest1, _ := requestBody["token"].(string)
	fmt.Println("\n GetClaims token: ",token)
	fmt.Println("\n GetClaims tokenTest1: ",tokenTest1)


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
