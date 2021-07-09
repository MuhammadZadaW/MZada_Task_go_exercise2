package router

import (
	"github.com/gin-gonic/gin"
	"task_go/model"
)

type UserHandler struct {
	UserUsecase model.UserUsecase
}

func NewUsersHandler(r *gin.RouterGroup, us model.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: us,
	}

	r.POST("/user", handler.AddUser)
	r.DELETE("/user/:id", handler.DeleteUser)
	r.PUT("/user/:id", handler.UpdateUser) 

	// r.GET("/user", handler.FindUsers)
}

func (a *UserHandler) AddUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)

	users, _ := a.UserUsecase.AddUser(c.Request.Context(), user)
	c.JSON(200, gin.H{
		"id": users.ID,
		"name": users.Name,
		"job": users.Job,
		"createdAt": users.CreatedAt,
	})
}

func (a *UserHandler) DeleteUser(c *gin.Context) {
	a.UserUsecase.DeleteUser(c.Request.Context(), c.Param("id"))
	c.JSON(204, gin.H{})
}

func (a *UserHandler) UpdateUser(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)

	users, _ := a.UserUsecase.UpdateUser(c.Request.Context(), c.Param("id"), user)
	c.JSON(200, gin.H{
		"name": users.Name,
		"job": users.Job,
		"updatedAt": users.UpdatedAt,
	})
}