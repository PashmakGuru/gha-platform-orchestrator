package clusters

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/gosimple/slug"
)

type ClustersConfig struct {
	Clusters []Cluster `json:"clusters"`
}

type Cluster struct {
	Name                  string `json:"name"`
	ResourceGroupName     string `json:"resource_group_name"`
	ResourceGroupLocation string `json:"resource_group_location"`
	Environment           string `json:"environment"`
}

func NewClustersConfig() *ClustersConfig {
	return &ClustersConfig{
		Clusters: []Cluster{},
	}
}

func ReadClusterConfigFile(file string) (*ClustersConfig, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result *ClustersConfig

	err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		panic(err)
	}

	return result, nil
}

func (f *ClustersConfig) Save(path string) error {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, []byte(jsonData), 0644)
}

func (c *ClustersConfig) AddCluster(
	name string,
	environment string,
	location string,
) error {
	resourceGroupName := slug.Make(fmt.Sprintf("cluster-solution-%s-%s", name, environment))

	for _, cluster := range c.Clusters {
		if cluster.Name == name {
			return fmt.Errorf("a cluster with name '%s' already exists", name)
		}
	}

	c.Clusters = append(c.Clusters, Cluster{
		Name:                  name,
		ResourceGroupName:     resourceGroupName,
		ResourceGroupLocation: location,
		Environment:           environment,
	})

	return nil
}
