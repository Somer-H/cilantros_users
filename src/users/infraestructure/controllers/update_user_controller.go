package controllers

import (
	"strconv"
	application "users_api/src/users/application/use_cases"
	"users_api/src/users/domain/entities"

	"github.com/gin-gonic/gin"
)

type UpdateUserController struct {
	useCase *application.UpdateUserUseCase
}

func NewUpdateUserController(useCase *application.UpdateUserUseCase) *UpdateUserController {
    return &UpdateUserController{useCase: useCase}
}

func (uuc *UpdateUserController) UpdateUser(c *gin.Context){
	var user entities.UserToUpdate
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	idUser := c.Params.ByName("idUsers")
	if idUser == "" {
		c.JSON(400, gin.H{"error": "idUser no puede estar vacío"})
		return
	}
	id, err := strconv.Atoi(idUser)
	if err != nil {
        c.JSON(400, gin.H{"error": "idUser debe ser un número entero"})
        return
    }
    userToUpdate, err := uuc.useCase.Execute(id, user);
	if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, gin.H{"puts": 
               gin.H{
				"idUser": userToUpdate.IdUser,
                "username": userToUpdate.Username,
                "email": userToUpdate.Gmail,
                "role": userToUpdate.Role,
			   },
})
}