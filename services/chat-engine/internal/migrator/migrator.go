package migrator

import (
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"chat-worker/internal/config"
	"chat-worker/migrations"

	"github.com/gocql/gocql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/cassandra"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

// Run executes database migrations using embedded CQL files.
func Run(cfg *config.DatabaseConfig) error {
	if len(cfg.Hosts) == 0 {
		return fmt.Errorf("cassandra hosts list is empty")
	}

	hostList, port := parseHosts(cfg.Hosts)

	// 1. Ensure Keyspace exists before binding session to it
	if err := ensureKeyspace(hostList, port, cfg.Keyspace); err != nil {
		return fmt.Errorf("failed to ensure keyspace '%s': %w", cfg.Keyspace, err)
	}

	// 2. Open Cassandra session bound to keyspace
	cluster := gocql.NewCluster(hostList...)
	cluster.Port = port
	cluster.Keyspace = cfg.Keyspace
	cluster.Consistency = gocql.LocalQuorum
	cluster.Timeout = 10 * time.Second
	cluster.ConnectTimeout = 15 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		return fmt.Errorf("failed to create session for migration: %w", err)
	}
	defer session.Close()

	// 3. Initialize Cassandra migrate driver
	driver, err := cassandra.WithInstance(session, &cassandra.Config{
		KeyspaceName: cfg.Keyspace,
	})
	if err != nil {
		return fmt.Errorf("failed to initialize cassandra migration driver: %w", err)
	}

	// 4. Initialize iofs source driver with embedded files
	sourceDriver, err := iofs.New(migrations.FS, ".")
	if err != nil {
		return fmt.Errorf("failed to initialize iofs source driver: %w", err)
	}

	// 5. Run migrations
	m, err := migrate.NewWithInstance("iofs", sourceDriver, "cassandra", driver)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("cassandra migration failed: %w", err)
	}

	log.Printf("[Migrator] Cassandra schema migration completed successfully for keyspace '%s'", cfg.Keyspace)
	return nil
}

func ensureKeyspace(hosts []string, port int, keyspace string) error {
	cluster := gocql.NewCluster(hosts...)
	cluster.Port = port
	cluster.Timeout = 10 * time.Second
	cluster.ConnectTimeout = 15 * time.Second
	cluster.RetryPolicy = &gocql.SimpleRetryPolicy{NumRetries: 3}

	session, err := cluster.CreateSession()
	if err != nil {
		return err
	}
	defer session.Close()

	query := fmt.Sprintf(
		"CREATE KEYSPACE IF NOT EXISTS %s WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};",
		keyspace,
	)

	return session.Query(query).Exec()
}

func parseHosts(rawHosts []string) ([]string, int) {
	var hostList []string
	port := 9042

	for _, h := range rawHosts {
		if strings.Contains(h, ":") {
			host, portStr, err := net.SplitHostPort(h)
			if err == nil {
				hostList = append(hostList, host)
				if parsedPort, err := strconv.Atoi(portStr); err == nil && parsedPort > 0 {
					port = parsedPort
				}
				continue
			}
		}
		hostList = append(hostList, h)
	}

	return hostList, port
}
