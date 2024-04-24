package repository

import (
    "github.com/filmons/go-ingnerd-backend/domain/user"
)

// UserRepository handles the operations with the user entity.
type UserRepository struct {
    // You might have other properties here
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository() *UserRepository {
    return &UserRepository{}
}

// Example function that uses the User struct
func (repo *UserRepository) FindAllUsers() []user.User {
    // Example implementation; actual implementation may vary
    return []user.User{}
}

