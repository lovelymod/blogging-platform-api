package bootstrap

import (
	"blogging-platform-api/internal/entity"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Application struct {
	DB     *gorm.DB
	S3     *s3.Client
	Config *entity.Config
	Cors   gin.HandlerFunc
}

func App() Application {
	gin.SetMode(gin.ReleaseMode)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	config := &entity.Config{
		SUPABASE_HOST:     os.Getenv("SUPABASE_HOST"),
		SUPABASE_USER:     os.Getenv("SUPABASE_USER"),
		SUPABASE_PASSWORD: os.Getenv("SUPABASE_PASSWORD"),
		SUPABASE_DB:       os.Getenv("SUPABASE_DB"),
		SUPABASE_PORT:     os.Getenv("SUPABASE_PORT"),

		R2_BUCKET_NAME:      os.Getenv("R2_BUCKET_NAME"),
		R2_PUBLIC_URL:       os.Getenv("R2_PUBLIC_URL"),
		R2_TOKEN:            os.Getenv("R2_TOKEN"),
		R2_ACCOUNT_ID:       os.Getenv("R2_ACCOUNT_ID"),
		R2_ACCESSKEY_ID:     os.Getenv("R2_ACCESSKEY_ID"),
		R2_ACCESSKEY_SECRET: os.Getenv("R2_ACCESSKEY_SECRET"),
		R2_S3_API:           os.Getenv("R2_S3_API"),

		HASH_COST: os.Getenv("HASH_COST"),

		ACCESS_TOKEN_SECRET:  os.Getenv("ACCESS_TOKEN_SECRET"),
		REFRESH_TOKEN_SECRET: os.Getenv("REFRESH_TOKEN_SECRET"),
	}

	app := &Application{
		Config: config,
	}
	app.DB = SetupDatabase(config)
	app.S3 = SetUpS3(config)

	app.Cors = cors.Default()
	return *app
}
