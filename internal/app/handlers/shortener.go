package handlers

import (
	"errors"
	"github.com/HeadGardener/link-shortener/internal/app/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) shortenLink(c echo.Context) error {
	var inputLink models.InputLink
	if err := c.Bind(&inputLink); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid data to create link")
		return err
	}

	userID, ok := c.Get(userCtx).(string)
	if !ok {
		newErrResponse(c, http.StatusBadRequest, "conversion mistake")
		return errors.New("conversion mistake")
	}

	link, err := h.service.Shortener.CreateLink(inputLink, userID)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, link)
}
