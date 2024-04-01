package cmd

import "github.com/spf13/cobra"

var serviceCmd = &cobra.Command{
	Use: "start",
	Short: "Start the restaurant service",

	Run: startServiceFunc,
}

func init() {
	rootCmd.AddCommand(serviceCmd)
}


func startServiceFunc(cmd *cobra.Command, args []string) {
}
