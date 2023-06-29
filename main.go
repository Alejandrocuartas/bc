package main

import (
	"github.com/Alejandrocuartas/bc/app"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app.RunServer()
}
