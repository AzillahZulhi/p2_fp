package handler

import (
	"errors"
	"net/http"
	"p2-fp/config"
	"p2-fp/model"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func GetBooks(c echo.Context) error {
	var Books []model.Book

	Result := config.DB.Find(&Books)
	if errors.Is(Result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, model.ErrorMessage{
			Message: "Books not found",
			Status:  http.StatusNotFound,
		})
	}

	if Result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Failed to fetch books data",
			Status:  http.StatusInternalServerError,
		})
	}

	SuccessMessage := model.SuccessMessageBook{
		Message: "Show Posts datas",
		Datas:   Books,
		Status:  http.StatusOK,
	}
	return c.JSON(http.StatusOK, SuccessMessage)
}

func GetBook(c echo.Context) error {
	var Book model.Book
	id := c.Param("id")

	Result := config.DB.First(&Book, id)
	if errors.Is(Result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, model.ErrorMessage{
			Message: "Book not found",
			Status:  http.StatusNotFound,
		})
	}

	if Result.Error != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Failed to fetch books data",
			Status:  http.StatusInternalServerError,
		})
	}

	SuccessMessage := model.SuccessMessageBook{
		Message: "Show Posts datas",
		Data:    &Book,
		Status:  http.StatusOK,
	}
	return c.JSON(http.StatusOK, SuccessMessage)
}
