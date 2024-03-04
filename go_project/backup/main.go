package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/schema"

	db "sqlc"
)

var templates *template.Template

// type Form struct {
// 	ViewerName string
// 	Email      string
// 	Phone      string
// }

func main() {
	templates = template.Must(template.ParseGlob("./templates/*.html"))

	// repo := db.dbMain()
	// handlers := db.NewHandlers(repo)

	r := setupRoutes()
	server := setupServer(r)
	startServer(server)
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

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		handler.ServeHTTP(writer, request)
	})
}

func setupRoutes() *chi.Mux {

	r := chi.NewRouter()

	r.Use(Logger)
	r.Use(middleware.Logger)

	// Serve public files for any path
	r.Get("/*", publicHandler)

	// Handle the root path
	r.Get("/", pageHandler("home"))

	r.Route("/lienhe", func(r chi.Router) {
		r.Get("/", pageHandler("lienhe"))
		r.Post("/", formHandler)
		r.Get("/thankyou", pageHandler("thankyou"))
	})

	return r
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	err := templates.ExecuteTemplate(w, filename, data)
	if err != nil {
		log.Printf("Error rendering template %s: %v\n", filename, err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}
}

// GET files from public
func publicHandler(w http.ResponseWriter, r *http.Request) {
	http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))).ServeHTTP(w, r)
}

// PATHS handlers
func pageHandler(pageName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, pageName, nil)
	}
}

// Form handler
func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusInternalServerError)
		return
	}

	msg := &db.SubmitFormTxParams{}

	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(msg, r.PostForm)
	if decodeErr != nil {
		log.Println("Error mapping parsed form data to struct:", decodeErr)
		http.Error(w, "Error processing form data", http.StatusInternalServerError)
		return
	}

	if msg.Validate() == false {
		log.Println(msg)
		render(w, "lienhe", msg)
		return
	}

	repo := &db.SQLRepo{}
	fmt.Printf("repo main %v", repo)

	if _, err := repo.SubmitFormTx(r.Context(), *msg); err != nil {
		log.Print(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/thankyou", http.StatusSeeOther)
}

func setupServer(r *chi.Mux) *http.Server {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "9595"
	}

	addr := fmt.Sprintf(":%s", port)
	return &http.Server{
		Addr:         addr,
		Handler:      logging(r),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  15 * time.Second,
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
