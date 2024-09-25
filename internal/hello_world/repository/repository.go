package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"gitlab.com/beabys/go-template/internal/domain/model"
	"gitlab.com/beabys/go-template/pkg/database"
	"gitlab.com/beabys/go-template/pkg/logger"
)

const (
	tableName = "hello"
)

func NewHelloRepository(log logger.Logger, db database.Database) *RepositoryHelloWorld {
	return &RepositoryHelloWorld{log, db}
}

func (r *RepositoryHelloWorld) SaveHelloWorld(ctx context.Context, m *model.HelloWorld) error {
	v, err := time.Now().UTC().MarshalText()
	if err != nil {
		msg := "not able to Marshall date into text"
		r.logger.Error(msg, err)
		return errors.New(msg)
	}
	dbConn := r.Db.GetDBConnection().(*sql.DB)
	q := sq.
		Insert(tableName).
		Columns("response").
		Values(string(v))
	result, err := q.RunWith(dbConn).ExecContext(ctx)
	if err != nil {
		msg := "error saving in db"
		r.logger.Error(msg, err)
		return fmt.Errorf(msg+": %w", err)
	}
	lastInsert, err := result.LastInsertId()
	if err != nil {
		msg := "error getting last inserted"
		r.logger.Error(msg, err)
		return fmt.Errorf(msg+": %w", err)
	}
	r.logger.Info(fmt.Sprintf("last inserted: %d", lastInsert))
	m.Hello = string(v)
	return nil
}
