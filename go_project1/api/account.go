package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	web "web"
)

type userData struct {
	Token string
	// Other fields as needed
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n Inside userHandler \n")

	// Extract the token from the URL
	userID := chi.URLParam(r, "userID")
	fmt.Printf("\n userHandler token %v\n", userID)

	// Render the HTML page with the token
	// data := userData{ID: id}
	// fmt.Printf("\n userHandler data %v\n", data)

	web.Render(w, "admin", nil)
	w.WriteHeader(200)

}
