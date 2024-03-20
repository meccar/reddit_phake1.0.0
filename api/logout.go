package api

import (
	"net/http"

	token "token"

	"github.com/gin-gonic/gin"
)

func (server *Server) logoutHandler(c *gin.Context) {
	tokenString := token.TokenFromCookie(c.Request)

	err := server.DbHandler.DeleteSession(c.Request.Context(), tokenString)
	if err != nil || tokenString == "" {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	token.DeleteJWTCookie(c, tokenString)
	c.JSON(http.StatusOK, tokenString)
}
