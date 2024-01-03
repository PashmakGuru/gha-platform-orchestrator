package cmd_test

import (
	"testing"

	"github.com/PashmakGuru/platform-cloud-resources/manager/cmd"
	"github.com/PashmakGuru/platform-cloud-resources/manager/tests"
	"github.com/stretchr/testify/assert"
)

func Test_FronthubTransformerCommand(t *testing.T) {
	cmd.RootCmd.SetArgs([]string{
		"fronthub:transform",
		"--input-file",
		"../tests/assets/fronthub-transformer-input-1.json",
		"--output-file",
		"/tmp/output-file.json",
	})
	cmd.RootCmd.Execute()

	actual, err := tests.ReadFile("/tmp/output-file.json")
	assert.NoError(t, err)

	expected, err := tests.ReadFile("../tests/assets/fronthub-transformer-processed-1.json")
	assert.NoError(t, err)

	assert.JSONEq(t, actual, expected)
}
