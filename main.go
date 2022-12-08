package main

import (
	"ever-book/internal/database"
	"ever-book/internal/server"
)

func init() {
	// initial database connection
	db := database.New()
	db.GetConnection()
	db.AutoMigrate()
}

func main() {
	server.Run()
}