package service

import (
	"crud_application/src/configuration/logger"
	"crud_application/src/configuration/rest_err"
	"crud_application/src/model"
	"fmt"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))
	fmt.Println(ud)
	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())
	return nil
}
