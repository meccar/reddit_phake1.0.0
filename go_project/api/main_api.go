package api

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/joho/godotenv"
)

var (
	DB_URL, DB_DRIVER, JWT_SECRET_KEY, PORT string
)

func init() {
	err := godotenv.Load("../environment.env")
	if err != nil {
		log.Fatal(err)
	}

	DB_URL = os.Getenv("DB_URL")
	fmt.Printf("\n DB_URL : %+v\n", DB_URL)

	DB_DRIVER = os.Getenv("DB_DRIVER")
	fmt.Printf("\n DB_DRIVER : %+v\n", DB_DRIVER)

	PORT = os.Getenv("PORT")
	fmt.Printf("\n PORT : %+v\n", PORT)

	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	fmt.Printf("\n DB_URL : %+v\n", JWT_SECRET_KEY)

}

func GenerateAuthToken() *jwtauth.JWTAuth {
	tokenAuth := jwtauth.New("HS256", []byte(JWT_SECRET_KEY), nil)
	return tokenAuth
}

func main_api() {
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)
}
