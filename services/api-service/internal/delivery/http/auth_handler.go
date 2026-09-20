package http

import (
	"errors"
	"net/http"

	"api-service/internal/delivery/http/dto"
	"api-service/internal/delivery/http/middleware"
	"api-service/internal/usecase"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.authUsecase.Register(c.Request.Context(), req.Username, req.Password, req.Name, req.Phone, req.Address)
	if err != nil {
		if errors.Is(err, usecase.ErrUserAlreadyExists) {
			middleware.SendError(c, http.StatusConflict, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusCreated, "Register successfull", dto.ToUserResponse(user))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.authUsecase.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCredentials) {
			middleware.SendError(c, http.StatusUnauthorized, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Login successfull", tokens)
}
