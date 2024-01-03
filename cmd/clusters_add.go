/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/clusters"
	"github.com/spf13/cobra"
)

var ClustersAddCmd = &cobra.Command{
	Use:   "clusters:add",
	Short: "Modify the DNS zone",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := mustGetStringFlag(cmd, "config-file")

		config, err := clusters.ReadClusterConfigFile(configPath)
		if err != nil {
			log.Panicln("unable to read config", err)
		}

		err = config.AddCluster(
			mustGetStringFlag(cmd, "name"),
			mustGetStringFlag(cmd, "environment"),
			mustGetStringFlag(cmd, "location"),
		)
		if err != nil {
			log.Panicln("unable to add cluster", err)
		}

		err = config.Save(configPath)
		if err != nil {
			log.Panicln("unable to write config file: "+configPath, err)
		}

		log.Println("all done.")
	},
}

func init() {
	RootCmd.AddCommand(ClustersAddCmd)
	ClustersAddCmd.Flags().StringP("config-file", "", "", "Clusters JSON configuration file")
	ClustersAddCmd.Flags().StringP("name", "", "", "Name of the cluster")
	ClustersAddCmd.Flags().StringP("environment", "", "", "Environment")
	ClustersAddCmd.Flags().StringP("location", "", "", "Resource group location on Azure")
}
