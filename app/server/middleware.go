package server

import (
	"fmt"
	"net/http"

	"github.com/Alejandrocuartas/bc/app/types"
)

func ShowMiddleware() types.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Middleware working")
			f(w, r)
		}
	}
}
