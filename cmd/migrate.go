package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/db"
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

	_, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Connection Error", err)
	}
}

func migrateDownFunc(cmd *cobra.Command, args []string) {

}




