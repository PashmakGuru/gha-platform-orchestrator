package modifier

import "github.com/PashmakGuru/platform-cloud-resources/manager/modules/fronthub"

func AddDnsZone(input fronthub.Fronthub, domain string) *fronthub.Fronthub {
	input.Zones = append(input.Zones, fronthub.Zones{
		Domain:    domain,
		Endpoints: []fronthub.Endpoints{},
	})

	return &input
}

func DeleteDnsZone(input fronthub.Fronthub, domain string) *fronthub.Fronthub {
	input.Zones = append(input.Zones, fronthub.Zones{
		Domain:    domain,
		Endpoints: []fronthub.Endpoints{},
	})

	return &input
}
