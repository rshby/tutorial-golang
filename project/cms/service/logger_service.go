package service

import (
	"cms/model/dto"
	"cms/model/entity"
	"cms/repository"
	"context"
)

// create object service
type LoggerService struct {
	LoggerRepository *repository.LoggerRepository
}

// function provider
func NewLoggerService(loggerRepo *repository.LoggerRepository) *LoggerService {
	return &LoggerService{
		LoggerRepository: loggerRepo,
	}
}

// method Insert
func (l *LoggerService) Insert(ctx context.Context, request *dto.LoggerInsertRequest) (*entity.Logger, error) {
	// create input logger
	loggerInput := entity.Logger{
		IpAddress:  request.IpAddress,
		UrlPath:    request.UrlPath,
		Method:     request.Method,
		StatusCode: request.StatusCode,
		Status:     request.Status,
		Duration:   request.Duration,
	}

	// call procedure Insert in repository
	resultInsert, err := l.LoggerRepository.Insert(ctx, &loggerInput)

	// jika ada error ketika insert logger
	if err != nil {
		return nil, err
	}

	// success insert -> get data by id
	response, err := l.LoggerRepository.GetById(ctx, resultInsert.ID)
	if err != nil {
		return nil, err
	}

	// success get data by id
	return response, nil
}
