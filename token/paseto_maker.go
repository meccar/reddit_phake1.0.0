package token

import (
	"fmt"
	"net/http"
	"errors"
	"time"
	"crypto/ed25519"

	util "util"

	"github.com/rs/zerolog/log"
	"github.com/vk-rv/pvx"
	"github.com/gin-gonic/gin"
)
// "encoding/hex"

// "aidanwoods.dev/go-paseto"

// func CreateToken(username string, role string, duration time.Duration, w http.ResponseWriter) error {

// 	payload, err := NewPayload(username, role, duration)
// 	if err != nil {
// 		log.Error().Err(err)
// 		return err
// 	}
// 	token := paseto.NewToken()
// 	token.SetAudience("audience")
// 	token.SetJti("identifier")
// 	token.SetIssuer("issuer")
// 	token.SetSubject("subject")
// 	token.SetString("username", payload.Username)
// 	token.SetString("role", payload.Role)

// 	secretKey := paseto.NewV4AsymmetricSecretKey()
// 	fmt.Println("\n secretKey: ", secretKey)

// 	publicKey := secretKey.Public()
// 	fmt.Println("\n publicKey: ", publicKey)

// 	signed := token.V4Sign(secretKey, nil)
// 	fmt.Println("\n signed: ", signed)

// 	SetPasetoCookie(w, signed, role, int(1*time.Minute.Seconds()))

// 	parser := paseto.NewParserWithoutExpiryCheck()

// 	parsetoken, err := parser.ParseV4Public(publicKey, signed, nil)
// 	if err != nil {
// 		log.Error().Err(err)
// 		return err
// 	}
// 	fmt.Println("\n parsetoken: ", parsetoken)
// 	fmt.Println("\n string(token.ClaimsJSON()): ", string(token.ClaimsJSON()))
// 	fmt.Println("\n string(token.Footer()): ", string(token.Footer()))

// 	return err
// }


// k, err := hex.DecodeString(TokenSecret)
// if err != nil {
	// return err 
// }
// fmt.Println("\n k: ", k)
// 
// symK := pvx.NewSymmetricKey(k, pvx.Version4)
// fmt.Println("\n symK: ", symK)
// 
// pv4 := pvx.NewPV4Local()
// fmt.Println("\n pv4: ", pv4)

// token, err := pv4.Encrypt(symK, claims, pvx.WithAssert([]byte("test")))
// if err != nil {
	// return err
// }
// fmt.Println("\n <<< after Encrypt 1 token: ", token)
// 
// cc := MyClaims{}
// 
// err = pv4.
	// Decrypt(token, symK, pvx.WithAssert([]byte("test"))).
	// ScanClaims(&cc)
// if err != nil {
	// return err 
// }
// work with cc claims ...
// 
// or without assert
// token, err := pv4.Encrypt(symK, claims)
// if err != nil {
	// return err
// }
// fmt.Println("\n <<< after Encrypt 2 token: ", token)
// 
// err = pv4.Decrypt(token, symK).ScanClaims(&cc)

type AdditionalClaims struct {
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Date time.Time `json:"date"`
}

type MyClaims struct {
	RegisteredClaims pvx.RegisteredClaims
	AdditionalClaims
} 

func (c *MyClaims) Valid() error {

	validationErr := &pvx.ValidationError{}
	
	// first, check the validity of registered claims
	if err := c.RegisteredClaims.Valid(); err != nil {
		errors.As(err, &validationErr)
	}
	
	//  then, perform custom validation
	
	
	return nil 
	
}

func CreateToken(username string, role string, duration time.Duration, c *gin.Context) error {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	TokenSecret = config.TokenSymmetricKey
	now := time.Now()
	claims := &MyClaims{
		RegisteredClaims: pvx.RegisteredClaims{
			IssuedAt:   pvx.TimePtr(now.Add(time.Minute * 60)),
			NotBefore:  pvx.TimePtr(now.Add(time.Minute * 60)),
			Expiration: pvx.TimePtr(now.AddDate(0, 0, -1)),
		},
		AdditionalClaims: AdditionalClaims{Username: username, Role: role, Date: time.Now().Add(time.Minute * 60)},
	}

	pv4 := pvx.NewPV4Public()

	publicKey, privateKey, _ := ed25519.GenerateKey(nil)
	fmt.Println("\n publicKey: ", publicKey)
	fmt.Println("\n privateKey: ", privateKey)

	sk := pvx.NewAsymmetricSecretKey(privateKey, pvx.Version4)
	fmt.Println("\n sk: ", sk)

	pk := pvx.NewAsymmetricPublicKey(publicKey, pvx.Version4)
	fmt.Println("\n pk: ", pk)


	token, err := pv4.Sign(sk, claims, pvx.WithAssert([]byte("test")))
	if err != nil {
		log.Fatal().Err(err)
	}
	fmt.Println("\n <<< after Sign token: ", token)



	c.Set("publicKey", publicKey)
	SetPasetoCookie(c, token, role, int(duration.Seconds()))
	return err
}

func SetPasetoCookie(c *gin.Context, token, role string, duration int) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "paseto",
		HttpOnly: true,
		MaxAge:   duration,
		// SameSite: http.SameSiteLaxMode,
		SameSite: http.SameSiteStrictMode,
		// Uncomment below for HTTPS:
		// Secure: true,
		Value: token,
		Path:  "/",
	})
}

func pasetoFromCookie(r *http.Request) string {
	cookie, err := r.Cookie("paseto")
	if err != nil {
		return ""
	}
	return cookie.Value
}

// // VerifyToken checks if the token is valid or not
// func VerifyPaseto(c *gin.Context) {
//   	signed := pasetoFromCookie(c.Request)
// 	fmt.Println("signed: ",signed)

// 	publicKey, exists := c.Get("publicKey")
//     if !exists {
//         // Handle case where public key is not found
//         c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Public key not found"})
//         return
//     }

// 	if err := pv4.Verify(token, pk, pvx.WithAssert([]byte("test"))).ScanClaims(claims); err != nil {
// 		errors.Errorf("can't verify paseto token, err is %v", err.Err())
// 	}
// 	if tk.HasFooter() {
// 		errors.Errorf("footer was not passed to the library")
// 	}

// 	c.Next()

// }

func VerifyPaseto(pv4 *pvx.ProtoV4Public) gin.HandlerFunc  {
    return func(c *gin.Context) {
        token := pasetoFromCookie(c.Request)
        if token == "" { 
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }
		fmt.Println("\n <<< after pasetoFromCookie: ", token)

        
        // Get the public key from the request context
		value, exists := c.Get("publicKey")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Missing Credentials"})
		}
		fmt.Println("\n <<< publicKey: ", value)
        
		publicKey, ok := value.(*pvx.AsymPublicKey)
		if !ok {
			// Handle the case where the value stored as "publicKey" is not of type *pvx.AsymPublicKey
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid public key type"})
			return
		}
		fmt.Println("\n <<< publicKey: ", publicKey)


        // Verify the token
        if err := pv4.Verify(token, publicKey, pvx.WithAssert([]byte("test"))).Err(); err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
            return
        }
		// if tk.HasFooter() {
		// 	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Token not founded in footer"})
		// 	return
		// }

        // If token is valid, call the next handler
		c.Next()

    }
}

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
