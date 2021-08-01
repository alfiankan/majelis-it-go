package helper

import (
	"embed"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitMysqlDB(t *testing.T) {
	db := InitMysqlDB()

	err := db.Ping()
	log.Println(err)
	assert.NoError(t, err)
}

//go:embed db.sql
var sqld embed.FS
func TestMigrate(t *testing.T) {

	sql, err := sqld.ReadFile("db.sql")
	log.Println(err)
	log.Println(string(sql))
	db := InitMysqlDB()
	res, err := db.Exec(string(sql))

	log.Println(res, err)
	
}