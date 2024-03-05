package token

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"net/http"
	"time"

	util "util"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/rs/zerolog/log"
	"github.com/gin-gonic/gin"

)

var (
	// JWT_SECRET_KEY string
	TokenAuthHS256        *JWTAuth
	TokenAuthRS256        *JWTAuth
	TokenSecret           string
// 	PrivateKeyRS256String = `-----BEGIN RSA PRIVATE KEY-----
// MIIBOwIBAAJBALxo3PCjFw4QjgOX06QCJIJBnXXNiEYwDLxxa5/7QyH6y77nCRQy
// J3x3UwF9rUD0RCsp4sNdX5kOQ9PUyHyOtCUCAwEAAQJARjFLHtuj2zmPrwcBcjja
// IS0Q3LKV8pA0LoCS+CdD+4QwCxeKFq0yEMZtMvcQOfqo9x9oAywFClMSlLRyl7ng
// gQIhAOyerGbcdQxxwjwGpLS61Mprf4n2HzjwISg20cEEH1tfAiEAy9dXmgQpDPir
// C6Q9QdLXpNgSB+o5CDqfor7TTyTCovsCIQDNCfpu795luDYN+dvD2JoIBfrwu9v2
// ZO72f/pm/YGGlQIgUdRXyW9kH13wJFNBeBwxD27iBiVj0cbe8NFUONBUBmMCIQCN
// jVK4eujt1lm/m60TlEhaWBC3p+3aPT2TqFPUigJ3RQ==
// -----END RSA PRIVATE KEY-----
// `

// 	PublicKeyRS256String = `-----BEGIN PUBLIC KEY-----
// MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBALxo3PCjFw4QjgOX06QCJIJBnXXNiEYw
// DLxxa5/7QyH6y77nCRQyJ3x3UwF9rUD0RCsp4sNdX5kOQ9PUyHyOtCUCAwEAAQ==
// -----END PUBLIC KEY-----
// `
)

func init() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	TokenSecret = config.TokenSymmetricKey

	// Decode the PEM-encoded RSA private key
	privateKeyBlock, _ := pem.Decode([]byte(config.PrivateKeyRS256))

	// Parse the decoded private key into an RSA private key object
	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)

	// Check for errors during private key parsing
	if err != nil {
		log.Error().Err(err)
		return
	}

	// Decode the PEM-encoded RSA public key
	publicKeyBlock, _ := pem.Decode([]byte(config.PublicKeyRS256))

	// Parse the decoded public key into an RSA public key object
	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		log.Error().Err(err)
		return
	}

	// Initialize a JWTAuth object with the RSA private and public keys
	TokenAuthRS256 = New(jwa.RS256.String(), privateKey, publicKey)

	// Initialize a JWTAuth object with the HSA key
}

func (t *JWTAuth) MakeToken(id, username, role string) (string, error) {
	// Define the claims for the JWT token
	claims := map[string]interface{}{
		"id":       id,
		"username": username,
		"role":     role,
	}
	SetIssuedNow(claims)
	SetExpiryIn(claims, 1*time.Minute)

	// claims, err := NewPayload(username, role, 1*time.Minute)
	// if err != nil {
	// 	return "", claims, err
	// }

	// Encode the claims into a JWT token string
	_, tokenString, err := TokenAuthRS256.Encode(claims)

	// Check for errors during encoding
	if err != nil {
		log.Error().Err(err).Msg("Failed to encode claims")
		return "", err
	}

	// Return the JWT token string
	return tokenString, nil
}

func (t *JWTAuth) SetJWTCookie(w http.ResponseWriter, token, role string, duration int) {
	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		HttpOnly: true,
		MaxAge:   duration,
		// SameSite: http.SameSiteLaxMode,
		SameSite: http.SameSiteStrictMode,
		// Uncomment below for HTTPS:
		// Secure: true,
		Value: token,
		Path: "/",
	})
}

func DeleteJWTCookie(c *gin.Context, token string) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "jwt",
		Value:   token,
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
		Path:    "/",
	})
}

func GetClaims(r *http.Request) (map[string]interface{}, error) {
	// Get token string from cookie
	tokenString := TokenFromCookie(r)

	// Decode token
	token, err := TokenAuthRS256.Decode(tokenString)
	if err != nil {
		return nil, err
	}

	// Convert token to claims
	tokenClaims, err := token.AsMap(context.Background())
	if err != nil {
		return nil, err
	}

	return tokenClaims, nil
}

// func JWTVerifierMiddleware(ja *JWTAuth) gin.HandlerFunc {
//     verifier := Verifier(ja)
//     return func(c *gin.Context) {
//         // Convert the context to http.Handler
//         h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//             c.Next()
//         })
//         // Pass the converted handler to the verifier
//         verifier(h).ServeHTTP(c.Writer, c.Request)
//     }
// }

// func JWTAuthenticatorMiddleware(ja *JWTAuth) gin.HandlerFunc {
//     authenticator := Authenticator(ja)
//     return func(c *gin.Context) {
//         // Convert the context to http.Handler
//         h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//             c.Next()
//         })
//         // Pass the converted handler to the authenticator
//         authenticator(h).ServeHTTP(c.Writer, c.Request)
//     }
// }


// func LoggedInRedirector(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token, _, _ := FromContext(r.Context())

// 		if token != nil && jwt.Validate(token) == nil {
// 			http.Redirect(w, r, "/profile", 302)
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }

// func UnloggedInRedirector(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		token, _, _ := FromContext(r.Context())

// 		if token == nil || jwt.Validate(token) != nil {
// 			http.Redirect(w, r, "/login", 302)
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }
