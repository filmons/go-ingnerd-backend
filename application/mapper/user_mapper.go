
package mapper

import (
    "github.com/filmons/go-ingnerd-backend/application/dto"
    "github.com/filmons/go-ingnerd-backend/domain/user"
)

// ToUserDTO maps a user domain model to a user DTO.
func ToUserDTO(u user.User) dto.UserDTO {
    return dto.UserDTO{
        ID:   u.ID,
        Name: u.Name,
    }
}
