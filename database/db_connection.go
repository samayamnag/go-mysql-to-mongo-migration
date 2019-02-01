package database

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnect(dbName string) (db *sql.DB)  {
	var dbUser, dbPwd, dbHost string
	var dbPort int
	dbUser = "root"
	dbPwd = ""
	dbHost = "localhost"
	dbPort = 3306

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", dbUser, dbPwd, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", connStr)

	if err != nil {
		panic(err.Error())
	}

	return db
}