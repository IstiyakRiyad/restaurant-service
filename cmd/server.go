package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/IstiyakRiyad/restaurant-service/db"
	"github.com/IstiyakRiyad/restaurant-service/internal/restaurant"
	transportHttp "github.com/IstiyakRiyad/restaurant-service/transport/http"
)

var serviceCmd = &cobra.Command{
	Use: "start",
	Short: "Start the restaurant service",

	Run: startServiceFunc,
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}


func startServiceFunc(cmd *cobra.Command, args []string) {
	// Create a database connection
	db, err := db.NewDatabase()
	if err != nil {
		log.Fatal("Database Connection Problem: ", err)
	}

	// Create a service with passing the database connection to it
	service := restaurant.NewRestaurantService(db)

	// Create a http transport layer and pass it the service
	transport := transportHttp.NewHandler(service)

	// Start the server 
	if err := transport.Serve(); err != nil {
		log.Fatal("Server Cound not start: ", err)
	}
}



