package view

import (
	"crud_application/src/controller/model/response"
	"crud_application/src/model"
)

// Nosso domain é privado - nesse caso entra a camada de view
// ela tem que saber exportar e converter os objetos
// no caso esta importando o domain e transformando ele em objeto de response
func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    userDomain.GetId(),
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
