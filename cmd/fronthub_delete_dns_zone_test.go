package cmd_test

import (
	"os"
	"testing"

	"github.com/PashmakGuru/platform-cloud-resources/manager/cmd"
	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub"
	"github.com/stretchr/testify/assert"
)

func Test_FronthubDeleteDnsZoneCommand(t *testing.T) {
	t.Run("deletes a domain", func(t *testing.T) {
		finalFile := "/tmp/deleted-dns-zone.json"

		config := fronthub.NewFronthub()
		config.AddDnsZone("alpha.com")
		config.AddDnsZone("beta.com")
		config.Save(finalFile)

		defer os.Remove(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"fronthub:delete-dns-zone",
			"--config-file",
			finalFile,
			"--domain",
			"alpha.com",
		})
		cmd.RootCmd.Execute()

		config, err := fronthub.ReadFronthubConfig(finalFile)
		assert.NoError(t, err)

		assert.Len(t, config.Zones, 1)
		assert.Equal(t, config.Zones[0].Domain, "beta.com")
	})

	t.Run("can't delete a non-existence domain", func(t *testing.T) {
		finalFile := "/tmp/deleted-dns-zone.json"

		config := fronthub.NewFronthub()
		config.AddDnsZone("beta.com")
		config.Save(finalFile)

		defer os.Remove(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"fronthub:delete-dns-zone",
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
		assert.Equal(t, config.Zones[0].Domain, "beta.com")
	})
}
