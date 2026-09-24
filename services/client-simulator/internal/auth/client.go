package auth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type BotSession struct {
	UserID      string
	Username    string
	AccessToken string
}

type AuthClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewAuthClient(baseURL string) *AuthClient {
	return &AuthClient{
		baseURL: strings.TrimRight(baseURL, "/"),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

type registerPayload struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type loginPayload struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type apiResponse struct {
	StatusCode int             `json:"statusCode"`
	Message    string          `json:"message"`
	Data       json.RawMessage `json:"data"`
}

type loginResponseData struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (c *AuthClient) EnsureBotAuth(username, password string) (*BotSession, error) {
	// 1. Thử login trước
	session, err := c.login(username, password)
	if err == nil {
		return session, nil
	}

	// 2. Nếu login thất bại, thử register
	_ = c.register(username, password)

	// 3. Login lại sau khi register
	return c.login(username, password)
}

func (c *AuthClient) register(username, password string) error {
	body, _ := json.Marshal(registerPayload{
		UserName: username,
		Password: password,
		Name:     fmt.Sprintf("Bot %s", username),
	})

	resp, err := c.httpClient.Post(fmt.Sprintf("%s/auth/register", c.baseURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusConflict {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("register failed with status %d: %s", resp.StatusCode, string(b))
	}
	return nil
}

func (c *AuthClient) login(username, password string) (*BotSession, error) {
	body, _ := json.Marshal(loginPayload{
		UserName: username,
		Password: password,
	})

	resp, err := c.httpClient.Post(fmt.Sprintf("%s/auth/login", c.baseURL), "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("login failed status %d: %s", resp.StatusCode, string(respBytes))
	}

	var rawResp apiResponse
	var token string
	if err := json.Unmarshal(respBytes, &rawResp); err == nil && len(rawResp.Data) > 0 {
		var tokenData loginResponseData
		if err := json.Unmarshal(rawResp.Data, &tokenData); err == nil && tokenData.AccessToken != "" {
			token = tokenData.AccessToken
		}
	}

	if token == "" {
		// Fallback parse directly
		var tokenData loginResponseData
		if err := json.Unmarshal(respBytes, &tokenData); err == nil && tokenData.AccessToken != "" {
			token = tokenData.AccessToken
		}
	}

	if token == "" {
		return nil, fmt.Errorf("could not extract accessToken from response: %s", string(respBytes))
	}

	userID, err := extractUserIDFromJWT(token)
	if err != nil {
		return nil, fmt.Errorf("failed to extract sub/user_id from jwt: %w", err)
	}

	return &BotSession{
		UserID:      userID,
		Username:    username,
		AccessToken: token,
	}, nil
}

func extractUserIDFromJWT(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid jwt format")
	}

	payloadPart := parts[1]
	// Handle Base64 URL padding
	if l := len(payloadPart) % 4; l > 0 {
		payloadPart += strings.Repeat("=", 4-l)
	}

	payloadBytes, err := base64.URLEncoding.DecodeString(payloadPart)
	if err != nil {
		return "", err
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(payloadBytes, &claims); err != nil {
		return "", err
	}

	if sub, ok := claims["sub"].(string); ok && sub != "" {
		return sub, nil
	}
	if uid, ok := claims["user_id"].(string); ok && uid != "" {
		return uid, nil
	}

	return "", fmt.Errorf("no sub or user_id found in jwt claims")
}
