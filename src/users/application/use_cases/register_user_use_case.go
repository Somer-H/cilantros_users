package application

import (
	"fmt"
	"users_api/src/users/domain/entities"
	"users_api/src/users/domain/repositories"
)

type RegisterUserUseCase struct {
	UserRepository repositories.UserRepository
}

func NewRegisterUserUseCase(userRepository repositories.UserRepository) *RegisterUserUseCase {
    return &RegisterUserUseCase{UserRepository: userRepository}
}

func (u *RegisterUserUseCase) Execute(user entities.User) (*entities.User, error) {
	if(user.Gmail == "" || user.Password == "" || user.Role == ""){
     return nil, fmt.Errorf("todos los campos son obligatorios")
	}
    userCreated, err := u.UserRepository.RegisterUser(user)
	if err != nil {
		return nil, err
	}
	return userCreated, nil
}