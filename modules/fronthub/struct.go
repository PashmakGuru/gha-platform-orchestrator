package fronthub

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gosimple/slug"
)

type Fronthub struct {
	Zones []Zones `json:"zones"`
}
type Endpoints struct {
	Id      string `json:"id"`
	URL     string `json:"url"`
	Cluster string `json:"cluster"`
}
type Zones struct {
	Domain    string      `json:"domain"`
	Endpoints []Endpoints `json:"endpoints"`
}

func NewFronthub() *Fronthub {
	return &Fronthub{
		Zones: []Zones{},
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

	id := slug.Make(fmt.Sprintf("%s-%s", url, cluster))

	for key, zone := range f.Zones {
		if zone.Domain == domain {
			for _, endpoint := range zone.Endpoints {
				if endpoint.Id == id {
					return fmt.Errorf("the endpoint already exists for the domain")
				}
			}
			f.Zones[key].Endpoints = append(f.Zones[key].Endpoints, Endpoints{
				Id:      id,
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

func (f *Fronthub) DeleteEndpoint(id string) error {
	altered := false

	for zKey, zone := range f.Zones {
		for eKey, endpoint := range zone.Endpoints {
			if endpoint.Id == id {
				f.Zones[zKey].Endpoints = append(
					f.Zones[zKey].Endpoints[:eKey],
					f.Zones[zKey].Endpoints[eKey+1:]...,
				)
				altered = true
			}
		}
	}

	if !altered {
		return fmt.Errorf("unable to found domain or endpoint to delete")
	}

	return nil
}
