package controller

import (
	"crud_application/src/configuration/rest_err"
	"crud_application/src/controller/model/request"
	"crud_application/src/controller/model/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

// c é o contexto da requisição é tudo que vem da requisição do objeto
func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest
	var userResponse response.UserResponse
	/*
	 ShouldBindJSON -> Esse método pega o body da requisição e joga dentro do objeto que colocar
	 O & é para jogar dentro do ponteiro, do endereço de mémoria userRequest

	 No trecho a seguir:
	 Aqui criamos a variavel error, para vermos se ShouldBindoJSON executou e
	 se ele for diferente de nulo vai ser criado o objeto RestErr
	*/
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			// Sprintf -> Para concatenar as strings
			fmt.Sprintf("There are incorrect fields, error=%s\n", err.Error()))

		//c.JSON -> Fala para retornar um JSON para a pessoal que fez a requisição
		c.JSON(restErr.Code, restErr)
	}

	c.JSON(200, userResponse)
	fmt.Println(userRequest)
}
