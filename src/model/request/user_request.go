//Pasta destinada para criar objetos de entrada e saída da aplicação
//Nesse arquivo serão criados os objetos de request
//Para receber infos do usuário

package request

type UserRequest struct {
	Email    string
	Password string
	Name     string
	Age      int8
}
