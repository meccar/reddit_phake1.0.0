package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/gin-gonic/gin"
	db "sqlc"
	jwtauth "token"
	util "util"
)
// "github.com/go-chi/chi"


var (
	PORT string
)

// type Server struct {
// 	Config         util.Config
// 	Router         *chi.Mux
// 	DbHandler      *db.Handlers
// 	TokenAuthRS256 *jwtauth.JWTAuth
// }
type Server struct {
	Config         util.Config
	Router         *gin.Engine
	DbHandler      *db.Handlers
	TokenAuthRS256 *jwtauth.JWTAuth
}

func CreateServer(Config util.Config, dbHandler *db.Handlers) (*Server, error) {
	server := &Server{
		Config:         Config,
		Router:         gin.Default(),
		DbHandler:      dbHandler,
		TokenAuthRS256: jwtauth.TokenAuthRS256,
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

		defer func() {
			if r := recover(); r != nil {
				log.Error().Interface("panic", r).Msg("Recovered from panic")
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) SetupServer() *http.Server {
	addr := fmt.Sprintf(":%s", PORT)
	return &http.Server{
		Addr:              addr,
		Handler:           logging(server.Router),
		// IdleTimeout:       30 * time.Second,
		// ReadHeaderTimeout: 5 * time.Second,
		// ReadTimeout:       10 * time.Second,
		// WriteTimeout:      10 * time.Second,
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


