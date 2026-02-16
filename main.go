package main

import (
	"blogging-platform-api/internal/bootstrap"
	"blogging-platform-api/internal/delivery/routes"
	"blogging-platform-api/internal/delivery/routes/handler"
	"blogging-platform-api/internal/repository"
	"blogging-platform-api/internal/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	app := bootstrap.App() // โหลด DB/Config มาไว้ในตัวแปรเดียว

	// สร้าง Layer ต่างๆ
	blogRepo := repository.NewBlogRepository(app.DB)
	blogUsecase := usecase.NewBlogUsecase(blogRepo, time.Second*2)
	blogHandler := handler.NewBlogHandler(blogUsecase)

	h := &routes.Handlers{
		BlogHandler: blogHandler,
	}

	r := gin.Default()

	r.Use(app.Cors)

	routes.SetupRoutes(r, h)

	r.Run(":8080")
}
