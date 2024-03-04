package db

import (
	"context"
	"regexp"
	"strings"

	util "util"
)

// "fmt"
// "log"
type SubmitFormTxParams struct {
	submitFormParams
	Errors map[string]string
}

type SubmitFormTxResult struct {
	Form *Form
}

var rxEmail = regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
var rxPhone = regexp.MustCompile(`^\d{10}$|^\d{11}$|^(\+\d{1,3})?\d{10,11}$`)

// func generateRandomID() (int32, error) {
// 	// Generate a random big.Int
// 	randomBig, err := rand.Int(rand.Reader, new(big.Int).SetInt64(1<<31-1))
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Convert the random big.Int to int32
// 	randomID := int32(randomBig.Int64())

// 	return randomID, nil
// }

func (msg *SubmitFormTxParams) ValidateForm() bool {
	msg.Errors = make(map[string]string)

	matchEmail := rxEmail.Match([]byte(msg.Email))

	matchPhone := rxPhone.Match([]byte(msg.Phone))

	if strings.TrimSpace(msg.ViewerName) == "" {
		msg.Errors["ViewerName"] = "Please enter a name"
	}
	if !matchEmail {
		msg.Errors["Email"] = "Please enter a valid email address"
	}

	if !matchPhone {
		msg.Errors["Phone"] = "Please enter a valid phone number"
	}

	return len(msg.Errors) == 0
}

func (h *Handlers) SubmitFormTx(ctx context.Context, arg SubmitFormTxParams) (SubmitFormTxResult, error) {
	var (
		result   SubmitFormTxResult
		randomID string
	)
	// fmt.Printf("\n SubmitFormTx submitFormParams %+v\n", arg.submitFormParams)
	// fmt.Printf("\n SubmitFormTx ctx %+v\n", ctx)
	randomID = util.GenerateID()

	// for {
	// 	randomID = util.GenerateID()

	// 	rows, _ := h.queries.getFormsID(ctx, randomID)
	// 	if rows == 0 {
	// 		break
	// 	}
	// }

	// // Submit the form to the database
	params := submitFormParams{
		ID:         randomID,
		ViewerName: arg.ViewerName,
		Email:      arg.Email,
		Phone:      arg.Phone,
	}

	// fmt.Printf("\n loop rows %v\n", rows)
	// fmt.Printf("\n loop len(rows) %v\n", len(rows))

	// fmt.Printf("\n SubmitFormTx params.ID 1 %v\n", params.ID)

	Form, err := h.queries.submitForm(ctx, params)

	// fmt.Printf("\n SubmitFormTx form ViewerName %v\n", Form.ViewerName)
	// fmt.Printf("\n SubmitFormTx form Email %v\n", Form.Email)
	// fmt.Printf("\n SubmitFormTx form Phone %v\n", Form.Phone)
	// fmt.Printf("\n SubmitFormTx form CreatedAt %v\n", Form.CreatedAt)
	// fmt.Printf("\n SubmitFormTx err %+v\n", err)

	if err != nil {
		return result, err
	}

	result.Form = &Form
	// fmt.Printf("\n SubmitFormTx result.Form %v\n", result.Form)
	return result, err

}
