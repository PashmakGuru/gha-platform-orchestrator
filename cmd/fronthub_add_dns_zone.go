/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub"
	"github.com/spf13/cobra"
)

var FronthubDnsZoneCmd = &cobra.Command{
	Use:   "fronthub:add-dns-zone",
	Short: "Modify the DNS zone",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := mustGetStringFlag(cmd, "config-file")
		domain := mustGetStringFlag(cmd, "domain")

		config, err := fronthub.ReadFronthubConfig(configPath)
		if err != nil {
			log.Panicln("unable to read config", err)
		}

		err = config.AddDnsZone(domain)
		if err != nil {
			log.Panicln("unable to add dns zone", err)
		}

		err = config.Save(configPath)
		if err != nil {
			log.Panicln("unable to write config file: "+configPath, err)
		}

		log.Println("all done.")
	},
}

func init() {
	RootCmd.AddCommand(FronthubDnsZoneCmd)
	FronthubDnsZoneCmd.Flags().StringP("config-file", "", "", "Domain of the DNS zone")
	FronthubDnsZoneCmd.Flags().StringP("domain", "", "", "Domain of the DNS zone")
	FronthubDnsZoneCmd.Flags().BoolP("delete", "", false, "Should the given zone be deleted?")
}
