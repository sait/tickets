package main

import (
	"fmt"

	"github.com/Flix14/soporte-tickets/api/db"
	"github.com/Flix14/soporte-tickets/api/models"
	"github.com/Flix14/soporte-tickets/api/router"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//DB d
var DB *gorm.DB

func main() {
	DB, err := db.InitConection()
	defer DB.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	DB.AutoMigrate(&models.Usuario{})
	DB.AutoMigrate(&models.Agente{})
	DB.AutoMigrate(&models.Ticket{})
	DB.AutoMigrate(&models.Mensaje{})

	r := gin.Default()

	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:8080"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))

	router.InitializeRoutes(r)

	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:8080"}

	// r.Use(cors.New(config))
	r.Run(":3000")
}
