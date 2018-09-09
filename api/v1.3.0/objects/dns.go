package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Dns is a generated struct representing the Sophos Dns Endpoint
// GET /api/nodes/dns
type Dns struct {
	AllowedNetworks []string      `json:"allowed_networks"`
	Axfr            []interface{} `json:"axfr"`
	Dnssec          int64         `json:"dnssec"`
	Email           string        `json:"email"`
	EmptyZones      string        `json:"empty_zones"`
	FwdDynamic      int64         `json:"fwd_dynamic"`
	FwdStatic       []string      `json:"fwd_static"`
	RecheckInterval int64         `json:"recheck_interval"`
	Routes          []string      `json:"routes"`
}

var defsDns = map[string]sophos.RestObject{
	"DnsAxfr":  &DnsAxfr{},
	"DnsGroup": &DnsGroup{},
	"DnsRoute": &DnsRoute{},
}

// RestObjects implements the sophos.Node interface and returns a map of Dns's Objects
func (Dns) RestObjects() map[string]sophos.RestObject { return defsDns }

// GetPath implements sophos.RestGetter
func (*Dns) GetPath() string { return "/api/nodes/dns" }

// RefRequired implements sophos.RestGetter
func (*Dns) RefRequired() (string, bool) { return "", false }

var defDns = &sophos.Definition{Description: "dns", Name: "dns", Link: "/api/definitions/dns"}

// Definition returns the /api/definitions struct of Dns
func (Dns) Definition() sophos.Definition { return *defDns }

// ApiRoutes returns all known Dns Paths
func (Dns) ApiRoutes() []string {
	return []string{
		"/api/objects/dns/axfr/",
		"/api/objects/dns/axfr/{ref}",
		"/api/objects/dns/axfr/{ref}/usedby",
		"/api/objects/dns/group/",
		"/api/objects/dns/group/{ref}",
		"/api/objects/dns/group/{ref}/usedby",
		"/api/objects/dns/route/",
		"/api/objects/dns/route/{ref}",
		"/api/objects/dns/route/{ref}/usedby",
	}
}

// References returns the Dns's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Dns) References() []string {
	return []string{
		"REF_DnsAxfr",
		"REF_DnsGroup",
		"REF_DnsRoute",
	}
}

// DnsAxfr is an Sophos Endpoint subType and implements sophos.RestObject
type DnsAxfr []interface{}

// GetPath implements sophos.RestObject and returns the DnsAxfr GET path
// Returns all available dns/axfr objects
func (*DnsAxfr) GetPath() string { return "/api/objects/dns/axfr/" }

// RefRequired implements sophos.RestObject
func (*DnsAxfr) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the DnsAxfr DELETE path
// Creates or updates the complete object axfr
func (*DnsAxfr) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/dns/axfr/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the DnsAxfr PATCH path
// Changes to parts of the object axfr types
func (*DnsAxfr) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/dns/axfr/%s", ref)
}

// PostPath implements sophos.RestObject and returns the DnsAxfr POST path
// Create a new dns/axfr object
func (*DnsAxfr) PostPath() string {
	return "/api/objects/dns/axfr/"
}

// PutPath implements sophos.RestObject and returns the DnsAxfr PUT path
// Creates or updates the complete object axfr
func (*DnsAxfr) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/dns/axfr/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*DnsAxfr) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dns/axfr/%s/usedby", ref)
}

// DnsGroup is an Sophos Endpoint subType and implements sophos.RestObject
type DnsGroup []interface{}

// GetPath implements sophos.RestObject and returns the DnsGroup GET path
// Returns all available dns/group objects
func (*DnsGroup) GetPath() string { return "/api/objects/dns/group/" }

// RefRequired implements sophos.RestObject
func (*DnsGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the DnsGroup DELETE path
// Creates or updates the complete object group
func (*DnsGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/dns/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the DnsGroup PATCH path
// Changes to parts of the object group types
func (*DnsGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/dns/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the DnsGroup POST path
// Create a new dns/group object
func (*DnsGroup) PostPath() string {
	return "/api/objects/dns/group/"
}

// PutPath implements sophos.RestObject and returns the DnsGroup PUT path
// Creates or updates the complete object group
func (*DnsGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/dns/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*DnsGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dns/group/%s/usedby", ref)
}

// DnsRoutes is an Sophos Endpoint subType and implements sophos.RestObject
type DnsRoutes []DnsRoute

// DnsRoute is a generated Sophos object
type DnsRoute struct {
	Locked    string   `json:"_locked"`
	Reference string   `json:"_ref"`
	_type     string   `json:"_type"`
	Comment   string   `json:"comment"`
	Name      string   `json:"name"`
	Prefix    string   `json:"prefix"`
	Status    bool     `json:"status"`
	Targets   []string `json:"targets"`
}

// GetPath implements sophos.RestObject and returns the DnsRoutes GET path
// Returns all available dns/route objects
func (*DnsRoutes) GetPath() string { return "/api/objects/dns/route/" }

// RefRequired implements sophos.RestObject
func (*DnsRoutes) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the DnsRoutes GET path
// Returns all available route types
func (d *DnsRoute) GetPath() string { return fmt.Sprintf("/api/objects/dns/route/%s", d.Reference) }

// RefRequired implements sophos.RestObject
func (d *DnsRoute) RefRequired() (string, bool) { return d.Reference, true }

// DeletePath implements sophos.RestObject and returns the DnsRoute DELETE path
// Creates or updates the complete object route
func (*DnsRoute) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/dns/route/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the DnsRoute PATCH path
// Changes to parts of the object route types
func (*DnsRoute) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/dns/route/%s", ref)
}

// PostPath implements sophos.RestObject and returns the DnsRoute POST path
// Create a new dns/route object
func (*DnsRoute) PostPath() string {
	return "/api/objects/dns/route/"
}

// PutPath implements sophos.RestObject and returns the DnsRoute PUT path
// Creates or updates the complete object route
func (*DnsRoute) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/dns/route/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*DnsRoute) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dns/route/%s/usedby", ref)
}

// GetType implements sophos.Object
func (d *DnsRoute) GetType() string { return d._type }