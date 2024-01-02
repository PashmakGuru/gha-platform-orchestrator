package fronthub

import (
	"encoding/json"
	"fmt"
	"os"
)

type Fronthub struct {
	Zones    []Zones    `json:"zones"`
	Clusters []Clusters `json:"clusters"`
}
type Endpoints struct {
	URL     string `json:"url"`
	Cluster string `json:"cluster"`
}
type Zones struct {
	Domain    string      `json:"domain"`
	Endpoints []Endpoints `json:"endpoints"`
}
type Clusters struct {
	Name          string `json:"name"`
	Subscription  string `json:"subscription"`
	ResourceGroup string `json:"resource_group"`
	IPNamePrefix  string `json:"ip_name_prefix"`
}

func (f *Fronthub) Save(path string) error {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, []byte(jsonData), 0644)
}

func (f *Fronthub) AddDnsZone(domain string) error {
	for _, zone := range f.Zones {
		if zone.Domain == domain {
			return fmt.Errorf("domain already exists")
		}
	}

	f.Zones = append(f.Zones, Zones{
		Domain:    domain,
		Endpoints: []Endpoints{},
	})

	return nil
}