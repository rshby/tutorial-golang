package repository

import (
	"cms/model/entity"
	"context"
	"database/sql"
	"errors"
)

// create object ContentRepository
type ContentRepository struct {
	DB     *sql.DB
	Entity *entity.Content
}

// function provider to create repository
func NewContentRepository(db *sql.DB) *ContentRepository {
	return &ContentRepository{
		DB:     db,
		Entity: &entity.Content{},
	}
}

// method insert
func (c *ContentRepository) Insert(ctx context.Context, entity *entity.Content) (*entity.Content, error) {
	query := "INSERT INTO contents (title, picture_url, textfill, account_id) VALUES (?, ?, ?, ?)"
	result, err := c.DB.ExecContext(ctx, query, entity.Title, entity.PictureUrl, entity.TextFill, entity.AccountId)

	// jika ada keslahan ketika insert
	if err != nil {
		return nil, err
	}

	// get ID inserted
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// get data by ID
	content, err := c.GetById(ctx, id)
	if err != nil {
		c.DeleteById(ctx, id)
		return nil, err
	}

	// success insert and get inserted-data
	return content, nil
}

// method get all data contents
func (c *ContentRepository) GetAll(ctx context.Context) ([]entity.Content, error) {
	// query select all
	query := "SELECT id, title, picture_url, textfill, `like`, dislike, average_rating, created_at, account_id FROM contents"
	rows, err := c.DB.QueryContext(ctx, query)

	// jika ada kesalahan ketika query data
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// create object
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

	// jika datanya 0 -> not found
	if len(contents) == 0 {
		return nil, errors.New("record not found")
	}

	// success get all data contents
	return contents, nil
}

// method get by ID
func (c *ContentRepository) GetById(ctx context.Context, id int64) (*entity.Content, error) {
	query := "SELECT id, title, picture_url, textfill, `like`, dislike, average_rating, created_at, account_id FROM contents WHERE id=?"
	row, err := c.DB.QueryContext(ctx, query, id)

	// jika ada error ketika query
	if err != nil {
		return nil, err
	}

	defer row.Close()

	// create object
	content := entity.Content{}
	if row.Next() {
		// scan hasil query ke object content
		err := row.Scan(&content.ID, &content.Title, &content.PictureUrl, &content.TextFill, &content.Like, &content.Dislike, &content.AverageRating, &content.CreatedAt, &content.AccountId)

		// jika ada keslahan ketika scan
		if err != nil {
			return nil, err
		}
	} else {
		// jika data tidak ditemukan -> not found
		return nil, errors.New("record not found")
	}

	// success get data content by ID -> return
	return &content, nil
}

// method Delete by ID
func (c *ContentRepository) DeleteById(ctx context.Context, id int64) error {
	query := "DELETE FROM contents WHERE id=?"
	_, err := c.DB.ExecContext(ctx, query, id)

	// jika ada kesalahan ketika delete
	if err != nil {
		return err
	}

	// success delete contents from database
	return nil
}

// method Update Content by ID
func (c *ContentRepository) UpdateById(ctx context.Context, entity *entity.Content) error {
	// query update data
	query := "UPDATE contents SET title=?, picture_url=?, textfill=? WHERE id=?"
	_, err := c.DB.ExecContext(ctx, query, entity.Title, entity.PictureUrl, entity.TextFill, entity.ID)

	// jika ada kesalahan ketika update
	if err != nil {
		return err
	}

	// success update data
	return nil
}

// method Update Like by id
func (c *ContentRepository) UpdateLike(ctx context.Context, like int, id int64) error {
	query := "UPDATE contents SET `like` = ? WHERE id = ?"
	_, err := c.DB.ExecContext(ctx, query, like, id)

	// jika ada kesalahan ketika like
	if err != nil {
		return err
	}

	// success update like
	return nil
}

// method update dislike by id
func (c *ContentRepository) UpdateDislike(ctx context.Context, dislike int, id int64) error {
	// query update
	query := "UPDATE contents SET dislike = ? WHERE id = ?"
	_, err := c.DB.ExecContext(ctx, query, dislike, id)

	// jika ada kesalahan ketika update
	if err != nil {
		return err
	}

	// success update dislike by id
	return nil
}

// method Get by Username
func (c *ContentRepository) GetByUsername(ctx context.Context, username string) ([]entity.Content, error) {
	// query get contents
	query := "SELECT c.id, c.title, c.picture_url, c.textfill, c.`like`, c.dislike, c.average_rating, c.created_at, c.account_id FROM contents c INNER JOIN accounts a ON a.id=c.account_id WHERE a.username=?"
	rows, err := c.DB.QueryContext(ctx, query, username)

	// jika ada kesalahan ketika get data
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	// create object
	contents := []entity.Content{}
	for rows.Next() {
		// scan hasil query ke object
		content := entity.Content{}
		err = rows.Scan(&content.ID, &content.Title, &content.PictureUrl, &content.TextFill, &content.Like, &content.Dislike, &content.AverageRating, &content.CreatedAt, &content.AccountId)
		if err != nil {
			return nil, err
		}

		// append ke contents
		contents = append(contents, content)
	}

	// jika datanya tidak ada
	if len(contents) == 0 {
		return nil, errors.New("record not found")
	}

	// success get data
	return contents, nil
}

// method UpdateRating
func (c *ContentRepository) UpdateRating(ctx context.Context, rating float64, id int64) error {
	// query update
	query := "UPDATE contents SET average_rating = ? WHERE id = ?"

	// execute query
	_, err := c.DB.ExecContext(ctx, query, rating, id)

	// jika ada kesalahan ketika update
	if err != nil {
		return err
	}

	// success update
	return nil
}
