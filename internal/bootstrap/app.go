package bootstrap

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Application struct {
	DB         *gorm.DB
	CorsConfig gin.HandlerFunc
}

func App() Application {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := &Application{}
	app.DB = SetupDatabase() // เรียกฟังก์ชันจาก database.go

	app.CorsConfig = cors.New(cors.DefaultConfig())
	return *app
}
