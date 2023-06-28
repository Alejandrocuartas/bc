package types

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Cafeteria struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Products []Product `json:"products"`
}

type Product struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Cafeteria_id int    `json:"cafeteria_id"`
}

func (u *Cafeteria) ToJson() ([]byte, error) {
	return json.Marshal(u)
}

func (p *Product) ToJson() ([]byte, error) {
	return json.Marshal(p)
}
