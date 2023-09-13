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

// create object DislikeService
type DislikeService struct {
	LikeRepository    *repository.LikeRepository
	DislikeRepository *repository.DislikeRepository
	AccountRepository *repository.AccountRepository
	ContentRepository *repository.ContentRepository
	Validate          *validator.Validate
}

// function provider to create new object DislikeService
func NewDislikeService(likeRepo *repository.LikeRepository, dislikeRepo *repository.DislikeRepository, accRepo *repository.AccountRepository, contentRepo *repository.ContentRepository, validate *validator.Validate) *DislikeService {
	return &DislikeService{
		LikeRepository:    likeRepo,
		DislikeRepository: dislikeRepo,
		AccountRepository: accRepo,
		ContentRepository: contentRepo,
		Validate:          validate,
	}
}

// method Dislike
func (d *DislikeService) Dislike(ctx context.Context, request *dto.DislikeRequest) (*dto.DislikeResponse, error) {
	// validasi request body
	err := d.Validate.Struct(*request)

	// jika ada error ketika validasi -> gagal validasi required
	if err != nil {
		return nil, err
	}

	// cek account_id apakah ada di database
	_, err = d.AccountRepository.GetById(ctx, request.AccountId)
	if err != nil {
		return nil, errors.New("account id not found in database")
	}

	// cek content_id apakah ada di database
	_, err = d.ContentRepository.GetById(ctx, request.ContentId)
	if err != nil {
		return nil, errors.New("content id not found in database")
	}

	// cek data dislike apakah sudah ada
	_, err = d.DislikeRepository.GetByAccountIdAndContentId(ctx, request.AccountId, request.ContentId)
	if err == nil {
		// jika sudah ada
		return nil, errors.New("you have disliked this content")
	}

	// create dislike input
	dislikeInput := entity.Dislike{
		AccountId: request.AccountId,
		ContentId: request.ContentId,
	}

	// call procedure dislike in repository
	_, err = d.DislikeRepository.Insert(ctx, &dislikeInput)

	// jika ada kesalahan ketika insert dislike
	if err != nil {
		return nil, err
	}

	// get total dislike
	totalDislike, err := d.DislikeRepository.GetTotalDislikes(ctx, request.ContentId)
	if err != nil {
		return nil, err
	}

	// update dislike by content_id
	err = d.ContentRepository.UpdateDislike(ctx, *totalDislike, request.ContentId)
	if err != nil {
		return nil, err
	}

	// hapus data yang sama di tabel likes -> jika ada
	d.LikeRepository.Delete(ctx, request.AccountId, request.ContentId)

	// get total likes
	totalLikes, err := d.LikeRepository.GetTotalLikes(ctx, request.ContentId)
	if err != nil {
		return nil, err
	}

	// update like by content_id
	err = d.ContentRepository.UpdateLike(ctx, *totalLikes, request.ContentId)
	if err != nil {
		return nil, err
	}

	// get data account by account_id
	account, err := d.AccountRepository.GetById(ctx, request.AccountId)
	if err != nil {
		return nil, err
	}

	// get data content by content_id
	content, err := d.ContentRepository.GetById(ctx, request.ContentId)
	if err != nil {
		return nil, err
	}

	// mapping response
	response := dto.DislikeResponse{
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

	// success dislike
	return &response, nil
}

// method undislike
func (d *DislikeService) Undislike(ctx context.Context, request *dto.DislikeRequest) (*dto.DislikeResponse, error) {
	// validasi  request
	err := d.Validate.Struct(*request)

	// jika ada error ketika validasi -> gagal validasi
	if err != nil {
		return nil, err
	}

	// cek account_id apakah ada di database
	_, err = d.AccountRepository.GetById(ctx, request.AccountId)
	if err != nil {
		// account_id tidak ada di database
		return nil, errors.New("account_id not found in database")
	}

	// cek content_id apakah ada di database
	_, err = d.ContentRepository.GetById(ctx, request.ContentId)
	if err != nil {
		// content_id tidak ada di database
		return nil, errors.New("content_id not found in database")
	}

	// cek apakah data dislike sudah tidak ada di database
	_, err = d.DislikeRepository.GetByAccountIdAndContentId(ctx, request.AccountId, request.ContentId)
	if err != nil {
		// jika datanya sudah tidak ada
		return nil, errors.New("you dont dislike this content")
	}

	// --- lolos semua validasi ---
	// hapus dislike by account_id and content_id
	err = d.DislikeRepository.Delete(ctx, request.AccountId, request.ContentId)

	// jika ada error ketika proses delete
	if err != nil {
		return nil, err
	}

	// hitung totalDislike
	totalDislike, err := d.DislikeRepository.GetTotalDislikes(ctx, request.ContentId)
	if err != nil {
		return nil, err
	}

	// update dislike by content_id
	err = d.ContentRepository.UpdateDislike(ctx, *totalDislike, request.ContentId)
	if err != nil {
		return nil, err
	}

	// get data account
	account, err := d.AccountRepository.GetById(ctx, request.AccountId)
	if err != nil {
		return nil, err
	}

	// get data content
	content, err := d.ContentRepository.GetById(ctx, request.ContentId)
	if err != nil {
		return nil, err
	}

	// mapping to dto response
	response := dto.DislikeResponse{
		Account: &dto.AccountDetail{
			Id:        account.ID,
			Email:     account.Email,
			Username:  account.Username,
			CreatedAt: account.CreatedAt.Format("2006-01-02 15:04:0s5"),
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

	// return
	return &response, nil
}

// method get contens disliked by username
func (d *DislikeService) GetContentDislikedByUsername(ctx context.Context, username string) ([]*dto.ContentDetail, error) {
	wg := &sync.WaitGroup{}

	// cek username apakah ada di database
	account, err := d.AccountRepository.GetByUsername(ctx, username)

	// jika data tidak ditemukan
	if err != nil {
		return nil, err
	}

	// call procedure get content disliked by account_id
	contents, err := d.DislikeRepository.GetContentDislikedByAccountId(ctx, account.ID)

	// jika data tidak ditemukan / ada error
	if err != nil {
		return nil, err
	}

	// create response variabel
	response := []*dto.ContentDetail{}

	// looping each contents
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

	// waitt all goroutines done
	wg.Wait()

	// success get contents
	return response, nil
}
