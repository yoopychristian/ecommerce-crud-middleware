package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB, error) {

	dbUser := ReadEnv("dbUser", "your user")
	dbPass := ReadEnv("dbPass", "your password")
	dbHost := ReadEnv("dbHost", "youd localhost")
	dbPort := ReadEnv("dbPort", "your port")
	dbName := ReadEnv("dbName", "your database")

	loadData := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, _ := sql.Open("mysql", loadData)

	err := db.Ping()
	if err != nil {
		log.Print(err)
	}
	return db, nil
}
