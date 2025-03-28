package application

import (
	"users_api/src/users/domain/entities"
	"users_api/src/users/domain/repositories"
)

type UpdateUserUseCase struct {
	UserRepository repositories.UserRepository
}

func NewUpdateUser(userRepository repositories.UserRepository) *UpdateUserUseCase {
    return &UpdateUserUseCase{UserRepository: userRepository}
}

func (u *UpdateUserUseCase) Execute(idUser int ,userToUpdate entities.UserToUpdate) (*entities.User, error) {
    userUpdated, err := u.UserRepository.UpdateUser(idUser, userToUpdate)
    if err != nil {
		return &entities.User{}, err
	}
	return userUpdated, nil
}