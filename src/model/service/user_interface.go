package service

import (
	"crud_application/src/configuration/rest_err"
	"crud_application/src/model"
)

// Função para inicializar nosso service
func NewuserDomainService() UserDomainService {
	return &userDomainService{}
}

// é necessário um objeto para poder implementar nossos métodos
// é necessário um objeto que implemente isso ao nosso service
type userDomainService struct {
	// é possível colocar a depedência para instanciar um service
}

// A intenção dessa interface é fazer com que o controller altere nela e não direto na userDomain
// Os métodos nesse primeiro momento pertencem a interface -> isso faz com uqe quando a gente chame a interface só chame as funções
// e nã os valores
type UserDomainService interface {
	// Criação dos métodos que o domain vai executar
	CreateUser(model.UserDomainInterface) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
	UpdateUser(string, model.UserDomainInterface) *rest_err.RestErr
	FindUser(string) (*model.UserDomainInterface, *rest_err.RestErr)
}
