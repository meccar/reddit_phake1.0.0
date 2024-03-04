package api

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"
)

func (server *Server) MountMiddleware() {
	server.Router.Use(Logger)
	// fmt.Printf("\n logging r %v\n", Logger)

	server.Router.Use(middleware.Logger)
	server.Router.Use(middleware.Recoverer)
	server.Router.Use(middleware.RequestID)
	server.Router.Use(middleware.RealIP)
	server.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
}

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		log.Info().Msg(request.URL.Path)
		handler.ServeHTTP(writer, request)
	})
}

// func render(w http.ResponseWriter, filename string, data interface{}) {
// 	err := util.Config.Templates.ExecuteTemplate(w, filename, data)
// 	if err != nil {
// 		log.Printf("Error rendering template %s: %v\n", filename, err)
// 		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
// 		return
// 	}
// }
