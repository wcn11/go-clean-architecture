package controllers

import (
	"encoding/json"
	"fmt"
	"full/pkg/dto"
	"full/pkg/usecase/user"
	"log"
	"net/http"
)

type IndexController struct {
	userService user.IUserService
}

func NewIndexController(userService user.IUserService) *IndexController {
	return &IndexController{userService}
}

func (c *IndexController) Index(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	ctx := r.Context()
	userId, err := c.userService.GetUserById(ctx, 1)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		// Assuming user.Name is a string
		fmt.Fprintf(w, "user Name: %s", userId) // Format response with user name
	}
}

func (c *IndexController) Create(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	var createUserRequest dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	userResponse, err := c.userService.CreateUser(ctx, &createUserRequest)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(userResponse)
		if err != nil {
			http.Error(w, "Invalid Return Data", http.StatusInternalServerError)
			return
		}
	}
}
