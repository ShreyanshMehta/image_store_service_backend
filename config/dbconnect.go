package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
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

func DatabaseSchemaInit() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	driver := os.Getenv("DB_DRIVER")
	db, err := sql.Open(driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		user, password, host, port))
	if err != nil {
		log.Fatal(err)
		return
	}
	_, err = db.Query(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name))
	db, err = sql.Open(driver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		user, password, host, port, name))

	log.Printf("Database '%s' created. \n", os.Getenv("DB_NAME"))

	_, err = db.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `image_store_service_albums` (" +
		"`id` int NOT NULL AUTO_INCREMENT, " +
		"`album_id` varchar(100) NOT NULL, " +
		"`album_name` varchar(100) NOT NULL, " +
		"`is_active` tinyint NOT NULL DEFAULT '1', " +
		"`created_at` varchar(50) NOT NULL, " +
		"`modified_at` varchar(50) NOT NULL, " +
		"PRIMARY KEY (`id`)) " +
		"ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci"))
	log.Println("Table 'image_store_service_albums' created.")

	_, err = db.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `image_store_service_images` (" +
		"`id` int NOT NULL AUTO_INCREMENT, " +
		"`image_id` varchar(100) NOT NULL, " +
		"`image_name` varchar(100) NOT NULL, " +
		"`album_id` varchar(100) NOT NULL, " +
		"`created_at` varchar(50) NOT NULL, " +
		"`is_active` tinyint NOT NULL DEFAULT '1', " +
		"PRIMARY KEY (`id`)) " +
		"ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci"))
	log.Println("Table 'image_store_service_images' created.")

	log.Println("Database schema initialised successfully.")
}
