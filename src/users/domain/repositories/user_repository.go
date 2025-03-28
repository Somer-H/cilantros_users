package repositories

import "users_api/src/users/domain/entities"

type UserRepository interface {
    RegisterUser(user entities.User) (*entities.User, error)
    FindUserByUsername(username string) (*entities.User, error)
    UpdateUser(idUser int, userToUpdate entities.UserToUpdate) (*entities.User, error)
}
