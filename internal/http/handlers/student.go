package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sonu31/student-api/internal/types"
	"github.com/sonu31/student-api/internal/utils/response"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJosn(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err != nil {
			response.WriteJosn(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		//required Valideation
		if err := validator.New().Struct(student); err != nil {

			validaterErrs := err.(validator.ValidationErrors)
			response.WriteJosn(w, http.StatusBadRequest, response.ValidationError(validaterErrs))
			return
		}

		slog.Info("careteing a studnet")
		response.WriteJosn(w, http.StatusCreated, map[string]string{"sucess": "Ok"})

	}

}
