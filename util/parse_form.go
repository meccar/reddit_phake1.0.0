package util

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

func ParseForm(r *http.Request, msg interface{}) error {
	err := r.ParseForm()
	if err != nil {
		fmt.Printf("\n ParseForm err %v \n", err)
		return err
	}

	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(msg, r.PostForm)
	if decodeErr != nil {
		fmt.Printf("\n ParseForm err %v \n", decodeErr)

		return decodeErr
	}

	return nil
}

func ValidateForm(msg interface{}, rules map[string]interface{}) (string, error) {
	errs := make(map[string]string)
	validate := validator.New(validator.WithRequiredStructEnabled())

	for field, rule := range rules {
		fieldValue := reflect.ValueOf(msg).Elem().FieldByName(field)


		if !fieldValue.IsValid() {
			errs[field] = fmt.Sprintf("Field %s not found in struct", field)
			continue
		}

		fieldStr := fmt.Sprintf("%v", fieldValue.Interface())

		if err := validate.Var(fieldStr, rule.(string)); err != nil {

			// Handle specific validation errors
			switch field {
			case "ViewerName":
				errs[field] = "Invalid Name"
			case "Email":
				errs[field] = "Invalid Email"
			case "Phone":
				errs[field] = "Invalid Phone"
			case "Username":
				errs[field] = "Invalid Username"
			case "Password":
				errs[field] = "Invalid Password"
			default:
				errs[field] = err.Error()
			}
		}
	}

	// Check if there are any errors in the validation
	if len(errs) > 0 {
		// Return the first error message
		for _, v := range errs {
			return v, nil
		}
	}

	return "", nil
}
