package api

import (
	"fmt"
	"net/http"

	db "sqlc"
	util "util"

	"github.com/gin-gonic/gin"
)

func (server *Server) commentHandler(c *gin.Context) {
	// Parse the contact form data
	msg, err := commentForm(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Submit the form to the database
	submit, err := server.DbHandler.CreateCommentTx(c.Request.Context(), *msg)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("submit: ", submit)
	// Set the ID of the submitted form in the message
	// msg.ID = submit.Comment.ID

	// Return a success response with the submitted form data
	c.JSON(http.StatusOK, msg)
}

func commentForm(r *http.Request) (*db.CreateCommentTxParams, error) {
	msg := &db.CreateCommentTxParams{}
	err := util.ParseForm(r, msg)
	return msg, err
}
