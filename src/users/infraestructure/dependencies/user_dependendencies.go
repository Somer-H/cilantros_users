package dependencies

import (
	"users_api/src/users/application"
	"users_api/src/users/infraestructure"
	"users_api/src/users/infraestructure/controllers"
	"users_api/src/users/infraestructure/routes"

	"github.com/gin-gonic/gin"
)

func InitUsers(r *gin.Engine) {
    ps, _ := infraestructure.NewMySQL();
	register_user_use_case := application.NewRegisterUserUseCase(ps)
	register_user_controller := controllers.NewRegisterUserController(register_user_use_case)
	routes.UserRouter(r, register_user_controller)
}