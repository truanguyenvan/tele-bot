package router

import (
	"net/http"
	"telegram-bot/api"

	// "telegram-bot/api/middlewares"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	//create groups
	telegram := e.Group("/telegram")

	//set groupRoutes
	api.BotChat(telegram)
	return e
}
