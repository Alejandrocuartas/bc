package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Alejandrocuartas/bc/app/db"
	"github.com/Alejandrocuartas/bc/app/responses"
	"github.com/Alejandrocuartas/bc/app/services"
	"github.com/Alejandrocuartas/bc/app/types"
)

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

func GetCafeteriasHandler(w http.ResponseWriter, r *http.Request) {
	database, e := db.ConnectDB()
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error connecting db: %v", e)
		return
	}
	cafes, error := services.GetCafeterias(database)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting cafeterias: %v", error)
		return
	}
	response := responses.CafeteriaResponse{
		Cafes: cafes,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to marshal JSON response: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var product types.Product
	err := decoder.Decode(&product)
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
	error := services.CreateProduct(database, product)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error creating product: %v", error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "Product created successfully.")
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	db, e := db.ConnectDB()
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error connecting db: %v", e)
		return
	}
	defer db.Close()
	products, error := services.GetProducts(db)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting cafeterias: %v", error)
		return
	}
	response := responses.ProductResponse{
		Products: products,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Failed to marshal JSON response: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func GetPopulatedCafeterias(w http.ResponseWriter, r *http.Request) {
	db, e := db.ConnectDB()
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error connecting db: %v", e)
		return
	}
	defer db.Close()
	cafeterias, err := services.GetCafeteriasWithProducts(db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting cafeterias: %v", err)
		return
	}

	jsonData, error := json.MarshalIndent(cafeterias, "", "  ")
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error json marshal: %v", error)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
