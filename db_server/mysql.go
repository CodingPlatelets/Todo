package db_server

import (
	"Todo/config"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MySqlDb *gorm.DB
var MySqlError error

func init() {
	dbConfig := config.GetDbConfig()
	// set db_server dsn
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
		dbConfig["username"],
		dbConfig["password"],
		dbConfig["hostname"],
		dbConfig["port"],
		dbConfig["database"],
		dbConfig["charset"],
		dbConfig["parseTime"],
	)

	MySqlDb, MySqlError = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})

	db, err := MySqlDb.DB()
	if err != nil {
		log.Print(err.Error())
		return
	}
	db.SetMaxIdleConns(dbConfig["maxIdleConns"].(int))
	db.SetMaxOpenConns(dbConfig["maxOpenConns"].(int))

	if MySqlError != nil {
		panic("database open error! " + MySqlError.Error())
	}

}
