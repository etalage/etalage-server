package controller

import (
	"log"
	"net/http"

	"github.com/etalage/etalage-server/model"
	"github.com/etalage/etalage-server/repo"
	"github.com/etalage/etalage-server/repo/repointerface"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserManager repointerface.UserManager
}

func (u UserController) New() (controller *UserController) {
	controller = new(UserController)
	controller.inject()
	return controller
}

func (u *UserController) inject() {
	u.UserManager = repo.UserRepo{}
}

func (controller UserController) AddUser(context *gin.Context) {
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
		})
	}
	log.Println("user:", user.UserName, " password:", user.Password)
	if controller.UserManager.AddUser(user) > 0 {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello:" + user.UserName,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "sorry," + user.UserName,
	})
}
