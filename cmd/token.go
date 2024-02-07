/*
Copyright © 2024 Donald Gifford dgifford06@gmail.com
*/

package cmd

import (
	"context"
	"fmt"

	"github.com/donaldgifford/nester/src"
	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generates a token used for API calls",
	Long: `
  This command generates a token used for API calls. It caches it locally in 
  the file specified in your .nester.yaml configuration file. This is to be ran the first time
  you setup the tool so that it can generate a refresh token to use past the timeout without forcing
  you to re-authenticate.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("token called")
		ctx := context.Background()
		config := src.InitConfig(ctx)
		src.Authenticate(config)
	},
}

func init() {
	initCmd.AddCommand(tokenCmd)
}
