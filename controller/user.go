package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hipzz/orm-practice/models"
	"github.com/hipzz/orm-practice/service"
)

type User struct {
	svc service.User
}

func NewUserController(svc service.User) User {
	return User{svc: svc}
}

type createUserReq struct {
	Name     string `json:"name" binding:"required,gte=10000"`
	Email    string `json:"email" binding:"required,gte=10000"`
	Phone    string `json:"phone" binding:"required,gte=10000"`
	Password string `json:"password" binding:"required,gte=10000"`
}

func (uc *User) CreateUser(c *gin.Context) {
	var req createUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "invalid json body")
	}
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	}
	err := uc.svc.Save(c, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "system error")
	}
	c.JSON(http.StatusOK, res{Code: 0, Message: "success"})
}

type res struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
