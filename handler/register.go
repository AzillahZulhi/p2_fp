package handler

import (
	"net/http"
	"p2-fp/config"
	"p2-fp/model"
	"regexp"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(c echo.Context) error {
	NewUser := new(model.User)
	if err := c.Bind(NewUser); err != nil {
		ErrorMessage := model.ErrorMessage{
			Message: "Invalid body input",
			Status:  http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, ErrorMessage)
	}

	if NewUser.Email == "" {
		ErrorMessage := model.ErrorMessage{
			Message: "Email cannot be empty",
			Status:  http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, ErrorMessage)
	}

	if NewUser.Password == "" {
		ErrorMessage := model.ErrorMessage{
			Message: "Password cannot be empty",
			Status:  http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, ErrorMessage)
	}

	if NewUser.Fullname == "" {
		ErrorMessage := model.ErrorMessage{
			Message: "Full name cannot be empty",
			Status:  http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, ErrorMessage)
	}

	if NewUser.Role != "admin" && NewUser.Role != "member" {
		ErrorMessage := model.ErrorMessage{
			Message: "Invalid role input",
			Status:  http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, ErrorMessage)
	}

	EmailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !EmailRegex.MatchString(NewUser.Email) {
		ErrorMessage := model.ErrorMessage{
			Message: "Invalid Email format",
			Status:  http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, ErrorMessage)
	}

	HashedPassword, err := bcrypt.GenerateFromPassword([]byte(NewUser.Password), bcrypt.DefaultCost)
	if err != nil {
		ErrorMessage := model.ErrorMessage{
			Message: "Failed to hash password",
			Status:  http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, ErrorMessage)
	}
	NewUser.Password = string(HashedPassword)

	var ExistUser model.User
	if err := config.DB.Where("email = ?", NewUser.Email).First(&ExistUser).Error; err == nil {
		ErrorMessage := model.ErrorMessage{
			Message: "Email already exists",
			Status:  http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, ErrorMessage)
	}

	Result := config.DB.Create(NewUser)
	if Result.Error != nil {
		ErrorMessage := model.ErrorMessage{
			Message: "Failed to insert data",
			Status:  http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, ErrorMessage)
	}

	SuccessMessage := model.SuccessMessageUser{
		Message: "Successfully Registered!",
		Status:  http.StatusCreated,
		Data:    NewUser,
	}
	return c.JSON(http.StatusCreated, SuccessMessage)
}
