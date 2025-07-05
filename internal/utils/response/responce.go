package response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ResponceD struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	StatusOK    = "Ok"
	StatusError = "Error"
)

func WriteJosn(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "applicaiton/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}

func GeneralError(err error) ResponceD {
	return ResponceD{
		Status: StatusError,
		Error:  err.Error(),
	}
}

func ValidationError(errs validator.ValidationErrors) ResponceD {
	var errMsgs []string
	for _, err := range errs {
		switch err.ActualTag() {
		case "reuired":
			errMsgs = append(errMsgs, fmt.Sprintf("field %s is required field", err.Field()))
		default:
			errMsgs = append(errMsgs, fmt.Sprint("field %s is invalid field", err.Field()))
		}

	}

	return ResponceD{
		Status: StatusError,
		Error:  strings.Join(errMsgs, ","),
	}

}
