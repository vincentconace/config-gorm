package main

import (
	"github.com/vincentcoance/config-gorm/db"
)

func main() {
	dbConfig := db.Configure(".", ".env")
	db.DB = dbConfig.InitMysqlDB()
}
