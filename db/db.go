package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)


type DataBase struct {
	Client *sql.DB
}

func CreateDBIfNotExist() {
	conURL := fmt.Sprintf(
		"postgres://%s:%s@%s/?sslmode=disable",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_HOST"),
	)

	db, err := sql.Open("postgres", conURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Postgresql schema modification query doesn't support parameters
	checkIfDBExists := fmt.Sprintf(
		`SELECT FROM pg_database WHERE datname = '%s'`, 
		viper.GetString("DB_NAME"),
	)
	row := db.QueryRow(checkIfDBExists)

	err = row.Scan(); 
	if err == nil {
		fmt.Println("Database Already exists")
		return
	}

	// Database nof exist so Create the database
	if err == sql.ErrNoRows {
		fmt.Println("Database Does not exists")
		createDBSQL := fmt.Sprintf(`CREATE DATABASE %s`, viper.GetString("DB_NAME"))

		_, err = db.Exec(createDBSQL)
		if err != nil {
			panic(err)
		}
		fmt.Println("Database Created")
		return
	}

	panic(err)
}

func NewDatabase() (*DataBase, error) {
	conURL := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_NAME"),
	)

	db, err := sql.Open("postgres", conURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Database is connected")

	return &DataBase{
		Client: db,
	}, nil
}

