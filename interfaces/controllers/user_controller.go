package controllers

import (
    "net/http"
    "github.com/filmons/go-ingnerd-backend/application/services"
)

// UserController handles HTTP requests related to Users.
type UserController struct {
    UserService *services.UserService
}

// NewUserController creates a new controller for user endpoints.
func NewUserController(us *services.UserService) *UserController {
    return &UserController{UserService: us}
}

// GetUser handles GET requests for user data.
func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
    // Implementation for handling a user retrieval
}
