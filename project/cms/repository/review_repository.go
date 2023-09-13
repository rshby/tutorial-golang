package repository

import (
	"cms/model/dto"
	"cms/model/entity"
	"context"
	"database/sql"
	"errors"
)

// crete object ReviewRepository
type ReviewRepository struct {
	DB     *sql.DB
	Entity *entity.Review
}

// function provider to create obejct ReviewRepository
func NewReviewRepository(db *sql.DB) *ReviewRepository {
	return &ReviewRepository{
		DB:     db,
		Entity: &entity.Review{},
	}
}

// method Insert reviews
func (r *ReviewRepository) Insert(ctx context.Context, entity *entity.Review) (*entity.Review, error) {
	// query insert
	query := "INSERT INTO reviews (rating, comment, account_id, content_id) VALUES (?, ?, ?, ?)"

	// execute
	result, err := r.DB.ExecContext(ctx, query, entity.Rating, entity.Comment, entity.AccountId, entity.ContentId)

	// jika ada kesalahan ketika insert
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// get data by ID
	review, err := r.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	// success insert and get data inserted
	return review, nil
}

// method get by ID
func (r *ReviewRepository) GetById(ctx context.Context, id int64) (*entity.Review, error) {
	// query select
	query := "SELECT id, rating, comment, created_at, account_id, content_id FROM reviews WHERE id = ?"

	// execute
	row, err := r.DB.QueryContext(ctx, query, id)

	// jika ada error ketika query
	if err != nil {
		return nil, err
	}

	defer row.Close()

	// create object entity
	review := entity.Review{}
	if row.Next() {
		// scan hasil query ke object
		err := row.Scan(&review.ID, &review.Rating, &review.Comment, &review.CreatedAt, &review.AccountId, &review.ContentId)
		if err != nil {
			return nil, err
		}
	} else {
		// jika datanya kosong / not found
		return nil, errors.New("record not found")
	}

	// success get data review by ID
	return &review, nil
}

// method delete review by ID
func (r *ReviewRepository) Delete(ctx context.Context, id int64) error {
	// query delete
	query := "DELETE FROM reviews WHERE id = ?"

	// execute query
	_, err := r.DB.ExecContext(ctx, query, id)

	// jika ada kesalahan ketika delete
	if err != nil {
		return err
	}

	// success delete reviews by ID
	return nil
}

// method count n by content_id
func (r *ReviewRepository) CountRatingByContentId(ctx context.Context, contentId int64) (*int, error) {
	// query select COUNT
	query := "SELECT COUNT(1) FROM reviews WHERE content_id = ?"

	// execute
	row, err := r.DB.QueryContext(ctx, query, contentId)

	// jika ada kesalahan ketika query
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var countRating int
	if row.Next() {
		// scan hasil query ke variabel
		err = row.Scan(&countRating)
		if err != nil {
			return nil, err
		}
	} else {
		// jika datanya tidak ada
		return nil, errors.New("record not found")
	}

	// success count rating by content_id
	return &countRating, nil
}

// method get SUM rating by content_id
func (r *ReviewRepository) SumRatingByContentId(ctx context.Context, contentId int64) (*int, error) {
	// query select sum
	query := "SELECT SUM(rating) FROM reviews WHERE content_id = ?"

	// execute query
	row, err := r.DB.QueryContext(ctx, query, contentId)

	// jika ada kesalahan ketika query
	if err != nil {
		return nil, err
	}

	defer row.Close()

	var sumRating int

	if row.Next() {
		// scan hasil query
		err = row.Scan(&sumRating)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("record not found")
	}

	// success sum rating
	return &sumRating, nil
}

// method check review by content_id, account_id, ID
func (r *ReviewRepository) CheckReview(ctx context.Context, input *dto.DeleteReviewRequest) (*entity.Review, error) {
	// query select
	query := "SELECT id, rating, comment, created_at, account_id, content_id FROM reviews WHERE id = ? AND account_id = ? AND content_id = ?"

	// execute query
	row, err := r.DB.QueryContext(ctx, query, input.ReviewId, input.AccountId, input.ContentId)

	// jika ada kesalahan ketika query
	if err != nil {
		return nil, err
	}

	defer row.Close()

	// create variabel
	review := entity.Review{}
	if row.Next() {
		// scan hasil query ke variabel
		err = row.Scan(&review.ID, &review.Rating, &review.Comment, &review.CreatedAt, &review.AccountId, &review.ContentId)
		if err != nil {
			return nil, err
		}
	} else {
		// jika datanya not found / kosong
		return nil, errors.New("record not found")
	}

	// success get data
	return &review, nil
}
