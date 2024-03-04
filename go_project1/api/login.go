package api

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"

	"net/http"
	db "sqlc"
	web "web"

	"github.com/gorilla/schema"
)

// util "util"

func (server *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		return
	}

	msg := &db.LoginAccountTxParams{}

	fmt.Printf("\n loginHandler msg 1 %v\n", msg)

	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(msg, r.PostForm)
	fmt.Printf("\n loginHandler r.PostForm %v\n", r.PostForm)

	if decodeErr != nil {
		log.Error().Err(decodeErr).Msg("Error mapping parsed form data to struct")
		http.Error(w, "Error processing form data", http.StatusInternalServerError)
		return
	}

	// if password, ok := r.PostForm["Password"]; ok {
	// 	// Hash the password using bcrypt
	// 	hashedPassword, _ := util.HashPassword(password[0])
	// 	// Set the hashed password in msg.Bcrypt
	// 	msg.Password = string(hashedPassword)
	// }

	fmt.Printf("\n loginHandler msg 2 %v\n", msg)
	// fmt.Printf("\n loginHandler msg.Bcrypt %v\n", msg.Bcrypt)
	// fmt.Printf("\n loginHandler msg.Bcrypt %v\n", &msg.Bcrypt)
	// fmt.Printf("\n loginHandler msg.Password %v\n", msg.Password)

	if !server.DbHandler.VerifyLogin(r.Context(), msg) {
		// log.Info().Msg(msg)
		web.Render(w, "login", msg)
		return
	}
	userID, err := server.DbHandler.GetID(r.Context(), msg)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// fmt.Printf("\n loginHandler msg after verify %v\n", msg)
	token, err := server.TokenAuthHS256.MakeToken(msg.Username)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Printf("\n loginHandler token %v\n", token)

	server.TokenAuthHS256.SetJWTCookie(w, token)

	// http.Redirect(w, r, "/login/admin", http.StatusSeeOther)
	// redirectURL := fmt.Sprintf("/login/admin/%s", msg.Username)
	// http.Redirect(w, r, redirectURL, http.StatusSeeOther)

	fmt.Printf("\n loginHandler ID %v\n", userID)
	response := map[string]string{"ID": userID}
	fmt.Printf("\n loginHandler response %v\n", response)

	jsonResponse, err := json.Marshal(response)
	fmt.Printf("\n loginHandler jsonResponse %v\n", jsonResponse)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Printf("\n loginHandler response %v\n", response)

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(jsonResponse)

	redirectURL := fmt.Sprintf("/admin/%s", userID)

	fmt.Printf("\n loginHandler redirectURL %v\n", redirectURL)

	http.Redirect(w, r, redirectURL, http.StatusSeeOther)
}
