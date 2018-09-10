package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Dyndns is a generated struct representing the Sophos Dyndns Endpoint
// GET /api/nodes/dyndns
type Dyndns struct {
	Rules []interface{} `json:"rules"`
}

var _ sophos.Endpoint = &Dyndns{}

var defsDyndns = map[string]sophos.RestObject{
	"DyndnsDyndns": &DyndnsDyndns{},
	"DyndnsGroup":  &DyndnsGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of Dyndns's Objects
func (Dyndns) RestObjects() map[string]sophos.RestObject { return defsDyndns }

// GetPath implements sophos.RestGetter
func (*Dyndns) GetPath() string { return "/api/nodes/dyndns" }

// RefRequired implements sophos.RestGetter
func (*Dyndns) RefRequired() (string, bool) { return "", false }

var defDyndns = &sophos.Definition{Description: "dyndns", Name: "dyndns", Link: "/api/definitions/dyndns"}

// Definition returns the /api/definitions struct of Dyndns
func (Dyndns) Definition() sophos.Definition { return *defDyndns }

// ApiRoutes returns all known Dyndns Paths
func (Dyndns) ApiRoutes() []string {
	return []string{
		"/api/objects/dyndns/dyndns/",
		"/api/objects/dyndns/dyndns/{ref}",
		"/api/objects/dyndns/dyndns/{ref}/usedby",
		"/api/objects/dyndns/group/",
		"/api/objects/dyndns/group/{ref}",
		"/api/objects/dyndns/group/{ref}/usedby",
	}
}

// References returns the Dyndns's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Dyndns) References() []string {
	return []string{
		"REF_DyndnsDyndns",
		"REF_DyndnsGroup",
	}
}

// DyndnsDyndnss is an Sophos Endpoint subType and implements sophos.RestObject
type DyndnsDyndnss []DyndnsDyndns

// DyndnsDyndns represents a UTM DynDNS mapping
type DyndnsDyndns struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	// Type can be one of: []string{"dns-o-matic", "dnsdynamic", "dnspark", "dtdns", "dyndns", "dyndns-custom", "easydns", "freedns", "namecheap", "no-ip", "opendns", "selfhost", "strato", "zoneedit"}
	// Type default value is "dyndns"
	Type    string        `json:"type"`
	Aliases []interface{} `json:"aliases"`
	// Hostname default value is ""
	Hostname string `json:"hostname"`
	// Label default value is ""
	Label string `json:"label"`
	// Mx description: (HOSTNAME)
	// Mx default value is ""
	Mx string `json:"mx"`
	// Password default value is ""
	Password string `json:"password"`
	// Interface description: REF(interface/*)
	Interface string `json:"interface"`
	// Record can be one of: []string{"a", "aaaa", "both"}
	// Record default value is "a"
	Record  string `json:"record"`
	User    string `json:"user"`
	Comment string `json:"comment"`
	// Status default value is false
	Status bool `json:"status"`
	// Strategy can be one of: []string{"if", "web"}
	// Strategy default value is "if"
	Strategy string `json:"strategy"`
	// Wildcard default value is false
	Wildcard bool `json:"wildcard"`
	// Backupmx default value is false
	Backupmx bool   `json:"backupmx"`
	Mxpri    int    `json:"mxpri"`
	Name     string `json:"name"`
}

var _ sophos.RestGetter = &DyndnsDyndns{}

// GetPath implements sophos.RestObject and returns the DyndnsDyndnss GET path
// Returns all available dyndns/dyndns objects
func (*DyndnsDyndnss) GetPath() string { return "/api/objects/dyndns/dyndns/" }

// RefRequired implements sophos.RestObject
func (*DyndnsDyndnss) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the DyndnsDyndnss GET path
// Returns all available dyndns types
func (d *DyndnsDyndns) GetPath() string {
	return fmt.Sprintf("/api/objects/dyndns/dyndns/%s", d.Reference)
}

// RefRequired implements sophos.RestObject
func (d *DyndnsDyndns) RefRequired() (string, bool) { return d.Reference, true }

// DeletePath implements sophos.RestObject and returns the DyndnsDyndns DELETE path
// Creates or updates the complete object dyndns
func (*DyndnsDyndns) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/dyndns/dyndns/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the DyndnsDyndns PATCH path
// Changes to parts of the object dyndns types
func (*DyndnsDyndns) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/dyndns/dyndns/%s", ref)
}

// PostPath implements sophos.RestObject and returns the DyndnsDyndns POST path
// Create a new dyndns/dyndns object
func (*DyndnsDyndns) PostPath() string {
	return "/api/objects/dyndns/dyndns/"
}

// PutPath implements sophos.RestObject and returns the DyndnsDyndns PUT path
// Creates or updates the complete object dyndns
func (*DyndnsDyndns) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/dyndns/dyndns/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*DyndnsDyndns) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dyndns/dyndns/%s/usedby", ref)
}

// DyndnsGroups is an Sophos Endpoint subType and implements sophos.RestObject
type DyndnsGroups []DyndnsGroup

// DyndnsGroup represents a UTM group
type DyndnsGroup struct {
	Locked     string `json:"_locked"`
	ObjectType string `json:"_type"`
	Reference  string `json:"_ref"`
	Comment    string `json:"comment"`
	Name       string `json:"name"`
}

var _ sophos.RestGetter = &DyndnsGroup{}

// GetPath implements sophos.RestObject and returns the DyndnsGroups GET path
// Returns all available dyndns/group objects
func (*DyndnsGroups) GetPath() string { return "/api/objects/dyndns/group/" }

// RefRequired implements sophos.RestObject
func (*DyndnsGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the DyndnsGroups GET path
// Returns all available group types
func (d *DyndnsGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/dyndns/group/%s", d.Reference)
}

// RefRequired implements sophos.RestObject
func (d *DyndnsGroup) RefRequired() (string, bool) { return d.Reference, true }

// DeletePath implements sophos.RestObject and returns the DyndnsGroup DELETE path
// Creates or updates the complete object group
func (*DyndnsGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/dyndns/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the DyndnsGroup PATCH path
// Changes to parts of the object group types
func (*DyndnsGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/dyndns/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the DyndnsGroup POST path
// Create a new dyndns/group object
func (*DyndnsGroup) PostPath() string {
	return "/api/objects/dyndns/group/"
}

// PutPath implements sophos.RestObject and returns the DyndnsGroup PUT path
// Creates or updates the complete object group
func (*DyndnsGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/dyndns/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*DyndnsGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dyndns/group/%s/usedby", ref)
}
