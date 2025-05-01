package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of steGo",
	Long:  `All software has versions. This is steGo's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("steGo version %s\n", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
