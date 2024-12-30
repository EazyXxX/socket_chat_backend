package main

import (
	"socket_chat_backend/internal/repository"
	env "socket_chat_backend/pkg/environment"
	"socket_chat_backend/types"

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

	defer db.Close()

	// repos := repository.NewRepository(db)

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
