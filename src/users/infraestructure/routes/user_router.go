package routes

import (
	"users_api/src/users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, register_user_controller *controllers.RegisterUserController, login_user_controller *controllers.LoginUserController){
	v1 := r.Group("/v1/users")
	{
        v1.POST("/register", register_user_controller.Register)
		v1.POST("/login", login_user_controller.Login)
    }
	
}