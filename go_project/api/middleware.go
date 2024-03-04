package api

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
)

func (server *Server) MountMiddleware() {
	server.Router.Use(Logger)
	server.Router.Use(middleware.Logger)
	server.Router.Use(middleware.Recoverer)
	server.Router.Use(middleware.RequestID)
}

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		handler.ServeHTTP(writer, request)
	})
}
