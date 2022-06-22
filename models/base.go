package models

import (
	"net/http"

	"github.com/labstack/echo"
)

type BaseRespone struct {
	Status    bool        `json:"status"`
	ErrorCode interface{} `json:"error_code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// newBaseRespone - newBaseRespone
func NewBaseRespone(c echo.Context, status bool, code int, data interface{}, errorCode ...interface{}) error {

	if len(errorCode) > 0 {
		return c.JSON(code, BaseRespone{Status: status, ErrorCode: errorCode[0], Message: http.StatusText(code), Data: data})
	}
	return c.JSON(code, BaseRespone{Status: status, Message: http.StatusText(code), Data: data})

}
