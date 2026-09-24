package http

import (
	"errors"
	"net/http"

	"api-service/internal/delivery/http/dto"
	"api-service/internal/delivery/http/middleware"
	"api-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	user, err := h.userUsecase.GetProfile(c.Request.Context(), userID)
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Lấy thông tin cá nhân thành công", dto.ToUserResponse(user))
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	var req dto.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	updated, err := h.userUsecase.UpdateProfile(c.Request.Context(), userID, req.Name, req.Phone, req.Address)
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Cập nhật thông tin cá nhân thành công", dto.ToUserResponse(updated))
}

func (h *UserHandler) SearchUsers(c *gin.Context) {
	query := c.Query("q")
	users, err := h.userUsecase.SearchUsers(c.Request.Context(), query)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	respList := make([]*dto.UserResponse, 0, len(users))
	for i := range users {
		respList = append(respList, dto.ToUserResponse(&users[i]))
	}

	middleware.SendSuccess(c, http.StatusOK, "Tìm kiếm người dùng thành công", respList)
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	idStr := c.Param("id")
	targetID, err := uuid.Parse(idStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID người dùng không hợp lệ")
		return
	}

	user, err := h.userUsecase.FindByID(c.Request.Context(), targetID)
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			middleware.SendError(c, http.StatusNotFound, "Không tìm thấy người dùng")
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Lấy thông tin người dùng thành công", dto.ToUserResponse(user))
}
