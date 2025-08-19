package validantion

import (
	"crud_application/src/configuration/rest_err"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

// Criar as váriaveis de translate e validação
// Criando a variavel Validate (global) para utilizar ela no controller
var (
	// Inicializar o validator
	Validate = validator.New()
	transl   ut.Translator
)

// Inicialização do pacote passando qual a linguagem que quer transformar
// Função initi vai sempre rodar quando carregar o pacote
// Configura o validator do Gin Gonic para usar traduções de mensagens de erro em inglês
func init() {
	//Relização do binding para essa struct do validator
	// Binding é do gin gonic
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()        //carrega o locale em inglês = instância da lingua
		unt := ut.New(en, en) //cria a instância da tradução
		transl, _ := unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl) //registra as traduções
	}
}

// Função para a validação
// Essa função vai receber um erro da validação do gin gonic
// e será convertido para o objeto de erro que temos
func ValidateUserError(
	validation_err error,
) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors
	// Validar os tipos dos erros
	//1 - validar se ele mandou o tipo errado = Unmarshall
	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
		//validar se é um erro de validação
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{} //Array de causas

		// aqui é usado o bindig para fazer um casting e converter esse tipo error
		for _, e := range validation_err.(validator.ValidationErrors) {
			// Objeto cause que será populado com o validation_err que iremos receber
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestErrorValidationError("Some fields are invalid", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
