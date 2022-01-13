package models

type GatewaysSchema struct {
	SerialNumber        string
	OrgID               int
	SiteID              int
	NodeID              string
	MeshAddr            int
	MacID               string
	MeshConfiguredState int
	CloudStatus         int
	ExtServicesState    string
}
