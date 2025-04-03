package dependencies

import (
	"log"
	"os"
	"users_api/src/users/application/use_cases"
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
	bs := adapters.NewBcrypt()
	register_user_use_case := application.NewRegisterUserUseCase(ps, bs)
	register_user_controller := controllers.NewRegisterUserController(register_user_use_case)
	login_user_use_case := application.NewLoginUserUseCase(ps, tm, bs)
	login_user_controller := controllers.NewLoginUserController(login_user_use_case)
	update_use_case := application.NewUpdateUserUseCase(ps)
	update_controller := controllers.NewUpdateUserController(update_use_case)
	routes.UserRouter(r, key, register_user_controller, login_user_controller, update_controller)
}