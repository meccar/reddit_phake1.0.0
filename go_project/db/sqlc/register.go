package db

// import (
// 	"context"
// 	"log"
// 	"regexp"
// 	"strings"
// )

// type createAccountTxParams struct {
// 	createAccountParams
// 	Errors map[string]string
// }

// type createAccountTxResult struct {
// 	Account *Account
// }

// var (
// 	rxAdmin    = regexp.MustCompile(`^[a-zA-Z0-9._-]+@tafviet.com$`)
// 	rxEmail    = regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
// 	rxPassword = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*()-_+=])[a-zA-Z0-9!@#$%^&*()-_+=]{12,}$
// `)
// )

// func (msg *createAccountTxParams) Validate() bool {
// 	msg.Errors = make(map[string]string)

// 	matchEmail := rxEmail.Match([]byte(msg.Email))
// 	matchAdmin := rxAdmin.Match([]byte(msg.Email))
// 	matchPassword := rxPassword.Match([]byte(msg.Password))

// 	if matchAdmin {
// 		msg.Role = "admin"
// 	}
// 	if !matchEmail {
// 		msg.Errors["Email"] = "Please enter a valid email address"
// 	}

// 	if !matchPassword {
// 		msg.Errors["Password"] = "Please enter a valid password"
// 	}

// 	return len(msg.Errors) == 0
// }

// func (h *Handlers) createAccountTx(ctx context.Context, arg createAccountTxParams) (createAccountTxResult, error) {
// 	var result createAccountTxResult

// 	// Submit the account to the database
// 	params := createAccountParams{
// 		ID:       arg.ID,
// 		Role:     arg.Role,
// 		Email:    arg.Email,
// 		Password: arg.Password,
// 	}
// 	rows, err := h.queries.getFormsID(ctx)
// 	if err != nil {
// 		log.Fatal("Cannot ping database:", err)
// 	}
// 	for i := 0; i < len(rows); i++ {
// 		// fmt.Printf("\n loop ID %v\n", rows[i])
// 		if params.ID == rows[i] {
// 			// fmt.Printf("\n loop params.ID %v\n", params.ID)
// 			params.ID++
// 		}

// 	}
// 	Account, err := h.queries.createAccount(ctx, params)

// 	if err != nil {
// 		return result, err
// 	}

// 	result.Account = &Account
// 	return result, err
// }
