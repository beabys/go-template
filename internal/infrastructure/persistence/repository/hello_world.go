package repository

import (
	"context"
	"fmt"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/beabys/go-template/internal/domain/example/model"
	"github.com/beabys/go-template/pkg/database"
	"github.com/beabys/go-template/pkg/logger"
)

const (
	tableName = "hello"
)

type HelloWorldRepository struct {
	logger logger.Logger
	Db     database.Database
}

func NewHelloWorldRepository(log logger.Logger, db database.Database) *HelloWorldRepository {
	return &HelloWorldRepository{logger: log, Db: db}
}

func (r *HelloWorldRepository) SaveHelloWorld(ctx context.Context, m *model.HelloWorld) error {
	dbConn := r.Db.GetDBImpl()
	mysqlConn, ok := dbConn.(*database.Mysql)
	if !ok {
		return fmt.Errorf("invalid database connection")
	}

	v, err := time.Now().UTC().MarshalText()
	if err != nil {
		return fmt.Errorf("failed to marshal timestamp: %w", err)
	}

	q := sq.
		Insert(tableName).
		Columns("response").
		Values(string(v))

	result, err := q.RunWith(mysqlConn.SqlDB).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to save hello world to database: %w", err)
	}

	lastInsert, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	r.logger.Info("saved hello world", logger.LogField{Key: "last_insert_id", Value: lastInsert})

	return nil
}

func (r *HelloWorldRepository) GetHelloWorld(ctx context.Context, id model.HelloWorldID) (*model.HelloWorld, error) {
	dbConn := r.Db.GetDBImpl()
	mysqlConn, ok := dbConn.(*database.Mysql)
	if !ok {
		return nil, fmt.Errorf("invalid database connection")
	}

	q := sq.
		Select("id", "response", "created_at", "updated_at").
		From(tableName).
		Where("id = ?", id)

	var resp string
	var createdAt, updatedAt time.Time

	err := q.RunWith(mysqlConn.SqlDB).QueryRowContext(ctx).Scan(&id, &resp, &createdAt, &updatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get hello world: %w", err)
	}

	return &model.HelloWorld{
		ID:      id,
		Message: resp,
		Timestamps: model.Timestamps{
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
	}, nil
}
