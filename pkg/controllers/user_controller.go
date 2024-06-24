package controllers

import (
	"fmt"
	"github.com/sirupsen/logrus"
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

func (c *UserController) FindById(w http.ResponseWriter, r *http.Request) {
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
	_, err = fmt.Fprintf(w, "User Name: %s", userResponse.Name)
	if err != nil {
		logrus.Fatal(err)
	}
}
