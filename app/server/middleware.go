package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Alejandrocuartas/bc/app/types"
)

func CheckName() types.Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			bodyBuffer, error1 := io.ReadAll(r.Body)
			if error1 != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "error on middleware buffer: %v", error1)
				return
			}
			bufferedBody := io.NopCloser(bytes.NewBuffer(bodyBuffer))
			clone := r.Clone(r.Context())
			clone.Body = bufferedBody
			var cafeteria types.Cafeteria
			decoder := json.NewDecoder(clone.Body)
			err := decoder.Decode(&cafeteria)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "error on middleware: %v", err)
				return
			}
			if cafeteria.Name == "" {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "error: Name could not be empty.")
				return
			}
			r.Body = bufferedBody
			f(w, r)
		}
	}
}
