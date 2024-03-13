package api

import (
	"fmt"
	"net/http"

	db "sqlc"
	util "util"

	"github.com/gin-gonic/gin"
)

func (server *Server) communityHandler(c *gin.Context) {
	// Parse the contact form data
	msg, err := communityForm(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Submit the form to the database
	submit, err := server.DbHandler.CreateCommunityTx(c.Request.Context(), *msg)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Println("communityHandler submit: ", submit)
	// msg.ID = submit.Community.ID

	// Return a success response with the submitted form data
	c.JSON(http.StatusOK, msg)
}

func communityForm(r *http.Request) (*db.CreateCommunityTxParams, error) {
	msg := &db.CreateCommunityTxParams{}
	err := util.ParseForm(r, msg)
	return msg, err
}

func (server *Server) SearchKeyCommunity(c *gin.Context) ([]db.SearchCommunityNameRow, error) {
	searchQuery := c.Query("q")
	fmt.Println("SearchCommunity searchQuery: ", searchQuery)
	data, err := server.DbHandler.SearchCommunity(c.Request.Context(), searchQuery)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return nil, err
	}

	return data, nil
	// c.JSON(http.StatusOK, data)
}
