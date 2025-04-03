package controllers

import (
	"net/http"
	application "users_api/src/users/application/use_cases"
	"users_api/src/users/domain/entities"
	"github.com/gin-gonic/gin"
)

type RegisterUserController struct {
	useCase *application.RegisterUserUseCase
}

func NewRegisterUserController(useCase *application.RegisterUserUseCase) *RegisterUserController {
    return &RegisterUserController{useCase: useCase}
}

func (ruc *RegisterUserController) Register(c *gin.Context) {
    var user entities.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
        return
    }
	userCreated, err := ruc.useCase.Execute(user)
	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusCreated, gin.H{
        "data": 
            gin.H{
                "idUser": userCreated.IdUser,
                "username": userCreated.Username,
                "email": userCreated.Gmail,
                "role": userCreated.Role,
            },
    })
}