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

	server.TokenAuthRS256.MakeToken((msg.ID).String(), msg.Username, role, c.Writer)
	// Generate and sign the JWT token
	// token, err := server.TokenAuthRS256.MakeToken((msg.ID).String(), msg.Username, role, c.Writer)
	// if err != nil {
		// c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		// return
	// }

	tokens.CreateToken(msg.Username, role, duration, c.Writer)
	// Create a token with payload and sign it
	// signed, payload, err := tokens.CreateToken(msg.Username, role, duration)
	// if err != nil {
		// c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		// return
	// }

	// Print the signed token and payload (for debugging purposes)
	// fmt.Println(signed)
	// fmt.Println(payload)

	// Set the JWT token as a cookie in the response
	// server.TokenAuthRS256.SetJWTCookie(c.Writer, token, role, int(duration.Seconds()))

	// Create a session for the user
	msg.Session, err = server.DbHandler.CreateSession(c.Request.Context(), msg.Username, role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Set the Authorization header in the response
	// c.Writer.Header().Set("Authorization", "Bearer "+token)

	// Respond with the login message
	c.JSON(http.StatusOK, msg)
}

// loginForm parses the login form data from the HTTP request
func loginForm(r *http.Request) (*db.LoginAccountTxParams, error) {
	msg := &db.LoginAccountTxParams{}
	err := util.ParseForm(r, msg)
	return msg, err
}
