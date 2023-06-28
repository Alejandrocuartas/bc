package app

import (
	"github.com/Alejandrocuartas/bc/app/handler"
	s "github.com/Alejandrocuartas/bc/app/server"
)

func RunServer() {
	server := s.NewServer(":3000")
	server.Handle("GET", "/", handler.HandleRoot)
	server.Handle("POST", "/api", handler.HandleHome)
	server.Handle("POST", "/api/cafeteria/create", handler.PostCafeteria)
	//server.Handle("POST", "/api/cafeteria/create", server.AddMidleware(handler.PostCafeteria, s.CheckName()))
	server.Listen()
}
