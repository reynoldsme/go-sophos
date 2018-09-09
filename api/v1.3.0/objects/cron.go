package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Cron is a generated struct representing the Sophos Cron Endpoint
// GET /api/nodes/cron
type Cron struct {
	CronAt    CronAt    `json:"cron_at"`
	CronGroup CronGroup `json:"cron_group"`
}

var _ sophos.Endpoint = &Cron{}

var defsCron = map[string]sophos.RestObject{
	"CronAt":    &CronAt{},
	"CronGroup": &CronGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of Cron's Objects
func (Cron) RestObjects() map[string]sophos.RestObject { return defsCron }

// GetPath implements sophos.RestGetter
func (*Cron) GetPath() string { return "/api/nodes/cron" }

// RefRequired implements sophos.RestGetter
func (*Cron) RefRequired() (string, bool) { return "", false }

var defCron = &sophos.Definition{Description: "cron", Name: "cron", Link: "/api/definitions/cron"}

// Definition returns the /api/definitions struct of Cron
func (Cron) Definition() sophos.Definition { return *defCron }

// ApiRoutes returns all known Cron Paths
func (Cron) ApiRoutes() []string {
	return []string{
		"/api/objects/cron/at/",
		"/api/objects/cron/at/{ref}",
		"/api/objects/cron/at/{ref}/usedby",
		"/api/objects/cron/group/",
		"/api/objects/cron/group/{ref}",
		"/api/objects/cron/group/{ref}/usedby",
	}
}

// References returns the Cron's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Cron) References() []string {
	return []string{
		"REF_CronAt",
		"REF_CronGroup",
	}
}

// CronAt is an Sophos Endpoint subType and implements sophos.RestObject
type CronAt []interface{}

var _ sophos.RestObject = &CronAt{}

// GetPath implements sophos.RestObject and returns the CronAt GET path
// Returns all available cron/at objects
func (*CronAt) GetPath() string { return "/api/objects/cron/at/" }

// RefRequired implements sophos.RestObject
func (*CronAt) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the CronAt DELETE path
// Creates or updates the complete object at
func (*CronAt) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/cron/at/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the CronAt PATCH path
// Changes to parts of the object at types
func (*CronAt) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/cron/at/%s", ref)
}

// PostPath implements sophos.RestObject and returns the CronAt POST path
// Create a new cron/at object
func (*CronAt) PostPath() string {
	return "/api/objects/cron/at/"
}

// PutPath implements sophos.RestObject and returns the CronAt PUT path
// Creates or updates the complete object at
func (*CronAt) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/cron/at/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*CronAt) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/cron/at/%s/usedby", ref)
}

// CronGroup is an Sophos Endpoint subType and implements sophos.RestObject
type CronGroup []interface{}

var _ sophos.RestObject = &CronGroup{}

// GetPath implements sophos.RestObject and returns the CronGroup GET path
// Returns all available cron/group objects
func (*CronGroup) GetPath() string { return "/api/objects/cron/group/" }

// RefRequired implements sophos.RestObject
func (*CronGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the CronGroup DELETE path
// Creates or updates the complete object group
func (*CronGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/cron/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the CronGroup PATCH path
// Changes to parts of the object group types
func (*CronGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/cron/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the CronGroup POST path
// Create a new cron/group object
func (*CronGroup) PostPath() string {
	return "/api/objects/cron/group/"
}

// PutPath implements sophos.RestObject and returns the CronGroup PUT path
// Creates or updates the complete object group
func (*CronGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/cron/group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*CronGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/cron/group/%s/usedby", ref)
}
