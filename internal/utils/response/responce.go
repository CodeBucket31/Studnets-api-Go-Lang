package response

import (
	"encoding/json"
	"net/http"
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
