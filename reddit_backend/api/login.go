package api

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gin-gonic/gin"

	tokens "token"
	errorHandler "error"
	db "sqlc"
	util "util"
)

// loginHandler handles the login request
func (server *Server) loginHandler(c *gin.Context) {
	// Set the duration for which the token will be valid
	duration := 120 * time.Second

	// Parse the login form data
	msg, err := loginForm(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Validate the login credentials
	if msgValidationResult := server.DbHandler.VerifyLogin(c.Request.Context(), msg); msgValidationResult != nil {
		status := errorHandler.StatusHandler(msgValidationResult)
		c.AbortWithStatusJSON(status, errorResponse(msgValidationResult))
		return
	}

	// Hash the password before storing or comparing
	msg.Password, err = util.HashPassword(msg.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Retrieve the user's ID and role
	msg.ID, _ = uuid.Parse(server.DbHandler.GetID(c.Request.Context(), msg))
	role := server.DbHandler.GetRole(c.Request.Context(), msg)

	// Generate and sign the JWT token
	err = server.TokenAuthRS256.MakeToken((msg.ID).String(), msg.Username, role, c.Writer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Generate and sign the Paseto token
	err = tokens.CreateToken(msg.Username, role, duration, c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Create a session for the user
	_, err = server.DbHandler.CreateSession(c.Request.Context(), msg.Username, role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Respond with the login message
	c.JSON(http.StatusOK, msg)

	c.Next()
}

// loginForm parses the login form data from the HTTP request
func loginForm(r *http.Request) (*db.LoginAccountTxParams, error) {
	msg := &db.LoginAccountTxParams{}
	err := util.ParseForm(r, msg)
	return msg, err
}
