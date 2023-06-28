package services

import (
	"database/sql"
	"fmt"

	"github.com/Alejandrocuartas/bc/app/types"
)

func CreateProduct(db *sql.DB, product types.Product) error {
	_, err := db.Exec("INSERT INTO products (name, price, cafeteria_id) VALUES ($1, $2, $3)", product.Name, product.Price, product.Cafeteria_id)
	if err != nil {
		return err
	}
	return nil
}

func GetProducts(db *sql.DB) ([]types.Product, error) {
	rows, err := db.Query("SELECT id, name, price, cafeteria_id FROM products")
	if err != nil {
		return nil, fmt.Errorf("failed to query products: %v", err)
	}
	defer rows.Close()

	products := make([]types.Product, 0)
	for rows.Next() {
		var product types.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Cafeteria_id)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product row: %v", err)
		}
		products = append(products, product)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating product rows: %v", err)
	}

	return products, nil
}
