package role

import (
	"net/http"
	"sentinel-apm-api/utils/response"
	"sentinel-apm-api/utils/validation"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoleHandler interface {
	CreateRole(c *gin.Context)
}

type roleHandler struct {
	roleService RoleService
}

func RouteRegister(router *gin.RouterGroup, db *gorm.DB) {
	roleRepo := NewRoleRepository(db)
	roleService := NewRoleService(roleRepo)
	h := &roleHandler{roleService: roleService}

	roleGroup := router.Group("/role")
	roleGroup.POST("/create", h.CreateRole)
}

func (h *roleHandler) CreateRole(c *gin.Context) {
	var req RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusBadRequest,
			Error:  validation.FormateValidationError(err),
		})
		return
	}

	roleResponse, err := h.roleService.CreateRole(req)
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
		Data:   roleResponse,
	})
}
