package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)


type DataBase struct {
	Client *sql.DB
}


func NewDatabase() (*DataBase, error) {
	conURL := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable", 
		os.Getenv("DBUSER"), 
		os.Getenv("DBPASS"),
		os.Getenv("DBNAME"),
	)

	db, err := sql.Open("postgres", conURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database is connected")

	return &DataBase{
		Client: db,
	}, nil
}




