package fronthub

import (
	"encoding/json"
	"io"
	"os"
)

func ReadFronthubConfig(file string) (*Fronthub, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result *Fronthub

	err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		panic(err)
	}

	return result, nil
}
