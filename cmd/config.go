/*
Copyright © 2024 Donald Gifford dgifford06@gmail.com
*/
package cmd

import (
	"fmt"

	"github.com/donaldgifford/nester/src"
	"github.com/spf13/cobra"
)

// configCmd generates a default config file
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Generate a config file for nester",
	Long:  `Creates a default config file in the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
		src.NewConfig()
	},
}

func init() {
	initCmd.AddCommand(configCmd)
}
