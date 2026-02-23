package handler

import (
	"blogging-platform-api/internal/entity"
	"blogging-platform-api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	usecase entity.UserUsercase
}

func NewUserHandler(usecase entity.UserUsercase) entity.UserHandler {
	return &userHandler{
		usecase: usecase,
	}
}

func (u *userHandler) GetUserProfile(c *gin.Context) {
	username := c.Param("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, &entity.Resp{
			Message: "username_not_provided",
			Success: false,
		})
		return
	}

	existUser, err := u.usecase.GetUserProfile(username)
	if err != nil {
		httpErrStatus := utils.GetHttpErrStatus(err)
		c.JSON(httpErrStatus, &entity.Resp{
			Message: err.Error(),
			Success: false,
		})
		return
	}

	c.JSON(http.StatusOK, &entity.Resp{
		Data:    existUser,
		Success: true,
	})
}
