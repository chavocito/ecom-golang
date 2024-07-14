package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func InitDB(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Printf("Error opening mysql database %s", err.Error())
		return nil, err
	}
	return db, nil
}
