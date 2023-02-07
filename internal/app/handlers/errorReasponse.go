package handlers

import "github.com/labstack/echo/v4"

type Error struct {
	Message string `json:"error message"`
}

func newErrResponse(c echo.Context, statusCode int, errMsg string) {
	if err := c.JSON(statusCode, Error{
		Message: errMsg,
	}); err != nil {
		c.Error(err)
	}
}
