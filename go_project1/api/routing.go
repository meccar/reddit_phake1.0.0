package api

import (
	"github.com/go-chi/chi"

	jwtauth "token"
	web "web"
)

func (server *Server) SetupRoutes() {

	// POST public routes
	// server.Router.Group(func(r chi.Router) {
	// 	r.Post("/lienhe", server.formHandlerWrapper())
	// 	r.Post("/login", server.loginHandlerWrapper())
	// 	r.Post("/register", server.registerHandlerWrapper())
	// })
	//
	server.Router.Group(func(r chi.Router) {
		r.Route("/lienhe", func(r chi.Router) {
			r.Get("/", GetHandler("lienhe"))
			// r.Post("/", server.formHandlerWrapper())
		})
		r.Route("/login", func(r chi.Router) {
			r.Get("/", GetHandler("login"))
			r.Post("/", server.loginHandlerWrapper())
		})
		r.Route("/register", func(r chi.Router) {
			r.Get("/", GetHandler("register"))
			r.Post("/", server.registerHandlerWrapper())
		})
		r.Get("/thankyou", GetHandler("thankyou"))
	})

	// GET private routes
	server.Router.Group(func(r chi.Router) {
		r.Use(
			jwtauth.Verifier(server.TokenAuthHS256),
			jwtauth.Authenticator(server.TokenAuthHS256),
		)
		// fmt.Printf("\n Inside /admin/{userID} \n")
		r.Get("/admin/{id}", userHandler)
	})

	server.Router.Group(func(r chi.Router) {
		r.Route("/api/v1/form", func(r chi.Router) {
			// r.Get("/", web.Form)
			r.Post("/", server.formHandlerWrapper())
		})
	})

	// GET public routes
	server.Router.Group(func(r chi.Router) {
		r.Get("/", GetHandler("home"))
		// r.Get("/lienhe", GetHandler("lienhe"))
		// r.Get("/login", GetHandler("login"))
		// r.Get("/register", GetHandler("register"))
		r.Get("/thankyou", GetHandler("thankyou"))
		r.Get("/*", publicHandler)
	})
}
