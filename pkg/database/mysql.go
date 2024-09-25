package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New() *Mysql {
	return &Mysql{}
}

// SetConfigs is Setter to Set a MysqlConfig struct
func (m *Mysql) SetConfigs(c *MysqlConfig) *Mysql {
	m.config = c
	return m
}

func (m *Mysql) SetSqlDB(s *sql.DB) *Mysql {
	m.SqlDB = s
	return m
}

func (m *Mysql) GetDBConnection() any {
	return m.SqlDB
}

// Connect Create the connection with Mysql Adapter
func (m *Mysql) Connect() error {

	multistatements := ""
	if m.config.IsMultiStatements {
		multistatements = "&multiStatements=true"
	}

	var stringConnection = fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?parseTime=True%s", m.config.Username, m.config.Password, m.config.Host, m.config.Port, m.config.DBName, multistatements)

	db, err := sqlx.Connect("mysql", stringConnection)
	if err != nil {
		return err
	}
	if db == nil {
		if err != nil {
			for i := 0; i < m.config.ConnectionRetries; i++ {
				time.Sleep(time.Second * time.Duration(i+1))
				db, err = sqlx.Connect("mysql", stringConnection)
				if err != nil {
					continue
				}
				break
			}
		}
	}

	if db == nil {
		return fmt.Errorf("error connecting to mysql")
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	maxIdleConns := 10
	if m.config.MaxIdleConns > 0 {
		maxIdleConns = m.config.MaxIdleConns
	}
	db.DB.SetMaxIdleConns(maxIdleConns)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	maxOpenConn := 10
	if m.config.MaxOpenConn > 0 {
		maxOpenConn = m.config.MaxOpenConn
	}
	db.DB.SetMaxOpenConns(maxOpenConn)

	// // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	if m.config.ConnMaxLifetime > 0 {
		duration := m.config.ConnMaxLifetime
		timeDuration := time.Duration(time.Duration.Seconds(duration))
		db.DB.SetConnMaxLifetime(timeDuration)
	}

	m.HashKey = m.config.HashKey
	m.SetSqlDB(db.DB)

	if err := m.Ping(); err != nil {
		return err
	}

	m.DB = db

	return nil
}

// Ping verifies a connection to the database is still alive
func (m *Mysql) Ping() error {
	return m.SqlDB.Ping()
}

// Close closes the database and prevents new queries from starting.
func (m *Mysql) Close() error {
	if err := m.SqlDB.Close(); err != nil {
		return err
	}
	return nil
}
