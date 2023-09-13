package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"strings"
	"training/app/model/dto"
	"training/app/service"
)

type BookHandler struct {
	BookService service.IBookService
}

func NewBookHander(bookService service.IBookService) *BookHandler {
	return &BookHandler{
		BookService: bookService,
	}
}

// method to get book by Id
func (b *BookHandler) GetBookById(c *gin.Context) {
	// get id from param
	id := c.Param("id")
	title := c.Param("title")

	bookId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status_code": http.StatusBadRequest,
			"message":     "cant convert id to int",
		})
		return
	}

	book := &dto.BookResponse{
		Id:    bookId,
		Title: title,
	}

	c.JSON(http.StatusOK, book)
}

// method to get book detail
func (b *BookHandler) GetBookDetail(c *gin.Context) {
	// get data from query
	title := c.Query("title")
	genre := c.Query("genre")

	book := &dto.BookResponse{
		Id:    1,
		Title: title,
		Genre: genre,
	}

	c.JSON(http.StatusOK, book)
}

// method to insert book
func (b *BookHandler) InsertBook(c *gin.Context) {
	var request dto.InsertBookRequest2

	// decode data from request_body
	err := c.ShouldBindJSON(&request)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			errorMessage := []string{}
			for _, msg := range err.(validator.ValidationErrors) {
				errMessage := fmt.Sprintf("error on field: %v, condition: %v", msg.Field(), msg.ActualTag())
				errorMessage = append(errorMessage, errMessage)
			}

			response := &dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    strings.Join(errorMessage, ", "),
			}

			c.JSON(http.StatusBadRequest, response)
			return
		case *json.UnmarshalTypeError:
			response := &dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    err.Error(),
			}

			c.JSON(http.StatusBadRequest, response)
			return
		default:
			response := &dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    err.Error(),
			}

			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	// call procedure insert in service
	book, err := b.BookService.InsertBook(&request)
	//	log.Println("error insert book:", err.Error())
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			c.JSON(http.StatusNotFound, &dto.ApiMessage{
				StatusCode: http.StatusNotFound,
				Status:     "not found",
				Message:    err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, &dto.ApiMessage{
			StatusCode: http.StatusInternalServerError,
			Status:     "internal server error",
			Message:    err.Error(),
		})
		return
	}

	response := &dto.ApiMessage{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success insert",
		Data:       &book,
	}
	c.JSON(http.StatusOK, response)
}

// method to get all books
func (b *BookHandler) GetAllBooks(c *gin.Context) {
	// call procedure in service
	books, err := b.BookService.GetAllBooks()
	if err != nil {
		responseError := &dto.ApiMessage{
			StatusCode: http.StatusNotFound,
			Status:     "not found",
			Message:    err.Error(),
		}
		c.JSON(http.StatusNotFound, &responseError)
		return
	}

	// success get all books
	response := &dto.ApiMessage{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success get all books",
		Data:       &books,
	}

	c.JSON(http.StatusOK, &response)
}

// method get book by id
func (b *BookHandler) GetBookById2(c *gin.Context) {
	// get data from params
	id := c.Param("id")

	bookId, err := strconv.Atoi(id)
	if err != nil {
		response := &dto.ApiMessage{
			StatusCode: http.StatusBadRequest,
			Status:     "bad request",
			Message:    "cant convert id to int",
		}
		c.JSON(http.StatusBadRequest, &response)
		return
	}

	// call procedure get by id in service
	book, err := b.BookService.GetBookById(bookId)
	if err != nil {
		// not found
		response := &dto.ApiMessage{
			StatusCode: http.StatusNotFound,
			Status:     "not found",
			Message:    err.Error(),
		}

		c.JSON(http.StatusNotFound, &response)
		return
	}

	// success get book by id
	response := &dto.ApiMessage{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success get data book by id",
		Data:       book,
	}
	c.JSON(http.StatusOK, &response)
}

// method delete book by id
func (b *BookHandler) DeleteBookById(c *gin.Context) {
	// get id from param
	id := c.Param("id")

	bookId, err := strconv.Atoi(id)
	if err != nil {
		response := &dto.ApiMessage{
			StatusCode: http.StatusBadRequest,
			Status:     "bad request",
			Message:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, &response)
		return
	}

	err = b.BookService.DeleteBookById(bookId)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			// not found
			response := &dto.ApiMessage{
				StatusCode: http.StatusNotFound,
				Status:     "not found",
				Message:    err.Error(),
			}
			c.JSON(http.StatusNotFound, &response)
			return
		}

		response := &dto.ApiMessage{
			StatusCode: http.StatusBadRequest,
			Status:     "bad request",
			Message:    err.Error(),
		}
		c.JSON(http.StatusBadRequest, &response)
		return
	}

	// success delete
	response := &dto.ApiMessage{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success delete book",
	}
	c.JSON(http.StatusOK, &response)
}

// method to update data book by id
func (b *BookHandler) UpdateById(c *gin.Context) {
	// get data from request_body
	var request dto.UpdateBookRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			var errMessage []string
			for _, value := range err.(validator.ValidationErrors) {
				errMessage = append(errMessage, fmt.Sprintf("error field: %v, condition: %v", value.Field(), value.ActualTag()))
			}
			response := &dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    strings.Join(errMessage, ", "),
			}

			c.JSON(http.StatusBadRequest, response)
			return
		default:
			response := &dto.ApiMessage{
				StatusCode: http.StatusBadRequest,
				Status:     "bad request",
				Message:    err.Error(),
			}

			c.JSON(http.StatusBadRequest, response)
			return
		}
	}

	// call procedure update in service
	newBook, err := b.BookService.UpdateBookById(&request)
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "not found") {
			// not found
			response := &dto.ApiMessage{
				StatusCode: http.StatusNotFound,
				Status:     "not found",
				Message:    err.Error(),
			}

			c.JSON(http.StatusNotFound, response)
			return
		}

		response := &dto.ApiMessage{
			StatusCode: http.StatusInternalServerError,
			Status:     "internal server error",
			Message:    err.Error(),
		}

		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// success update
	response := &dto.ApiMessage{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success update data book",
		Data:       newBook,
	}

	c.JSON(http.StatusOK, response)
}

// method to Get data books by user_id
func (b *BookHandler) GetByUserId(c *gin.Context) {
	// get id from param
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ApiMessage{
			StatusCode: http.StatusBadRequest,
			Status:     "bad request",
			Message:    "cant convert user_id to int",
		})
		return
	}

	// call procedure GetByUserId in service layer
	books, err := b.BookService.GetByUserId(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, &dto.ApiMessage{
			StatusCode: http.StatusNotFound,
			Status:     "not found",
			Message:    err.Error(),
		})
		return
	}

	// success get data by user_id
	c.JSON(http.StatusOK, &dto.ApiMessage{
		StatusCode: http.StatusOK,
		Status:     "ok",
		Message:    "success get data book by user_id",
		Data:       books,
	})
}
