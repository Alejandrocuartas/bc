package db

import "fmt"

func CreateDB() {
	db, err := ConnectDB()
	if err != nil {
		fmt.Printf("Error connecting db")
	}

	_, er := db.Exec(`CREATE TABLE IF NOT EXISTS cafeterias (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255)
	);`)
	if er != nil {
		fmt.Printf("error creating db 1: %v", er)
	}

	// Create products table
	_, e := db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		price INT,
		cafeteria_id INT,
		FOREIGN KEY (cafeteria_id) REFERENCES cafeterias (id)
	);`)
	if e != nil {
		fmt.Printf("error creating db 2: %v", e)
	}
}
