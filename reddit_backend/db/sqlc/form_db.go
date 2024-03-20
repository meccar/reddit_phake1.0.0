package db

import (
	"context"
	"errors"
	util "util"

	"github.com/google/uuid"
)

type SubmitFormTxParams struct {
	submitFormParams
}

type SubmitFormTxResult struct {
	Form *Form
}

var FormRules = map[string]interface{}{
	"ViewerName": "required,min=4,max=35",
	"Email":      "required,email",
	"Phone":      "required,min=10,max=11",
}

func (msg *SubmitFormTxParams) ValidateForm() error {
	errorMsg, err := util.ValidateForm(msg, FormRules)
	if err != nil {
		return err
	}

	if errorMsg != "" {
		return errors.New(errorMsg)
	}

	// Return a success message indicating validation success
	return nil
}

func (h *Handlers) SubmitFormTx(ctx context.Context, arg SubmitFormTxParams) (SubmitFormTxResult, error) {
	var result SubmitFormTxResult

	err := h.execTx(ctx, func(q *Queries) error {
		var err error

		ranID, _ := uuid.NewRandom()


		// // Submit the form to the database
		params := submitFormParams{
			ID:         ranID,
			ViewerName: arg.ViewerName,
			Email:      arg.Email,
			Phone:      arg.Phone,
		}

		Form, err := q.submitForm(ctx, params)

		if err != nil {
			return err
		}

		result.Form = &Form
		return err
	})
	return result, err
}
