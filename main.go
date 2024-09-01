package main

import (
	"gotodo-app/controller"
	"gotodo-app/db"
	"gotodo-app/repository"
	"gotodo-app/router"
	"gotodo-app/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
