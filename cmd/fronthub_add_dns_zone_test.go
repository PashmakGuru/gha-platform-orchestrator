package cmd_test

import (
	"os"
	"testing"

	"github.com/PashmakGuru/platform-cloud-resources/manager/cmd"
	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub"
	"github.com/stretchr/testify/assert"
)

func Test_FronthubAddDnsZoneCommand(t *testing.T) {
	t.Run("adds new domain", func(t *testing.T) {
		finalFile := "/tmp/added-dns-zone.json"
		fronthub.NewFronthub().Save(finalFile)

		defer os.Remove(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"fronthub:add-dns-zone",
			"--config-file",
			finalFile,
			"--domain",
			"alpha.com",
		})
		cmd.RootCmd.Execute()

		config, err := fronthub.ReadFronthubConfig(finalFile)
		assert.NoError(t, err)

		assert.Len(t, config.Zones, 1)
		assert.Equal(t, config.Zones[0].Domain, "alpha.com")
		assert.Empty(t, config.Zones[0].Endpoints)
	})

	t.Run("doesn't add the domain if already exists", func(t *testing.T) {
		finalFile := "/tmp/added-dns-zone.json"
		fronthub.NewFronthub().Save(finalFile)

		defer os.Remove(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"fronthub:add-dns-zone",
			"--config-file",
			finalFile,
			"--domain",
			"alpha.com",
		})
		cmd.RootCmd.Execute()

		cmd.RootCmd.SetArgs([]string{
			"fronthub:add-dns-zone",
			"--config-file",
			finalFile,
			"--domain",
			"alpha.com",
		})

		assert.Panics(t, func() {
			cmd.RootCmd.Execute()
		})

		config, err := fronthub.ReadFronthubConfig(finalFile)
		assert.NoError(t, err)

		assert.Len(t, config.Zones, 1)
	})
}
