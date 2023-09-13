package repository

import (
	"cms/model/entity"
	"context"
	"database/sql"
	"errors"
)

// create object DislikeRepository
type DislikeRepository struct {
	DB     *sql.DB
	Entity *entity.Dislike
}

// function provider to create new object DislikeRepo
func NewDislikeRepository(db *sql.DB) *DislikeRepository {
	return &DislikeRepository{
		DB:     db,
		Entity: &entity.Dislike{},
	}
}

// method Insert Dislike
func (d *DislikeRepository) Insert(ctx context.Context, entity *entity.Dislike) (*entity.Dislike, error) {
	// query insert dislike
	query := "INSERT INTO dislikes (account_id, content_id) VALUES (?, ?)"
	result, err := d.DB.ExecContext(ctx, query, entity.AccountId, entity.ContentId)

	// jika ada error ketika insert dislike
	if err != nil {
		return nil, err
	}

	// get id inserted
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	entity.ID = id

	// success insert dislike
	return entity, nil
}

// method Delete
func (d *DislikeRepository) Delete(ctx context.Context, accId int64, contentId int64) error {
	// query delete dislike
	query := "DELETE FROM dislikes WHERE account_id = ? AND content_id = ?"
	_, err := d.DB.ExecContext(ctx, query, accId, contentId)

	// jika ada kesalahan saat delete
	if err != nil {
		return err
	}

	// success delete data dislike
	return nil
}

// method Get dislike by account_id and content_id
func (d *DislikeRepository) GetByAccountIdAndContentId(ctx context.Context, accId int64, contentId int64) (*entity.Dislike, error) {
	// query get data
	query := "SELECT id, account_id, content_id FROM dislikes WHERE account_id = ? AND content_id = ?"
	row, err := d.DB.QueryContext(ctx, query, accId, contentId)

	// jika ada kesalahan ketika get data
	if err != nil {
		return nil, err
	}

	defer row.Close()

	// create object
	dislike := entity.Dislike{}

	if row.Next() {
		// scan hasil query ke object
		err = row.Scan(&dislike.ID, &dislike.AccountId, &dislike.ContentId)
		if err != nil {
			return nil, err
		}
	} else {
		// jika data tidak ditemukan
		return nil, errors.New("record not found")
	}

	// success get data dislike
	return &dislike, nil
}

// method Count Total likes by content_id
func (d *DislikeRepository) GetTotalDislikes(ctx context.Context, contentId int64) (*int, error) {
	// query count
	query := "SELECT COUNT(1) FROM dislikes WHERE content_id = ?"
	row, err := d.DB.QueryContext(ctx, query, contentId)

	// jika ada error ketika query
	if err != nil {
		return nil, err
	}

	defer row.Close()

	// create variabel untuk menampung hasil query
	var totalDislikes int

	if row.Next() {
		// scan hasil query ke variabel penampung
		err = row.Scan(&totalDislikes)
		if err != nil {
			return nil, err
		}
	} else {
		// jika data tidak ditemukan
		return nil, errors.New("record not found")
	}

	// success get total dislikes
	return &totalDislikes, nil
}

// method get contents disliked by username
func (d *DislikeRepository) GetContentDislikedByAccountId(ctx context.Context, accountId int64) ([]entity.Content, error) {
	// query select join
	query := "SELECT c.id, c.title, c.picture_url, c.textfill, c.`like`, c.dislike, c.average_rating, c.created_at, c.account_id FROM contents c INNER JOIN dislikes d ON d.content_id = c.id WHERE d.account_id = ?"

	// execute
	rows, err := d.DB.QueryContext(ctx, query, accountId)

	// jika ada error ketika query
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// siapkan variabel penampung
	contents := []entity.Content{}
	for rows.Next() {
		// scan hasil query ke object
		content := entity.Content{}
		err := rows.Scan(&content.ID, &content.Title, &content.PictureUrl, &content.TextFill, &content.Like, &content.Dislike, &content.AverageRating, &content.CreatedAt, &content.AccountId)
		if err != nil {
			return nil, err
		}

		// append ke variabel contents
		contents = append(contents, content)
	}

	// jika data kosong
	if len(contents) == 0 {
		return nil, errors.New("record not found")
	}

	// success get data contents disliked by account_id
	return contents, nil
}
