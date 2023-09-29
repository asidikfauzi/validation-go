package auth

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	modelAuth "test-prepare/model/web/auth"
	modelResponse "test-prepare/model/web/response"
	"test-prepare/repository/utils"
	serviceAuth "test-prepare/service/auth"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var (
		requestLogin modelAuth.LoginRequest
		err          error
	)

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestLogin)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, err.Error(), err)
		return
	}
	defer r.Body.Close()

	validate := validator.New()
	err = validate.Struct(requestLogin)
	if err != nil {
		validationError := err.(validator.ValidationErrors)
		errorMessage := make([]interface{}, len(validationError))
		for i, fieldError := range validationError {
			data := modelResponse.ErrorValidate{
				Field:   fieldError.Field(),
				Message: utils.GetErrorMessage(fieldError),
			}
			errorMessage[i] = data
		}

		utils.Response(w, http.StatusUnprocessableEntity, "Unprocessable Entity", errorMessage)
		return
	}

	response, err := serviceAuth.LoginService(w, requestLogin)
	if err != nil {
		return
	}

	utils.Response(w, http.StatusOK, response, "")
	return
}
