package services

import (
	"database/sql"

	"github.com/Alejandrocuartas/bc/app/types"
)

func CreateCafeteria(db *sql.DB, cafeteria types.Cafeteria) error {
	_, err := db.Exec("INSERT INTO cafeterias (name) VALUES ($1)", cafeteria.Name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCafeterias(db *sql.DB, cafeteria types.Cafeteria) error {
	_, err := db.Exec("INSERT INTO cafeterias (name) VALUES ($1)", cafeteria.Name)
	if err != nil {
		return err
	}
	return nil
}
