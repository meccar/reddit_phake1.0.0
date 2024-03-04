package db

import (
	"context"
	"fmt"
)

// "golang.org/x/crypto/bcrypt"
// "database/sql"
// "errors"
// "fmt"
// "log"
type LoginAccountTxParams struct {
	loginAccountParams
	// Password string `schema:"Password"`
	Errors map[string]string
}

type SessionTxResult struct {
	Session *Session
}

func (h *Handlers) verifyUsername(ctx context.Context, username string) bool {
	// Retrieve the hashed password from the database

	// fmt.Printf("\n verifyUsername ctx %v\n", ctx)
	// fmt.Printf("\n verifyUsername username %v\n", username)
	// fmt.Printf("\n verifyUser password %v\n", password)

	_, err := h.queries.authUsername(ctx, username)

	// fmt.Printf("\n verifyUsername err %v\n", err)

	if err != nil {

		// fmt.Printf("\n verifyUsername false \n")
		return false
	}

	return true
}

func (h *Handlers) verifyPassword(ctx context.Context, username string, password string) bool {
	// Retrieve the hashed password from the database

	// fmt.Printf("\n verifyPassword ctx %v\n", ctx)
	// fmt.Printf("\n verifyPassword username %v\n", username)
	// fmt.Printf("\n verifyPassword password %v\n", password)

	_, err := h.queries.authPassword(ctx, username)

	// fmt.Printf("\n verifyPassword err %v\n", err)
	// fmt.Printf("\n verifyPassword password %v\n", row)

	if err != nil {

		// fmt.Printf("\n verifyPassword false \n")
		return false
	}
	// err = bcrypt.CompareHashAndPassword([]byte(row), []byte(password))
	// fmt.Printf("\n verifyPassword password err %v\n", err)

	return true
}

func (h *Handlers) VerifyLogin(ctx context.Context, arg *LoginAccountTxParams) bool {

	arg.Errors = make(map[string]string)

	// fmt.Printf("\n VerifyLogin Username %v\n", arg.Username)
	// fmt.Printf("\n VerifyLogin Password %v\n", arg.Password)
	bcryptStr, ok := arg.Bcrypt.(string)
	if !ok {
		// Handle the case where arg.Bcrypt is not a string
		arg.Errors["PasswordLogin"] = "Invalid password"
		return false
	}
	fmt.Printf("\n VerifyLogin bcryptStr %v\n", bcryptStr)

	if !h.verifyUsername(ctx, arg.Username) {
		arg.Errors["UsernameLogin"] = "Invalid username"
	}
	if !h.verifyPassword(ctx, arg.Username, bcryptStr) {
		arg.Errors["PasswordLogin"] = "Invalid password"
	}

	return len(arg.Errors) == 0
}
