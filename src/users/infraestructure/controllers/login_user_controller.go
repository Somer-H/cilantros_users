package controllers

import (
	"users_api/src/users/application"
	"users_api/src/users/application/entities"

	"github.com/gin-gonic/gin"
)

type LoginUserController struct {
	useCase *application.LoginUserUseCase
}

func NewLoginUserController(useCase *application.LoginUserUseCase)  *LoginUserController {
    return &LoginUserController{useCase: useCase}
}
func (luc *LoginUserController) Login(c *gin.Context) {
    var userLog entities.UserToLog
    if err := c.ShouldBindJSON(&userLog); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    login, err := luc.useCase.LoginUser(userLog)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.Header("Authorization", "Bearer "+login.TokenLog)

    c.JSON(200, gin.H{
        "message": "Login successful",
        "user": gin.H{
            "username": login.Username,
            "id":       login.ID,
        },
    })
}