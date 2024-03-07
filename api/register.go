package api

import (
	"fmt"
	"net/http"

	errorHandler "error"
	db "sqlc"
	util "util"
	mail "mail"

	"github.com/gin-gonic/gin"
)
func (server *Server) registerHandler(c *gin.Context) {
	msg, err := registerForm(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if msgValidationResult := msg.ValidateRegister(c.Request.Context()); msgValidationResult != nil {
		fmt.Printf("\n <- ValidateRegister msgValidationResult %v\n", msgValidationResult)
		status := errorHandler.StatusHandler(msgValidationResult)
		c.AbortWithStatusJSON(status, errorResponse(msgValidationResult))
		return
	}

	msg.Password, err = util.HashPassword(msg.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if register, err := server.DbHandler.CreateAccountTx(c.Request.Context(), *msg); err != nil {
		status := errorHandler.StatusHandler(err)
		c.AbortWithStatusJSON(status, errorResponse(err))
		return
	} else {
		msg.ID = register.Account.ID
	}

	mail.StartMail(msg.Username, util.RandomString(6))
	c.JSON(http.StatusOK, msg)

}

func registerForm(r *http.Request) (*db.CreateAccountTxParams, error) {
	msg := &db.CreateAccountTxParams{}
	err := util.ParseForm(r, msg)
	return msg, err
}
