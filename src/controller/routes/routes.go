package routes

// Instalação da biblioteca de router https://github.com/gin-gonic/gin
// Gin é otimizado para alta performance, utilizando o roteador httprouter para lidar com as requisições HTTP de forma eficiente.
import (
	"crud_application/src/controller"

	"github.com/gin-gonic/gin"
)

// Função de inicialização de rotas
// A ideia é inicializar o router no main.go
// Passar o grupo de Router no initRoutes para a gente atralar as rotas da aplicação
// Pra não ter que inicializar a aplicação nesse arquivo
func InitRoute(r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
	//:var - recebimento de parametro
	r.GET("/getUserById/:userId", userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser/", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)
}
