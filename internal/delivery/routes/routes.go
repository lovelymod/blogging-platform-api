package routes

import (
	"blogging-platform-api/internal/entity"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	BlogHandler entity.BlogHandler
}

func SetupRoutes(r *gin.Engine, h *Handlers) {
	// Public Routes
	api := r.Group("/api")
	{

		api.GET("/blogs", h.BlogHandler.GetAll)
		api.GET("/blog/:id", h.BlogHandler.GetByID)
		api.POST("/blog", h.BlogHandler.Create)
		api.PUT("/blog/:id", h.BlogHandler.Update)
		api.DELETE("/blog/:id", h.BlogHandler.Delete)
	}

	// api := r.GroupogHandler.GetAll)

	// Protected Routes (ตัวอย่างถ้ามี Middleware)
	// api.Use(middleware.AuthMiddleware()).POST("/blogs", blogHandler.Create)
}
