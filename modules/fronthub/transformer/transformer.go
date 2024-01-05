package transformer

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub"
	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/portlabs"
)

type UrlParts struct {
	Hostname   string
	SubDomain  string
	RootDomain string
	Path       string
}

func Transform(input fronthub.Fronthub, portClient *portlabs.PortClient) (*FronthubModuleInput, error) {
	output := NewFronthubModuleInput()

	for _, zone := range input.Zones {
		output.Zones = append(output.Zones, zone.Domain)
		for _, friendlyEndpoint := range zone.Endpoints {
			urlParts := parseUrl(friendlyEndpoint.URL)

			id := friendlyEndpoint.Id
			idCompressed := strings.ReplaceAll(id, "-", "")

			output.Endpoints = append(output.Endpoints, id)
			output.OriginGroups = append(output.OriginGroups, id)
			output.RuleSets = append(output.RuleSets, idCompressed)

			output.Routes[id] = Route{
				EndpointName:    id,
				OriginGroupName: id,
				OriginNames:     []string{id},
				RuleSetNames:    []string{idCompressed},
				PatternsToMatch: []string{urlParts.Path},
				UseAzureDomain:  true,
			}

			targetCluster, err := portClient.GetCluster(friendlyEndpoint.Cluster)
			if err != nil {
				return nil, fmt.Errorf("failed to get the cluster %s from port: %v", friendlyEndpoint.Cluster, err)
			}

			// var targetCluster fronthub.Clusters
			// for _, cluster := range input.Clusters {
			// 	if cluster.Name == friendlyEndpoint.Cluster {
			// 		targetCluster = cluster
			// 	}
			// }

			output.PublicIPOrigins[id] = PublicIpOrigin{
				OriginGroupName:      id,
				PipResourceGroupName: targetCluster.Properties.AzureResourceGroupName,
				PipNamePrefix:        "kubernetes-",
				OriginHostHeader:     urlParts.Hostname,
			}
		}
	}

	return &output, nil
}

// parseUrl processes a URL into subdomain, domain, and routing path
func parseUrl(inputUrl string) UrlParts {
	parsedUrl, err := url.Parse("https://" + inputUrl)
	if err != nil {
		return UrlParts{}
	}

	hostname := parsedUrl.Hostname()
	var subdomain, domain string

	// Split the hostname by dots
	parts := strings.Split(hostname, ".")
	if len(parts) >= 3 {
		subdomain = strings.Join(parts[:len(parts)-2], ".")
		domain = strings.Join(parts[len(parts)-2:], ".")
	} else {
		domain = hostname
	}

	// Extract the path
	path := parsedUrl.Path

	return UrlParts{
		Hostname:   hostname,
		SubDomain:  subdomain,
		RootDomain: domain,
		Path:       path,
	}
}
