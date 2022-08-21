package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hipzz/orm-practice/controller"
)

func NewRouter(userController controller.User) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	v1 := r.Group("v1")
	{
		v1.POST("/users", userController.CreateUser)
	}
	return r
}
