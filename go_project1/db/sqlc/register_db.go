package db

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"

	"regexp"
	util "util"
)

type CreateAccountTxParams struct {
	createAccountParams
	Errors map[string]string
}

type createAccountTxResult struct {
	Account *Account
}

var (
	rxAdmin    = regexp.MustCompile(`^[a-zA-Z0-9._-]+@tafviet.com$`)
	rxUsername = regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
	// rxPassword = regexp.MustCompile(`^[a-zA-Z0-9!@#$%^&*()-_+=]*[a-z][a-zA-Z0-9!@#$%^&*()-_+=]*[A-Z][a-zA-Z0-9!@#$%^&*()-_+=]*[0-9][a-zA-Z0-9!@#$%^&*()-_+=]*[!@#$%^&*()-_+=][a-zA-Z0-9!@#$%^&*()-_+=]{11,}$`)
)

func (register *CreateAccountTxParams) ValidateRegister() bool {
	register.Errors = make(map[string]string)

	matchEmail := rxUsername.Match([]byte(register.Username))
	matchAdmin := rxAdmin.Match([]byte(register.Username))
	// matchPassword := rxPassword.Match([]byte(register.Password))

	fmt.Printf("\n ValidateRegister register %v\n", register)

	if matchAdmin {
		register.Role = "admin"
	} else {
		register.Role = "user"
	}

	if !matchEmail {
		register.Errors["UsernameRegister"] = "Please enter a valid email address"
	}

	// if !matchPassword {
	// 	register.Errors["PasswordRegister"] = "Please enter a valid password"
	// }

	return len(register.Errors) == 0
}

func (h *Handlers) CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (createAccountTxResult, error) {
	var (
		result   createAccountTxResult
		randomID string
	)

	// Submit the account to the database

	if _, err := h.queries.authUsername(ctx, arg.Username); err == nil {
		fmt.Printf("\n VerifyLogin err %v\n", err)
		log.Info().Msg("Email is already registered")
	}

	// fmt.Printf("\n VerifyLogin Username %v\n", rows)

	// for i := 0; i < len(rows); i++ {
	// 	// fmt.Printf("\n loop ID %v\n", rows[i])
	// 	if params.ID == rows[i] {
	// 		// fmt.Printf("\n loop params.ID %v\n", params.ID)
	// 		params.ID++
	// 	}

	// }
	randomID = util.GenerateID()

	// for {
	// randomID = util.GenerateID()
	//
	// rows, _ := h.queries.getAccountID(ctx, randomID)
	// if rows == 0 {
	// break
	// }
	// }
	fmt.Printf("\n loop randomID %v\n", randomID)

	params := createAccountParams{
		ID:       randomID,
		Role:     arg.Role,
		Username: arg.Username,
		Password: arg.Password,
	}

	Account, err := h.queries.createAccount(ctx, params)

	if err != nil {
		return result, err
	}

	result.Account = &Account
	return result, nil
}
