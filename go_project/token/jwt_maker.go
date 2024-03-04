package token

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// var tokenAuth JWTAuth
var (
	JWT_SECRET_KEY string
	TokenAuthHS256 *JWTAuth
	TokenAuthRS256 *JWTAuth
	TokenSecret    []byte
)

func init() {
	err := godotenv.Load("./environment.env")
	if err != nil {
		log.Fatal(err)
	}

	JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
	// fmt.Printf("\n JWT_SECRET_KEY : %+v\n", JWT_SECRET_KEY)
	TokenSecret = []byte(JWT_SECRET_KEY)
	TokenAuthHS256 = New(jwa.HS256.String(), TokenSecret, nil, jwt.WithAcceptableSkew(30*time.Second))
	fmt.Printf("TokenAuthHS256 is %s\n\n", TokenAuthHS256)

	// _, tokenString, _ := TokenAuthHS256.Encode(map[string]interface{}{"user_id": 123})
	// fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

}

func (t *JWTAuth) MakeToken(name string) (string, error) {
	claims := map[string]interface{}{"username": name}
	_, tokenString, err := t.Encode(claims)
	if err != nil {
		log.Printf("Error encoding token: %v", err)
		return "", err
	}
	fmt.Printf("tokenString %s\n\n", tokenString)
	return tokenString, nil
}

func (t *JWTAuth) SetJWTCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		SameSite: http.SameSiteLaxMode,
		// Uncomment below for HTTPS:
		// Secure: true,
		Name:  "jwt",
		Value: token,
	})
}

func LoggedInRedirector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, _ := FromContext(r.Context())

		if token != nil && jwt.Validate(token) == nil {
			http.Redirect(w, r, "/profile", 302)
		}

		next.ServeHTTP(w, r)
	})
}

func UnloggedInRedirector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, _ := FromContext(r.Context())

		if token == nil || jwt.Validate(token) != nil {
			http.Redirect(w, r, "/login", 302)
		}

		next.ServeHTTP(w, r)
	})
}
