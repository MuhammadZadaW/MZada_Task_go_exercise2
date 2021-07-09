package main

import (
	"github.com/gin-gonic/gin"
	"task_go/config"
	"github.com/gin-gonic/contrib/cors"

	"task_go/repository"
	"task_go/usecase"
	"task_go/router"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	config.SetupModels()
	db := config.GetDBConnection()
	port := config.GetPortConnection()

	repoUser := repository.NewPsqlUserRepository(db)
	entityUser := usecase.NewUserUsecase(repoUser)

	api := r.Group("/api")

	router.NewUsersHandler(api, entityUser)

	r.Run(port)
}