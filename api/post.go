package api

import (
	"net/http"

	db "sqlc"
	util "util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *Server) postHandler(c *gin.Context) {
	msg, err := postForm(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	// msg.UserID = c.GetHeader("id")
	id := c.GetHeader("id")

	// Parse the string value into a UUID
	msg.UserID, err = uuid.Parse(id)
	if err != nil {
		// Handle error if the string is not a valid UUID
	}

	communityResult, err := server.SearchKeyCommunity(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	msg.CommunityID = communityResult[0].ID

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
