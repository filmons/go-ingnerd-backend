package services

import (
"github.com/filmons/go-ingnerd-backend/domain/user"
) 

// UserService handles business logic related to users.
type UserService struct {
	repository user.Repository
    // Dependency injection of repositories if needed
}

// NewUserService creates a new instance of UserService.
func NewUserService() *UserService {
    return &UserService{}
}
