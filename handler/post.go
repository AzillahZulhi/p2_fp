package handler

import (
	"net/http"
	"p2-fp/config"
	"p2-fp/model"

	"github.com/labstack/echo"
)

func CreateCart(c echo.Context) error {
	NewCart := new(model.Cart)
	if err := c.Bind(NewCart); err != nil {
		ErrorMessage := model.ErrorMessage{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, ErrorMessage)
	}

	if NewCart.User_id == 0 || NewCart.Book_id == 0 {
		ErrorMessage := model.ErrorMessage{
			Message: "User_id and Book_id cannot be empty",
			Status:  http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, ErrorMessage)
	}

	if NewCart.Quantity <= 0 {
		ErrorMessage := model.ErrorMessage{
			Message: "Quantity cannot be null or negative",
			Status:  http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, ErrorMessage)
	}

	var book model.Book
	if err := config.DB.First(&book, NewCart.Book_id).Error; err != nil {
		ErrorMessage := model.ErrorMessage{
			Message: "Failed to get book information",
			Status:  http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, ErrorMessage)
	}

	total := float64(NewCart.Quantity) * book.Rent_cost

	NewCart.Total_price = total

	Result := config.DB.Create(NewCart)
	if Result.Error != nil {
		ErrorMessage := model.ErrorMessage{
			Message: "Failed to insert data",
			Status:  http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, ErrorMessage)
	}

	SuccessMessage := model.SuccessMessageCart{
		Message: "Succesfully added to cart",
		Status:  http.StatusCreated,
		Data:    NewCart,
	}
	return c.JSON(http.StatusCreated, SuccessMessage)

}
