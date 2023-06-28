package main

import (
	"github.com/Alejandrocuartas/bc/app"
	"github.com/Alejandrocuartas/bc/app/db"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.CreateDB()
	app.RunServer()
}
