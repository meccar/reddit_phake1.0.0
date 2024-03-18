package db

import (
	"context"
	"errors"
	"fmt"
	"strings"

	util "util"

	"github.com/google/uuid"
)

type CreateAccountTxParams struct {
	createAccountParams
}

type createAccountTxResult struct {
	Account *Account
}

var RegisterRules = map[string]interface{}{
	"Username": "required,email",
	"Password": "required,min=12,max=64",
}

func isAdmin(email string) bool {
	return strings.HasSuffix(email, "@google.com")
}

func (register *CreateAccountTxParams) ValidateRegister(ctx context.Context) error {

	if isAdmin(register.Username) {
		register.Role = "admin"
	} else {
		register.Role = "user"
	}

	errorMsg, err := util.ValidateForm(register, RegisterRules)
	if err != nil {
		// Handle validation error
		return err
	} else if errorMsg != "" {
		// Handle specific error message
		return errors.New(errorMsg)
	}

	// Return a success message indicating validation success
	return nil
}

func (h *Handlers) CreateAccountTx(ctx context.Context, arg CreateAccountTxParams) (createAccountTxResult, error) {
	var result createAccountTxResult

	err := h.execTx(ctx, func(q *Queries) error {
		// Submit the account to the database
		if _, err := q.authUsername(ctx, arg.Username); err == nil {
			return fmt.Errorf("Email is already registered")
		}

		ranID, _ := uuid.NewRandom()

		base64Encoded := util.GetImageBinary("https://tafviet.com/wp-content/uploads/2024/03/profile-picture.jpg")
		// If the photo is empty, set it to the base64-encoded image data
		if len(arg.Photo) == 0 {
			arg.Photo = []byte(base64Encoded)
		}

		params := createAccountParams{
			ID:       ranID,
			Role:     arg.Role,
			Username: arg.Username,
			Password: arg.Password,
			Photo:    arg.Photo,
		}

		Account, err := q.createAccount(ctx, params)
		if err != nil {
			return err
		}

		result.Account = &Account
		return err
	})
	return result, err
}
