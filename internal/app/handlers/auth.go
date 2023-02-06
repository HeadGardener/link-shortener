package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) signUp(c echo.Context) error {
	return c.String(http.StatusOK, "all right")
}

func (h *Handler) signIn(c echo.Context) error {
	return nil
}
