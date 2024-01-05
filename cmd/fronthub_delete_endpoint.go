/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub"
	"github.com/spf13/cobra"
)

var FronthubDeleteEndpointCmd = &cobra.Command{
	Use:   "fronthub:delete-endpoint",
	Short: "Delete a domain's endpoint",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := mustGetStringFlag(cmd, "config-file")

		config, err := fronthub.ReadFronthubConfig(configPath)
		if err != nil {
			log.Panicln("unable to read config", err)
		}

		err = config.DeleteEndpoint(
			mustGetStringFlag(cmd, "id"),
		)
		if err != nil {
			log.Panicln("unable to delete endpoint", err)
		}

		err = config.Save(configPath)
		if err != nil {
			log.Panicln("unable to write config file: "+configPath, err)
		}

		log.Println("all done.")
	},
}

func init() {
	RootCmd.AddCommand(FronthubDeleteEndpointCmd)
	FronthubDeleteEndpointCmd.Flags().StringP("config-file", "", "", "Domain of the DNS zone")
	FronthubDeleteEndpointCmd.Flags().StringP("id", "", "", "Endpoint ID")
}
