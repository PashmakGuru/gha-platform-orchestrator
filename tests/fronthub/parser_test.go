package fronthub_test

import (
	"io"
	"os"
	"testing"

	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub"
	"github.com/stretchr/testify/assert"
)

func TestFronthub(t *testing.T) {
	t.Run("processing single input", func(t *testing.T) {
		jsonFile, _ := os.Open("../assets/fronthub-processed-1.json")
		expectedOutput, _ := io.ReadAll(jsonFile)

		output, err := fronthub.Process([]string{
			"../assets/fronthub-input-1.json",
		})

		assert.NoError(t, err)
		assert.JSONEq(t, output, string(expectedOutput))
	})

	t.Run("processing multiple input", func(t *testing.T) {
		jsonFile, _ := os.Open("../assets/fronthub-processed-2.json")
		expectedOutput, _ := io.ReadAll(jsonFile)

		output, err := fronthub.Process([]string{
			"../assets/fronthub-input-2-1.json",
			"../assets/fronthub-input-2-2.json",
		})

		assert.NoError(t, err)
		assert.JSONEq(t, output, string(expectedOutput))
	})
}
