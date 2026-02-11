package bootstrap

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Config struct {
	SUPABASE_HOST     string
	SUPABASE_USER     string
	SUPABASE_PASSWORD string
	SUPABASE_DB       string
	SUPABASE_PORT     string
}

type Application struct {
	DB     *gorm.DB
	Config *Config
	Cors   gin.HandlerFunc
}

func App() Application {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Println("Error loading .env file")
	}

	env := &Config{
		SUPABASE_HOST:     os.Getenv("SUPABASE_HOST"),
		SUPABASE_USER:     os.Getenv("SUPABASE_USER"),
		SUPABASE_PASSWORD: os.Getenv("SUPABASE_PASSWORD"),
		SUPABASE_DB:       os.Getenv("SUPABASE_DB"),
		SUPABASE_PORT:     os.Getenv("SUPABASE_PORT"),
	}

	app := &Application{
		Config: env,
	}
	app.DB = SetupDatabase(env) // เรียกฟังก์ชันจาก database.go

	app.Cors = cors.Default()
	return *app
}
