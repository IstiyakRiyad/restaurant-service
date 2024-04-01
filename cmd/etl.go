package cmd

import (
	"fmt"

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
	restaurants, users := etl.ExtractData()

	fmt.Println(restaurants)
	fmt.Println(users)
}

