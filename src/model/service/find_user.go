package service

import (
	"crud_application/src/configuration/rest_err"
	"crud_application/src/model"
)

// func (u *userDomainService) -> Aqui é o receiver -> significa que findUser é um método do tipo UserdomainService
// FindUser(id string) -> Aqui é o que esse método recebe
//(*model.UserDomainInterface, *rest_err.RestErr) -> Aqui é o retorno desse método

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
