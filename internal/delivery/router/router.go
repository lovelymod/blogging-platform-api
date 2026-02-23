package router

import (
	"blogging-platform-api/internal/entity"
	"blogging-platform-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	BlogHandler entity.BlogHandler
	AuthHandler entity.AuthHandler
	UserHandler entity.UserHandler
}

func SetupRoutes(r *gin.Engine, h *Handlers, config *entity.Config) {
	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		// Auth
		auth.POST("/register", h.AuthHandler.Register)
		auth.POST("/login", h.AuthHandler.Login)
		auth.POST("/refresh-token", h.AuthHandler.RefreshToken)
	}

	private := api.Group("/")
	private.Use(middleware.AuthMiddleware(config))

	{
		// Auth
		{
			private.POST("/auth/logout", h.AuthHandler.Logout)
		}

		// Blog
		blog := private.Group("/blogs")
		{
			blog.GET("", h.BlogHandler.GetAll)
			blog.GET("/:id", h.BlogHandler.GetByID)
			blog.POST("", h.BlogHandler.Create)
			blog.PUT("/:id", h.BlogHandler.Update)
			blog.DELETE("/:id", h.BlogHandler.Delete)
		}

		// User
		{
			private.GET("/profile/:username", h.UserHandler.GetUserProfile)
		}
	}

}
