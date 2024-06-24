package routes

import (
	"github.com/gorilla/mux"
	"go-clean-architecture/pkg/config"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	userController := config.InitUserInjector()

	router.HandleFunc("/", userController.FindById).Methods("GET")
	// Add other routes as needed

	return router
}
