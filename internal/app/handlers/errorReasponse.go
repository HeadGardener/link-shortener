package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Error struct {
	Message string `json:"error message"`
}

func newErrResponse(c echo.Context, statusCode int, errMsg string) error {
	logrus.Errorf(errMsg)
	return c.JSON(statusCode, Error{
		Message: errMsg,
	})
}
