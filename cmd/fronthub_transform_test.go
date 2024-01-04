package cmd_test

import (
	"testing"

	"github.com/PashmakGuru/platform-cloud-resources/manager/cmd"
	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/portlabs"
	"github.com/PashmakGuru/platform-cloud-resources/manager/tests"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func Test_FronthubTransformerCommand(t *testing.T) {
	defer gock.Off()

	gock.New("https://api.getport.io").
		Post("/v1/auth/access_token").
		MatchType("json").
		JSON(portlabs.AccessTokenRequest{
			ClientID:     "test-client-id",
			ClientSecret: "test-client-secret",
		}).
		Reply(200).
		JSON(portlabs.AccessTokenResponse{
			Ok:          true,
			AccessToken: "my-access-token",
		})

	gock.New("https://api.getport.io").
		Get("/v1/blueprints/clusters/entities/cluster-alpha").
		MatchHeader("Authorization", "my-access-token").
		Reply(200).
		JSON(portlabs.ClusterEntity{
			Properties: portlabs.ClusterEntityProperties{
				AzureResourceGroupName: "cluster-solution-cluster-alpha-testing",
			},
			Relations: portlabs.ClusterEntityRelations{
				Environment: "testing",
			},
		})

	cmd.RootCmd.SetArgs([]string{
		"fronthub:transform",
		"--input-file",
		"../tests/assets/fronthub-transformer-input-1.json",
		"--output-file",
		"/tmp/output-file.json",
		"--port-client-id",
		"test-client-id",
		"--port-client-secret",
		"test-client-secret",
	})
	cmd.RootCmd.Execute()

	actual, err := tests.ReadFile("/tmp/output-file.json")
	assert.NoError(t, err)

	expected, err := tests.ReadFile("../tests/assets/fronthub-transformer-processed-1.json")
	assert.NoError(t, err)

	assert.JSONEq(t, actual, expected)
	assert.True(t, gock.IsDone(), "There are pending requests from gock!")
}
