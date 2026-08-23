package handler

import (
	"errors"
	"net/http"
	"ws-gateway/internal/connection"
	"ws-gateway/internal/domain"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func validateToken(tokenStr string, secret string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if sub, ok := claims["sub"].(string); ok {
			return sub, nil
		}
	}

	return "", jwt.ErrInvalidKeyType
}

func HandleWebSocket(hub *connection.Hub, jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.URL.Query().Get("token")
		deviceID := r.URL.Query().Get("device_id")

		if deviceID == "" {
			http.Error(w, "missing device id", http.StatusBadRequest)
			return
		}

		if tokenStr == "" {
			http.Error(w, "missing token", http.StatusBadRequest)
			return
		}

		userID, err := validateToken(tokenStr, jwtSecret)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Failed to establish websocket connection", http.StatusInternalServerError)
			return
		}

		client := &connection.Client{
			UserID:   userID,
			DeviceID: deviceID,
			Hub:      hub,
			SendChan: make(chan *domain.WSMessage, 256),
			Conn:     conn,
		}

		hub.RegisterClient(client)

		go client.ReadPump()
		go client.WritePump()
	}
}
