package service

import (
	"cms/model/dto"
	"cms/model/entity"
	"cms/repository"
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/go-playground/validator/v10"
)

// create object struct ContentService
type ContentService struct {
	ContentRepository *repository.ContentRepository
	AccountRepository *repository.AccountRepository
	UserRepository    *repository.UserRepository
	Validate          *validator.Validate
}

// function provider to create new object ContentService
func NewContentService(contentRepo *repository.ContentRepository, accRepo *repository.AccountRepository, userRepo *repository.UserRepository, validate *validator.Validate) *ContentService {
	return &ContentService{
		ContentRepository: contentRepo,
		AccountRepository: accRepo,
		UserRepository:    userRepo,
		Validate:          validate,
	}
}

// method insert
func (c *ContentService) Insert(ctx context.Context, request *dto.CreateContentRequest) (*dto.CreateContentResponse, error) {
	// validasi request
	err := c.Validate.Struct(*request)

	// jika gagal validasi
	if err != nil {
		return nil, err
	}

	// create request object entity
	contentInput := entity.Content{
		Title:      request.Title,
		PictureUrl: request.PictureUrl,
		TextFill:   request.TextFill,
		AccountId:  request.AccountId,
	}

	// call procedure Insert in repository
	result, err := c.ContentRepository.Insert(ctx, &contentInput)

	// jika ada error ketika proses insert
	if err != nil {
		return nil, err
	}

	// get account by account_id
	account, err := c.AccountRepository.GetById(ctx, result.AccountId)
	if err != nil {
		c.ContentRepository.DeleteById(ctx, result.ID)
		return nil, err
	}

	// get user by id
	user, err := c.UserRepository.GetById(ctx, account.UserId)
	if err != nil {
		c.ContentRepository.DeleteById(ctx, result.ID)
		return nil, err
	}

	// create object response
	response := dto.CreateContentResponse{
		Content: &dto.ContentDetail{
			ID:            result.ID,
			Title:         result.Title,
			PictureUrl:    result.PictureUrl,
			TextFill:      result.TextFill,
			Like:          result.Like,
			Dislike:       result.Dislike,
			AverageRating: result.AverageRating,
			CreatedAt:     result.CreatedAt.Format("2006-01-02 15:04:05"),
		},
		CreatedBy: &dto.ContentCreatedBy{
			Email:     account.Email,
			Username:  account.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			FullName:  fmt.Sprintf("%v %v", user.FirstName, user.LastName),
		},
	}

	// sucess insert -> return
	return &response, nil
}

// method update content by Id
func (c *ContentService) Update(ctx context.Context, request *dto.UpdateContentRequest) (*dto.UpdateContentResponse, error) {
	// validate request
	err := c.Validate.Struct(*request)

	// jika ada error ketika validasi -> gagal validasi required
	if err != nil {
		return nil, err
	}

	// cek data by id apakah ada di database
	cekData, err := c.ContentRepository.GetById(ctx, request.Id)

	// jika ID data yang akan diupdate tidak ada di database
	if err != nil {
		return nil, err
	}

	// jika ada -> create input object to update
	contentUpdate := entity.Content{
		ID:         cekData.ID,
		Title:      request.Title,
		PictureUrl: request.PictureUrl,
		TextFill:   request.TextFill,
	}

	// call procedure Update in repository
	err = c.ContentRepository.UpdateById(ctx, &contentUpdate)

	// jika ada error ketika update
	if err != nil {
		return nil, err
	}

	// success update -> get data by ID
	content, err := c.ContentRepository.GetById(ctx, contentUpdate.ID)
	if err != nil {
		return nil, err
	}

	// mapping to dto
	response := dto.UpdateContentResponse{
		Id:            content.ID,
		Title:         content.Title,
		PictureUrl:    content.PictureUrl,
		TextFill:      content.TextFill,
		Like:          content.Like,
		Dislike:       content.Dislike,
		AverageRating: content.AverageRating,
		CreatedAt:     content.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	// return
	return &response, nil
}

// method get all contents
func (c *ContentService) GetAll(ctx context.Context) ([]*dto.ContentDetail, error) {
	wg := &sync.WaitGroup{}

	// call procedure get all in repository
	results, err := c.ContentRepository.GetAll(ctx)

	// jika ada kesalahan ketika get data
	if err != nil {
		return nil, err
	}

	// create variabel response
	response := []*dto.ContentDetail{}

	// looping each data -> append to response
	for _, data := range results {
		wg.Add(1)
		go func(item entity.Content) {
			defer wg.Done()

			// append to response
			response = append(response, &dto.ContentDetail{
				ID:            item.ID,
				Title:         item.Title,
				PictureUrl:    item.PictureUrl,
				TextFill:      item.TextFill,
				Like:          item.Like,
				Dislike:       item.Dislike,
				AverageRating: item.AverageRating,
				CreatedAt:     item.CreatedAt.Format("2006-01-02 15:04:05"),
				AccountId:     item.AccountId,
			})
		}(data)
	}

	// waitt all goroutines done
	wg.Wait()

	// success append to response -> return
	return response, nil
}

// method get content by id
func (c *ContentService) GetById(ctx context.Context, id int64) (*dto.ContentDetail, error) {
	// call procedure GetBYId in repository
	result, err := c.ContentRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	// success get data -> mapping to dto
	content := dto.ContentDetail{
		ID:            result.ID,
		Title:         result.Title,
		PictureUrl:    result.PictureUrl,
		TextFill:      result.TextFill,
		Like:          result.Like,
		Dislike:       result.Dislike,
		AverageRating: result.AverageRating,
		CreatedAt:     result.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	// return
	return &content, nil
}

// method get all contents by creator(username)
func (c *ContentService) GetByUsername(ctx context.Context, username string) ([]*dto.ContentDetail, error) {
	wg := &sync.WaitGroup{}
	mtx := &sync.RWMutex{}

	// cek username apakah ada di database
	_, err := c.AccountRepository.GetByUsername(ctx, username)

	// jika username tidak ada di database
	if err != nil {
		return nil, errors.New("username not found in database")
	}

	// call procedure GetByUsername in repository
	contents, err := c.ContentRepository.GetByUsername(ctx, username)

	// jika ada kesalahan ketika get data
	if err != nil {
		return nil, err
	}

	// create response object
	response := []*dto.ContentDetail{}

	for _, item := range contents {
		wg.Add(1)
		go func(data entity.Content) {
			defer mtx.Unlock()
			defer wg.Done()

			mtx.Lock()

			// mapping to object content
			content := dto.ContentDetail{
				ID:            data.ID,
				Title:         data.Title,
				PictureUrl:    data.PictureUrl,
				TextFill:      data.TextFill,
				Like:          data.Like,
				Dislike:       data.Dislike,
				AverageRating: data.AverageRating,
				CreatedAt:     data.CreatedAt.Format("2006-01-02 15:04:05"),
				AccountId:     data.AccountId,
			}

			// append to
			response = append(response, &content)
		}(item)
	}

	// wait all thread done
	wg.Wait()

	// return
	return response, nil
}
