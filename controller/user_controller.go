package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	AddUser gin.HandlerFunc
}

func (u UserController) New() (controller *UserController) {
	controller = new(UserController)
	controller.AddUser = addUser
	return controller
}

func addUser(context *gin.Context) {
	time.Now()
	context.JSON(http.StatusOK, gin.H{
		"message": time.Now(),
	})
}
