package repository

import (
	"cms/model/entity"
	"context"
	"database/sql"
	"errors"
)

// create object loggerRepository
type LoggerRepository struct {
	DB     *sql.DB
	Entity *entity.Logger
}

// function provider to create new object loggerRepository
func NewLoggerRepository(db *sql.DB) *LoggerRepository {
	return &LoggerRepository{
		DB:     db,
		Entity: &entity.Logger{},
	}
}

// method insert
func (l *LoggerRepository) Insert(ctx context.Context, entity *entity.Logger) (*entity.Logger, error) {
	query := "INSERT INTO logger (ip_address, url_path, method, status_code, status, duration) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := l.DB.ExecContext(ctx, query, entity.IpAddress, entity.UrlPath, entity.Method, entity.StatusCode, entity.Status, entity.Duration)

	// jika ada error ketika proses insert
	if err != nil {
		return nil, err
	}

	// success insert -> get id inserted
	id, err := result.LastInsertId()

	// jika ada error ketika proses get LastInsertId
	if err != nil {
		return nil, err
	}

	// mapping id to entity.ID
	entity.ID = id

	// return
	return entity, nil
}

// method Get by Id
func (l *LoggerRepository) GetById(ctx context.Context, id int64) (*entity.Logger, error) {
	query := "SELECT id, ip_address, url_path, method, status_code, status, duration, created_at FROM logger WHERE id=? LIMIT 1"
	row, err := l.DB.QueryContext(ctx, query, id)

	// jika ada error ketika proses get data by id
	if err != nil {
		return nil, err
	}

	defer row.Close()

	// create initiate object
	logger := entity.Logger{}

	if row.Next() {
		// scan hasil query ke object logger
		err := row.Scan(&logger.ID, &logger.IpAddress, &logger.UrlPath, &logger.Method, &logger.StatusCode, &logger.Status, &logger.Duration, &logger.CreatedAt)
		if err != nil {
			return nil, err
		}

		// success scan
		return &logger, nil
	} else {
		// jika tidak ada datanya
		return nil, errors.New("record not found")
	}
}
