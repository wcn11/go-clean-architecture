package routes

import (
	"full/pkg/controllers"
	middleware2 "full/pkg/middleware"
	"full/pkg/usecase/user"
	"github.com/gorilla/mux"
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
