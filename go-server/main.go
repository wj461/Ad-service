package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName(".env")
	viper.AddConfigPath("../")
	viper.SetConfigType("dotenv")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Fatal("Error reading config file, ", err)
	}
}

func main() {
	// logrus.SetReportCaller(true)

	logrus.Info("HTTP server is starting...")

	// restfulHost := viper.GetString("RESTFUL_HOST")
	// restfulPort := viper.GetString("RESTFUL_PORT")
	dbDatabase := viper.GetString("DB_DATABASE")
	dbUser := viper.GetString("POSTGRES_USER")
	dbPassword := viper.GetString("POSTGRES_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	// if go not run in docker, host will be localhost
	if dbHost == "" {
		dbHost = "localhost"
	}

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbDatabase),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	defer db.Close()
}
