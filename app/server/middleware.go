package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Alejandrocuartas/bc/app/types"
)

func CheckAuth() types.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking Authentication")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Logging() types.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		}
	}
}
