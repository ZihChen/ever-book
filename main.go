package main

import (
	"embed"
	"ever-book/app/global/settings"
	_ "ever-book/docs"
	"ever-book/internal/database"
	"ever-book/internal/server"
)

//go:embed env/*
var f embed.FS

func init() {
	// Loading env
	settings.Load(f)

	// initial database connection
	db := database.New()
	db.GetConnection()
	db.AutoMigrate()
}

func main() {
	server.Run()
}
