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

func NewFronthub() *Fronthub {
	return &Fronthub{
		Zones:    []Zones{},
		Clusters: []Clusters{},
	}
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

func (f *Fronthub) DeleteDnsZone(domain string) error {
	found := false

	for key, zone := range f.Zones {
		if zone.Domain == domain {
			f.Zones = append(f.Zones[:key], f.Zones[key+1:]...)
			found = true
		}
	}

	if !found {
		return fmt.Errorf("the domain doesn't exist")
	}

	return nil
}

func (f *Fronthub) AddEndpoint(domain string, url string, cluster string) error {
	altered := false

	for key, zone := range f.Zones {
		if zone.Domain == domain {
			for _, endpoint := range zone.Endpoints {
				if endpoint.URL == url {
					return fmt.Errorf("the endpoint already exists for the domain")
				}
			}
			f.Zones[key].Endpoints = append(f.Zones[key].Endpoints, Endpoints{
				URL:     url,
				Cluster: cluster,
			})
			altered = true
		}
	}

	if !altered {
		return fmt.Errorf("unable to found domain '%s' to add its endpoint", domain)
	}

	return nil
}
