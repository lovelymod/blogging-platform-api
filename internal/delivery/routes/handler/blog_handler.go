package handler

import (
	"blogging-platform-api/internal/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	Usecase entity.BlogUsecase
}

func (h *BlogHandler) Create(c *gin.Context) {
	var blog entity.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Usecase.Create(c.Request.Context(), &blog); err != nil {
		c.JSON(http.StatusInternalServerError, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	c.JSON(http.StatusCreated, &entity.Resp{
		Data:    blog,
		Success: true,
	})
}

func (h *BlogHandler) GetAll(c *gin.Context) {
	blogs, err := h.Usecase.GetAll(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return

	}

	c.JSON(http.StatusOK, &entity.Resp{
		Data:    blogs,
		Success: true,
	})
}

func (h *BlogHandler) Delete(c *gin.Context) {
	deleteID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	if err := h.Usecase.Delete(c.Request.Context(), uint(deleteID)); err != nil {
		c.JSON(http.StatusInternalServerError, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return

	}

	c.JSON(http.StatusOK, &entity.Resp{
		Success: true,
	})
}
