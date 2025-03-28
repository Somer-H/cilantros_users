package application

import (
	"fmt"
	"users_api/src/users/application/entities"
	repository "users_api/src/users/application/reposoitory"
	"users_api/src/users/domain/repositories"

	"golang.org/x/crypto/bcrypt"
)

type LoginUserUseCase struct {
	db           repositories.UserRepository
	tokenManager repository.TokenManager
}

func NewLoginUserUseCase(db repositories.UserRepository, tokenManager repository.TokenManager) *LoginUserUseCase {
	return &LoginUserUseCase{
		db:           db,
		tokenManager: tokenManager,
	}
}

func (uc *LoginUserUseCase) LoginUser(userNew entities.UserToLog) (entities.UserLog, error) {
	user, err := uc.db.FindUserByUsername(userNew.Username)
	if err != nil {
		return entities.UserLog{}, fmt.Errorf("usuario o contraseña incorrectos")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userNew.Password))
	if err != nil {
		return entities.UserLog{}, fmt.Errorf("usuario o contraseña incorrectos")
	}
	token, err := uc.tokenManager.GenerateToken(user.Username, user.Role)
	if err != nil {
		return entities.UserLog{}, fmt.Errorf("error generating token")
	}
	return entities.UserLog{
		TokenLog: token,
		Username: user.Username,
		ID:       user.IdUser,
	}, nil
}
