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

func (h *BlogHandler) GetAll(c *gin.Context) {
	filter := &entity.BlogFilter{
		Title:    c.Query("title"),
		Category: c.Query("category"),
		Tags:     []uint{},
	}

	tagsStr := c.QueryArray("tags")

	for _, v := range tagsStr {
		tagID, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return
		}

		filter.Tags = append(filter.Tags, uint(tagID))

	}

	blogs, err := h.Usecase.GetAll(c.Request.Context(), filter)

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

func (h *BlogHandler) GetByID(c *gin.Context) {
	blogID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	blog, err := h.Usecase.GetByID(c.Request.Context(), uint(blogID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	c.JSON(http.StatusOK, &entity.Resp{
		Data:    blog,
		Success: true,
	})
}

func (h *BlogHandler) Create(c *gin.Context) {
	var req entity.CreateBlogRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	tags := make([]entity.Tag, 0)
	for _, v := range req.Tags {
		tags = append(tags, entity.Tag{ID: v})
	}

	blog := entity.Blog{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Tags:     tags,
	}

	createdBlog, err := h.Usecase.Create(c.Request.Context(), &blog)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	c.JSON(http.StatusCreated, &entity.Resp{
		Data:    createdBlog,
		Success: true,
	})
}

func (h *BlogHandler) Update(c *gin.Context) {
	blogID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	var req entity.UpdateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	tags := make([]entity.Tag, 0)
	for _, v := range req.Tags {
		tags = append(tags, entity.Tag{ID: v})
	}

	updatedData := entity.Blog{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Tags:     tags,
	}

	updatedBlog, err := h.Usecase.Update(c.Request.Context(), uint(blogID), &updatedData)

	if err != nil {
		c.JSON(http.StatusInternalServerError, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	c.JSON(http.StatusOK, &entity.Resp{
		Data:    updatedBlog,
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
