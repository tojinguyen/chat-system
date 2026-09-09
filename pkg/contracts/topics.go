package contracts

import "fmt"

const (
	// SubjectInbound is the default subject where gateways publish client messages
	SubjectInbound = "chat.inbound"

	// SubjectNotification is the subject consumed by notification worker
	SubjectNotification = "chat.notification"

	// SubjectGlobal is the broadcast channel for system-wide announcements
	SubjectGlobal = "chat.global"
)

// GatewayNodeSubject returns the dedicated subject for a specific WebSocket Gateway node
func GatewayNodeSubject(nodeID string) string {
	return fmt.Sprintf("chat.gateway.%s", nodeID)
}

// ClanSubject returns the dedicated subject for a specific clan/guild channel
func ClanSubject(clanID string) string {
	return fmt.Sprintf("chat.clan.%s", clanID)
}

// RegionSubject returns the dedicated subject for a specific geographic region
func RegionSubject(regionID string) string {
	return fmt.Sprintf("chat.region.%s", regionID)
}
