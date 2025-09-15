package godbpool

import (
	"context"
	"database/sql"
	"time"
)

type Options struct {
	MinConnections  int
	MaxConnections  int
	MaxIdleTime     time.Duration
	MaxOpenLifeTime time.Duration
}

// GetPool returns a pointer to a standard sql.DB object,
// configured with the provided options,
// and warms up the database connections.
func GetPool(db *sql.DB, options Options) (*sql.DB, error) {
	db.SetMaxIdleConns(options.MinConnections)
	db.SetMaxOpenConns(options.MaxConnections)
	db.SetConnMaxIdleTime(options.MaxIdleTime)
	db.SetConnMaxLifetime(options.MaxOpenLifeTime)

	if options.MinConnections == 0 {
		return db, nil
	}

	connections := make([]*sql.Conn, 0, options.MinConnections)

	for range options.MinConnections {
		connection, connectionErr := db.Conn(context.Background())
		if connectionErr != nil {
			return nil, connectionErr
		}

		connections = append(connections, connection)
	}

	for idx := range connections {
		connections[idx].Close()
	}

	return db, nil
}
