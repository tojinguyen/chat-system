package conversation

import (
	"bytes"
	"client-simulator/internal/auth"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Target struct {
	ConversationID string
	PartnerID      string
}

type ConversationResponse struct {
	ID      string `json:"id"`
	Members []struct {
		UserID string `json:"userId"`
	} `json:"members"`
}

type apiResponse struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       json.RawMessage `json:"data"`
}

type Manager struct {
	baseURL    string
	httpClient *http.Client
}

func NewManager(baseURL string) *Manager {
	return &Manager{
		baseURL: strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// EnsureDirectConversations đảm bảo mỗi bot có đủ số lượng direct conversation với các bot khác song song
func (m *Manager) EnsureDirectConversations(bots []*auth.BotSession, desiredConvsPerBot int) map[string][]Target {
	botTargets := make(map[string][]Target)
	totalBots := len(bots)

	if desiredConvsPerBot >= totalBots {
		desiredConvsPerBot = totalBots - 1
	}
	if desiredConvsPerBot <= 0 {
		desiredConvsPerBot = 1
	}

	log.Printf("[Conversation] Setting up conversation graph concurrently: %d bots, target ~%d convs/bot...",
		totalBots, desiredConvsPerBot)

	var mu sync.Mutex
	var wg sync.WaitGroup
	sem := make(chan struct{}, 30) // 30 concurrent queries

	for i, currentBot := range bots {
		if currentBot == nil {
			continue
		}
		wg.Add(1)
		go func(idx int, b *auth.BotSession) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()

			existingMap := make(map[string]string)

			// 1. Lấy danh sách conversation hiện có của bot
			existingConvos, err := m.fetchUserConversations(b.AccessToken)
			if err == nil {
				for _, convo := range existingConvos {
					for _, member := range convo.Members {
						if member.UserID != b.UserID && member.UserID != "" {
							existingMap[member.UserID] = convo.ID
						}
					}
				}
			}

			// 2. Nếu chưa đủ desiredConvsPerBot, tạo thêm với các bot khác
			for step := 1; step <= desiredConvsPerBot; step++ {
				partnerIdx := (idx + step) % totalBots
				if partnerIdx == idx {
					continue
				}
				partnerBot := bots[partnerIdx]
				if partnerBot == nil {
					continue
				}

				if _, exists := existingMap[partnerBot.UserID]; exists {
					continue
				}

				// Gọi POST /conversations/direct để tạo
				convoID, err := m.createDirectConversation(b.AccessToken, partnerBot.UserID)
				if err == nil && convoID != "" {
					existingMap[partnerBot.UserID] = convoID
				}
			}

			// Gom danh sách target cho bot hiện tại
			targets := make([]Target, 0, len(existingMap))
			for partnerID, convoID := range existingMap {
				targets = append(targets, Target{
					ConversationID: convoID,
					PartnerID:      partnerID,
				})
			}

			mu.Lock()
			botTargets[b.UserID] = targets
			mu.Unlock()
		}(i, currentBot)
	}
	wg.Wait()

	log.Printf("[Conversation] Conversation graph established successfully in parallel.")
	return botTargets
}

func (m *Manager) fetchUserConversations(token string) ([]ConversationResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/conversations", m.baseURL), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var convos []ConversationResponse
	var rawResp apiResponse
	if err := json.Unmarshal(respBytes, &rawResp); err == nil && len(rawResp.Data) > 0 {
		_ = json.Unmarshal(rawResp.Data, &convos)
	}

	if len(convos) == 0 {
		_ = json.Unmarshal(respBytes, &convos)
	}

	return convos, nil
}

func (m *Manager) createDirectConversation(token, partnerID string) (string, error) {
	body, _ := json.Marshal(map[string]string{
		"partnerId": partnerID,
	})

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/conversations/direct", m.baseURL), bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	req.Header.Set("Content-Type", "application/json")

	resp, err := m.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var created ConversationResponse
	var rawResp apiResponse
	if err := json.Unmarshal(respBytes, &rawResp); err == nil && len(rawResp.Data) > 0 {
		_ = json.Unmarshal(rawResp.Data, &created)
	}

	if created.ID == "" {
		_ = json.Unmarshal(respBytes, &created)
	}

	return created.ID, nil
}
