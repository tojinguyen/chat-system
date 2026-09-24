package http

import (
	"errors"
	"net/http"

	"api-service/internal/delivery/http/dto"
	"api-service/internal/delivery/http/middleware"
	"api-service/internal/domain"
	"api-service/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ConversationHandler struct {
	convoUsecase usecase.ConversationUsecase
}

func NewConversationHandler(convoUsecase usecase.ConversationUsecase) *ConversationHandler {
	return &ConversationHandler{convoUsecase: convoUsecase}
}

func (h *ConversationHandler) GetMyConversations(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	convos, err := h.convoUsecase.GetUserConversations(c.Request.Context(), userID)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp := make([]*dto.ConversationResponse, 0, len(convos))
	for i := range convos {
		resp = append(resp, dto.ToConversationResponse(&convos[i]))
	}

	middleware.SendSuccess(c, http.StatusOK, "Lấy danh sách cuộc hội thoại thành công", resp)
}

func (h *ConversationHandler) GetOrCreateDirect(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	var req dto.CreateDirectConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	partnerID, err := uuid.Parse(req.PartnerID)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "partnerId không hợp lệ")
		return
	}

	convo, err := h.convoUsecase.GetOrCreateDirectConversation(c.Request.Context(), userID, partnerID)
	if err != nil {
		if errors.Is(err, usecase.ErrTargetNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Khởi tạo cuộc hội thoại 1-1 thành công", dto.ToConversationResponse(convo))
}

func (h *ConversationHandler) CreateGroup(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	var req dto.CreateGroupConversationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	memberUUIDs := make([]uuid.UUID, 0, len(req.MemberIDs))
	for _, m := range req.MemberIDs {
		if uid, err := uuid.Parse(m); err == nil {
			memberUUIDs = append(memberUUIDs, uid)
		}
	}

	convo, err := h.convoUsecase.CreateGroupConversation(c.Request.Context(), userID, req.Name, req.IconURL, memberUUIDs)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusCreated, "Tạo nhóm chat thành công", dto.ToConversationResponse(convo))
}

func (h *ConversationHandler) GetConversationByID(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	idStr := c.Param("id")
	convoID, err := uuid.Parse(idStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID cuộc hội thoại không hợp lệ")
		return
	}

	convo, err := h.convoUsecase.GetConversationByID(c.Request.Context(), convoID, userID)
	if err != nil {
		if errors.Is(err, usecase.ErrConversationNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Lấy chi tiết cuộc hội thoại thành công", dto.ToConversationResponse(convo))
}

func (h *ConversationHandler) AddMember(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	convoIDStr := c.Param("id")
	convoID, err := uuid.Parse(convoIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID cuộc hội thoại không hợp lệ")
		return
	}

	var req dto.AddMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	targetMemberID, err := uuid.Parse(req.UserID)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "userId không hợp lệ")
		return
	}

	role := req.Role
	if role == "" {
		role = domain.MemberRoleMember
	}

	member, err := h.convoUsecase.AddMember(c.Request.Context(), convoID, userID, targetMemberID, role)
	if err != nil {
		if errors.Is(err, usecase.ErrConversationNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, usecase.ErrNotMember) {
			middleware.SendError(c, http.StatusForbidden, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp := dto.ConversationMemberResponse{
		ID:             member.ID.String(),
		ConversationID: member.ConversationID.String(),
		UserID:         member.UserID.String(),
		Role:           string(member.Role),
		JoinedAt:       member.JoinedAt,
		User:           dto.ToUserResponse(member.User),
	}

	middleware.SendSuccess(c, http.StatusCreated, "Thêm thành viên vào nhóm thành công", resp)
}

func (h *ConversationHandler) RemoveMember(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	convoIDStr := c.Param("id")
	convoID, err := uuid.Parse(convoIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID cuộc hội thoại không hợp lệ")
		return
	}

	targetIDStr := c.Param("targetUserId")
	targetID, err := uuid.Parse(targetIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "targetUserId không hợp lệ")
		return
	}

	success, err := h.convoUsecase.RemoveMember(c.Request.Context(), convoID, userID, targetID)
	if err != nil {
		if errors.Is(err, usecase.ErrConversationNotFound) {
			middleware.SendError(c, http.StatusNotFound, err.Error())
			return
		}
		if errors.Is(err, usecase.ErrNotMember) || errors.Is(err, usecase.ErrNotAdmin) {
			middleware.SendError(c, http.StatusForbidden, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	middleware.SendSuccess(c, http.StatusOK, "Xóa thành viên hoặc rời nhóm thành công", gin.H{"success": success})
}

func (h *ConversationHandler) UpdateReadReceipt(c *gin.Context) {
	userID := middleware.GetCurrentUserID(c)
	convoIDStr := c.Param("id")
	convoID, err := uuid.Parse(convoIDStr)
	if err != nil {
		middleware.SendError(c, http.StatusBadRequest, "ID cuộc hội thoại không hợp lệ")
		return
	}

	var req dto.UpdateReadReceiptRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	receipt, err := h.convoUsecase.UpdateReadReceipt(c.Request.Context(), convoID, userID, req.LastMessageSeen)
	if err != nil {
		if errors.Is(err, usecase.ErrNotMember) {
			middleware.SendError(c, http.StatusForbidden, err.Error())
			return
		}
		middleware.SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	resp := gin.H{
		"userId":          receipt.UserID.String(),
		"lastMessageSeen": receipt.LastMessageSeen,
		"lastReadAt":      receipt.LastReadAt,
	}

	middleware.SendSuccess(c, http.StatusOK, "Cập nhật trạng thái đã đọc thành công", resp)
}
