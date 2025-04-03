package application

import (
	"fmt"
	"users_api/src/users/application/entities"
	repository "users_api/src/users/application/reposoitory"
	"users_api/src/users/domain/repositories"
	"users_api/src/users/domain/services"
)

type LoginUserUseCase struct {
	db           repositories.IUserRepository
	tokenManager repository.TokenManager
	bs services.IBcrypService
}

func NewLoginUserUseCase(db repositories.IUserRepository, tokenManager repository.TokenManager, bs services.IBcrypService) *LoginUserUseCase {
	return &LoginUserUseCase{
		db:           db,
		tokenManager: tokenManager,
		bs: bs,
	}
}

func (uc *LoginUserUseCase) LoginUser(userNew entities.UserToLog) (entities.UserLog, error) {
	if(userNew.Password == "" || userNew.Username == "") {
		return entities.UserLog{}, fmt.Errorf("los campos de contraseña y nombre de usuario son obligatorios")
	}
	user, err := uc.db.FindUserByUsername(userNew.Username)
	if err != nil {
		return entities.UserLog{}, fmt.Errorf("usuario o contraseña incorrectos")
	}
    validate := uc.bs.ComparePasswords(user.Password, userNew.Password)
	if !validate {
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
