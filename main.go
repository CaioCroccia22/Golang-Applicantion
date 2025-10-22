package main

import (
	"context"
	"crud_application/src/configuration/database/mongodb"
	"crud_application/src/configuration/logger"
	"crud_application/src/controller/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Biblioteca godot env -> para visualizar as váriavies do arquivo .env
func main() {
	// Inicialização do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Chamando o pacote de conexão com o banco

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error, trying to connect to database, erro=%s \n")
		err.Error()
		return
	}

	// Inicialização de dependências -> Passado para o arquivo init_dependencies
	// Com a criação do arquivo user_controller.go será necessário criar a instância para uma estrutura unica do service
	// essa estrutura sera chamada no controller
	// service := service.NewuserDomainService()
	// userController := controller.NewUserControllerInterface(service)
	userController := initDependencies(database)

	// Inicialização de rotas
	router := gin.Default() //gin.Default inicializa o roteador com logger e middlewares de recovery
	/*
		Os roteadores podem ser inicializados de duas maneiras diferentes:
		 - gin.New() -> criará uma novo roteador gin.
		 - gin.Default() -> toda fez que sua aplicação estiver dando um panic em algum lugar, ele vai dar um default
		A diferença é que gin.New() inicializa um roteador sem nenhum middleware enquanto o
	*/

	// Criação de logs
	logger.Logger.Info("About to start user application") //>> {"level":"info","time":"2025-08-24T14:36:48.321-0300","message":"About to start user application"}

	routes.InitRoute(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
