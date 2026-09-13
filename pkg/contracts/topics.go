package contracts

import "fmt"

// GatewayNodeSubject returns the dedicated subject for a specific WebSocket Gateway node
func GatewayNodeSubject(nodeID string) string {
	return fmt.Sprintf("chat.gateway.%s", nodeID)
}
