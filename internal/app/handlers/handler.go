package handlers

import (
	"github.com/HeadGardener/link-shortener/internal/app/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *echo.Echo {
	router := echo.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	shortener := router.Group("shortener", h.identifyUser)
	{
		link := shortener.Group("/link")
		{
			link.POST("/", h.shortenLink)
		}
	}

	return router
}
