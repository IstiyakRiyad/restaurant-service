package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/IstiyakRiyad/restaurant-service/db"
	"github.com/IstiyakRiyad/restaurant-service/etl"
)

var etlCmd = &cobra.Command{
	Use: "etl",
	Short: "extract, transform and load the data",
	Run: etlFunc,
}

func init() {
	rootCmd.AddCommand(etlCmd)
}


func etlFunc(cmd *cobra.Command, args []string) {
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Connection Error", err)
	}

	// Extract the data from json files
	etl := etl.NewETL(db)

	// Start the etl process
	etl.Start()
}

