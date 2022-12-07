package main

import (
	"ever-book/internal/database"
	"ever-book/internal/server"
)

func init() {
	db := database.New()
	db.GetConnection()
}

func main() {
	server.Run()
}
