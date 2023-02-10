package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

const (
	userCtx = "userID"
)

func (h *Handler) identifyUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if header == "" {
			return newErrResponse(c, http.StatusBadRequest, "empty auth header")
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			return newErrResponse(c, http.StatusBadRequest, "invalid auth header")
		}

		if len(headerParts[1]) == 0 {
			return newErrResponse(c, http.StatusBadRequest, "jwt token is empty")
		}

		userID, err := h.service.Authorization.ParseToken(headerParts[1])
		if err != nil {
			return newErrResponse(c, http.StatusUnauthorized, err.Error())
		}

		c.Set(userCtx, userID)
		err = next(c)
		if err != nil {
			return newErrResponse(c, http.StatusUnauthorized, err.Error())
		}

		return nil
	}
}
