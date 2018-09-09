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

// DyndnsDyndns is an Sophos Endpoint subType and implements sophos.RestObject
type DyndnsDyndns []interface{}

var _ sophos.RestObject = &DyndnsDyndns{}

// GetPath implements sophos.RestObject and returns the DyndnsDyndns GET path
// Returns all available dyndns/dyndns objects
func (*DyndnsDyndns) GetPath() string { return "/api/objects/dyndns/dyndns/" }

// RefRequired implements sophos.RestObject
func (*DyndnsDyndns) RefRequired() (string, bool) { return "", false }

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

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*DyndnsDyndns) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dyndns/dyndns/%s/usedby", ref)
}

// DyndnsGroup is an Sophos Endpoint subType and implements sophos.RestObject
type DyndnsGroup []interface{}

var _ sophos.RestObject = &DyndnsGroup{}

// GetPath implements sophos.RestObject and returns the DyndnsGroup GET path
// Returns all available dyndns/group objects
func (*DyndnsGroup) GetPath() string { return "/api/objects/dyndns/group/" }

// RefRequired implements sophos.RestObject
func (*DyndnsGroup) RefRequired() (string, bool) { return "", false }

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

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*DyndnsGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/dyndns/group/%s/usedby", ref)
}
