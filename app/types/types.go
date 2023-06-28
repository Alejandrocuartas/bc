package types

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type Cafeteria struct {
	Name string `json:"name"`
}

type Product struct {
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Cafeteria_id string `json:"cafeteria_id"`
}

func (u *Cafeteria) ToJson() ([]byte, error) {
	return json.Marshal(u)
}

func (p *Product) ToJson() ([]byte, error) {
	return json.Marshal(p)
}
