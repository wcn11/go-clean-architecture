package controllers

import (
	"fmt"
	"go-clean-architecture/pkg/usecase/user"
	"log"
	"net/http"
)

type UserController struct {
	UserService user.IUserService
}

func NewUserController(userService user.IUserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (c *UserController) Index(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("Panic: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}()

	ctx := r.Context()
	userResponse, err := c.UserService.GetUserById(ctx, 1)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Assuming userResponse.Name is a string
	fmt.Fprintf(w, "User Name: %s", userResponse.Name)
}
