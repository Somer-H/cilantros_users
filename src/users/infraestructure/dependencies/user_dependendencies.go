package dependencies

import (
	"log"
	"os"
	"users_api/src/users/application"
	"users_api/src/users/infraestructure/adapters"
	"users_api/src/users/infraestructure/controllers"
	"users_api/src/users/infraestructure/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitUsers(r *gin.Engine) {
    ps, _ := adapters.NewMySQL();
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }
	key := os.Getenv("SECRET_KEY")
	tm := adapters.NewJWTManager(key)
	register_user_use_case := application.NewRegisterUserUseCase(ps)
	register_user_controller := controllers.NewRegisterUserController(register_user_use_case)
	login_user_use_case := application.NewLoginUserUseCase(ps, tm)
	login_user_controller := controllers.NewLoginUserController(login_user_use_case)
	routes.UserRouter(r, key, register_user_controller, login_user_controller)
}