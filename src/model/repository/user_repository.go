package repository

import (
	"crud_application/src/configuration/rest_err"
	"crud_application/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

// Criação do usuário no banco
func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_err.RestErr)
}
