package db

import (
	"context"
	"errors"

	util "util"
)

type LoginAccountTxParams struct {
	loginAccountParams
	// Session Session
}

func (h *Handlers) VerifyLogin(ctx context.Context, arg *LoginAccountTxParams) error {

	if !h.VerifyUsername(ctx, arg.Username) {
		return errors.New("Username is Incorrect")
	}
	if !h.verifyPassword(ctx, arg.Username, arg.Password) {
		return errors.New("Password is Incorrect")
	}

	// Return a success message indicating validation success
	return nil
}

func (h *Handlers) VerifyUsername(ctx context.Context, username string) bool {
	// Retrieve the hashed password from the database
	_, err := h.Queries.authUsername(ctx, username)
	if err != nil {
		return false
	}

	return true
}

func (h *Handlers) verifyPassword(ctx context.Context, username string, password string) bool {
	// Retrieve the hashed password from the database
	checkPassword, err := h.Queries.authPassword(ctx, username)
	if err != nil {
		return false
	}

	if util.CheckPasswordHash(password, checkPassword) {
		return true
	}
	return false
}

func (h *Handlers) GetID(ctx context.Context, arg *LoginAccountTxParams) string {
	id, err := h.Queries.getAccountIDbyUsername(ctx, arg.Username)
	if err != nil {
		return ""
	}
	return id.String()
}

func (h *Handlers) GetRole(ctx context.Context, arg *LoginAccountTxParams) string {
	role, err := h.Queries.getAccountRolebyUsername(ctx, arg.Username)
	if err != nil {
		return ""
	}
	return string(role)
}

// func (h *Handlers) VerifyEmail(ctx context.Context, arg *CreateVerifyEmailTxParams) string {
// 	email, err := h.queries.CreateVerifyEmail(ctx, arg)
// 	if err != nil {
// 		return ""
// 	}
// 	return email
// }
