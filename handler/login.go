package handler

import (
	"errors"
	"net/http"
	"p2-fp/config"
	"p2-fp/helper"
	"p2-fp/model"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginHandler(c echo.Context) error {
	LoginCredential := new(model.LoginCredentials)
	if err := c.Bind(LoginCredential); err != nil {
		ErrorMessage := model.ErrorMessage{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, ErrorMessage)
	}

	if LoginCredential.Email == "" || LoginCredential.Password == "" {
		return c.JSON(http.StatusBadRequest, model.ErrorMessage{
			Message: "Email or Password cannot be empty",
			Status:  http.StatusBadRequest,
		})
	}

	var User model.User
	if err := config.DB.Where("email = ?", LoginCredential.Email).First(&User).Error; err == nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, model.ErrorMessage{
				Message: "User not found",
				Status:  http.StatusNotFound,
			})
		}
		return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Database error",
			Status:  http.StatusInternalServerError,
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(LoginCredential.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, model.ErrorMessage{
			Message: "Incorrect Password",
			Status:  http.StatusUnauthorized,
		})
	}

	token, err := helper.GenerateToken(User.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, model.ErrorMessage{
			Message: "Failed to generate token",
			Status:  http.StatusInternalServerError,
		})
	}

	c.Response().Header().Set("Authorization", "Bearer "+token)

	SuccessMessage := model.SuccessMessageUser{
		Message: "Successfully Login",
		Data:    &User,
		Token:   token,
		Status:  http.StatusOK,
	}
	return c.JSON(http.StatusOK, SuccessMessage)

}
