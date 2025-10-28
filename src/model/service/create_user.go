package service

import (
	"crud_application/src/configuration/logger"
	"crud_application/src/configuration/rest_err"
	"crud_application/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init createUser model",
		zap.String("journey", "createUser"))

	// encripta a senha
	userDomain.EncryptPassword()

	// Chama o repositori para criar o usuário
	userDomainRepository, err := ud.UserRepository.CreateUser(userDomain)

	if err != nil {
		logger.Info("Error trying to call repository", zap.String("journey", "createUser"))
		return userDomainRepository, err
	}
	logger.Info(
		"Create service executed sucessfully",
		zap.String("userId", userDomain.GetId()),
		zap.String("journey", "createUser"))
	return userDomainRepository, nil
}
