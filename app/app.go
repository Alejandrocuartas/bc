package app

import (
	"github.com/Alejandrocuartas/bc/app/handler"
	s "github.com/Alejandrocuartas/bc/app/server"
)

func RunServer() {
	server := s.NewServer(":3000")
	server.Handle("GET", "/", handler.HandleRoot)
	server.Handle("POST", "/api", server.AddMidleware(handler.HandleHome, s.CheckAuth(), s.Logging()))
	server.Handle("POST", "/create", handler.UserPostRequest)
	server.Listen()
}
