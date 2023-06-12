package db

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

type Client struct {
	*sql.DB
}

var client *Client

func GetDatabaseInstance() *Client {
	if client != nil {
		return client
	}
	db := createClient()
	client = &Client{db}
	return client
}

func createClient() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USERNAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}
