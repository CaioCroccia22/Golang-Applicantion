package repository

import (
	"context"
	"crud_application/src/configuration/logger"
	"crud_application/src/configuration/rest_err"
	"crud_application/src/model"
	"os"
)

// Criação do registro de usuário no banco
const (
	MONGO_USER_DB = "MONGO_USER_DB"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init createUser repository")

	collection_name := os.Getenv("MONGO_USER_DB")

	// Collection do banco
	collection := ur.databaseConnection.Collection(collection_name)

	// Pegar os atributos do domain no formato JSON
	value, err := userDomain.GetJSONValue()

	if err != nil {
		return nil, rest_err.InternalServerError(err.Error())
	}

	// Pela fato do userDomainInterface que a gente recebe ser uma interface, teriamos que cria um objeto getName e afins
	// Então vamos pegar do userDomain
	// Inserção dos valores transformados em JSON na collection
	result, err := collection.InsertOne(context.Background(), value)

	if err != nil {
		return nil, rest_err.InternalServerError(err.Error())
	}

	//
	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil

}
