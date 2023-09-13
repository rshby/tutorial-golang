package repository

import (
	"13_repository_pattern/entity"
	"context"
)

// buat interface untuk comment repository
type CommentRepositoryInterface interface {
	Insert(ctx context.Context, inputData entity.Comment) (entity.Comment, error)
	FindById(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}
