// Arquivo destinado para inicialização de dependencias
// Para retirnar do main.go
package main

import (
	"crud_application/src/controller"
	"crud_application/src/model/repository"
	"crud_application/src/model/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
	database *mongo.Database,
) controller.UserControllerInterface {
	// Requisição: controller -> service -> domain -> entity

	// camada de dados
	repo := repository.NewUserRepository(database)
	// Criação do service
	service := service.NewUserDomainService(repo)

	userController := controller.NewUserControllerInterface(service)

	return userController

}
