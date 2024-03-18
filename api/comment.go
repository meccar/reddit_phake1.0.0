package api

import (
	"fmt"
	"net/http"

	db "sqlc"
	token "token"
	util "util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (server *Server) CreateComment(c *gin.Context) {
	// Parse the contact form data
	msg, err := commentForm(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	claim, err := token.GetClaims(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	msg.UserID, err = server.DbHandler.GetAccountIDbyUsername(c, claim["Username"].(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	msg.PostID, err = uuid.Parse(c.Param("post_id"))
	if err != nil {
		// Handle the error if the post ID is not a valid UUID
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Submit the form to the database
	submit, err := server.DbHandler.CreateCommentTx(c.Request.Context(), *msg)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("submit: ", submit)

	// Return a success response with the submitted form data
	c.JSON(http.StatusOK, msg)
}

func commentForm(r *http.Request) (*db.CreateCommentTxParams, error) {
	msg := &db.CreateCommentTxParams{}
	err := util.ParseForm(r, msg)
	return msg, err
}

// func (server *Server) commentHandler(c *gin.Context) {
// }
