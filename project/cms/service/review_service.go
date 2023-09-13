package service

import (
	"cms/model/dto"
	"cms/model/entity"
	"cms/repository"
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
)

// create object ReviewService
type ReviewService struct {
	ReviewRepository  *repository.ReviewRepository
	AccountRepository *repository.AccountRepository
	ContentRepository *repository.ContentRepository
	Validate          *validator.Validate
}

// function provider to create new object ReviewService
func NewReviewService(reviewRepo *repository.ReviewRepository, accRepo *repository.AccountRepository, contentRepo *repository.ContentRepository, validate *validator.Validate) *ReviewService {
	return &ReviewService{
		ReviewRepository:  reviewRepo,
		AccountRepository: accRepo,
		ContentRepository: contentRepo,
		Validate:          validate,
	}
}

// method Insert
func (r *ReviewService) Insert(ctx context.Context, request *dto.CreateReviewRequest) (*dto.CreateReviewResponse, error) {
	// validasi required
	err := r.Validate.Struct(*request)

	// jika ada error kesalahan validasi
	if err != nil {
		return nil, err
	}

	// cek account_id apakah ada di database
	_, err = r.AccountRepository.GetById(ctx, request.AccountId)
	if err != nil {
		return nil, errors.New("record with account_id not found in database")
	}

	// cek content_id apakah ada di database
	_, err = r.ContentRepository.GetById(ctx, request.ContentId)
	if err != nil {
		return nil, errors.New("record with content_id not found in database")
	}

	// create review entity
	reviewInput := entity.Review{
		Rating:    request.Rating,
		Comment:   request.Comment,
		AccountId: request.AccountId,
		ContentId: request.ContentId,
	}

	// call procedure insert in repository
	result, err := r.ReviewRepository.Insert(ctx, &reviewInput)

	// jika ada kesalahan ketika input
	if err != nil {
		return nil, err
	}

	// get sum rating by content_id
	sumRating, err := r.ReviewRepository.SumRatingByContentId(ctx, request.ContentId)
	if err != nil {
		// delete review
		r.ReviewRepository.Delete(ctx, result.ID)

		// return error
		return nil, err
	}

	// get count reviews by content_id
	countRating, err := r.ReviewRepository.CountRatingByContentId(ctx, request.ContentId)
	if err != nil {
		// delete reviews from database
		r.ReviewRepository.Delete(ctx, result.ID)

		// return error
		return nil, err
	}

	// hitung average_rating nya
	averageRating := float64(*sumRating) / float64(*countRating)

	// update rating
	err = r.ContentRepository.UpdateRating(ctx, averageRating, request.ContentId)
	if err != nil {
		// delete reviews from database
		r.ReviewRepository.Delete(ctx, result.ID)

		// return error
		return nil, err
	}

	// get data content
	content, err := r.ContentRepository.GetById(ctx, request.ContentId)
	if err != nil {
		// hapus reviews dari database
		r.ReviewRepository.Delete(ctx, result.ID)

		// return error
		return nil, err
	}

	// get data account
	account, err := r.AccountRepository.GetById(ctx, request.AccountId)
	if err != nil {
		// hapus reviews dari database
		r.ReviewRepository.Delete(ctx, result.ID)

		// return error
		return nil, err
	}

	// get data reviews
	// create response -> mapping to response dto
	response := dto.CreateReviewResponse{
		Id: result.ID,
		Review: &dto.ReviewDetail{
			Id:        result.ID,
			Rating:    result.Rating,
			Comment:   result.Comment,
			CreatedAt: result.CreatedAt.Format("2006-01-02 15:04:05"),
			AccountId: result.AccountId,
			ContentId: result.ContentId,
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
			AccountId:     content.AccountId,
		},
		Account: &dto.AccountDetail{
			Id:        account.ID,
			Email:     account.Email,
			Username:  account.Username,
			CreatedAt: account.CreatedAt.Format("2006-01-02 15:04:05"),
		},
	}

	// success insert reviews -> return
	return &response, nil
}

// method delete review
func (r *ReviewService) Delete(ctx context.Context, request *dto.DeleteReviewRequest) error {
	// validasi request required
	err := r.Validate.Struct(*request)

	// jika ada kesalahan ketika validasi required
	if err != nil {
		return err
	}

	// check apakah review_id ada di database
	_, err = r.ReviewRepository.GetById(ctx, request.ReviewId)
	if err != nil {
		return errors.New("record with review_id not found in database")
	}

	// cek apakah account_id ada di database
	_, err = r.AccountRepository.GetById(ctx, request.AccountId)
	if err != nil {
		return errors.New("record with account_id not found in database")
	}

	// cek apakah content_id ada di database
	_, err = r.ContentRepository.GetById(ctx, request.ContentId)
	if err != nil {
		return errors.New("record with content_id not found in database")
	}

	// cek apakah data pada request ada di database
	review, err := r.ReviewRepository.CheckReview(ctx, request)
	if err != nil {
		return err
	}

	// delete data
	err = r.ReviewRepository.Delete(ctx, review.ID)
	if err != nil {
		return err
	}

	// success delete review
	return nil
}
