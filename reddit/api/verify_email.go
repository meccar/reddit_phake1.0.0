package api
import (
  "net/http"

  db "sqlc"

  "github.com/gin-gonic/gin"
)

func (server *Server) VerifyEmailHandler(c *gin.Context) {
	// Parse query parameters
	queryParams := c.Request.URL.Query()
	
	email := queryParams.Get("email")
	secret_code := queryParams.Get("secret_code")

	response, err := server.DbHandler.VerifyEmail(c.Request.Context(), &db.CreateVerifyEmailParams{
		Username:   email,
		SecretCode: secret_code,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
    	return 
  	}

	  c.JSON(http.StatusOK, response)
}