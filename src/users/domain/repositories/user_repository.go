package repositories

import "users_api/src/users/domain/entities"

type UserRepository interface {
    RegisterUser(user entities.User) (*entities.User, error)
   //LoginUser(username string, password string) (string, int, error)
}
