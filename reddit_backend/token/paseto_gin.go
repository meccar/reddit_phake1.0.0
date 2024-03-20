package token

// import (
// 	"context"
// 	"errors"
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/vk-rv/pvx"
// )

// type PasetoAuth struct {
// 	TokenSecret string
// }

// var (
// 	TokenCtxKey = &contextKey{"Token"}
// 	ErrorCtxKey = &contextKey{"Error"}
// )

// var (
// 	ErrUnauthorized = errors.New("token is unauthorized")
// 	ErrExpired      = errors.New("token is expired")
// 	ErrNBFInvalid   = errors.New("token nbf validation failed")
// 	ErrIATInvalid   = errors.New("token iat validation failed")
// 	ErrNoTokenFound = errors.New("no token found")
// )

// func NewPasetoAuth(tokenSecret string) *PasetoAuth {
// 	return &PasetoAuth{
// 		TokenSecret: tokenSecret,
// 	}
// }

// func Verifier(pa *PasetoAuth) gin.HandlerFunc {
// 	return Verify(pa, TokenFromHeader, TokenFromCookie)
// }

// func Verify(pa *PasetoAuth, findTokenFns ...func(r *http.Request) string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx := c.Request.Context()
// 		token, err := VerifyRequest(pa, c.Request, findTokenFns...)
// 		ctx = NewContext(ctx, token, err)
// 		c.Request = c.Request.WithContext(ctx)
// 		c.Next()
// 	}
// }

// func VerifyRequest(pa *PasetoAuth, r *http.Request, findTokenFns ...func(r *http.Request) string) (pvx.Token, error) {
// 	var tokenString string

// 	// Extract token string from the request by calling token find functions in
// 	// the order they were provided. Further extraction stops if a function
// 	// returns a non-empty string.
// 	for _, fn := range findTokenFns {
// 		tokenString = fn(r)
// 		if tokenString != "" {
// 			break
// 		}
// 	}
// 	if tokenString == "" {
// 		return nil, ErrNoTokenFound
// 	}

// 	return VerifyToken(pa, tokenString)
// }

// func VerifyToken(pa *PasetoAuth, tokenString string) (pvx.Token, error) {
// 	// Verify the PASETO token
// 	pv4 := pvx.NewPV4Public()
// 	token, err := pv4.Verify(tokenString, []byte(pa.TokenSecret))
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Valid!
// 	return token, nil
// }

// func CreatePasetoToken(username string, role string, duration time.Duration, c *gin.Context) error {
// 	now := time.Now()
// 	expiry := now.Add(duration)

// 	claims := pvx.Claims{
// 		"username": username,
// 		"role":     role,
// 		"exp":      expiry,
// 	}

// 	pv4 := pvx.NewPV4Public()
// 	token, err := pv4.Sign(claims, []byte(TokenSecret))
// 	if err != nil {
// 		return err
// 	}

// 	SetPasetoCookie(c, token, role, int(duration.Seconds()))

// 	return nil
// }

// func NewContext(ctx context.Context, t pvx.Token, err error) context.Context {
// 	ctx = context.WithValue(ctx, TokenCtxKey, t)
// 	ctx = context.WithValue(ctx, ErrorCtxKey, err)
// 	return ctx
// }

// func FromContext(ctx context.Context) (pvx.Token, map[string]interface{}, error) {
// 	token, _ := ctx.Value(TokenCtxKey).(pvx.Token)

// 	var err error
// 	var claims map[string]interface{}

// 	if token != nil {
// 		claims = token.Claims()
// 	} else {
// 		claims = make(map[string]interface{})
// 	}

// 	err, _ = ctx.Value(ErrorCtxKey).(error)

// 	return token, claims, err
// }

// type contextKey struct {
// 	name string
// }

// func (k *contextKey) String() string {
// 	return "paseto context value " + k.name
// }

// func TokenFromCookie(r *http.Request) string {
// 	cookie, err := r.Cookie("paseto")
// 	if err != nil {
// 		return ""
// 	}
// 	return cookie.Value
// }

// func TokenFromHeader(r *http.Request) string {
// 	bearer := r.Header.Get("Authorization")
// 	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
// 		return bearer[7:]
// 	}
// 	return ""
// }

// func TokenFromQuery(r *http.Request) string {
// 	return r.URL.Query().Get("paseto")
// }
