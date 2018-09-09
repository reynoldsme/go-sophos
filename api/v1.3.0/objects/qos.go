package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Qos is a generated struct representing the Sophos Qos Endpoint
// GET /api/nodes/qos
type Qos struct {
	Advanced struct {
		Ecn                 int64 `json:"ecn"`
		KeepClassAfterEncap int64 `json:"keep_class_after_encap"`
	} `json:"advanced"`
	Interfaces []string `json:"interfaces"`
}

var _ sophos.Endpoint = &Qos{}

var defsQos = map[string]sophos.RestObject{
	"QosApplicationSelector":  &QosApplicationSelector{},
	"QosGroup":                &QosGroup{},
	"QosIngressRule":          &QosIngressRule{},
	"QosInterface":            &QosInterface{},
	"QosRule":                 &QosRule{},
	"QosTrafficSelector":      &QosTrafficSelector{},
	"QosTrafficSelectorGroup": &QosTrafficSelectorGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of Qos's Objects
func (Qos) RestObjects() map[string]sophos.RestObject { return defsQos }

// GetPath implements sophos.RestGetter
func (*Qos) GetPath() string { return "/api/nodes/qos" }

// RefRequired implements sophos.RestGetter
func (*Qos) RefRequired() (string, bool) { return "", false }

var defQos = &sophos.Definition{Description: "qos", Name: "qos", Link: "/api/definitions/qos"}

// Definition returns the /api/definitions struct of Qos
func (Qos) Definition() sophos.Definition { return *defQos }

// ApiRoutes returns all known Qos Paths
func (Qos) ApiRoutes() []string {
	return []string{
		"/api/objects/qos/application_selector/",
		"/api/objects/qos/application_selector/{ref}",
		"/api/objects/qos/application_selector/{ref}/usedby",
		"/api/objects/qos/group/",
		"/api/objects/qos/group/{ref}",
		"/api/objects/qos/group/{ref}/usedby",
		"/api/objects/qos/ingress_rule/",
		"/api/objects/qos/ingress_rule/{ref}",
		"/api/objects/qos/ingress_rule/{ref}/usedby",
		"/api/objects/qos/interface/",
		"/api/objects/qos/interface/{ref}",
		"/api/objects/qos/interface/{ref}/usedby",
		"/api/objects/qos/rule/",
		"/api/objects/qos/rule/{ref}",
		"/api/objects/qos/rule/{ref}/usedby",
		"/api/objects/qos/traffic_selector/",
		"/api/objects/qos/traffic_selector/{ref}",
		"/api/objects/qos/traffic_selector/{ref}/usedby",
		"/api/objects/qos/traffic_selector_group/",
		"/api/objects/qos/traffic_selector_group/{ref}",
		"/api/objects/qos/traffic_selector_group/{ref}/usedby",
	}
}

// References returns the Qos's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Qos) References() []string {
	return []string{
		"REF_QosApplicationSelector",
		"REF_QosGroup",
		"REF_QosIngressRule",
		"REF_QosInterface",
		"REF_QosRule",
		"REF_QosTrafficSelector",
		"REF_QosTrafficSelectorGroup",
	}
}

// QosApplicationSelector is an Sophos Endpoint subType and implements sophos.RestObject
type QosApplicationSelector []interface{}

var _ sophos.RestObject = &QosApplicationSelector{}

// GetPath implements sophos.RestObject and returns the QosApplicationSelector GET path
// Returns all available qos/application_selector objects
func (*QosApplicationSelector) GetPath() string { return "/api/objects/qos/application_selector/" }

// RefRequired implements sophos.RestObject
func (*QosApplicationSelector) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the QosApplicationSelector DELETE path
// Creates or updates the complete object application_selector
func (*QosApplicationSelector) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/application_selector/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosApplicationSelector PATCH path
// Changes to parts of the object application_selector types
func (*QosApplicationSelector) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/application_selector/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosApplicationSelector POST path
// Create a new qos/application_selector object
func (*QosApplicationSelector) PostPath() string {
	return "/api/objects/qos/application_selector/"
}

// PutPath implements sophos.RestObject and returns the QosApplicationSelector PUT path
// Creates or updates the complete object application_selector
func (*QosApplicationSelector) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/application_selector/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosApplicationSelector) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/application_selector/%s/usedby", ref)
}

// QosGroup is an Sophos Endpoint subType and implements sophos.RestObject
type QosGroup []interface{}

var _ sophos.RestObject = &QosGroup{}

// GetPath implements sophos.RestObject and returns the QosGroup GET path
// Returns all available qos/group objects
func (*QosGroup) GetPath() string { return "/api/objects/qos/group/" }

// RefRequired implements sophos.RestObject
func (*QosGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the QosGroup DELETE path
// Creates or updates the complete object group
func (*QosGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosGroup PATCH path
// Changes to parts of the object group types
func (*QosGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosGroup POST path
// Create a new qos/group object
func (*QosGroup) PostPath() string {
	return "/api/objects/qos/group/"
}

// PutPath implements sophos.RestObject and returns the QosGroup PUT path
// Creates or updates the complete object group
func (*QosGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/group/%s/usedby", ref)
}

// QosIngressRule is an Sophos Endpoint subType and implements sophos.RestObject
type QosIngressRule []interface{}

var _ sophos.RestObject = &QosIngressRule{}

// GetPath implements sophos.RestObject and returns the QosIngressRule GET path
// Returns all available qos/ingress_rule objects
func (*QosIngressRule) GetPath() string { return "/api/objects/qos/ingress_rule/" }

// RefRequired implements sophos.RestObject
func (*QosIngressRule) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the QosIngressRule DELETE path
// Creates or updates the complete object ingress_rule
func (*QosIngressRule) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/ingress_rule/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosIngressRule PATCH path
// Changes to parts of the object ingress_rule types
func (*QosIngressRule) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/ingress_rule/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosIngressRule POST path
// Create a new qos/ingress_rule object
func (*QosIngressRule) PostPath() string {
	return "/api/objects/qos/ingress_rule/"
}

// PutPath implements sophos.RestObject and returns the QosIngressRule PUT path
// Creates or updates the complete object ingress_rule
func (*QosIngressRule) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/ingress_rule/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosIngressRule) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/ingress_rule/%s/usedby", ref)
}

// QosInterfaces is an Sophos Endpoint subType and implements sophos.RestObject
type QosInterfaces []QosInterface

// QosInterface is a generated Sophos object
type QosInterface struct {
	Locked            string        `json:"_locked"`
	Reference         string        `json:"_ref"`
	_type             string        `json:"_type"`
	Comment           string        `json:"comment"`
	Downlink          int64         `json:"downlink"`
	DownlinkOptimizer bool          `json:"downlink_optimizer"`
	IngressRules      []interface{} `json:"ingress_rules"`
	Interface         string        `json:"interface"`
	Name              string        `json:"name"`
	Rules             []interface{} `json:"rules"`
	Status            bool          `json:"status"`
	Uplink            int64         `json:"uplink"`
	UplinkLimit       bool          `json:"uplink_limit"`
	UplinkOptimizer   bool          `json:"uplink_optimizer"`
}

var _ sophos.RestGetter = &QosInterface{}

// GetPath implements sophos.RestObject and returns the QosInterfaces GET path
// Returns all available qos/interface objects
func (*QosInterfaces) GetPath() string { return "/api/objects/qos/interface/" }

// RefRequired implements sophos.RestObject
func (*QosInterfaces) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the QosInterfaces GET path
// Returns all available interface types
func (q *QosInterface) GetPath() string {
	return fmt.Sprintf("/api/objects/qos/interface/%s", q.Reference)
}

// RefRequired implements sophos.RestObject
func (q *QosInterface) RefRequired() (string, bool) { return q.Reference, true }

// DeletePath implements sophos.RestObject and returns the QosInterface DELETE path
// Creates or updates the complete object interface
func (*QosInterface) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/interface/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosInterface PATCH path
// Changes to parts of the object interface types
func (*QosInterface) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/interface/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosInterface POST path
// Create a new qos/interface object
func (*QosInterface) PostPath() string {
	return "/api/objects/qos/interface/"
}

// PutPath implements sophos.RestObject and returns the QosInterface PUT path
// Creates or updates the complete object interface
func (*QosInterface) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/interface/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosInterface) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/interface/%s/usedby", ref)
}

// GetType implements sophos.Object
func (q *QosInterface) GetType() string { return q._type }

// QosRule is an Sophos Endpoint subType and implements sophos.RestObject
type QosRule []interface{}

var _ sophos.RestObject = &QosRule{}

// GetPath implements sophos.RestObject and returns the QosRule GET path
// Returns all available qos/rule objects
func (*QosRule) GetPath() string { return "/api/objects/qos/rule/" }

// RefRequired implements sophos.RestObject
func (*QosRule) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the QosRule DELETE path
// Creates or updates the complete object rule
func (*QosRule) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/rule/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosRule PATCH path
// Changes to parts of the object rule types
func (*QosRule) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/rule/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosRule POST path
// Create a new qos/rule object
func (*QosRule) PostPath() string {
	return "/api/objects/qos/rule/"
}

// PutPath implements sophos.RestObject and returns the QosRule PUT path
// Creates or updates the complete object rule
func (*QosRule) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/rule/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosRule) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/rule/%s/usedby", ref)
}

// QosTrafficSelector is an Sophos Endpoint subType and implements sophos.RestObject
type QosTrafficSelector []interface{}

var _ sophos.RestObject = &QosTrafficSelector{}

// GetPath implements sophos.RestObject and returns the QosTrafficSelector GET path
// Returns all available qos/traffic_selector objects
func (*QosTrafficSelector) GetPath() string { return "/api/objects/qos/traffic_selector/" }

// RefRequired implements sophos.RestObject
func (*QosTrafficSelector) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the QosTrafficSelector DELETE path
// Creates or updates the complete object traffic_selector
func (*QosTrafficSelector) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosTrafficSelector PATCH path
// Changes to parts of the object traffic_selector types
func (*QosTrafficSelector) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosTrafficSelector POST path
// Create a new qos/traffic_selector object
func (*QosTrafficSelector) PostPath() string {
	return "/api/objects/qos/traffic_selector/"
}

// PutPath implements sophos.RestObject and returns the QosTrafficSelector PUT path
// Creates or updates the complete object traffic_selector
func (*QosTrafficSelector) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosTrafficSelector) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector/%s/usedby", ref)
}

// QosTrafficSelectorGroup is an Sophos Endpoint subType and implements sophos.RestObject
type QosTrafficSelectorGroup []interface{}

var _ sophos.RestObject = &QosTrafficSelectorGroup{}

// GetPath implements sophos.RestObject and returns the QosTrafficSelectorGroup GET path
// Returns all available qos/traffic_selector_group objects
func (*QosTrafficSelectorGroup) GetPath() string { return "/api/objects/qos/traffic_selector_group/" }

// RefRequired implements sophos.RestObject
func (*QosTrafficSelectorGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the QosTrafficSelectorGroup DELETE path
// Creates or updates the complete object traffic_selector_group
func (*QosTrafficSelectorGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector_group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the QosTrafficSelectorGroup PATCH path
// Changes to parts of the object traffic_selector_group types
func (*QosTrafficSelectorGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector_group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the QosTrafficSelectorGroup POST path
// Create a new qos/traffic_selector_group object
func (*QosTrafficSelectorGroup) PostPath() string {
	return "/api/objects/qos/traffic_selector_group/"
}

// PutPath implements sophos.RestObject and returns the QosTrafficSelectorGroup PUT path
// Creates or updates the complete object traffic_selector_group
func (*QosTrafficSelectorGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector_group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*QosTrafficSelectorGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/qos/traffic_selector_group/%s/usedby", ref)
}
