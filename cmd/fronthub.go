/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// frontHubCmd represents the frontHub command
var frontHubCmd = &cobra.Command{
	Use:   "fronthub",
	Short: "Parse front-hub configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("frontHub called")
	},
}

func init() {
	rootCmd.AddCommand(frontHubCmd)
}
