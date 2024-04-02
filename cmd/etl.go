package cmd

import (
	"github.com/spf13/cobra"
	"gitlab.com/IstiyakRiyad/technical-assessment-pathao/etl"
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
	// Extract the data from json files
	restaurants, users := etl.ExtractData()

	// Transform the data
	etl.TransformData(restaurants, users)
}

