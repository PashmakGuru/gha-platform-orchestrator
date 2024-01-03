package cmd_test

import (
	"os"
	"testing"

	"github.com/PashmakGuru/platform-cloud-resources/manager/cmd"
	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/clusters"
	"github.com/stretchr/testify/assert"
)

func Test_ClustersDeleteCommand(t *testing.T) {
	t.Run("deletes a cluster", func(t *testing.T) {
		finalFile := "/tmp/deleted-cluster.json"

		config := clusters.NewClustersConfig()
		config.AddCluster("alpha", "testing", "West US")
		config.AddCluster("beta", "testing", "West US")
		config.Save(finalFile)

		defer os.Remove(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"clusters:delete",
			"--config-file",
			finalFile,
			"--name",
			"alpha",
		})
		cmd.RootCmd.Execute()

		config, err := clusters.ReadClusterConfigFile(finalFile)
		assert.NoError(t, err)

		assert.Len(t, config.Clusters, 1)
		assert.Equal(t, config.Clusters[0].Name, "beta")
	})

	t.Run("can't delete a non-existence cluster", func(t *testing.T) {
		finalFile := "/tmp/deleted-cluster.json"

		config := clusters.NewClustersConfig()
		config.AddCluster("alpha", "testing", "West US")
		config.Save(finalFile)

		defer os.Remove(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"clusters:delete",
			"--config-file",
			finalFile,
			"--name",
			"beta",
		})

		assert.Panics(t, func() {
			cmd.RootCmd.Execute()
		})

		config, err := clusters.ReadClusterConfigFile(finalFile)
		assert.NoError(t, err)

		assert.Len(t, config.Clusters, 1)
		assert.Equal(t, config.Clusters[0].Name, "alpha")
	})
}
