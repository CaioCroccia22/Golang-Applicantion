package controller

import (
	"crud_application/src/configuration/logger"
	"crud_application/src/configuration/validantion"
	"crud_application/src/controller/model/request"
	"crud_application/src/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	// Comunicação do controller com o model
	UserDomainInterface model.UserDomainInterface
)

// c é o contexto da requisição é tudo que vem da requisição do objeto
func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest
	// var userResponse response.UserResponse
	/*
	 ShouldBindJSON -> Esse método pega o body da requisição e joga dentro do objeto que colocar
	 O & é para jogar dentro do ponteiro, do endereço de mémoria userRequest

	 No trecho a seguir:
	 Aqui criamos a variavel error, para vermos se ShouldBindoJSON executou e
	 se ele for diferente de nulo vai ser criado o objeto RestErr
	*/
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		// Sprintf -> Para concatenar as strings
		// fmt.Sprintf("There are incorrect fields, error=%s\n", err.Error())
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"),
		)
		errRest := validantion.ValidateUserError(err)

		//c.JSON -> Fala para retornar um JSON para a pessoa que fez a requisição
		c.JSON(errRest.Code, errRest)
		return
	}

	// Inicialização do nosso domain(constructor) do model
	domain := model.NewUserDomain(userRequest.Email, userRequest.Name, userRequest.Password, userRequest.Age)
	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	// response := response.UserResponse{
	// 	ID:    "test",
	// 	Email: userRequest.Email,
	// 	Name:  userRequest.Name,
	// 	Age:   userRequest.Age,
	// }

	// c.JSON(200, response)

	c.String(http.StatusOK, "Tudo certo")

	logger.Info("User create sucessfully", zap.String("journey", "createUser"))
}
