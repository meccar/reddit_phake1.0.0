package api

import (
	"fmt"
	"net/http"

	db "sqlc"
	util "util"

	"github.com/gin-gonic/gin"
)

func (server *Server) replyHandler(c *gin.Context) {
	// Parse the contact form data
	msg, err := replyForm(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Submit the form to the database
	submit, err := server.DbHandler.CreateReplyTx(c.Request.Context(), *msg)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("submit: ", submit)
	// Set the ID of the submitted form in the message
	// msg.ID = submit.Reply.ID

	// Return a success response with the submitted form data
	c.JSON(http.StatusOK, msg)
}

func replyForm(r *http.Request) (*db.CreateReplyTxParams, error) {
	msg := &db.CreateReplyTxParams{}
	err := util.ParseForm(r, msg)
	return msg, err
}
