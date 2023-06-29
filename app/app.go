package app

import (
	"os"
	"github.com/Alejandrocuartas/bc/app/handler"
	s "github.com/Alejandrocuartas/bc/app/server"
)

func RunServer() {
	port := os.Getenv("PORT")
	server := s.NewServer(":" + port)
	server.Handle("POST", "/api/cafeteria", server.AddMidleware(handler.PostCafeteria, s.ShowMiddleware()))
	server.Handle("GET", "/api/cafeteria", handler.GetCafeteriasHandler)
	server.Handle("GET", "/api/product", handler.GetProductsHandler)
	server.Handle("GET", "/api/cafeteria/populate", handler.GetPopulatedCafeterias)
	server.Handle("POST", "/api/product", handler.PostProduct)
	server.Listen()
}
