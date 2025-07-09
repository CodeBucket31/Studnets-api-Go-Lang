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

func UpdateItme(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("UpdateItme Index ")

		// type StudentTest struct {
		// 	ID    int
		// 	Name  string
		// 	Email string
		// }
		// var s1list []StudentTest

		// var _ []StudentTest

		// if students1 == nil {
		// 	response.WriteJosn(w, http.StatusBadGateway, "NUlll Ata")
		// }

		// s1list := StudentTest{ID: 2, Name: "RAj ", Email: "sonu@gmial.con"}

		// for i := 0; i < 40; i++ {
		// 	students1 = append(students1, s1list)
		// }

		// numbers := []int{1, 2, 3, 4, 5}

		//    sk int[]int{2,3,4,5,6}

		// 1. Get ID from URL
		// idStr := URLParam(r, "id") // Assuming chi router used
		// id, err := strconv.ParseInt(idStr, 10, 64)
		// if err != nil {
		// 	http.Error(w, "Invalid student ID", http.StatusBadRequest)
		// 	return
		// }

		// 2. Decode request body
		// var student types.Student
		// err = json.NewDecoder(r.Body).Decode(&student)
		// if err != nil {
		// 	http.Error(w, "Invalid JSON", http.StatusBadRequest)
		// 	return
		// }

		// 3. Validation
		// if strings.TrimSpace(student.Name) == "" || strings.TrimSpace(student.Email) == "" {
		// 	http.Error(w, "Name and Email are required", http.StatusBadRequest)
		// 	return
		// }

		// 4. Call DB update
		// err = storage.UpdateStudentByID(id, student)
		// if err != nil {
		// 	http.Error(w, "Failed to update student", http.StatusInternalServerError)
		// 	return
		// }

		// json.NewEncoder(w).Encode(map[string]string{
		// 	"message": "Student updated successfully",
		// })

		// response.WriteJosn(w, http.StatusOK, sk)

	}
}
