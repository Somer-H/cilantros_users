package routes

import (
	"users_api/src/users/infraestructure/controllers"
	"users_api/src/users/infraestructure/service"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine, key string, register_user_controller *controllers.RegisterUserController, login_user_controller *controllers.LoginUserController, updateUserController *controllers.UpdateUserController) {
	v1 := r.Group("/v1/users")
	v1.POST("/login", login_user_controller.Login)
	protectedRoutesSuperUser := v1.Group("/superuser")
	protectedRoutesSuperUser.Use(service.RoleMiddleware(key, []string{"superuser"}))
	protectedRoutesSuperUser.POST("/register", register_user_controller.Register)
    protectedRoutesAllUsers := v1.Group("/allUsers")
	protectedRoutesAllUsers.Use(service.RoleMiddleware(key, []string{"normaluser", "superuser", "premiumuser"}))
	protectedRoutesAllUsers.PUT("/update/:id", updateUserController.UpdateUser)
}