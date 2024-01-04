package transformer

import (
	"encoding/json"
	"os"
)

type FronthubModuleInput struct {
	Zones           []string        `json:"zones"`
	OriginGroups    []string        `json:"origin_groups"`
	PublicIPOrigins PublicIPOrigins `json:"public_ip_origins"`
	Endpoints       []string        `json:"endpoints"`
	RuleSets        []string        `json:"rule_sets"`
	Routes          Routes          `json:"routes"`
}
type PublicIpOrigin struct {
	OriginGroupName      string `json:"origin_group_name"`
	PipResourceGroupName string `json:"pip_resource_group_name"`
	PipNamePrefix        string `json:"pip_name_prefix"`
	OriginHostHeader     string `json:"origin_host_header"`
}
type PublicIPOrigins map[string]PublicIpOrigin

type Route struct {
	EndpointName    string   `json:"endpoint_name"`
	OriginGroupName string   `json:"origin_group_name"`
	OriginNames     []string `json:"origin_names"`
	RuleSetNames    []string `json:"rule_set_names"`
	PatternsToMatch []string `json:"patterns_to_match"`
	UseAzureDomain  bool     `json:"use_azure_domain"`
}

type Routes map[string]Route

func (f *FronthubModuleInput) Save(path string) error {
	jsonData, err := json.MarshalIndent(f, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, []byte(jsonData), 0644)
}

func NewFronthubModuleInput() FronthubModuleInput {
	return FronthubModuleInput{
		Zones:           []string{},
		Routes:          map[string]Route{},
		PublicIPOrigins: map[string]PublicIpOrigin{},
		OriginGroups:    []string{},
		Endpoints:       []string{},
		RuleSets:        []string{},
	}
}
