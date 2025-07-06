package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/sonu31/student-api/internal/storage"
	"github.com/sonu31/student-api/internal/types"
	"github.com/sonu31/student-api/internal/utils/response"
)

func Create(storage storage.Storage) http.HandlerFunc {
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

		lastid, errr := storage.CreateStudent(student.Name, student.Email, student.Age)

		slog.Info("user Cerated successfully", slog.String("userId", fmt.Sprint(lastid)))

		if errr != nil {
			response.WriteJosn(w, http.StatusInternalServerError, err)
			return

		}

		response.WriteJosn(w, http.StatusCreated, map[string]int64{"id": lastid})

	}

}

func GetById(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		slog.Info("getting a student", slog.String("id", id))

		intId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			response.WriteJosn(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		student, e := storage.GetStudentById(intId)

		if e != nil {
			slog.Error("erro getting user", slog.String("id", id))
			response.WriteJosn(w, http.StatusInternalServerError, response.GeneralError(e))
			return
		}

		response.WriteJosn(w, http.StatusOK, student)

	}

}

func GetList(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		slog.Info("Getting all Students")

		students, err := storage.GetStudents()
		if err != nil {
			response.WriteJosn(w, http.StatusInternalServerError, err)
			return
		}
		response.WriteJosn(w, http.StatusOK, students)
	}

}
