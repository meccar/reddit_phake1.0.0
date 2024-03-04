package api

import (
	"net/http"
	web "web"
)

// GET files from public
func publicHandler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
}

// PATHS handlers
func GetHandler(pageName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		web.Render(w, pageName, nil)
		// w.WriteHeader(200)
	}
}

func (server *Server) formHandlerWrapper() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.formHandler(w, r)
		// w.WriteHeader(200)
	}
}

func (server *Server) loginHandlerWrapper() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.loginHandler(w, r)
		// w.WriteHeader(200)
	}
}

func (server *Server) registerHandlerWrapper() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.registerHandler(w, r)
		// w.WriteHeader(200)
	}
}
