// Arquivo de padronização de erros -> Para todos os erros da aplicação estarem em um lugar
// Evitar usar o nome err para não dar conflito com o pacote nativo do go
package rest_err

import "net/http"

// Criando um objeto do tipo struct
type RestErr struct {
	// Campo mensagem vai ter a mensagem do erro
	//json -> Deixar o nosso objeto exportável para JSON
	Message string `json:"message"`
	// Campo da stack trace do que aconteceu
	Err string `json:"error"`
	// Campo code vai guardar qual o código da requisição que estamos dando para o cliente
	Code int `json:"code"`
	// Campo Cause do tipo Causes vai guardar quais as causas de erro da aplicação
	Causes []Causes `json:"causes"`
}

//Vai ser um array de um objeto
type Causes struct {
	// Dentro desse objeto vão ter os campos Field e Message
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Criando um constructor para o objeto
// *RestErr -> Significa retorno
func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

// Função implementada apatir de RestErr-> Para satisfazer a interface de error do go
func (r *RestErr) Error() string {
	return r.Message
}

// Função de badRequest
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestErrorValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func InternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Not found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}
