package services

import (
	"net/http"
	"telegram-bot/handlers"
	"telegram-bot/models"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// SendMessage - SendMessage
func SendMessage(c echo.Context) error {
	query := new(models.SendMessage)
	if err := c.Bind(query); err != nil {
		log.Info("services -- ", err)
		return models.NewBaseRespone(c, false, http.StatusBadRequest, nil)
	}
	if err := c.Validate(query); err != nil {
		log.Info("services -- ", err)
		return models.NewBaseRespone(c, false, http.StatusBadRequest, nil)
	}
	resp, err := handlers.SendMessageHandler(query)
	if err != nil {
		return models.NewBaseRespone(c, false, http.StatusNotImplemented, nil)
	}
	return models.NewBaseRespone(c, true, http.StatusOK, resp)

}
