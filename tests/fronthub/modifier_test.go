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
		fronthub.NewFronthub().Save(finalFile)

		defer deleteFile(finalFile)

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

		defer deleteFile(finalFile)

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

func Test_FronthubDeleteDnsZoneCommand(t *testing.T) {
	t.Run("deletes a domain", func(t *testing.T) {
		finalFile := "/tmp/deleted-dns-zone.json"

		config := fronthub.NewFronthub()
		config.AddDnsZone("alpha.com")
		config.AddDnsZone("beta.com")
		config.Save(finalFile)

		defer deleteFile(finalFile)

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

		defer deleteFile(finalFile)

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

func Test_FronthubAddEndpointCommand(t *testing.T) {
	t.Run("adds new endpoint", func(t *testing.T) {
		finalFile := "/tmp/data.json"

		config := fronthub.NewFronthub()
		config.AddDnsZone("alpha.com")
		config.AddDnsZone("beta.com")
		config.Save(finalFile)

		defer deleteFile(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"fronthub:add-endpoint",
			"--config-file",
			finalFile,
			"--domain",
			"alpha.com",
			"--url",
			"alpha.com/my-path/*",
			"--cluster",
			"cluster-alpha",
		})
		cmd.RootCmd.Execute()

		config, err := fronthub.ReadFronthubConfig(finalFile)
		assert.NoError(t, err)

		for _, zone := range config.Zones {
			if zone.Domain == "alpha.com" {
				assert.Len(t, zone.Endpoints, 1)
				assert.Equal(t, zone.Endpoints[0].URL, "alpha.com/my-path/*")
				assert.Equal(t, zone.Endpoints[0].Cluster, "cluster-alpha")
			} else {
				assert.Len(t, zone.Endpoints, 0)
			}
		}
	})

	t.Run("doesn't add the endpoint if its path already exists", func(t *testing.T) {
		finalFile := "/tmp/data.json"

		config := fronthub.NewFronthub()
		config.AddDnsZone("alpha.com")
		config.AddDnsZone("beta.com")
		config.AddEndpoint("alpha.com", "alpha.com/my-path/*", "cluster-alpha")
		config.Save(finalFile)

		defer deleteFile(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"fronthub:add-endpoint",
			"--config-file",
			finalFile,
			"--domain",
			"alpha.com",
			"--url",
			"alpha.com/my-path/*",
			"--cluster",
			"cluster-alpha",
		})

		assert.Panics(t, func() {
			cmd.RootCmd.Execute()
		})

		config, err := fronthub.ReadFronthubConfig(finalFile)
		assert.NoError(t, err)

		for _, zone := range config.Zones {
			if zone.Domain == "alpha.com" {
				assert.Len(t, zone.Endpoints, 1)
				assert.Equal(t, zone.Endpoints[0].URL, "alpha.com/my-path/*")
				assert.Equal(t, zone.Endpoints[0].Cluster, "cluster-alpha")
			} else {
				assert.Len(t, zone.Endpoints, 0)
			}
		}
	})

	t.Run("doesn't add the endpoint if the domain doesn't exist", func(t *testing.T) {
		finalFile := "/tmp/data.json"

		config := fronthub.NewFronthub()
		config.AddDnsZone("beta.com")
		config.Save(finalFile)

		defer deleteFile(finalFile)

		cmd.RootCmd.SetArgs([]string{
			"fronthub:add-endpoint",
			"--config-file",
			finalFile,
			"--domain",
			"alpha.com",
			"--url",
			"alpha.com/my-path/*",
			"--cluster",
			"cluster-alpha",
		})

		assert.Panics(t, func() {
			cmd.RootCmd.Execute()
		})

		config, err := fronthub.ReadFronthubConfig(finalFile)
		assert.NoError(t, err)

		assert.Len(t, config.Zones, 1)
		assert.Len(t, config.Zones[0].Endpoints, 0)
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
