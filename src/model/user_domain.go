package model

import (
	"crypto/md5"
	"encoding/hex"
)

// De modo que userDomain é privado os métodos da apsta service não enxergam eles
// Para isso criamos essa interface que traz os métodos referentes a nossos objetos
// O userDomain tem métodos e esses métodos são disponibilizados para a aplicação por essa interface
// type UserDomainInterface interface {
// 	GetEmail() string
// 	GetName() string
// 	GetAge() int8
// 	GetPassword() string

// 	SetID(string)

// 	GetJSONValue() (string, error)

// 	EncryptPassword()

// 	// EncryptPassword() []byte
// }

/*
Nosso construtor
*/

// func NewUserDomain(
// 	email, password, name string, age int8,
// ) UserDomainInterface {
// 	return &userDomain{
// 		Email:    email,
// 		Password: password,
// 		Name:     name,
// 		Age:      age,
// 	}
// }

func (ud *userDomain) SetID(id string) {
	ud.id = id
}

/*
UserDomain
Tem informações para seguir a regra de negócio
é a conversão do request para domain, ou vice versa
*/

/*Para deixar privado o acesso ao UserDomain e atributos será deixado minusculo */
// UserDomain mesmo sendo privado podemos utilizar os seus atributos se deixarmos em maiusculo
// Ex: Email
type userDomain struct {
	id       string
	email    string `JSON:"email"`
	password string `JSON:"password"`
	name     string `JSON:"name"`
	age      int8   `JSON:"age"`
}

func (ud *userDomain) GetID() string {
	return ud.id
}

// GetName implements UserDomainInterface.
func (ud *userDomain) GetName() string {
	return ud.name
}

func (ud *userDomain) GetEmail() string {
	return ud.email
}

func (ud *userDomain) GetAge() int8 {
	return ud.age
}

func (ud *userDomain) EncryptPassword() {
	// Cria a instância
	hash := md5.New()
	// Reset
	defer hash.Reset()
	// escre convertendo de string do objeto de response para um array de byte
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

func (ud *userDomain) GetPassword() string {
	return ud.password
}
