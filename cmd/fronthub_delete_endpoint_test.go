package cmd_test

import (
	"os"
	"testing"

	"github.com/PashmakGuru/platform-cloud-resources/manager/cmd"
	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub"
	"github.com/stretchr/testify/assert"
)

func Test_FronthubDeleteEndpointCommand(t *testing.T) {
	t.Run("deletes an endpoint", func(t *testing.T) {
		finalFile := "/tmp/data.json"

		config := fronthub.NewFronthub()
		config.AddDnsZone("alpha.com")
		config.AddEndpoint("alpha.com", "alpha.com/my-path/*", "cluster-alpha")
		config.AddEndpoint("alpha.com", "alpha.com/my-another/*", "cluster-alpha")
		config.Save(finalFile)

		defer os.Remove(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"fronthub:delete-endpoint",
			"--config-file",
			finalFile,
			"--id",
			"alpha-com-my-path-cluster-alpha",
		})
		cmd.RootCmd.Execute()

		config, err := fronthub.ReadFronthubConfig(finalFile)
		assert.NoError(t, err)

		assert.Len(t, config.Zones[0].Endpoints, 1)
		assert.Equal(t, config.Zones[0].Endpoints[0].Id, "alpha-com-my-another-cluster-alpha")
		assert.Equal(t, config.Zones[0].Endpoints[0].URL, "alpha.com/my-another/*")
		assert.Equal(t, config.Zones[0].Endpoints[0].Cluster, "cluster-alpha")
	})

	t.Run("doesn't delete the endpoint if the id doesn't exist", func(t *testing.T) {
		finalFile := "/tmp/data.json"

		config := fronthub.NewFronthub()
		config.AddDnsZone("alpha.com")
		config.AddEndpoint("alpha.com", "alpha.com/my-path/*", "cluster-alpha")
		config.Save(finalFile)

		defer os.Remove(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"fronthub:delete-endpoint",
			"--config-file",
			finalFile,
			"--id",
			"beta.com-my-path-cluster-beta",
		})

		assert.Panics(t, func() {
			cmd.RootCmd.Execute()
		})

		config, err := fronthub.ReadFronthubConfig(finalFile)
		assert.NoError(t, err)

		assert.Len(t, config.Zones[0].Endpoints, 1)
	})

}
