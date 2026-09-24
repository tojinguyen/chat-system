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

type FriendHandler struct {
	friendUsecase usecase.FriendUsecase
}

func NewFriendHandler(friendUsecase usecase.FriendUsecase) *FriendHandler {
	return &FriendHandler{friendUsecase: friendUsecase}
}

func (h *FriendHandler) GetFriendsList(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	list, err := h.friendUsecase.GetFriendsList(c.Request.Context(), userID)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp := make([]*dto.FriendResponse, 0, len(list))
	for i := range list {
		resp = append(resp, dto.ToFriendResponse(&list[i]))
	}

	middleware.SendSuccess(c, http.StatusOK, "Lấy danh sách bạn bè thành công", resp)
}

func (h *FriendHandler) GetPendingRequests(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	list, err := h.friendUsecase.GetPendingRequests(c.Request.Context(), userID)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp := make([]*dto.FriendResponse, 0, len(list))
	for i := range list {
		resp = append(resp, dto.ToFriendResponse(&list[i]))
	}

	middleware.SendSuccess(c, http.StatusOK, "Lấy danh sách lời mời kết bạn thành công", resp)
}

func (h *FriendHandler) SendFriendRequest(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	targetIDStr := c.Param("targetUserId")
	targetID, err := uuid.Parse(targetIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID người dùng không hợp lệ")
		return
	}

	result, err := h.friendUsecase.SendFriendRequest(c.Request.Context(), userID, targetID)
	if err != nil {
		if errors.Is(err, usecase.ErrTargetNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusCreated, "Gửi lời mời kết bạn thành công", dto.ToFriendResponse(result))
}

func (h *FriendHandler) AcceptFriendRequest(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	reqIDStr := c.Param("requestId")
	reqID, err := uuid.Parse(reqIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID lời mời không hợp lệ")
		return
	}

	result, err := h.friendUsecase.AcceptFriendRequest(c.Request.Context(), userID, reqID)
	if err != nil {
		if errors.Is(err, usecase.ErrRelationshipNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, usecase.ErrNotAddressee) {
			middleware.SendError(c, http.StatusForbidden, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Đã chấp nhận lời mời kết bạn", dto.ToFriendResponse(result))
}

func (h *FriendHandler) DeclineFriendRequest(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	reqIDStr := c.Param("requestId")
	reqID, err := uuid.Parse(reqIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID lời mời không hợp lệ")
		return
	}

	result, err := h.friendUsecase.DeclineFriendRequest(c.Request.Context(), userID, reqID)
	if err != nil {
		if errors.Is(err, usecase.ErrRelationshipNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, usecase.ErrNotAddressee) {
			middleware.SendError(c, http.StatusForbidden, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Đã từ chối lời mời kết bạn", dto.ToFriendResponse(result))
}

func (h *FriendHandler) BlockUser(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	targetIDStr := c.Param("targetUserId")
	targetID, err := uuid.Parse(targetIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID người dùng không hợp lệ")
		return
	}

	result, err := h.friendUsecase.BlockUser(c.Request.Context(), userID, targetID)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Đã chặn người dùng", dto.ToFriendResponse(result))
}

func (h *FriendHandler) Unfriend(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	relIDStr := c.Param("relationshipId")
	relID, err := uuid.Parse(relIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID mối quan hệ không hợp lệ")
		return
	}

	success, err := h.friendUsecase.UnfriendOrCancel(c.Request.Context(), userID, relID)
	if err != nil {
		if errors.Is(err, usecase.ErrRelationshipNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Hủy kết bạn hoặc hủy yêu cầu thành công", gin.H{"success": success})
}
