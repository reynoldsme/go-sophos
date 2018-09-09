package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Itfparams is a generated struct representing the Sophos Itfparams Endpoint
// GET /api/nodes/itfparams
type Itfparams struct {
	ItfparamsBridgePort           ItfparamsBridgePort           `json:"itfparams_bridge_port"`
	ItfparamsGroup                ItfparamsGroup                `json:"itfparams_group"`
	ItfparamsLinkAggregationGroup ItfparamsLinkAggregationGroup `json:"itfparams_link_aggregation_group"`
	ItfparamsPrimary              ItfparamsPrimary              `json:"itfparams_primary"`
	ItfparamsSecondary            ItfparamsSecondary            `json:"itfparams_secondary"`
}

var _ sophos.Endpoint = &Itfparams{}

var defsItfparams = map[string]sophos.RestObject{
	"ItfparamsBridgePort":           &ItfparamsBridgePort{},
	"ItfparamsGroup":                &ItfparamsGroup{},
	"ItfparamsLinkAggregationGroup": &ItfparamsLinkAggregationGroup{},
	"ItfparamsPrimary":              &ItfparamsPrimary{},
	"ItfparamsSecondary":            &ItfparamsSecondary{},
}

// RestObjects implements the sophos.Node interface and returns a map of Itfparams's Objects
func (Itfparams) RestObjects() map[string]sophos.RestObject { return defsItfparams }

// GetPath implements sophos.RestGetter
func (*Itfparams) GetPath() string { return "/api/nodes/itfparams" }

// RefRequired implements sophos.RestGetter
func (*Itfparams) RefRequired() (string, bool) { return "", false }

var defItfparams = &sophos.Definition{Description: "itfparams", Name: "itfparams", Link: "/api/definitions/itfparams"}

// Definition returns the /api/definitions struct of Itfparams
func (Itfparams) Definition() sophos.Definition { return *defItfparams }

// ApiRoutes returns all known Itfparams Paths
func (Itfparams) ApiRoutes() []string {
	return []string{
		"/api/objects/itfparams/bridge_port/",
		"/api/objects/itfparams/bridge_port/{ref}",
		"/api/objects/itfparams/bridge_port/{ref}/usedby",
		"/api/objects/itfparams/group/",
		"/api/objects/itfparams/group/{ref}",
		"/api/objects/itfparams/group/{ref}/usedby",
		"/api/objects/itfparams/link_aggregation_group/",
		"/api/objects/itfparams/link_aggregation_group/{ref}",
		"/api/objects/itfparams/link_aggregation_group/{ref}/usedby",
		"/api/objects/itfparams/primary/",
		"/api/objects/itfparams/primary/{ref}",
		"/api/objects/itfparams/primary/{ref}/usedby",
		"/api/objects/itfparams/secondary/",
		"/api/objects/itfparams/secondary/{ref}",
		"/api/objects/itfparams/secondary/{ref}/usedby",
	}
}

// References returns the Itfparams's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Itfparams) References() []string {
	return []string{
		"REF_ItfparamsBridgePort",
		"REF_ItfparamsGroup",
		"REF_ItfparamsLinkAggregationGroup",
		"REF_ItfparamsPrimary",
		"REF_ItfparamsSecondary",
	}
}

// ItfparamsBridgePort is an Sophos Endpoint subType and implements sophos.RestObject
type ItfparamsBridgePort []interface{}

var _ sophos.RestObject = &ItfparamsBridgePort{}

// GetPath implements sophos.RestObject and returns the ItfparamsBridgePort GET path
// Returns all available itfparams/bridge_port objects
func (*ItfparamsBridgePort) GetPath() string { return "/api/objects/itfparams/bridge_port/" }

// RefRequired implements sophos.RestObject
func (*ItfparamsBridgePort) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ItfparamsBridgePort DELETE path
// Creates or updates the complete object bridge_port
func (*ItfparamsBridgePort) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/bridge_port/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfparamsBridgePort PATCH path
// Changes to parts of the object bridge_port types
func (*ItfparamsBridgePort) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/bridge_port/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfparamsBridgePort POST path
// Create a new itfparams/bridge_port object
func (*ItfparamsBridgePort) PostPath() string {
	return "/api/objects/itfparams/bridge_port/"
}

// PutPath implements sophos.RestObject and returns the ItfparamsBridgePort PUT path
// Creates or updates the complete object bridge_port
func (*ItfparamsBridgePort) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/bridge_port/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfparamsBridgePort) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/bridge_port/%s/usedby", ref)
}

// ItfparamsGroup is an Sophos Endpoint subType and implements sophos.RestObject
type ItfparamsGroup []interface{}

var _ sophos.RestObject = &ItfparamsGroup{}

// GetPath implements sophos.RestObject and returns the ItfparamsGroup GET path
// Returns all available itfparams/group objects
func (*ItfparamsGroup) GetPath() string { return "/api/objects/itfparams/group/" }

// RefRequired implements sophos.RestObject
func (*ItfparamsGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ItfparamsGroup DELETE path
// Creates or updates the complete object group
func (*ItfparamsGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfparamsGroup PATCH path
// Changes to parts of the object group types
func (*ItfparamsGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfparamsGroup POST path
// Create a new itfparams/group object
func (*ItfparamsGroup) PostPath() string {
	return "/api/objects/itfparams/group/"
}

// PutPath implements sophos.RestObject and returns the ItfparamsGroup PUT path
// Creates or updates the complete object group
func (*ItfparamsGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfparamsGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/group/%s/usedby", ref)
}

// ItfparamsLinkAggregationGroups is an Sophos Endpoint subType and implements sophos.RestObject
type ItfparamsLinkAggregationGroups []ItfparamsLinkAggregationGroup

// ItfparamsLinkAggregationGroup is a generated Sophos object
type ItfparamsLinkAggregationGroup struct {
	Locked         string   `json:"_locked"`
	Reference      string   `json:"_ref"`
	_type          string   `json:"_type"`
	AdSelect       int64    `json:"ad_select"`
	ArpInterval    int64    `json:"arp_interval"`
	ArpIPTarget    string   `json:"arp_ip_target"`
	Comment        string   `json:"comment"`
	Downdelay      int64    `json:"downdelay"`
	EnforceMac     bool     `json:"enforce_mac"`
	ID             int64    `json:"id"`
	Itfhw          []string `json:"itfhw"`
	LacpRate       int64    `json:"lacp_rate"`
	Mac            string   `json:"mac"`
	Miimon         int64    `json:"miimon"`
	Mode           int64    `json:"mode"`
	Name           string   `json:"name"`
	Primary        string   `json:"primary"`
	Status         bool     `json:"status"`
	Updelay        int64    `json:"updelay"`
	UseCarrier     bool     `json:"use_carrier"`
	VirtualMac     string   `json:"virtual_mac"`
	XmitHashPolicy string   `json:"xmit_hash_policy"`
}

var _ sophos.RestGetter = &ItfparamsLinkAggregationGroup{}

// GetPath implements sophos.RestObject and returns the ItfparamsLinkAggregationGroups GET path
// Returns all available itfparams/link_aggregation_group objects
func (*ItfparamsLinkAggregationGroups) GetPath() string {
	return "/api/objects/itfparams/link_aggregation_group/"
}

// RefRequired implements sophos.RestObject
func (*ItfparamsLinkAggregationGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfparamsLinkAggregationGroups GET path
// Returns all available link_aggregation_group types
func (i *ItfparamsLinkAggregationGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/itfparams/link_aggregation_group/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *ItfparamsLinkAggregationGroup) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the ItfparamsLinkAggregationGroup DELETE path
// Creates or updates the complete object link_aggregation_group
func (*ItfparamsLinkAggregationGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/link_aggregation_group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfparamsLinkAggregationGroup PATCH path
// Changes to parts of the object link_aggregation_group types
func (*ItfparamsLinkAggregationGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/link_aggregation_group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfparamsLinkAggregationGroup POST path
// Create a new itfparams/link_aggregation_group object
func (*ItfparamsLinkAggregationGroup) PostPath() string {
	return "/api/objects/itfparams/link_aggregation_group/"
}

// PutPath implements sophos.RestObject and returns the ItfparamsLinkAggregationGroup PUT path
// Creates or updates the complete object link_aggregation_group
func (*ItfparamsLinkAggregationGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/link_aggregation_group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfparamsLinkAggregationGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/link_aggregation_group/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *ItfparamsLinkAggregationGroup) GetType() string { return i._type }

// ItfparamsPrimarys is an Sophos Endpoint subType and implements sophos.RestObject
type ItfparamsPrimarys []ItfparamsPrimary

// ItfparamsPrimary is a generated Sophos object
type ItfparamsPrimary struct {
	Locked                 string `json:"_locked"`
	Reference              string `json:"_ref"`
	_type                  string `json:"_type"`
	Address                string `json:"address"`
	Address6               string `json:"address6"`
	Comment                string `json:"comment"`
	DefaultGatewayAddress  string `json:"default_gateway_address"`
	DefaultGatewayAddress6 string `json:"default_gateway_address6"`
	DefaultGatewayStatus   bool   `json:"default_gateway_status"`
	DefaultGatewayStatus6  bool   `json:"default_gateway_status6"`
	Dhcpv6RapidCommit      bool   `json:"dhcpv6_rapid_commit"`
	DNSServer1             string `json:"dns_server_1"`
	DNSServer2             string `json:"dns_server_2"`
	DNSServer3             string `json:"dns_server_3"`
	DNSServer4             string `json:"dns_server_4"`
	GatewayType            string `json:"gateway_type"`
	GatewayType6           string `json:"gateway_type6"`
	Hostname               string `json:"hostname"`
	InterfaceAddress       string `json:"interface_address"`
	InterfaceBroadcast     string `json:"interface_broadcast"`
	InterfaceNetwork       string `json:"interface_network"`
	Name                   string `json:"name"`
	Netmask                int64  `json:"netmask"`
	Netmask6               int64  `json:"netmask6"`
	PdAddress6             string `json:"pd_address6"`
	PdNetmask6             int64  `json:"pd_netmask6"`
	PdResolved6            bool   `json:"pd_resolved6"`
	Resolved               bool   `json:"resolved"`
	Resolved6              bool   `json:"resolved6"`
	Six2four               bool   `json:"six2four"`
	Type                   string `json:"type"`
	Type6                  string `json:"type6"`
}

var _ sophos.RestGetter = &ItfparamsPrimary{}

// GetPath implements sophos.RestObject and returns the ItfparamsPrimarys GET path
// Returns all available itfparams/primary objects
func (*ItfparamsPrimarys) GetPath() string { return "/api/objects/itfparams/primary/" }

// RefRequired implements sophos.RestObject
func (*ItfparamsPrimarys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ItfparamsPrimarys GET path
// Returns all available primary types
func (i *ItfparamsPrimary) GetPath() string {
	return fmt.Sprintf("/api/objects/itfparams/primary/%s", i.Reference)
}

// RefRequired implements sophos.RestObject
func (i *ItfparamsPrimary) RefRequired() (string, bool) { return i.Reference, true }

// DeletePath implements sophos.RestObject and returns the ItfparamsPrimary DELETE path
// Creates or updates the complete object primary
func (*ItfparamsPrimary) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/primary/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfparamsPrimary PATCH path
// Changes to parts of the object primary types
func (*ItfparamsPrimary) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/primary/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfparamsPrimary POST path
// Create a new itfparams/primary object
func (*ItfparamsPrimary) PostPath() string {
	return "/api/objects/itfparams/primary/"
}

// PutPath implements sophos.RestObject and returns the ItfparamsPrimary PUT path
// Creates or updates the complete object primary
func (*ItfparamsPrimary) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/primary/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfparamsPrimary) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/primary/%s/usedby", ref)
}

// GetType implements sophos.Object
func (i *ItfparamsPrimary) GetType() string { return i._type }

// ItfparamsSecondary is an Sophos Endpoint subType and implements sophos.RestObject
type ItfparamsSecondary []interface{}

var _ sophos.RestObject = &ItfparamsSecondary{}

// GetPath implements sophos.RestObject and returns the ItfparamsSecondary GET path
// Returns all available itfparams/secondary objects
func (*ItfparamsSecondary) GetPath() string { return "/api/objects/itfparams/secondary/" }

// RefRequired implements sophos.RestObject
func (*ItfparamsSecondary) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ItfparamsSecondary DELETE path
// Creates or updates the complete object secondary
func (*ItfparamsSecondary) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/secondary/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ItfparamsSecondary PATCH path
// Changes to parts of the object secondary types
func (*ItfparamsSecondary) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/secondary/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ItfparamsSecondary POST path
// Create a new itfparams/secondary object
func (*ItfparamsSecondary) PostPath() string {
	return "/api/objects/itfparams/secondary/"
}

// PutPath implements sophos.RestObject and returns the ItfparamsSecondary PUT path
// Creates or updates the complete object secondary
func (*ItfparamsSecondary) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/secondary/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ItfparamsSecondary) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/itfparams/secondary/%s/usedby", ref)
}
