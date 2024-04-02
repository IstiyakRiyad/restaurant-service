package db

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
)


func (db *DataBase) MigrateUpDB() error {
	driver, err := postgres.WithInstance(db.Client, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		viper.GetString("MIGRATION_FILE_URL"),
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	// Migrating db
	if err := m.Up(); err != nil {
		if err.Error() == migrate.ErrNoChange.Error() {
			fmt.Println("The database is already in sync with the migrations.")
			return nil
		} else {
			if err := m.Steps(-1); err != nil {
				fmt.Println("Failed to migrate to previous version")
			}

			fmt.Println("Could not migrate: ", err)
			return err
		}
	}
	fmt.Println("Successfully migrated database")

	return nil 
}


func (db *DataBase) MigrateDownDB() error {
	driver, err := postgres.WithInstance(db.Client, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		viper.GetString("MIGRATION_FILE_URL"),
		"postgres",
		driver,
	)

	if err != nil {
		return err
	}


	// Migrating db
	if err := m.Steps(-1); err != nil {
		fmt.Println("Failed to migrate to previous version")
		return err
	}
	fmt.Println("Successfully migrated to previous version")

	return nil 
}


