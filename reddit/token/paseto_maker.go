package token

import (
	"fmt"
	"time"
	
	"aidanwoods.dev/go-paseto"
)

func CreateToken(username string, role string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, role, duration)
	if err != nil {
		return "", payload, err
	}
	token := paseto.NewToken()
	
	token.SetString("username", payload.Username)
	token.SetString("role", payload.Role)

	secretKey := paseto.NewV4AsymmetricSecretKey()
	publicKey := secretKey.Public() 
	fmt.Println(publicKey)
	
	signed := token.V4Sign(secretKey, nil)
	return signed, payload, err
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