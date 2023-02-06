package app

import (
	"github.com/HeadGardener/link-shortener/internal/app/handlers"
	"github.com/HeadGardener/link-shortener/internal/app/repository"
	"github.com/HeadGardener/link-shortener/internal/app/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	repos    *repository.Repository
	services *service.Service
	handler  *handlers.Handler
	echo     *echo.Echo
}

func New(conf repository.Config) (*App, error) {
	app := &App{}

	db, err := repository.NewMongoDB(conf)

	if err != nil {
		return nil, err
	}

	app.repos = repository.NewRepository(db)
	app.services = service.NewService(app.repos)
	app.handler = handlers.NewHandler(app.services)

	app.echo = app.handler.InitRoutes()
	app.echo.Use(middleware.Recover())

	return app, nil
}

func (app *App) Run(port string) error {

	return app.echo.Start(":" + port)
}
