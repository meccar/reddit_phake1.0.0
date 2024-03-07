package api

import (
	"net/http"

	db "sqlc"
	util "util"

	"github.com/gin-gonic/gin"
)

func (server *Server) postHandler(c *gin.Context) {
	msg, err := postForm(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	
	if submit, err := server.DbHandler.CreatePostTx(c.Request.Context(), *msg); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	} else {
		msg.ID = submit.Post.ID
	}

	c.JSON(http.StatusOK, msg)
}

func postForm(r *http.Request) (*db.CreatePostTxParams, error) {
	msg := &db.CreatePostTxParams{}
	err := util.ParseForm(r, msg)
	return msg, err
}
