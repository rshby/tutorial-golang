package repository

import (
	"cms/model/entity"
	"context"
	"database/sql"
	"errors"
)

// create object LikeRepository
type LikeRepository struct {
	DB     *sql.DB
	Entity *entity.Like
}

// function provider to create new object Like-Repository
func NewLikeRepository(db *sql.DB) *LikeRepository {
	return &LikeRepository{
		DB:     db,
		Entity: &entity.Like{},
	}
}

// method Insert
func (l *LikeRepository) Insert(ctx context.Context, entity *entity.Like) (*entity.Like, error) {
	// query insert
	query := "INSERT INTO likes (account_id, content_id) VALUES (?, ?)"
	result, err := l.DB.ExecContext(ctx, query, entity.AccountId, entity.ContentId)

	// jika ada kesalahan ketika insert
	if err != nil {
		return nil, err
	}

	// get inserted ID
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	entity.ID = id

	return entity, nil
}

// method Delete / unlike
func (l *LikeRepository) Delete(ctx context.Context, accountId int64, contentId int64) error {
	// query delete
	query := "DELETE FROM likes WHERE account_id=? AND content_id=?"
	_, err := l.DB.ExecContext(ctx, query, accountId, contentId)

	// jika ada kesalahan ketika delete
	if err != nil {
		return err
	}

	// success delete
	return nil
}

// method Get Like by account_id and content_id
func (l *LikeRepository) GetByAccountIdAndContentId(ctx context.Context, accId int64, contentId int64) (*entity.Like, error) {
	// query
	query := "SELECT id, account_id, content_id FROM likes WHERE account_id = ? AND content_id = ?"
	row, err := l.DB.QueryContext(ctx, query, accId, contentId)

	// jika ada kesalahan ketika get data
	if err != nil {
		return nil, err
	}

	defer row.Close()

	// create object
	like := entity.Like{}
	if row.Next() {
		// scan hasil query ke object
		err := row.Scan(&like.ID, &like.AccountId, &like.ContentId)
		if err != nil {
			return nil, err
		}
	} else {
		// jika data tidak ditemukan
		return nil, errors.New("record not found")
	}

	// success get data
	return &like, nil
}

// method get total like each content
func (l *LikeRepository) GetTotalLikes(ctx context.Context, contentId int64) (*int, error) {
	query := "SELECT COUNT(1) FROM likes WHERE content_id=?"
	row, err := l.DB.QueryContext(ctx, query, contentId)

	// jika ada error ketika query
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var totalCount int

	if row.Next() {
		err := row.Scan(&totalCount)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("content not found")
	}

	// success count
	return &totalCount, nil
}

// method get likes by account_id
func (l *LikeRepository) GetByAccountId(ctx context.Context, accountId int64) ([]entity.Content, error) {
	// query select
	query := "SELECT c.id, c.title, c.picture_url, c.textfill, c.`like`, c.dislike, c.average_rating, c.created_at, c.account_id FROM contents c INNER JOIN likes l ON l.content_id = c.id WHERE l.account_id = ?"

	// execute query
	rows, err := l.DB.QueryContext(ctx, query, accountId)

	// jika ada kesalahan ketika proses query select data
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// create variabel penampung
	contents := []entity.Content{}
	for rows.Next() {
		// scan hasil query ke object
		content := entity.Content{}
		err := rows.Scan(&content.ID, &content.Title, &content.PictureUrl, &content.TextFill, &content.Like, &content.Dislike, &content.AverageRating, &content.CreatedAt, &content.AccountId)
		if err != nil {
			return nil, err
		}

		// append to contents
		contents = append(contents, content)
	}

	// cek jika datanya kosong
	if len(contents) == 0 {
		return nil, errors.New("record not found")
	}

	// success get data
	return contents, nil
}
