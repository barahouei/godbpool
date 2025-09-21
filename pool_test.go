package godbpool

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	testDefaultOptions = Options{
		MinConnections:  10,
		MaxConnections:  100,
		MaxIdleTime:     time.Minute * 5,
		MaxOpenLifeTime: time.Minute * 30,
	}

	testNoMinConnectionsOptions = Options{
		MaxConnections:  100,
		MaxIdleTime:     time.Minute * 5,
		MaxOpenLifeTime: time.Minute * 30,
	}
)

// setupTestDB creates an in-memory SQLite database for testing.
func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, dbErr := sql.Open("sqlite3", ":memory:")
	if dbErr != nil {
		t.Fatalf("could not open a database connection: %v", dbErr)
	}

	return db
}

// TestGetPool tests the main functionality of GetPool function
// Verifies that connection pool is properly configured with correct min/max connections.
func TestGetPool(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	pool, poolErr := GetPool(db, testDefaultOptions)
	if poolErr != nil {
		t.Fatalf("could not warm up database connection pool: %v", poolErr)
	}

	poolStats := pool.Stats()

	if poolStats.OpenConnections != testDefaultOptions.MinConnections {
		t.Errorf("got %d minimum connections, but want %d minimum connections.", poolStats.OpenConnections, testDefaultOptions.MinConnections)
	}

	if poolStats.MaxOpenConnections != testDefaultOptions.MaxConnections {
		t.Errorf("got %d minimum connections, but want %d minimum connections.", poolStats.MaxOpenConnections, testDefaultOptions.MinConnections)
	}
}

// TestGetPoolNoDB tests error handling when nil database is provided
// Verifies that appropriate error is returned for nil input
func TestGetPoolNoDB(t *testing.T) {
	_, poolErr := GetPool(nil, testDefaultOptions)
	if poolErr != ErrNilDB {
		t.Errorf("want error: %v, got error: %v", ErrNilDB, poolErr)
	}
}

// TestGetPoolNoMinConnections tests GetPool functionality when MinConnections is zero
// Verifies that pool works correctly without pre-warming connections.
func TestGetPoolNoMinConnections(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	pool, poolErr := GetPool(db, testNoMinConnectionsOptions)
	if poolErr != nil {
		t.Fatalf("could not warm up database connection pool: %v", poolErr)
	}

	poolStats := pool.Stats()

	if poolStats.OpenConnections != testNoMinConnectionsOptions.MinConnections {
		t.Errorf("got %d minimum connections, but want %d minimum connections.", poolStats.OpenConnections, testNoMinConnectionsOptions.MinConnections)
	}

	if poolStats.MaxOpenConnections != testNoMinConnectionsOptions.MaxConnections {
		t.Errorf("got %d minimum connections, but want %d minimum connections.", poolStats.MaxOpenConnections, testNoMinConnectionsOptions.MinConnections)
	}
}
