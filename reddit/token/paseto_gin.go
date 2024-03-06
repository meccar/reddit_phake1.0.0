package token

// import (
// 	"context"
// 	"errors"
// 	"net/http"
// 	"strings"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"aidanwoods.dev/go-paseto"
// )

// type PasetoAuth struct {
// 	alg             jwa.SignatureAlgorithm
// 	signKey         interface{}
// 	verifyKey       interface{}
// 	verifier        paseto.NewParser
// 	// verifier        paseto.ParseV4Local
// 	validateOptions []paseto.v4Verify
// 	// validateLocalOptions []paseto.ParseV4Local
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
// 	ErrAlgoInvalid  = errors.New("algorithm mismatch")
// )

// func New(alg string, signKey interface{}, verifyKey interface{}, validateOptions ...paseto.v4Verify) *PasetoAuth {
// 	ja := &PasetoAuth{
// 		alg:             jwa.SignatureAlgorithm(alg),
// 		signKey:         signKey,
// 		verifyKey:       verifyKey,
// 		validateOptions: validateOptions,
// 	}

// 	if ja.verifyKey != nil {
// 		ja.verifier = jwt.WithKey(ja.alg, ja.verifyKey)
// 	} else {
// 		ja.verifier = jwt.WithKey(ja.alg, ja.signKey)
// 	}

// 	return ja
// }

// func Verifier(ja *PasetoAuth) gin.HandlerFunc {
// 	return Verify(ja, TokenFromHeader, TokenFromCookie)
// }

// func Verify(ja *PasetoAuth, findTokenFns ...func(r *http.Request) string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		ctx := c.Request.Context()
// 		token, err := VerifyRequest(ja, c.Request, findTokenFns...)
// 		ctx = NewContext(ctx, token, err)
// 		c.Request = c.Request.WithContext(ctx)
// 		c.Next()
// 	}
// }

// func VerifyRequest(ja *PasetoAuth, r *http.Request, findTokenFns ...func(r *http.Request) string) (jwt.Token, error) {
// 	var tokenString string

// 	// Extract token string from the request by calling token find functions in
// 	// the order they where provided. Further extraction stops if a function
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

// 	return VerifyToken(ja, tokenString)
// }

// func VerifyToken(ja *PasetoAuth, tokenString string) (paseto.Token, error) {
// 	// Decode & verify the token
// 	token, err := ja.Decode(tokenString)

// 	if err != nil {
// 		return token, ErrorReason(err)
// 	}

// 	if token == nil {
// 		return nil, ErrUnauthorized
// 	}

// 	if err := paseto.v4Verify(token, ja.validateOptions...); err != nil {
// 		return token, ErrorReason(err)
// 	}

// 	// Valid!
// 	return token, nil
// }

// func (ja *PasetoAuth) Encode(claims map[string]interface{}) (t paseto.Token, tokenString string, err error) {
// 	t = jwt.New()
// 	for k, v := range claims {
// 		t.Set(k, v)
// 	}
// 	payload, err := ja.sign(t)
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	tokenString = string(payload)
// 	return
// }

// func (ja *PasetoAuth) Decode(tokenString string) (paseto.Token, error) {
// 	return ja.parse([]byte(tokenString))
// }

// func (ja *PasetoAuth) ValidateOptions() []paseto.v4Verify {
// 	return ja.validateOptions
// }

// func (ja *PasetoAuth) sign(token paseto.Token) ([]byte, error) {
// 	return paseto.V4Sign(token, jwt.WithKey(ja.alg, ja.signKey))
// }

// func (ja *PasetoAuth) parse(payload []byte) (jwt.Token, error) {
// 	// we disable validation here because we use jwt.Validate to validate tokens
// 	return jwt.Parse(payload, ja.verifier, jwt.WithValidate(false))
// }

// func ErrorReason(err error) error {
// 	switch {
// 	case errors.Is(err, jwt.ErrTokenExpired()), err == ErrExpired:
// 		return ErrExpired
// 	case errors.Is(err, jwt.ErrInvalidIssuedAt()), err == ErrIATInvalid:
// 		return ErrIATInvalid
// 	case errors.Is(err, jwt.ErrTokenNotYetValid()), err == ErrNBFInvalid:
// 		return ErrNBFInvalid
// 	default:
// 		return ErrUnauthorized
// 	}
// }

// func Authenticator(ja *PasetoAuth) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		token, _, err := FromContext(c.Request.Context())

// 		if err != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
// 			return
// 		}

// 		if token == nil || jwt.Validate(token, ja.validateOptions...) != nil {
// 			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
// 			return
// 		}

// 		// Token is authenticated, pass it through
// 		c.Next()
// 	}
// }


// func NewContext(ctx context.Context, t jwt.Token, err error) context.Context {
// 	ctx = context.WithValue(ctx, TokenCtxKey, t)
// 	ctx = context.WithValue(ctx, ErrorCtxKey, err)
// 	return ctx
// }

// func FromContext(ctx context.Context) (jwt.Token, map[string]interface{}, error) {
// 	token, _ := ctx.Value(TokenCtxKey).(jwt.Token)

// 	var err error
// 	var claims map[string]interface{}

// 	if token != nil {
// 		claims, err = token.AsMap(context.Background())
// 		if err != nil {
// 			return token, nil, err
// 		}
// 	} else {
// 		claims = map[string]interface{}{}
// 	}

// 	err, _ = ctx.Value(ErrorCtxKey).(error)

// 	return token, claims, err
// }

// type contextKey struct {
// 	name string
// }

// func (k *contextKey) String() string {
// 	return "pasetoauth context value " + k.name
// }

// func TokenFromCookie(r *http.Request) string {
// 	cookie, err := r.Cookie("jwt")
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
// 	return r.URL.Query().Get("jwt")
// }

// // // UnixTime returns the given time in UTC milliseconds
// func UnixTime(tm time.Time) int64 {
// 	return tm.UTC().Unix()
// }

// // // EpochNow is a helper function that returns the NumericDate time value used by the spec
// func EpochNow() int64 {
// 	return time.Now().UTC().Unix()
// }

// // // ExpireIn is a helper function to return calculated time in the future for "exp" claim
// func ExpireIn(tm time.Duration) int64 {
// 	return EpochNow() + int64(tm.Seconds())
// }

// // // Set issued at ("iat") to specified time in the claims
// func SetIssuedAt(claims map[string]interface{}, tm time.Time) {
// 	claims["iat"] = tm.UTC().Unix()
// }

// // // Set issued at ("iat") to present time in the claims
// func SetIssuedNow(claims map[string]interface{}) {
// 	claims["iat"] = EpochNow()
// }

// // // Set expiry ("exp") in the claims
// func SetExpiry(claims map[string]interface{}, tm time.Time) {
// 	claims["exp"] = tm.UTC().Unix()
// }

// // // Set expiry ("exp") in the claims to some duration from the present time
// func SetExpiryIn(claims map[string]interface{}, tm time.Duration) {
// 	claims["exp"] = ExpireIn(tm)
// }