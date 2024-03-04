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
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"

	db "sqlc"
	jwtauth "token"
)

var (
	templates               *template.Template
	DB_URL, DB_DRIVER, PORT string
)

type Server struct {
	Router         *chi.Mux
	dbHandler      *db.Handlers
	TokenAuthHS256 *jwtauth.JWTAuth
}

func CreateServer(dbHandler *db.Handlers) *Server {
	return &Server{
		Router:         chi.NewRouter(),
		dbHandler:      dbHandler,
		TokenAuthHS256: jwtauth.TokenAuthHS256,
	}
}

func init() {
	err := godotenv.Load("./environment.env")
	if err != nil {
		log.Fatal(err)
	}

	DB_URL = os.Getenv("DB_URL")
	// fmt.Printf("\n DB_URL : %+v\n", DB_URL)

	DB_DRIVER = os.Getenv("DB_DRIVER")
	// fmt.Printf("\n DB_DRIVER : %+v\n", DB_DRIVER)

	PORT = os.Getenv("PORT")
	// fmt.Printf("\n PORT : %+v\n", PORT)
}

func main() {
	templates = template.Must(template.ParseGlob("./templates/*.html"))
	postgres, err := db.InitialiseDB(DB_URL)

	if err != nil {
		log.Fatal(err.Error())
	}

	queries := db.New(postgres)
	newHandler := db.NewHandlers(queries)

	// r := chi.NewRouter()
	server := CreateServer(newHandler)

	// fmt.Printf("\n main TokenAuthHS256 : %+v\n", jwtauth.TokenAuthHS256)
	// setupRoutes(r, newHandler, jwtauth.TokenAuthHS256)
	server.setupRoutes()
	httpServer := server.setupServer()
	startServer(httpServer)
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		req := fmt.Sprintf("%s %s", r.Method, r.URL)
		log.Println(req)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		// Calculate and log the elapsed time using time.Since
		elapsed := time.Since(start)
		log.Println(req, "completed in", elapsed)
	})
}

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.URL.Path)
		handler.ServeHTTP(writer, request)
	})
}

func (server *Server) setupRoutes() {
	server.Router.Use(Logger)
	server.Router.Use(middleware.Logger)

	// Serve public files for any path
	server.Router.Get("/*", publicHandler)

	// Handle the root path
	server.Router.Get("/", GetHandler("home"))

	server.Router.Route("/lienhe", func(r chi.Router) {
		r.Get("/", GetHandler("lienhe"))
		r.Post("/", server.formHandlerWrapper())

		r.Route("/thankyou", func(r chi.Router) {
			r.Get("/", GetHandler("thankyou"))
		})
	})

	server.Router.Route("/login", func(r chi.Router) {
		r.Get("/", GetHandler("login"))
		r.Post("/", server.loginHandlerWrapper())

		// fmt.Printf("\n Route login token %v\n", TokenAuthHS256)

		r.Route("/admin", func(r chi.Router) {
			// fmt.Printf("\n Route admin token %v\n", server.TokenAuthHS256)
			r.Use(jwtauth.Verifier(server.TokenAuthHS256))
			// r.Use(jwtauth.Authenticator(jwtauth.TokenAuthHS256))

			r.Get("/", GetHandler("admin"))
		})
	})

	server.Router.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Requested URL: %s\n", r.URL.Path)
	})
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
func GetHandler(pageName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		render(w, pageName, nil)
	}
}

// func (server *Server) handlerWrapper(, handlerFunc func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
// return func(w http.ResponseWriter, r *http.Request) {
// handlerFunc(w, r)
// }
// }

func (server *Server) formHandlerWrapper() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.formHandler(w, r)
	}
}

func (server *Server) formHandler(w http.ResponseWriter, r *http.Request) {
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

	if !msg.Validate() {
		log.Println(msg)
		render(w, "lienhe", msg)
		return
	}

	fmt.Printf("\n formHandler dbHandler %v\n", server.dbHandler)

	if _, err := server.dbHandler.SubmitFormTx(r.Context(), *msg); err != nil {
		log.Print(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/lienhe/thankyou", http.StatusSeeOther)
}

func (server *Server) loginHandlerWrapper() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.loginHandler(w, r)
	}
}

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
	if decodeErr != nil {
		log.Println("Error mapping parsed form data to struct:", decodeErr)
		http.Error(w, "Error processing form data", http.StatusInternalServerError)
		return
	}

	if password, ok := r.PostForm["Password"]; ok {
		// Hash the password using bcrypt
		hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(password[0]), bcrypt.DefaultCost)
		if hashErr != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		// Set the hashed password in msg.Bcrypt
		msg.Bcrypt = string(hashedPassword)
	}

	fmt.Printf("\n loginHandler msg 2 %v\n", msg)
	fmt.Printf("\n loginHandler msg.Bcrypt %v\n", msg.Bcrypt)
	fmt.Printf("\n loginHandler msg.Bcrypt %v\n", &msg.Bcrypt)
	// fmt.Printf("\n loginHandler msg.Password %v\n", msg.Password)

	if !server.dbHandler.VerifyLogin(r.Context(), msg) {
		log.Println(msg)
		render(w, "login", msg)
		return
	}
	token, err := server.TokenAuthHS256.MakeToken(msg.Username)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	server.TokenAuthHS256.SetJWTCookie(w, token)
	http.Redirect(w, r, "/login/admin", http.StatusSeeOther)
}

func (server *Server) setupServer() *http.Server {
	addr := fmt.Sprintf(":%s", PORT)
	return &http.Server{
		Addr:    addr,
		Handler: logging(server.Router),
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
