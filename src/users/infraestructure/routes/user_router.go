package routes

import (
	middlewares "users_api/src/users/application/middleWares"
	"users_api/src/users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, key string, register_user_controller *controllers.RegisterUserController, login_user_controller *controllers.LoginUserController) {
	v1 := r.Group("/v1/users")
	v1.POST("/login", login_user_controller.Login)
	protectedRoutes := v1.Group("/superuser")
	protectedRoutes.Use(middlewares.RoleMiddleware(key, "superuser"))
	protectedRoutes.POST("/register", register_user_controller.Register)
}