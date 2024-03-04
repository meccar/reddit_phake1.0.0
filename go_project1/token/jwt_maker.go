package token

import (
	"context"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"reflect"
	"time"
	util "util"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/rs/zerolog/log"
)

// "encoding/json"

// "log"

// var tokenAuth JWTAuth
var (
	// JWT_SECRET_KEY string
	TokenAuthHS256        *JWTAuth
	TokenAuthRS256        *JWTAuth
	TokenSecret           string
	PrivateKeyRS256String = `-----BEGIN RSA PRIVATE KEY-----
MIIBOwIBAAJBALxo3PCjFw4QjgOX06QCJIJBnXXNiEYwDLxxa5/7QyH6y77nCRQy
J3x3UwF9rUD0RCsp4sNdX5kOQ9PUyHyOtCUCAwEAAQJARjFLHtuj2zmPrwcBcjja
IS0Q3LKV8pA0LoCS+CdD+4QwCxeKFq0yEMZtMvcQOfqo9x9oAywFClMSlLRyl7ng
gQIhAOyerGbcdQxxwjwGpLS61Mprf4n2HzjwISg20cEEH1tfAiEAy9dXmgQpDPir
C6Q9QdLXpNgSB+o5CDqfor7TTyTCovsCIQDNCfpu795luDYN+dvD2JoIBfrwu9v2
ZO72f/pm/YGGlQIgUdRXyW9kH13wJFNBeBwxD27iBiVj0cbe8NFUONBUBmMCIQCN
jVK4eujt1lm/m60TlEhaWBC3p+3aPT2TqFPUigJ3RQ==
-----END RSA PRIVATE KEY-----
`

	PublicKeyRS256String = `-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBALxo3PCjFw4QjgOX06QCJIJBnXXNiEYw
DLxxa5/7QyH6y77nCRQyJ3x3UwF9rUD0RCsp4sNdX5kOQ9PUyHyOtCUCAwEAAQ==
-----END PUBLIC KEY-----
`
)

// func init() {
// 	config, err := util.Init()
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot load config")
// 	}

// 	// JWT_SECRET_KEY = os.Getenv("JWT_SECRET_KEY")
// 	// fmt.Printf("\n JWT_SECRET_KEY : %+v\n", JWT_SECRET_KEY)
// 	TokenSecret = []byte(config.JWTSecretKey)
// 	TokenAuthHS256 = New(jwa.HS256.String(), TokenSecret, nil, jwt.WithAcceptableSkew(30*time.Second))
// 	// fmt.Printf("TokenAuthHS256 is %v\n\n", TokenAuthHS256)

// 	// _, tokenString, _ := TokenAuthHS256.Encode(map[string]interface{}{"user_id": 123})
// 	// fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

// }

func init() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	TokenSecret = config.TokenSymmetricKey

	// TokenAuthHS256 = New(jwa.HS256.String(), TokenSecret, nil, jwt.WithAcceptableSkew(30*time.Second))
}

func (t *JWTAuth) MakeToken(name string) (string, error) {
	fmt.Printf("\n TokenSecret %v\n", TokenSecret)

	claims := map[string]interface{}{"username": name}
	fmt.Printf("\n MakeToken claims %v\n", claims)

	privateKeyBlock, _ := pem.Decode([]byte(PrivateKeyRS256String))
	fmt.Printf("\n MakeToken privateKeyBlock %v\n", privateKeyBlock)

	privateKey, err := x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	fmt.Printf("\n MakeToken privateKey %v\n", privateKey)

	if err != nil {
		log.Error().Err(err)
	}

	publicKeyBlock, _ := pem.Decode([]byte(PublicKeyRS256String))
	fmt.Printf("\n MakeToken publicKeyBlock %v\n", publicKeyBlock)

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	fmt.Printf("\n MakeToken publicKey %v\n", publicKey)

	if err != nil {
		log.Error().Err(err)
	}

	TokenAuthRS256 = New(jwa.RS256.String(), privateKey, publicKey)
	fmt.Printf("\n MakeToken TokenAuthRS256 %v\n", TokenAuthRS256)

	_, tokenString, err := TokenAuthRS256.Encode(claims)
	if err != nil {
		log.Error().Err(err).Msg("Failed to encode claims")
	}

	token, err := TokenAuthRS256.Decode(tokenString)
	fmt.Printf("\n MakeToken token %v\n", token)

	if err != nil {
		log.Error().Err(err).Msg("Failed to decode token string")
	}

	tokenClaims, err := token.AsMap(context.Background())
	fmt.Printf("\n MakeToken tokenClaims %v\n", tokenClaims)

	if err != nil {
		log.Error().Err(err)
	}

	if !reflect.DeepEqual(claims, tokenClaims) {
		log.Info().Msg("The decoded claims don't match the original ones\n")
	}

	// tokenJSON, err := json.Marshal(tokenClaims)
	// fmt.Printf("\n MakeToken tokenJSON %v\n", tokenJSON)

	// if err != nil {
	// 	log.Error().Err(err)
	// 	return "", err
	// }

	return tokenString, nil
}

// func (t *JWTAuth) MakeToken(name string) (string, error) {
// 	claims := map[string]interface{}{"username": name}
// 	_, tokenString, err := t.Encode(claims)
// 	if err != nil {
// 		return "", err
// 	}
// 	return tokenString, nil
// }

// func (t *JWTAuth) MakeToken(name string) (string, error) {
// 	claims := map[string]interface{}{"username": name}
// 	tokenString := newJwtToken(TokenSecret, claims)
// 	fmt.Printf("tokenString %s\n\n", tokenString)
// 	return tokenString, nil
// }

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

func newJwtToken(secret []byte, claims ...map[string]interface{}) string {
	token := jwt.New()
	if len(claims) > 0 {
		for k, v := range claims[0] {
			token.Set(k, v)
		}
	}

	tokenPayload, err := jwt.Sign(token, jwt.WithKey(jwa.HS256, secret))
	if err != nil {
		log.Error().Err(err).Msg("Error signing JWT")
	}
	return string(tokenPayload)
}

// func newAuthHeader(claims ...map[string]interface{}) http.Header {
// 	h := http.Header{}
// 	h.Set("Authorization", "BEARER "+newJwtToken(TokenSecret, claims...))
// 	return h
// }

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
