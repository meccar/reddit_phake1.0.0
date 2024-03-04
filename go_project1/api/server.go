package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"

	db "sqlc"
	jwtauth "token"
	util "util"
)

// "log"
var (
	PORT string
)

type Server struct {
	Config         util.Config
	Router         *chi.Mux
	DbHandler      *db.Handlers
	TokenAuthHS256 *jwtauth.JWTAuth
}

func CreateServer(Config util.Config, dbHandler *db.Handlers) (*Server, error) {
	// config, err := util.Init()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	server := &Server{
		Config:         Config,
		Router:         chi.NewRouter(),
		DbHandler:      dbHandler,
		TokenAuthHS256: jwtauth.TokenAuthHS256,
	}
	server.MountMiddleware()
	server.SetupRoutes()
	return server, nil
}

func init() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	PORT = config.HTTPServerAddress
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)

		defer func() {
			if r := recover(); r != nil {
				log.Error().Interface("panic", r).Msg("Recovered from panic")
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		log.Info().Msg(req)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
		// fmt.Printf("\n logging w %v\n", w)
		// fmt.Printf("\n logging r %v\n", r)

		// Calculate and log the elapsed time using time.Since
		elapsed := time.Since(start)
		log.Info().Msgf("%s completed in %s", req, elapsed)
	})
}

func (server *Server) SetupServer() *http.Server {
	addr := fmt.Sprintf(":%s", PORT)
	return &http.Server{
		Addr:    addr,
		Handler: logging(server.Router),
	}
}

func StartServer(server *http.Server) {
	// Log the server's address
	log.Info().Msgf("Server is running on http://localhost%s\n", server.Addr)

	// Start the server

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {

		log.Error().Err(err).Msg("Error starting server")
	}

	// Listen for interrupt signals to gracefully shut down the server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	// Log the shutdown process
	log.Info().Msg("Shutting down the server...")

	// Create a context with a timeout for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := server.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Error during server shutdown")
	} else {
		log.Info().Msg("Server gracefully stopped")
	}
}
