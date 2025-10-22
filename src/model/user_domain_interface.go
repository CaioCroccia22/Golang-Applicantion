// Criada só para fins de organização
// Pois o userdomaininterface e user domain estavam no mesmo arquivo
// De modo que userDomain é privado os métodos da apsta service não enxergam eles
// Para isso criamos essa interface que traz os métodos referentes a nossos objetos
// O userDomain tem métodos e esses métodos são disponibilizados para a aplicação por essa interface
package model

type UserDomainInterface interface {
	GetEmail() string
	GetName() string
	GetAge() int8
	GetPassword() string

	SetID(string)

	GetJSONValue() (string, error)

	EncryptPassword()

	// EncryptPassword() []byte
}

func NewUserDomain(
	email, password, name string, age int8,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}
