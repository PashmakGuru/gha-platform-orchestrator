/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub"
	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub/transformer"
	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/portlabs"
	"github.com/spf13/cobra"
)

// FronthubCmd represents the frontHub command
var FronthubCmd = &cobra.Command{
	Use:   "fronthub:transform",
	Short: "Transform friendly fronthub configuration",
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := mustGetStringFlag(cmd, "input-file")
		outputPath := mustGetStringFlag(cmd, "output-file")

		inputConfig, err := fronthub.ReadFronthubConfig(inputPath)
		if err != nil {
			log.Panicln("unable to read config", err)
		}

		portClient, err := portlabs.NewPortClient(
			mustGetStringFlag(cmd, "port-client-id"),
			mustGetStringFlag(cmd, "port-client-secret"),
		)
		if err != nil {
			log.Panicln("unable to initiate port client", err)
		}

		outputConfig, err := transformer.Transform(*inputConfig, portClient)
		if err != nil {
			log.Panicln("unable to transform config", err)
		}

		err = outputConfig.Save(outputPath)
		if err != nil {
			log.Panicln("unable to write output file: "+outputPath, err)
		}

		log.Println("all done.")
	},
}

func mustGetStringFlag(cmd *cobra.Command, flag string) string {
	content, err := cmd.Flags().GetString(flag)
	if err != nil {
		log.Panicln("unable to get the following flag: "+flag, err)
	}

	return content
}

func init() {
	RootCmd.AddCommand(FronthubCmd)
	FronthubCmd.Flags().StringP("input-file", "", "", "Comma separated list of manual input files")
	FronthubCmd.Flags().StringP("output-file", "", "", "Output directory containing the Port-fetched input file and the processed one")
	FronthubCmd.Flags().StringP("port-client-id", "", "", "Client ID of Port")
	FronthubCmd.Flags().StringP("port-client-secret", "", "", "Client Secret of Port")
}
