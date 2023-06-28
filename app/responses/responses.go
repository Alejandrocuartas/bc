package responses

import "github.com/Alejandrocuartas/bc/app/types"

type CafeteriaResponse struct {
	Cafes []types.Cafeteria `json:"cafes"`
}

type ProductResponse struct {
	Products []types.Product `json:"products"`
}
