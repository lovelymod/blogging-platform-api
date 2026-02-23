package main

import (
	"blogging-platform-api/internal/bootstrap"
	"blogging-platform-api/internal/delivery/handler"
	"blogging-platform-api/internal/delivery/router"
	"blogging-platform-api/internal/repository"
	"blogging-platform-api/internal/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	// Blog layer
	blogRepo := repository.NewBlogRepository(app.DB)
	blogUsecase := usecase.NewBlogUsecase(blogRepo, time.Second*2)
	blogHandler := handler.NewBlogHandler(blogUsecase)

	// Auth layer
	authRepo := repository.NewAuthRepository(app.DB)
	authUsecase := usecase.NewAuthUsecase(authRepo, time.Second*2, app.Config)
	authHandler := handler.NewAuthHandler(authUsecase)

	// User layer
	userRepo := repository.NewUserRepository(app.DB)
	userUsecase := usecase.NewUserUsecase(userRepo, time.Second*2)
	userHandler := handler.NewUserHandler(userUsecase)

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
