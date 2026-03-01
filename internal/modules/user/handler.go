package user

import (
	"net/http"
	"sentinel-apm-api/utils/response"
	"sentinel-apm-api/utils/validation"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
}

type userHandler struct {
	userService *userService
}

func RouteRegister(router *gin.RouterGroup, db *gorm.DB) {
	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo)
	h := &userHandler{userService: userService}

	userGroup := router.Group("/user")
	userGroup.POST("/", h.CreateUser)
}

func (h *userHandler) CreateUser(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusBadRequest,
			Error:  validation.FormateValidationError(err),
		})
		return
	}

	userId, err := h.userService.CreateUser(&req)
	if err != nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	response.HandleResponse(c, response.Response{
		Status: true,
		Code:   http.StatusOK,
		Data: gin.H{
			"user_id": userId,
		},
	})
}
