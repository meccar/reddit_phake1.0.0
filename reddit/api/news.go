package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

)
func (server *Server) handleNews(c *gin.Context) {
	// Retrieve posts from the database
	posts, err := server.DbHandler.GetPosts(c.Request.Context())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Write JSON response
	c.JSON(http.StatusOK, posts)
}
