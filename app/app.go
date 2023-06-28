package app

import (
	"github.com/Alejandrocuartas/bc/app/handler"
	s "github.com/Alejandrocuartas/bc/app/server"
)

func RunServer() {
	server := s.NewServer(":3000")
	server.Handle("POST", "/api/cafeteria", server.AddMidleware(handler.PostCafeteria, s.ShowMiddleware()))
	server.Handle("GET", "/api/cafeteria", handler.GetCafeteriasHandler)
	server.Handle("GET", "/api/product", handler.GetProductsHandler)
	server.Handle("GET", "/api/cafeteria/populate", handler.GetPopulatedCafeterias)
	server.Handle("POST", "/api/product", handler.PostProduct)
	server.Listen()
}
