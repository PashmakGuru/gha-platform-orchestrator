package portlabs

import "time"

type Entity[Properties any, Relations any] struct {
	Identifier string     `json:"identifier"`
	Title      string     `json:"title"`
	Icon       any        `json:"icon"`
	Blueprint  string     `json:"blueprint"`
	Team       []string   `json:"team"`
	Properties Properties `json:"properties"`
	Relations  Relations  `json:"relations"`
	CreatedAt  time.Time  `json:"createdAt"`
	CreatedBy  string     `json:"createdBy"`
	UpdatedAt  time.Time  `json:"updatedAt"`
	UpdatedBy  string     `json:"updatedBy"`
}

type ClusterEntity Entity[ClusterEntityProperties, ClusterEntityRelations]

type ClusterEntityProperties struct {
	AzureResourceGroupName     string `json:"azure_resource_group_name"`
	AzureResourceGroupLocation string `json:"azure_resource_group_location"`
}

type ClusterEntityRelations struct {
	Environment string `json:"environment"`
}
