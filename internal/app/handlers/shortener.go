package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) shortenLink(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"mw": "alive",
	})
}
