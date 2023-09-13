package repository

import (
	"13_repository_pattern/entity"
	repository "13_repository_pattern/repository/Interface"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

// buat struct yang menampung object repository untuk tabel comment
type commentRepository struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) repository.CommentRepositoryInterface {
	return &commentRepository{DB: db}
}

// buat function untuk insert -> implementasi dari interface
func (commentRepo *commentRepository) Insert(ctx context.Context, inputData entity.Comment) (entity.Comment, error) {
	query := "INSERT INTO comments (email, comment) VALUES (?, ?)"
	result, err := commentRepo.DB.ExecContext(ctx, query, inputData.Email, inputData.Comment)
	if err != nil {
		return inputData, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return inputData, err
	}

	inputData.Id = int32(id)
	return inputData, nil
}

// buat function untuk Get Data By Id -> implementasi dari interface
func (commentRepo *commentRepository) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := commentRepo.DB.QueryContext(ctx, query, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}

	defer rows.Close()
	if rows.Next() {
		// ada dataya
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		// tidak ada -> tidak ditemukan
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + " Tidak Ditemukan")
	}
}

// buat function untuk Get All Data -> implementasi dari interface
func (commentRepo *commentRepository) FindAll(ctx context.Context) ([]entity.Comment, error) {
	query := "SELECT id, email, comment FROM comments"
	rows, err := commentRepo.DB.QueryContext(ctx, query)
	comments := []entity.Comment{}
	comment := entity.Comment{}
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		// ada datanya
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
