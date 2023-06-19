package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/katerji/UserAuthKit/envs"
	"time"
)

type Client struct {
	*sql.DB
}

var instance *Client

func GetDbInstance() *Client {
	if instance == nil {
		instance, _ = getDbClient()
	}
	return instance
}

func getDbClient() (*Client, error) {
	dbHost := envs.GetInstance().GetDbHost()
	dbUser := envs.GetInstance().GetDbUser()
	dbPort := envs.GetInstance().GetDbPort()
	dbPass := envs.GetInstance().GetDbPassword()
	dbName := envs.GetInstance().GetDbName()

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return &Client{}, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return &Client{
		db,
	}, nil
}

func (client *Client) Fetch(query string, args ...any) []any {
	rows, err := client.Query(query, args...)
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

func (client *Client) Insert(query string, args ...any) (int, error) {
	rows, err := client.Prepare(query)
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

func (client *Client) FetchRows(query string, args ...any) ([]map[string]interface{}, error) {
	rows, err := client.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	result := []map[string]any{}
	for rows.Next() {
		row := make(map[string]interface{})
		values := make([]interface{}, len(columns))

		for i := range columns {
			values[i] = new(interface{})
		}

		err := rows.Scan(values...)
		if err != nil {
			return nil, err
		}

		for i, column := range columns {
			row[column] = *(values[i].(*interface{}))
		}

		result = append(result, row)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func (client *Client) Exec(query string, args ...any) bool {
	rows, err := client.Prepare(query)
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
