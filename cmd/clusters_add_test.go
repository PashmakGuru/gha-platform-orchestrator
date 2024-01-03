package cmd_test

import (
	"os"
	"testing"

	"github.com/PashmakGuru/platform-cloud-resources/manager/cmd"
	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/clusters"
	"github.com/stretchr/testify/assert"
)

func Test_ClustersAddCommand(t *testing.T) {
	t.Run("adds new cluster", func(t *testing.T) {
		finalFile := "/tmp/clusters_add_test.json"
		clusters.NewClustersConfig().Save(finalFile)

		defer os.Remove(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"clusters:add",
			"--config-file",
			finalFile,
			"--name",
			"my-cluster",
			"--environment",
			"testing",
			"--location",
			"East US",
		})
		cmd.RootCmd.Execute()

		config, err := clusters.ReadClusterConfigFile(finalFile)
		assert.NoError(t, err)

		assert.Len(t, config.Clusters, 1)
		assert.Equal(t, config.Clusters[0].Name, "my-cluster")
		assert.Equal(t, config.Clusters[0].Environment, "testing")
		assert.Equal(t, config.Clusters[0].ResourceGroupName, "cluster-solution-my-cluster-testing")
		assert.Equal(t, config.Clusters[0].ResourceGroupLocation, "East US")
	})

	t.Run("doesn't add if the cluster already exists", func(t *testing.T) {
		finalFile := "/tmp/clusters_add_test.json"
		config := clusters.NewClustersConfig()
		config.AddCluster("my-cluster", "production", "West US")
		config.Save(finalFile)

		defer os.Remove(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"clusters:add",
			"--config-file",
			finalFile,
			"--name",
			"my-cluster",
			"--environment",
			"testing",
			"--location",
			"East US",
		})

		assert.Panics(t, func() {
			cmd.RootCmd.Execute()
		})

		config, err := clusters.ReadClusterConfigFile(finalFile)
		assert.NoError(t, err)

		assert.Len(t, config.Clusters, 1)
		assert.Equal(t, config.Clusters[0].Name, "my-cluster")
		assert.Equal(t, config.Clusters[0].Environment, "production")
		assert.Equal(t, config.Clusters[0].ResourceGroupName, "cluster-solution-my-cluster-production")
		assert.Equal(t, config.Clusters[0].ResourceGroupLocation, "West US")
	})
}
