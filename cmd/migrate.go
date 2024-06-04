package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/IstiyakRiyad/restaurant-service/db"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Short: "Migrate the database",
}

var migrateUpCmd = &cobra.Command{
	Use: "up",
	Short: "Make the database migrate up",

	Run: migrateUpFunc,
}

var migrateDownCmd = &cobra.Command{
	Use: "down",
	Short: "Make the database migrate down",

	Run: migrateDownFunc,
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.AddCommand(migrateUpCmd)
	migrateCmd.AddCommand(migrateDownCmd)
}


func migrateUpFunc(cmd *cobra.Command, args []string) {
	// Create the database if not exists
	db.CreateDBIfNotExist()

	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Connection Error", err)
	}

	if err := db.MigrateUpDB(); err != nil {
		log.Fatal("Could not migrate db: ", err)
	}
}

func migrateDownFunc(cmd *cobra.Command, args []string) {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Connection Error", err)
	}

	if err := db.MigrateDownDB(); err != nil {
		log.Fatal("Could not migrate down the db: ", err)
	}
}




