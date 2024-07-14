package main

import (
	"database/sql"
	"log"

	"github.com/chavocito/ecom/cmd/api"
	"github.com/chavocito/ecom/config"
	"github.com/chavocito/ecom/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.InitDB(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		AllowAllFiles:        true,
	})

	if err != nil {
		log.Fatalf("Failed to initialize db with err: %v", err)
	}

	initStorage(db)
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start the server...")
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database connection successfully established...")
}
