package service

import (
	"cms/model/dto"
	"cms/model/entity"
	"cms/repository"
	"context"
	"errors"
	"sync"

	"github.com/go-playground/validator/v10"
)

// create object LikeService
type LikeService struct {
	LikeRepository    *repository.LikeRepository
	DislikeRepository *repository.DislikeRepository
	AccountRepository *repository.AccountRepository
	ContentRepository *repository.ContentRepository
	Validate          *validator.Validate
}

// function provider to create new object LikeService
func NewLikeService(likeRepo *repository.LikeRepository, accRepo *repository.AccountRepository, contentRepo *repository.ContentRepository, dislikeRepo *repository.DislikeRepository, validate *validator.Validate) *LikeService {
	return &LikeService{
		LikeRepository:    likeRepo,
		DislikeRepository: dislikeRepo,
		AccountRepository: accRepo,
		ContentRepository: contentRepo,
		Validate:          validate,
	}
}

// method Like
func (l *LikeService) Like(ctx context.Context, request *dto.LikeRequest) (*dto.LikeResponse, error) {
	// validasi request
	err := l.Validate.Struct(*request)

	// jika gagal validasi required
	if err != nil {
		return nil, err
	}

	// cek data ID account
	_, err = l.AccountRepository.GetById(ctx, request.AccountId)
	if err != nil {
		return nil, errors.New("account id not found in database")
	}

	// cek data ID content
	_, err = l.ContentRepository.GetById(ctx, request.ContentId)
	if err != nil {
		return nil, errors.New("content id not found in database")
	}

	// cek apakah sudah dilike
	_, err = l.LikeRepository.GetByAccountIdAndContentId(ctx, request.AccountId, request.ContentId)
	if err == nil {
		return nil, errors.New("you have liked this content before")
	}

	// create object input Like
	likeRequest := entity.Like{
		AccountId: request.AccountId,
		ContentId: request.ContentId,
	}

	// call procedure Insert in repository
	result, err := l.LikeRepository.Insert(ctx, &likeRequest)
	if err != nil {
		return nil, err
	}

	// get total likes each content
	totalLikes, err := l.LikeRepository.GetTotalLikes(ctx, result.ContentId)
	if err != nil {
		l.LikeRepository.Delete(ctx, request.AccountId, request.ContentId)
		return nil, err
	}

	// update like content
	err = l.ContentRepository.UpdateLike(ctx, *totalLikes, result.ContentId)
	if err != nil {
		l.LikeRepository.Delete(ctx, request.AccountId, request.ContentId)
		return nil, err
	}

	// delete dislike jika ada
	l.DislikeRepository.Delete(ctx, request.AccountId, request.ContentId)

	// get total dislike by content_id
	totalDislikes, err := l.DislikeRepository.GetTotalDislikes(ctx, request.ContentId)
	if err != nil {
		return nil, err
	}

	// update dislike by content_id
	err = l.ContentRepository.UpdateDislike(ctx, *totalDislikes, request.ContentId)
	if err != nil {
		return nil, err
	}

	// get content by id
	content, err := l.ContentRepository.GetById(ctx, result.ContentId)
	if err != nil {
		l.LikeRepository.Delete(ctx, request.AccountId, request.ContentId)
		return nil, err
	}

	// get account by id
	account, err := l.AccountRepository.GetById(ctx, result.AccountId)
	if err != nil {
		l.LikeRepository.Delete(ctx, request.AccountId, request.ContentId)
		return nil, err
	}

	// mapping result to dto
	response := dto.LikeResponse{
		Account: &dto.AccountDetail{
			Id:       account.ID,
			Email:    account.Email,
			Username: account.Username,
		},
		Content: &dto.ContentDetail{
			ID:            content.ID,
			Title:         content.Title,
			PictureUrl:    content.PictureUrl,
			TextFill:      content.TextFill,
			Like:          content.Like,
			Dislike:       content.Dislike,
			AverageRating: content.AverageRating,
		},
	}

	// return
	return &response, nil
}

// method unlike
func (l *LikeService) Unlike(ctx context.Context, request *dto.LikeRequest) (*dto.LikeResponse, error) {
	// validate request
	err := l.Validate.Struct(*request)
	if err != nil {
		return nil, err
	}

	// cek account ID
	_, err = l.AccountRepository.GetById(ctx, request.AccountId)
	if err != nil {
		return nil, errors.New("account id not found in database")
	}

	// cek content ID
	_, err = l.ContentRepository.GetById(ctx, request.ContentId)
	if err != nil {
		return nil, errors.New("content id not found in database")
	}

	// cek like by account_id and content_id
	_, err = l.LikeRepository.GetByAccountIdAndContentId(ctx, request.AccountId, request.ContentId)
	if err != nil {
		return nil, errors.New("you not like this content")
	}

	// call procedure Delete in repository
	err = l.LikeRepository.Delete(ctx, request.AccountId, request.ContentId)
	if err != nil {
		return nil, errors.New("you not like this content")
	}

	// get total likes
	totalLikes, err := l.LikeRepository.GetTotalLikes(ctx, request.ContentId)
	if err != nil {
		return nil, err
	}

	// update likes in content
	err = l.ContentRepository.UpdateLike(ctx, *totalLikes, request.ContentId)
	if err != nil {
		return nil, err
	}

	// success unlike and update likes in content
	// get data account
	account, err := l.AccountRepository.GetById(ctx, request.AccountId)
	if err != nil {
		return nil, err
	}

	// get content data
	content, err := l.ContentRepository.GetById(ctx, request.ContentId)
	if err != nil {
		return nil, err
	}

	// create response
	response := dto.LikeResponse{
		Account: &dto.AccountDetail{
			Id:       account.ID,
			Email:    account.Email,
			Username: account.Username,
		},
		Content: &dto.ContentDetail{
			ID:            content.ID,
			Title:         content.Title,
			PictureUrl:    content.PictureUrl,
			TextFill:      content.TextFill,
			Like:          content.Like,
			Dislike:       content.Dislike,
			AverageRating: content.AverageRating,
			CreatedAt:     content.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	return &response, nil
}

// method get likes by username
func (l *LikeService) GetContentLikedByUsername(ctx context.Context, username string) ([]*dto.ContentDetail, error) {
	wg := &sync.WaitGroup{}

	// cek username apakah ada di database
	account, err := l.AccountRepository.GetByUsername(ctx, username)

	// jika account tidak ada di database
	if err != nil {
		return nil, err
	}

	// get likes by account.ID
	contents, err := l.LikeRepository.GetByAccountId(ctx, account.ID)

	// jika tidak ada / error
	if err != nil {
		return nil, err
	}

	// create variabel response
	response := []*dto.ContentDetail{}

	// looping each content
	for _, content := range contents {
		wg.Add(1)
		go func(content entity.Content) {
			defer wg.Done()

			response = append(response, &dto.ContentDetail{
				ID:            content.ID,
				Title:         content.Title,
				PictureUrl:    content.PictureUrl,
				TextFill:      content.TextFill,
				Like:          content.Like,
				Dislike:       content.Dislike,
				AverageRating: content.AverageRating,
				CreatedAt:     content.CreatedAt.Format("2006-01-02 15:04:05"),
				AccountId:     content.AccountId,
			})
		}(content)
	}

	// wait all goroutines done
	wg.Wait()

	// success get contents -> return
	return response, nil
}
