package api

import (
	"telegram-bot/services"

	"github.com/labstack/echo"
)

func BotChat(e *echo.Group) {
	// Route / to handler function
	e.GET("/noti", services.SendMessage)
}
