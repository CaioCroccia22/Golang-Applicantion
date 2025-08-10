//Pasta destinada para criar objetos de entrada e saída da aplicação
//Nesse arquivo serão criados os objetos de request
//Para receber infos do usuário

package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,containsany='!@#$%&*'"`
	Name     string `json:"name" binding:"required,min=4,max=100"`
	Age      int8   `json:"age" binding:"required,numeric,min=1,max=140"`
}
