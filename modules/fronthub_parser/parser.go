package fronthub

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub_parser/friendly"
	"github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub_parser/processed"
	"github.com/gosimple/slug"
)

type UrlParts struct {
	Hostname   string
	SubDomain  string
	RootDomain string
	Path       string
}

func Process(inputFiles []string) (string, error) {
	output := processed.Fronthub{
		Routes:          map[string]processed.Route{},
		PublicIPOrigins: map[string]processed.PublicIpOrigin{},
	}

	for _, file := range inputFiles {
		input, err := readFile(file)
		if err != nil {
			return "", err
		}

		for _, zone := range input.Zones {
			output.Zones = append(output.Zones, zone.Domain)
			for _, friendlyEndpoint := range zone.Endpoints {
				urlParts := parseUrl(friendlyEndpoint.URL)

				id := slug.Make(fmt.Sprintf("%s-%s", friendlyEndpoint.URL, friendlyEndpoint.Cluster))
				idCompressed := strings.ReplaceAll(id, "-", "")

				output.Endpoints = append(output.Endpoints, id)
				output.OriginGroups = append(output.OriginGroups, id)
				output.RuleSets = append(output.RuleSets, idCompressed)

				output.Routes[id] = processed.Route{
					EndpointName:    id,
					OriginGroupName: id,
					OriginNames:     []string{id},
					RuleSetNames:    []string{idCompressed},
					PatternsToMatch: []string{urlParts.Path},
					UseAzureDomain:  true,
				}

				var targetCluster friendly.Clusters
				for _, cluster := range input.Clusters {
					if cluster.Name == friendlyEndpoint.Cluster {
						targetCluster = cluster
					}
				}

				if targetCluster.Name == "" {
					return "", fmt.Errorf("unable to find target cluster %s", friendlyEndpoint.Cluster)
				}

				output.PublicIPOrigins[id] = processed.PublicIpOrigin{
					OriginGroupName:      id,
					PipResourceGroupName: targetCluster.ResourceGroup,
					PipNamePrefix:        "kubernetes-",
					OriginHostHeader:     urlParts.Hostname,
				}
			}
		}
	}

	processedOutput, err := json.Marshal(output)
	if err != nil {
		return "", err
	}

	return string(processedOutput), nil
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

func readFile(file string) (*friendly.Fronthub, error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var result *friendly.Fronthub

	err = json.Unmarshal([]byte(byteValue), &result)

	if err != nil {
		panic(err)
	}

	return result, nil
}
