# godbpool

A lightweight database connection pool manager for Go that works with any database driver supporting the standard database/sql interface (PostgreSQL, MySQL, SQLite, etc.) and warms up connections.

## What It Does
🚀 Warms up the database connections

⚙️ Configurable connection limits

⏱️ Controls connection lifetimes

🔄 Returns standard *sql.DB objects

🧩 Database agnostic - works with any SQL driver

## Installation

```bash
go get github.com/barahouei/godbpool
```

## Quick Start

```go
options := godbpool.Options{
    MinConnections:  10,
    MaxConnections:  100,
    MaxIdleTime:     5 * time.Minute,
    MaxOpenLifeTime: 30 * time.Minute,
}

pool, err := godbpool.GetPool(db, options)
```
## Configuration

Option | Description
------------- | -------------
MinConnections | Minimum connections to pre-warm
MaxConnections | Maximum open connections
MaxIdleTime | Maximum idle time for connections
MaxOpenLifeTime | Maximum lifetime of connections

