package model

import (
	"crud_application/src/configuration/logger"
	"crud_application/src/configuration/rest_err"
	"fmt"

	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	fmt.Println(ud)
	ud.EncryptPassword()
	fmt.Println(ud)
	return nil
}
