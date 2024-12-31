package main

import (
	"context"
	"os"
	"os/signal"
	"socket_chat_backend/internal/delivery/httpServer"
	"socket_chat_backend/internal/handler"
	"socket_chat_backend/internal/repository"
	"socket_chat_backend/internal/service"
	env "socket_chat_backend/pkg/environment"
	"socket_chat_backend/types"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Setting a log JSON formatter for future use
	logrus.SetFormatter(new(logrus.TextFormatter))

	// Initialising viper config
	if err := initConfig(); err != nil {
		logrus.WithError(err).Fatal("Error initializing a config")
	}

	// Loading .env variables
	if err := godotenv.Load(); err != nil {
		logrus.WithError(err).Fatal("Error loading env variables")
	}

	// Establishing database connection
	db, err := repository.NewPostgresDB(types.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: env.GetEnv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to initialize db")
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// Server initialising
	srv := new(httpServer.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.WithError(err).Fatal("Error occured while running http server")
		}
	}()

	logrus.Info("socket_chat is started")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Info("socket_chat is shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.WithError(err).Fatal("Error occured on server shutting down")
	}

	if err := db.Close(); err != nil {
		logrus.WithError(err).Fatal("Error occured on database connection close")
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
