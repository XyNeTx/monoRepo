package interfaces

import "fiber-api/internal/models/dto"

type UserService interface {
	GetAllUsers() ([]dto.UserDTO, error)
	GetUserByEmail(email string) (dto.UserDTO, error)
	GetUserById(id int) (dto.UserDTO, error)
	CreateUsers(userDTO dto.UserDTO) error
	EditUsers(userDTO dto.UserDTO) (dto.UserDTO, error)
	DeleteUser(email string) error
	VerifyPassword(email string, password string) (dto.UserDTO, error)
	ChangePassword(email string, newPassword string) error
}
