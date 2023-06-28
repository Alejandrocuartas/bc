package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Alejandrocuartas/bc/app/db"
	"github.com/Alejandrocuartas/bc/app/services"
	"github.com/Alejandrocuartas/bc/app/types"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello to the coffee API.")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the API Endpoint")
}

func PostCafeteria(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var cafeteria types.Cafeteria
	err := decoder.Decode(&cafeteria)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	database, e := db.ConnectDB()
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error connecting db: %v", e)
		return
	}
	error := services.CreateCafeteria(database, cafeteria)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error creating cafeteria: %v", error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Cafeteria created successfully.")
}
