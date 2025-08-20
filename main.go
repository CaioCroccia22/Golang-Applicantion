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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	/*
		Os roteadores podem ser inicializados de duas maneiras diferentes:
		 - gin.New() -> criará uma novo roteador gin.
		 - gin.Default() -> toda fez que sua aplicação estiver dando um panic em algum lugar, ele vai dar um default
		A diferença é que gin.New() inicializa um roteador sem nenhum middleware enquanto o
		gin.Default inicializa o roteador com logger e middlewares de recovery
	*/
	router := gin.Default()
	logger.Logger.Info("Iniciando aplicação")

	routes.InitRoute(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
