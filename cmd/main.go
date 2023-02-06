package main

import (
	"github.com/HeadGardener/link-shortener/internal/app/repository"
	app "github.com/HeadGardener/link-shortener/internal/pkg"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("error while initializing config: %s", err.Error())
	}

	application, err := app.New(repository.Config{
		Host:   viper.GetString("db.host"),
		Port:   viper.GetString("db.port"),
		DBName: viper.GetString("db.name"),
	})
	if err != nil {
		logrus.Fatalf("error while creating server: %s", err.Error())
	}

	logrus.Info("app started...")

	if err := application.Run(viper.GetString("port")); err != nil {
		logrus.Fatalf("error while starting server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
