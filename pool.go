package godbpool

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNilDB error = errors.New("db cannot be nil")
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
	if db == nil {
		return nil, ErrNilDB
	}

	db.SetMaxIdleConns(options.MinConnections)
	db.SetMaxOpenConns(options.MaxConnections)
	db.SetConnMaxIdleTime(options.MaxIdleTime)
	db.SetConnMaxLifetime(options.MaxOpenLifeTime)

	if options.MinConnections == 0 {
		return db, nil
	}

	connections := make([]*sql.Conn, 0, options.MinConnections)

	defer closeConnections(connections)

	for range options.MinConnections {
		connection, connectionErr := db.Conn(context.Background())
		if connectionErr != nil {
			return nil, connectionErr
		}

		connections = append(connections, connection)
	}

	return db, nil
}

func closeConnections(connections []*sql.Conn) {
	for idx := range connections {
		if connections[idx] != nil {
			connections[idx].Close()
		}
	}
}
