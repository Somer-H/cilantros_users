package main

import (
	"time"
	"users_api/src/users/infraestructure/dependencies"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:  []string{"Content-Type", "Authorization"},
		ExposeHeaders: []string{"Authorization"},
		MaxAge:        12 * time.Hour,
	}))
	dependencies.InitUsers(r)
    if err := r.Run(); err != nil {
        panic(err)
    }
}
