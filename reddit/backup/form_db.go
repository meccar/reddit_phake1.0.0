package db

import (
	"context"
	"fmt"
	"regexp"
	"strings"
)

type SubmitFormTxParams struct {
	submitFormParams
	Errors      map[string]string
	AfterCreate func(form Form) error
}

type SubmitFormTxResult struct {
	Form Form
}

var rxEmail = regexp.MustCompile(`^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`)
var rxPhone = regexp.MustCompile(`^\d{10}$|^\d{11}$|^(\+\d{1,3})?\d{10,11}$`)

func (msg *SubmitFormTxParams) Validate() bool {
	msg.Errors = make(map[string]string)

	matchEmail := rxEmail.Match([]byte(msg.Email))

	matchPhone := rxPhone.Match([]byte(msg.Phone))

	if strings.TrimSpace(msg.ViewerName) == "" {
		msg.Errors["ViewerName"] = "Please enter a name"
	}
	if matchEmail == false {
		msg.Errors["Email"] = "Please enter a valid email address"
	}

	if matchPhone == false {
		msg.Errors["Phone"] = "Please enter a valid phone number"
	}

	return len(msg.Errors) == 0
}

// store *SQLStore
// msg *validateForm,
func (repo *SQLRepo) SubmitFormTx(ctx context.Context, arg SubmitFormTxParams) (SubmitFormTxResult, error) {
	var result SubmitFormTxResult

	fmt.Printf("\n arg SubmitFormTx %v\n", arg.submitFormParams)
	fmt.Printf("\n repo SubmitFormTx %v\n", repo)

	// Submit the form to the database
	err := repo.execTx(ctx, func(q *Queries) error {
		fmt.Printf("\n 2 %v\n", arg.submitFormParams)

		var err error
		submittedForm, err := q.submitForm(ctx, arg.submitFormParams)
		fmt.Printf("3 %v", result.Form)

		if err != nil {
			return err
		}

		result.Form = submittedForm

		// Call AfterCreate function if provided
		if arg.AfterCreate != nil {
			return arg.AfterCreate(result.Form)
		}

		return nil
	})
	return result, err
}
