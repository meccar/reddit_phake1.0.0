package api

import (
	"net/http"

	errorHandler "error"
	db "sqlc"
	util "util"
	
	"github.com/gin-gonic/gin"
)
func (server *Server) formHandler(c *gin.Context) {
	// Parse the contact form data
	msg, err := contactForm(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Validate the form data
	if msgValidationResult := msg.ValidateForm(); msgValidationResult != nil {
		status := errorHandler.StatusHandler(msgValidationResult)
		c.AbortWithStatusJSON(status, errorResponse(msgValidationResult))
		return
	}

	// Submit the form to the database
	submit, err := server.DbHandler.SubmitFormTx(c.Request.Context(), *msg)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Set the ID of the submitted form in the message
	msg.ID = submit.Form.ID

	// Return a success response with the submitted form data
	c.JSON(http.StatusOK, msg)
}

func contactForm(r *http.Request) (*db.SubmitFormTxParams, error) {
	msg := &db.SubmitFormTxParams{}
	err := util.ParseForm(r, msg)
	return msg, err
}
