package model

import (
	"Todo/db_server"
	"gorm.io/gorm"
)

var db *gorm.DB = db_server.MySqlDb

func init() {
	_ = db.AutoMigrate(&User{}, &TodoGroup{}, &TodoItem{})
}