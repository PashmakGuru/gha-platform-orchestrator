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
	Use:   "fronthub-prepare",
	Short: "Parse fronthub configuration",
	Run: func(cmd *cobra.Command, args []string) {
		mifs, err := cmd.Flags().GetString("manual-input-files")
		if err != nil {
			log.Panicln("unable to get the manual input files flag", err)
		}

		outputDir, err := cmd.Flags().GetString("output-dir")
		if err != nil {
			log.Panicln("unable to get the output dir flag", err)
		}

		err = os.MkdirAll(outputDir, os.ModePerm)
		if err != nil {
			log.Panicln("unable to create the output dir", outputDir, err)
		}

		manualInputFiles := strings.Split(mifs, ",")

		outputConfig, err := fronthub.Process(manualInputFiles)
		if err != nil {
			log.Panicln("unable to process the fronthub inputs", err)
		}

		err = os.WriteFile(filepath.Join(outputDir, "/processed.lock.json"), []byte(outputConfig), 0644)
		if err != nil {
			log.Panicln("unable to write processed.lock.json", err)
		}

		log.Println("all done.")
	},
}

func init() {
	rootCmd.AddCommand(frontHubCmd)
	frontHubCmd.Flags().StringP("manual-input-files", "", "", "Comma separated list of manual input files")
	frontHubCmd.Flags().StringP("output-dir", "", "", "Output directory containing the Port-fetched input file and the processed one")
}
