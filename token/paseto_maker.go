package token

import (
	"fmt"
	"time"
	"net/http"

	"github.com/rs/zerolog/log"
	"aidanwoods.dev/go-paseto"
)

func CreateToken(username string, role string, duration time.Duration, w http.ResponseWriter) error {

	payload, err := NewPayload(username, role, duration)
	if err != nil {
		log.Error().Err(err)
		return err
	}
	token := paseto.NewToken()
	token.SetAudience("audience")
	token.SetJti("identifier")
	token.SetIssuer("issuer")
	token.SetSubject("subject")
	token.SetString("username", payload.Username)
	token.SetString("role", payload.Role)

	secretKey := paseto.NewV4AsymmetricSecretKey()
	fmt.Println("\n secretKey: ",secretKey)

	publicKey := secretKey.Public() 
	fmt.Println("\n publicKey: ",publicKey)

	signed := token.V4Sign(secretKey, nil)
	fmt.Println("\n signed: ",signed)

	SetPasetoCookie(w, signed, role, int(1*time.Minute.Seconds()))

	parser := paseto.NewParserWithoutExpiryCheck()

	parsetoken, err := parser.ParseV4Public(publicKey, signed, nil)
	if err != nil {
		log.Error().Err(err)
		return err
	}
	fmt.Println("\n parsetoken: ",parsetoken)
	fmt.Println("\n string(token.ClaimsJSON()): ",string(token.ClaimsJSON()))
	fmt.Println("\n string(token.Footer()): ",string(token.Footer()))


	return err
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

func pasetoFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("paseto")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// VerifyToken checks if the token is valid or not
// func VerifyPaseto(r *http.Request) {
//   	signed := pasetoFromCookie(r)
// 	fmt.Println("signed: ",signed)
	
	// parser := paseto.NewParser()


	// token, err := parser.ParseV4Public(publicKey, signed, nil)
	// fmt.Println("token: ",token)

	// if err != nil {
	// 	log.Error().Err(err)
	// 	return 
	// }
// }

// publicKey, err := paseto.NewV4AsymmetricPublicKeyFromHex("1eb9dbbbbc047c03fd70604e0071f0987e16b28b757225c11f00415d0e20b1a2") // this wil fail if given key in an invalid format
// signed := "v4.public.eyJkYXRhIjoidGhpcyBpcyBhIHNpZ25lZCBtZXNzYWdlIiwiZXhwIjoiMjAyMi0wMS0wMVQwMDowMDowMCswMDowMCJ9v3Jt8mx_TdM2ceTGoqwrh4yDFn0XsHvvV_D0DtwQxVrJEBMl0F2caAdgnpKlt4p7xBnx1HcO-SPo8FPp214HDw.eyJraWQiOiJ6VmhNaVBCUDlmUmYyc25FY1Q3Z0ZUaW9lQTlDT2NOeTlEZmdMMVc2MGhhTiJ9"

// parser := paseto.NewParserWithoutExpiryCheck() // only used because this example token has expired, use NewParser() (which checks expiry by default)
// token, err := parser.ParseV4Public(publicKey, signed, nil) // this will fail if parsing failes, cryptographic checks fail, or validation rules fail

// // the following will succeed
// require.JSONEq(t,
//     "{\"data\":\"this is a signed message\",\"exp\":\"2022-01-01T00:00:00+00:00\"}",
//     string(token.ClaimsJSON()),
// )
// require.Equal(t,
//     "{\"kid\":\"zVhMiPBP9fRf2snEcT7gFTioeA9COcNy9DfgL1W60haN\"}",
//     string(token.Footer()),
// )
// require.NoError(t, err)