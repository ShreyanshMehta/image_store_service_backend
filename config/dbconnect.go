package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func ConnectDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	driver := os.Getenv("DB_DRIVER")
	db, err := sql.Open(driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, host, port, name))

	if err != nil {
		return nil, err
	}
	return db, nil
}
