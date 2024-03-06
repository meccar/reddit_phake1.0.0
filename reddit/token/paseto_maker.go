package token

import (
	"fmt"
	"time"
	"net/http"

	"github.com/rs/zerolog/log"
	"aidanwoods.dev/go-paseto"
)

func CreateToken(username string, role string, duration time.Duration, w http.ResponseWriter) {
	payload, err := NewPayload(username, role, duration)
	if err != nil {
		log.Error().Err(err)
		return 
	}
	token := paseto.NewToken()
	
	token.SetString("username", payload.Username)
	token.SetString("role", payload.Role)

	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public() 
	fmt.Println(publicKey)
	
	signed := token.V4Sign(secretKey, nil)

	SetPasetoCookie(w, signed, role, int(1*time.Minute.Seconds()))
	// return signed, payload, err
}

func SetPasetoCookie(w http.ResponseWriter, token, role string, duration int) {
	http.SetCookie(w, &http.Cookie{
		Name:     "paseto",
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

// VerifyToken checks if the token is valid or not
// func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
// 	payload := &Payload{}

// 	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
// 	if err != nil {
// 		return nil, ErrUnauthorized
// 	}

// 	err = payload.Valid()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return payload, nil
// }