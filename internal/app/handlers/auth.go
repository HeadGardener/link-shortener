package handlers

import (
	"github.com/HeadGardener/link-shortener/internal/app/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) signUp(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid data for sign up user")
		return err
	}

	id, err := h.service.Authorization.CreateUser(user)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, err.Error())
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) signIn(c echo.Context) error {
	var userInput models.UserInput

	if err := c.Bind(&userInput); err != nil {
		newErrResponse(c, http.StatusBadRequest, "invalid data for sign in user")
		return err
	}

	token, err := h.service.Authorization.GenerateToken(userInput)
	if err != nil {
		newErrResponse(c, http.StatusInternalServerError, "mistake in authorization")
		return err
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"token": token,
	})
}
