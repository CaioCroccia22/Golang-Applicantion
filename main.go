package main

import (
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

	routes.InitRoute(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
