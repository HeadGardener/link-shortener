package handlers

import (
	"github.com/HeadGardener/link-shortener/internal/app/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) shortenLink(c echo.Context) error {
	var inputLink models.InputLink
	if err := c.Bind(&inputLink); err != nil {
		return newErrResponse(c, http.StatusBadRequest, "invalid data to create link")
	}

	userID, ok := c.Get(userCtx).(string)
	if !ok {
		return newErrResponse(c, http.StatusBadRequest, "conversion mistake")
	}

	link, err := h.service.Shortener.CreateLink(inputLink, userID)
	if err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, link)
}

func (h *Handler) redirect(c echo.Context) error {
	identifier := c.Param("id")

	if identifier == "" {
		return newErrResponse(c, http.StatusBadRequest, "empty link identifier")
	}

	url, err := h.service.Shortener.Redirect(identifier)
	if err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.Redirect(http.StatusPermanentRedirect, url)
}

func (h *Handler) getAllLinks(c echo.Context) error {
	userID, ok := c.Get(userCtx).(string)
	if !ok {
		return newErrResponse(c, http.StatusBadRequest, "conversion mistake")
	}

	links, err := h.service.Shortener.GetAll(userID)
	if err != nil {
		return newErrResponse(c, http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, links)
}
