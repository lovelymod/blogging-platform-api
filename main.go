package main

import (
	"blogging-platform-api/internal/bootstrap"
	"blogging-platform-api/internal/delivery/handler"
	"blogging-platform-api/internal/delivery/router"
	"blogging-platform-api/internal/provider"
	"blogging-platform-api/internal/repository"
	"blogging-platform-api/internal/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	// Object storage
	s3Provider := provider.NewS3Provider(app.S3, app.Config.R2_PUBLIC_URL, app.Config.R2_BUCKET_NAME)

	// Blog layer
	blogRepo := repository.NewBlogRepository(app.DB)
	blogUsecase := usecase.NewBlogUsecase(blogRepo, time.Second*2)
	blogHandler := handler.NewBlogHandler(blogUsecase)

	// Auth layer
	authRepo := repository.NewAuthRepository(app.DB)
	authUsecase := usecase.NewAuthUsecase(authRepo, time.Second*2, app.Config, s3Provider)
	authHandler := handler.NewAuthHandler(authUsecase)

	// User layer
	userRepo := repository.NewUserRepository(app.DB)
	userUsecase := usecase.NewUserUsecase(userRepo, s3Provider, time.Second*2)
	userHandler := handler.NewUserHandler(userUsecase)

	// Group handler
	h := &router.Handlers{
		BlogHandler: blogHandler,
		AuthHandler: authHandler,
		UserHandler: userHandler,
	}

	r := gin.Default()

	r.Use(app.Cors)

	router.SetupRoutes(r, h, app.Config)

	r.Run(":8080")
}
