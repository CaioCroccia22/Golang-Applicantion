package controller

import (
	"crud_application/src/configuration/rest_err"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Erro catastrófico")
	c.JSON(err.Code, err)
}
