package auth

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	modelAuth "test-prepare/model/web/auth"
	modelResponse "test-prepare/model/web/response"
	"test-prepare/repository/utils"
	serviceAuth "test-prepare/service/auth"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var (
		requestRegister modelAuth.RequestRegister
		err             error
	)

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&requestRegister)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, err.Error(), err)
		return
	}

	validate := validator.New()
	err = validate.Struct(requestRegister)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := make([]interface{}, len(validationErrors))
		for i, fieldError := range validationErrors {
			data := modelResponse.ErrorValidate{
				Field:   fieldError.Field(),
				Message: utils.GetErrorMessage(fieldError),
			}
			errorMessage[i] = data
		}
		utils.Response(w, http.StatusUnprocessableEntity, "Unprocessable Entity", errorMessage)
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(requestRegister.Password), bcrypt.DefaultCost)
	requestRegister.Password = string(hashPassword)

	response, _ := serviceAuth.RegisterService(requestRegister)

	utils.Response(w, http.StatusCreated, "Register Successfully", response)
	return

}
