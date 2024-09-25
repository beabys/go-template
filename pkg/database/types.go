package database

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

type Database interface {
	Connect() error
	Ping() error
	Close() error
	GetDBConnection() any
}

// Mysql type to connect to Mysql using Gorm
type Mysql struct {
	DB      *sqlx.DB
	SqlDB   *sql.DB
	config  *MysqlConfig
	HashKey string
}

// MysqlConfig type to connect to Mysql using Gorm
type MysqlConfig struct {
	Username          string
	Password          string
	Host              string
	Port              int
	DBName            string
	LogSQL            bool
	MaxIdleConns      int
	MaxOpenConn       int
	IsMultiStatements bool
	ConnectionRetries int
	ConnMaxLifetime   time.Duration
	HashKey           string
}
