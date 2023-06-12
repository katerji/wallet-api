package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"sync"
	"time"
)

type Client struct {
	*sql.DB
}

var instance *Client
var once sync.Once

func GetDbInstance() *Client {
	once.Do(func() {
		instance, _ = getDbClient()
	})
	return instance
}

func getDbClient() (*Client, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USERNAME")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_DATABASE")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println(dataSourceName)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		fmt.Println(err)
		return &Client{}, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &Client{
		db,
	}, nil
}

func (closerDb *Client) Fetch(query string, args ...any) []interface{} {
	rows, err := closerDb.Query(query, args...)
	if err != nil {
		fmt.Println(err.Error())
	}
	results := make([]interface{}, 0)
	for rows.Next() {
		var row interface{}

		err = rows.Scan(&row)
		results = append(results, row)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}
	return results
}

func (closerDb *Client) Insert(query string, args ...any) (int, error) {
	rows, err := closerDb.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	exec, err := rows.Exec(args...)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	insertId, err := exec.LastInsertId()
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	return int(insertId), nil
}

func (closerDb *Client) Exec(query string, args ...any) bool {
	rows, err := closerDb.Prepare(query)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	_, err = rows.Exec(args...)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
