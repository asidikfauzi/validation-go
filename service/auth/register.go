package auth

import (
	"fmt"
	"log"
	"test-prepare/app"
	modelDomain "test-prepare/model/domain"
	modelAuth "test-prepare/model/web/auth"
)

func RegisterService(reqRegister modelAuth.RequestRegister) (interface{}, error) {
	var domainUser modelDomain.Users

	domainUser.Email = reqRegister.Email
	domainUser.Username = reqRegister.Username
	domainUser.Password = reqRegister.Password

	err := app.DB.Save(&domainUser).Error
	fmt.Println(err)
	if err != nil {
		log.Fatal("Register gagal di dbnya")
		return nil, err
	}

	return reqRegister, nil

}
