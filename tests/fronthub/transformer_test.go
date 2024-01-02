package fronthub_test

import (
	"io"
	"os"
	"testing"

	"github.com/PashmakGuru/platform-cloud-resources/manager/cmd"
	"github.com/stretchr/testify/assert"
)

func Test_FronthubTransformerCommand(t *testing.T) {
	cmd.RootCmd.SetArgs([]string{
		"fronthub:transform",
		"--input-file",
		"../assets/fronthub-transformer-input-1.json",
		"--output-file",
		"/tmp/output-file.json",
	})
	cmd.RootCmd.Execute()

	actual, err := readFile("/tmp/output-file.json")
	assert.NoError(t, err)

	expected, err := readFile("../assets/fronthub-transformer-processed-1.json")
	assert.NoError(t, err)

	assert.JSONEq(t, actual, expected)
}

func readFile(path string) (string, error) {
	content, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer content.Close()

	byteValue, err := io.ReadAll(content)

	return string(byteValue), err
}
