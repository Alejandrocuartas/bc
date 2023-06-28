package services

import (
	"database/sql"
	"fmt"

	"github.com/Alejandrocuartas/bc/app/types"
)

func CreateCafeteria(db *sql.DB, cafeteria types.Cafeteria) error {
	_, err := db.Exec("INSERT INTO cafeterias (name) VALUES ($1)", cafeteria.Name)
	if err != nil {
		return err
	}
	return nil
}

func GetCafeterias(db *sql.DB) ([]types.Cafeteria, error) {
	rows, err := db.Query("SELECT id, name FROM cafeterias")
	if err != nil {
		return nil, fmt.Errorf("failed to query cafeterias: %v", err)
	}
	defer rows.Close()

	cafeterias := make([]types.Cafeteria, 0)
	for rows.Next() {
		var cafeteria types.Cafeteria
		err := rows.Scan(&cafeteria.ID, &cafeteria.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cafeteria row: %v", err)
		}
		cafeterias = append(cafeterias, cafeteria)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating cafeteria rows: %v", err)
	}

	return cafeterias, nil
}

func GetCafeteriasWithProducts(db *sql.DB) ([]types.Cafeteria, error) {
	rows, err := db.Query("SELECT id, name FROM cafeterias")
	if err != nil {
		return nil, fmt.Errorf("failed to query cafeterias: %v", err)
	}
	defer rows.Close()

	cafeterias := make([]types.Cafeteria, 0)
	for rows.Next() {
		var cafeteria types.Cafeteria
		err := rows.Scan(&cafeteria.ID, &cafeteria.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan cafeteria row: %v", err)
		}
		cafeterias = append(cafeterias, cafeteria)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating cafeteria rows: %v", err)
	}

	rowsP, err := db.Query("SELECT id, name, price, cafeteria_id FROM products")
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %v", err)
	}
	defer rowsP.Close()

	products := make([]types.Product, 0)
	for rowsP.Next() {
		var product types.Product
		err := rowsP.Scan(&product.ID, &product.Name, &product.Price, &product.Cafeteria_id)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}
		products = append(products, product)
	}
	if err = rowsP.Err(); err != nil {
		return nil, fmt.Errorf("error iterating product rows: %v", err)
	}

	for i := 0; i < len(cafeterias); i++ {
		fmt.Println(cafeterias[i])
		for j := 0; j < len(products); j++ {
			if products[j].Cafeteria_id == cafeterias[i].ID {
				cafeterias[i].Products = append(cafeterias[i].Products, products[j])
			}
		}
	}

	return cafeterias, nil
}
