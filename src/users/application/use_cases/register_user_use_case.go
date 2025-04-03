package application

import (
	"fmt"
	"users_api/src/users/domain/entities"
	"users_api/src/users/domain/repositories"
	"users_api/src/users/domain/services"
)

type RegisterUserUseCase struct {
	UserRepository repositories.IUserRepository
	BcryptService services.IBcrypService
}

func NewRegisterUserUseCase(userRepository repositories.IUserRepository, bcryptService services.IBcrypService) *RegisterUserUseCase {
    return &RegisterUserUseCase{UserRepository: userRepository, BcryptService:bcryptService}
}

func (u *RegisterUserUseCase) Execute(user entities.User) (*entities.User, error) {
	if(user.Gmail == "" || user.Password == "" || user.Role == ""){
     return nil, fmt.Errorf("todos los campos son obligatorios")
	}
	hashedPassword, err := u.BcryptService.HashPassword(user.Password)
	if err != nil {
        return nil, fmt.Errorf("error al encriptar la contraseña: %v", err)
    }
	user.Password = hashedPassword
    userCreated, err := u.UserRepository.RegisterUser(user)
	if err != nil {
		return nil, err
	}
	return userCreated, nil
}