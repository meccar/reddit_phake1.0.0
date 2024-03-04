package api

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gin-gonic/gin"

	errorHandler "error"
	db "sqlc"
	util "util"
)

func (server *Server) loginHandler(c *gin.Context) {
	duration := 120 * time.Second

	msg, err := loginForm(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if msgValidationResult := server.DbHandler.VerifyLogin(c.Request.Context(), msg); msgValidationResult != nil {
		status := errorHandler.StatusHandler(msgValidationResult)
		c.AbortWithStatusJSON(status, errorResponse(msgValidationResult))

		return
	}

	msg.Password, err = util.HashPassword(msg.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	msg.ID, _ = uuid.Parse(server.DbHandler.GetID(c.Request.Context(), msg))
	role := server.DbHandler.GetRole(c.Request.Context(), msg)

	token, err := server.TokenAuthRS256.MakeToken((msg.ID).String(), msg.Username, role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	server.TokenAuthRS256.SetJWTCookie(c.Writer, token, role, int(duration.Seconds()))
	msg.Session, err = server.DbHandler.CreateSession(c.Request.Context(), msg.Username, role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.Writer.Header().Set("Authorization", "Bearer "+token)

	c.JSON(http.StatusOK, msg)
}

func loginForm(r *http.Request) (*db.LoginAccountTxParams, error) {
	msg := &db.LoginAccountTxParams{}
	err := util.ParseForm(r, msg)
	return msg, err
}
