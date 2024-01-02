/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	fronthub "github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub_parser"
	"github.com/spf13/cobra"
)

// frontHubCmd represents the frontHub command
var frontHubCmd = &cobra.Command{
	Use:   "fronthub-transform",
	Short: "Transform friendly fronthub configuration",
	Run: func(cmd *cobra.Command, args []string) {
		inputFiles := strings.Split(mustGetStringFlag("input-files"), ",")
		outputFile := mustGetStringFlag("output-file")

		outputConfig, err := fronthub.Process(inputFiles)
		if err != nil {
			log.Panicln("unable to transform config", err)
		}

		err = os.WriteFile(outputFile, []byte(outputConfig), 0644)
		if err != nil {
			log.Panicln("unable to write output file: "+outputFile, err)
		}

		log.Println("all done.")
	}
}

func mustGetStringFlag(name string) string {
	content, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Panicln("unable to get the following flag: "+flag, err)
	}

	return content
}

func init() {
	rootCmd.AddCommand(frontHubCmd)
	frontHubCmd.Flags().StringP("manual-input-files", "", "", "Comma separated list of manual input files")
	frontHubCmd.Flags().StringP("output-dir", "", "", "Output directory containing the Port-fetched input file and the processed one")
}
