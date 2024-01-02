/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub"
	"github.com/spf13/cobra"
)

var FronthubAddEndpointCmd = &cobra.Command{
	Use:   "fronthub:add-endpoint",
	Short: "Modify the DNS zone",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := mustGetStringFlag(cmd, "config-file")

		config, err := fronthub.ReadFronthubConfig(configPath)
		if err != nil {
			log.Panicln("unable to read config", err)
		}

		err = config.AddEndpoint(
			mustGetStringFlag(cmd, "domain"),
			mustGetStringFlag(cmd, "url"),
			mustGetStringFlag(cmd, "cluster"),
		)
		if err != nil {
			log.Panicln("unable to add endpoint", err)
		}

		err = config.Save(configPath)
		if err != nil {
			log.Panicln("unable to write config file: "+configPath, err)
		}

		log.Println("all done.")
	},
}

func init() {
	RootCmd.AddCommand(FronthubAddEndpointCmd)
	FronthubAddEndpointCmd.Flags().StringP("config-file", "", "", "Domain of the DNS zone")
	FronthubAddEndpointCmd.Flags().StringP("domain", "", "", "Domain of the DNS zone")
	FronthubAddEndpointCmd.Flags().StringP("url", "", "", "Endpoint URL")
	FronthubAddEndpointCmd.Flags().StringP("cluster", "", "", "Cluster ID on the IDP")
}
