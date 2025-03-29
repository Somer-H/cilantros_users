package routes

import (
	middlewares "users_api/src/users/application/middleWares"
	"users_api/src/users/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, key string, register_user_controller *controllers.RegisterUserController, login_user_controller *controllers.LoginUserController, updateUserController *controllers.UpdateUserController) {
	v1 := r.Group("/v1/users")
	v1.POST("/login", login_user_controller.Login)
	protectedRoutesSuperUser := v1.Group("/superuser")
	protectedRoutesSuperUser.Use(middlewares.RoleMiddleware(key, []string{"superuser"}))
	protectedRoutesSuperUser.POST("/register", register_user_controller.Register)
    protectedRoutesAllUsers := v1.Group("/allUsers")
	protectedRoutesAllUsers.Use(middlewares.RoleMiddleware(key, []string{"normaluser", "superuser", "premiumuser"}))
	protectedRoutesAllUsers.PUT("/update/:id", updateUserController.UpdateUser)
}