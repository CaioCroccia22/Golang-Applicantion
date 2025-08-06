package response

// Importante lembrar de adicionar propriedades de JSON
// nos objetos para que seja possivel converte o objeto para JSON e vice versa
// `json:"id"` -> GoAddTags
// BSON para o nome desse campo dentro do banco
type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}
