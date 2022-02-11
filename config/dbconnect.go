package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() (*sql.DB, error) {
	url := os.Getenv("DB_URL")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s)/%s",
		password, url, name))

	if err != nil {
		return nil, err
	}
	return db, nil
}

func DatabaseSchemaInit() {
	url := os.Getenv("DB_URL")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	db, err := sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s)/",
		password, url))

	HandleErr(err)
	_, err = db.Query(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name))

	HandleErr(err)
	log.Printf("Database '%s' created. \n", os.Getenv("DB_NAME"))

	db, err = sql.Open("mysql", fmt.Sprintf("root:%s@tcp(%s)/%s",
		password, url, name))

	HandleErr(err)
	_, err = db.Query(fmt.Sprintf("CREATE TABLE IF NOT EXISTS `image_store_service_albums` (" +
		"`id` int NOT NULL AUTO_INCREMENT, " +
		"`album_id` varchar(100) NOT NULL, " +
		"`album_name` varchar(100) NOT NULL, " +
		"`is_active` tinyint NOT NULL DEFAULT '1', " +
		"`created_at` varchar(50) NOT NULL, " +
		"`modified_at` varchar(50) NOT NULL, " +
		"PRIMARY KEY (`id`)) " +
		"ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci"))

	HandleErr(err)
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

	HandleErr(err)
	log.Println("Table 'image_store_service_images' created.")

	log.Println("Database schema initialised successfully.")
}
