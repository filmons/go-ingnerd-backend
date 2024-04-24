package dto

// TodoDTO is used to transfer todo data.
type TodoDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
