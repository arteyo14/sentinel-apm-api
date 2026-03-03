package auth

import (
	"net/http"
	"sentinel-apm-api/utils/context"
	"sentinel-apm-api/utils/cookie"
	"sentinel-apm-api/utils/response"
	"sentinel-apm-api/utils/validation"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler interface {
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type authHandler struct {
	authService AuthService
}

func RouteRegister(router *gin.RouterGroup, db *gorm.DB) {
	authRepo := NewAuthRepository(db)
	authService := NewAuthService(authRepo)
	h := &authHandler{authService: authService}

	authGroup := router.Group("/auth")
	authGroup.POST("/login", h.Login)
	authGroup.POST("/register", h.Register)
	authGroup.POST("/refresh-token", h.RefreshToken)
	authGroup.POST("/logout", h.Logout)
}

func (h *authHandler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusBadRequest,
			Error:  validation.FormateValidationError(err),
		})
		return
	}

	registerResponse, err := h.authService.Register(req)
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
		Data:   registerResponse,
	})
}

func (h *authHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusBadRequest,
			Error:  validation.FormateValidationError(err),
		})
		return
	}

	loginResponse, err := h.authService.Login(req)
	if err != nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	cookie.SetAuthCookies(c, loginResponse.AccessToken, loginResponse.RefreshToken)

	response.HandleResponse(c, response.Response{
		Status: true,
		Code:   http.StatusOK,
		Data:   loginResponse,
	})
}

func (h *authHandler) RefreshToken(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusBadRequest,
			Error:  validation.FormateValidationError(err),
		})
		return
	}

	session := context.GetSessionUser()
	if session == nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusInternalServerError,
			Error:  "no session found",
		})
		return
	}

	refreshTokenResponse, err := h.authService.RefreshToken(req, session)
	if err != nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	cookie.SetAuthCookies(c, refreshTokenResponse.AccessToken, refreshTokenResponse.RefreshToken)

	response.HandleResponse(c, response.Response{
		Status: true,
		Code:   http.StatusOK,
		Data:   refreshTokenResponse,
	})
}

func (h *authHandler) Logout(c *gin.Context) {
	session := context.GetSessionUser()
	if session == nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusInternalServerError,
			Error:  "no session found",
		})
		return
	}

	err := h.authService.Logout(session)
	if err != nil {
		response.HandleResponse(c, response.Response{
			Status: false,
			Code:   http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	cookie.ClearAuthCookies(c)

	response.HandleResponse(c, response.Response{
		Status: true,
		Code:   http.StatusOK,
		Data:   gin.H{},
	})
}
