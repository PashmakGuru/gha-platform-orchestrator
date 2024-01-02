package fronthub_test

import (
	"io"
	"os"
	"testing"

	fronthub "github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub_transformer"
	"github.com/stretchr/testify/assert"
)

func TestTransform(t *testing.T) {
	t.Run("transforming single input", func(t *testing.T) {
		jsonFile, _ := os.Open("../assets/fronthub-transformer-processed-1.json")
		expectedOutput, _ := io.ReadAll(jsonFile)

		output, err := fronthub.Transform([]string{
			"../assets/fronthub-transformer-input-1.json",
		})

		assert.NoError(t, err)
		assert.JSONEq(t, output, string(expectedOutput))
	})

	t.Run("transforming multiple input", func(t *testing.T) {
		jsonFile, _ := os.Open("../assets/fronthub-transformer-processed-2.json")
		expectedOutput, _ := io.ReadAll(jsonFile)

		output, err := fronthub.Transform([]string{
			"../assets/fronthub-transformer-input-2-1.json",
			"../assets/fronthub-transformer-input-2-2.json",
		})

		assert.NoError(t, err)
		assert.JSONEq(t, output, string(expectedOutput))
	})
}
