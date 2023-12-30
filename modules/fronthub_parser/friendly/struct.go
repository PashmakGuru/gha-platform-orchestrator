package friendly

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
