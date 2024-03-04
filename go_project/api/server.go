package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
)

type Server struct {
	config     util.Config
	handler    db.Handlers
	tokenMaker token.Maker
	Router     *chi.Mux
}

// type Server struct {
// 	Router    *chi.Mux
// 	queries   *db.Queries
// 	AuthToken *jwtauth.JWTAuth
// }

func CreateServer(config util.Config, handler db.Handlers) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		handler:    handler,
		tokenMaker: tokenMaker,
	}

	server.setupRoutes()
	return server, nil
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)
		log.Println(req)
		next.ServeHTTP(w, r)
		log.Println(req, "completed in", time.Now().Sub(start))
	})
}

func (server *Server) setupRoutes() {

	// r := chi.NewRouter()
	// login := JWT{}.new()

	// r.Use(login.Verifier())

	// Serve public files for any path
	server.Router.Get("/*", server.publicHandler)

	// Handle the root path
	server.Router.Get("/", GetHandler("home"))

	server.Router.Route("/lienhe", func(r chi.Router) {
		r.Get("/", GetHandler("lienhe"))
		r.Post("/", formHandlerWrapper(dbHandler))
	})

	server.Router.Route("/login", func(r chi.Router) {
		// r.Use(login.Authenticator())
		r.Get("/", GetHandler("login"))
		// r.Post("/", loginHandlerWrapper(dbHandler))
	})

	server.Router.Get("/thankyou", GetHandler("thankyou"))
}

func setupServer(r *chi.Mux) *http.Server {
	// port, ok := os.LookupEnv("PORT")
	// if !ok {
	// 	port = "9595"
	// }

	addr := fmt.Sprintf(":%s", PORT)
	return &http.Server{
		Addr:    addr,
		Handler: logging(r),
		// ReadTimeout:  15 * time.Second,
		// WriteTimeout: 15 * time.Second,
		// IdleTimeout:  15 * time.Second,
	}
}

func startServer(server *http.Server) {
	// Log the server's address
	log.Printf("Server is running on http://localhost%s\n", server.Addr)

	// Start the server
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %v", err)
	}

	// Listen for interrupt signals to gracefully shut down the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Log the shutdown process
	log.Println("Shutting down the server...")

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error during server shutdown: %v", err)
	} else {
		log.Println("Server gracefully stopped")
	}
	defer cancel()

}
