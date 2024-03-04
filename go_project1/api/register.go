package api

import (
	"fmt"
	"github.com/rs/zerolog/log"

	"net/http"
	db "sqlc"
	util "util"
	web "web"

	"github.com/gorilla/schema"
)

func (server *Server) registerHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		return
	}

	register := &db.CreateAccountTxParams{}
	// fmt.Printf("\n registerHandler register %v\n", register)

	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(register, r.PostForm)
	fmt.Printf("\n registerHandler r.PostForm %v\n", r.PostForm)

	if decodeErr != nil {
		log.Error().Err(decodeErr).Msg("Error mapping parsed form data to struct")
		http.Error(w, "Error processing form data", http.StatusInternalServerError)
		return
	}

	if password, ok := r.PostForm["Password"]; ok {
		// Hash the password using bcrypt
		hashedPassword, _ := util.HashPassword(password[0])

		// Set the hashed password in msg.Bcrypt
		register.Password = string(hashedPassword)
	}
	fmt.Printf("\n registerHandler register.Password %v\n", register.Password)

	if !register.ValidateRegister() {
		// log.Info().Msg(register)
		web.Render(w, "register", register)
		return
	}

	// fmt.Printf("\n formHandler dbHandler %v\n", server.dbHandler)

	if _, err := server.DbHandler.CreateAccountTx(r.Context(), *register); err != nil {
		// log.Info().Msg(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
