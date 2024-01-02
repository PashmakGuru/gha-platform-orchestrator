package fronthub_test

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
		copyFile("../assets/fronthub-empty.json", finalFile)

		defer deleteFile(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"fronthub:add-dns-zone",
			"--config-file",
			finalFile,
			"--domain",
			"my-test-domain.com",
		})
		cmd.RootCmd.Execute()

		config, err := fronthub.ReadFronthubConfig(finalFile)
		assert.NoError(t, err)

		assert.Len(t, config.Zones, 1)
		assert.Equal(t, config.Zones[0].Domain, "my-test-domain.com")
		assert.Empty(t, config.Zones[0].Endpoints)
		assert.Empty(t, config.Clusters)
	})

	t.Run("doesn't add the domain if already exists", func(t *testing.T) {
		finalFile := "/tmp/added-dns-zone.json"
		copyFile("../assets/fronthub-empty.json", finalFile)

		defer deleteFile(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"fronthub:add-dns-zone",
			"--config-file",
			finalFile,
			"--domain",
			"my-test-domain.com",
		})
		cmd.RootCmd.Execute()

		cmd.RootCmd.SetArgs([]string{
			"fronthub:add-dns-zone",
			"--config-file",
			finalFile,
			"--domain",
			"my-test-domain.com",
		})

		assert.Panics(t, func() {
			cmd.RootCmd.Execute()
		})

		config, err := fronthub.ReadFronthubConfig(finalFile)
		assert.NoError(t, err)

		assert.Len(t, config.Zones, 1)
	})
}

func deleteFile(path string) {
	os.Remove(path)
}

func copyFile(src string, dst string) error {
	data, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	err = os.WriteFile(dst, data, 0644)

	return nil
}
