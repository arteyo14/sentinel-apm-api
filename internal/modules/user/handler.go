package user

import (
	"net/http"
	"sentinel-apm-api/utils/response"
	"sentinel-apm-api/utils/validation"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserHandler interface {
	CreateUser(c *gin.Context)
}

type userHandler struct {
	userService UserService
}

func RouteRegister(router *gin.RouterGroup, db *gorm.DB) {
	userRepo := NewUserRepository(db)
	userService := NewUserService(userRepo)
	h := &userHandler{userService: userService}

	userGroup := router.Group("/user")
	userGroup.POST("/", h.CreateUser)
	userGroup.GET("/", h.GetUsers)
	userGroup.PUT("/:id", h.UpdateUser)
	userGroup.DELETE("/:id", h.DeleteUser)
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

func (h *userHandler) GetUsers(c *gin.Context) {
	users, err := h.userService.GetUsers()
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
		Data:   users,
	})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusBadRequest,
			Error:  validation.FormateValidationError(err),
		})
		return
	}

	userId := c.Param("id")
	userIdUUID, err := uuid.Parse(userId)
	if err != nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusBadRequest,
			Error:  "invalid user id format",
		})
		return
	}

	if err := h.userService.UpdateUser(userIdUUID, &req); err != nil {
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
			"message": "user updated successfully",
		},
	})
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	userIdUUID, err := uuid.Parse(userId)
	if err != nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusBadRequest,
			Error:  "invalid user id format",
		})
		return
	}

	if err := h.userService.DeleteUser(userIdUUID); err != nil {
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
			"message": "user deleted successfully",
		},
	})
}
