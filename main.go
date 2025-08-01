package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Biblioteca godot env -> para visualizar as váriavies do arquivo .env
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("PORT"))
}
