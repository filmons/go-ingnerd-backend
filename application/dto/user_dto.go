package dto

// UserDTO is a data transfer object representing a user.
type UserDTO struct {
    ID   uint   `json:"id"`
    Name string `json:"name"`
}