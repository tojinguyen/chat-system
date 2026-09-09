package nats

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// Connect establishes a robust connection to NATS server with automatic reconnects.
func Connect(natsURL string, clientName string) (*nats.Conn, error) {
	opts := []nats.Option{
		nats.Name(clientName),
		nats.MaxReconnects(-1),
		nats.ReconnectWait(2 * time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("[NATS] Client %s disconnected: %v", clientName, err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("[NATS] Client %s reconnected to %s", clientName, nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Printf("[NATS] Client %s connection closed permanently", clientName)
		}),
	}

	nc, err := nats.Connect(natsURL, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS (%s): %w", natsURL, err)
	}
	return nc, nil
}
