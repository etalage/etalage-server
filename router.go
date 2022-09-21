package main

import (
	C "github.com/etalage/etalage-server/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	routeBinder *gin.Engine
}

func (r *Router) BindRoute() (success bool) {
	if nil == &r {
		return false
	}
	userController := C.UserController{}.New()
	r.routeBinder.GET("/ping", userController.AddUser)

	return true
}
