package model

import (
	"crud_application/src/configuration/rest_err"
	"crypto/md5"
	"encoding/hex"
)

/*
Nosso construtor
*/

func NewUserDomain(
	email, password, name string, age int8,
) UserDomainInterface {
	return &UserDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

/*
UserDomain
Tem informações para seguir a regra de negócio
é a conversão do request para domain, ou vice versa
*/

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

// Create implements UserDomainInterface.
func (ud *UserDomain) Create() *rest_err.RestErr {
	panic("unimplemented")
}

// A intenção dessa interface é fazer com que o controller altere nela e não direto na userDomain
// Os métodos nesse primeiro momento pertencem a interface -> isso faz com uqe quando a gente chame a interface só chame as funções
// e nã os valores
type UserDomainInterface interface {
	// Crialção dos métodos que o domain vai executar
	CreateUser() *rest_err.RestErr
	UpdateUser(string) *rest_err.RestErr
	FindUser(string) *rest_err.RestErr
	DeleteUser(string) *rest_err.RestErr
}

func (ud *UserDomain) EncryptPassword() {
	// Cria a instância
	hash := md5.New()
	// Reset
	defer hash.Reset()
	// escre convertendo de string do objeto de response para um array de byte
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}
