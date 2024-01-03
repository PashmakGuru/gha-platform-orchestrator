/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/clusters"
	"github.com/spf13/cobra"
)

var ClustersDeleteCmd = &cobra.Command{
	Use:   "clusters:delete",
	Short: "Delete a cluster",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := mustGetStringFlag(cmd, "config-file")

		config, err := clusters.ReadClusterConfigFile(configPath)
		if err != nil {
			log.Panicln("unable to read config", err)
		}

		err = config.DeleteCluster(
			mustGetStringFlag(cmd, "name"),
		)
		if err != nil {
			log.Panicln("unable to delete cluster", err)
		}

		err = config.Save(configPath)
		if err != nil {
			log.Panicln("unable to write config file: "+configPath, err)
		}

		log.Println("all done.")
	},
}

func init() {
	RootCmd.AddCommand(ClustersDeleteCmd)
	ClustersDeleteCmd.Flags().StringP("config-file", "", "", "Clusters JSON configuration file")
	ClustersDeleteCmd.Flags().StringP("name", "", "", "Name of the cluster")
}
