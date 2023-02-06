package handlers

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) identifyUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
