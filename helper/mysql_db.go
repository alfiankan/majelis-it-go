package helper

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)



func InitMysqlDB() *sql.DB {
	
	db, err :=  sql.Open("mysql", "root:473550@tcp(localhost:3306)/mahasiswa")
	if err != nil {
		//error handling
		log.Println(err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Hour)
	db.SetConnMaxIdleTime(time.Minute * 30)

	return db
}