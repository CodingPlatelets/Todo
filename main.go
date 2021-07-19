package main

import (
	"Todo/db_server"
	"Todo/server"
	"log"

	"github.com/gin-gonic/gin"
)

var httpServer *gin.Engine

func main() {
	db, err := db_server.MySqlDb.DB()
	if err != nil {
		log.Print(err)
	}
	defer func() {
		db.Close()
	}()

	server.Run(httpServer)
}
