package routes

import (
	"github.com/gorilla/mux"
	"go-clean-architecture/pkg/config"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	indexController := config.InitInjector()

	router.HandleFunc("/", indexController.Index).Methods("GET")
	// Add other routes as needed

	return router
}
