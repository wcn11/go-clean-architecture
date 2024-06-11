package routes

import (
	"github.com/gorilla/mux"
	"go-clean-architecture/pkg/controllers"
	middleware2 "go-clean-architecture/pkg/middleware"
	"go-clean-architecture/pkg/usecase/user"
)

func InitRouter(userService user.IUserService) *mux.Router {
	router := mux.NewRouter()

	indexController := controllers.NewIndexController(userService)

	router.Use(middleware2.Logging)
	router.Use(middleware2.Recover)

	router.HandleFunc("/", indexController.Index).Methods("GET")
	router.HandleFunc("/user/create", indexController.Create).Methods("POST")

	return router
}
