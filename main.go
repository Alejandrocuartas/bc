package main

import (
	"github.com/Alejandrocuartas/bc/app"
	"github.com/Alejandrocuartas/bc/app/db"
	"github.com/joho/godotenv"
)

func main() {
	db.CreateDB()
	godotenv.Load()
	app.RunServer()
}
