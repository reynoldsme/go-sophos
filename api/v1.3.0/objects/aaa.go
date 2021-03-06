package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Aaa is a generated struct representing the Sophos Aaa Endpoint
// GET /api/nodes/aaa
type Aaa struct {
	AaaGroup AaaGroup `json:"aaa_group"`
	AaaUser  AaaUser  `json:"aaa_user"`
}

var _ sophos.Endpoint = &Aaa{}

var defsAaa = map[string]sophos.RestObject{
	"AaaGroup": &AaaGroup{},
	"AaaUser":  &AaaUser{},
}

// RestObjects implements the sophos.Node interface and returns a map of Aaa's Objects
func (Aaa) RestObjects() map[string]sophos.RestObject { return defsAaa }

// GetPath implements sophos.RestGetter
func (*Aaa) GetPath() string { return "/api/nodes/aaa" }

// RefRequired implements sophos.RestGetter
func (*Aaa) RefRequired() (string, bool) { return "", false }

var defAaa = &sophos.Definition{Description: "aaa", Name: "aaa", Link: "/api/definitions/aaa"}

// Definition returns the /api/definitions struct of Aaa
func (Aaa) Definition() sophos.Definition { return *defAaa }

// ApiRoutes returns all known Aaa Paths
func (Aaa) ApiRoutes() []string {
	return []string{
		"/api/objects/aaa/group/",
		"/api/objects/aaa/group/{ref}",
		"/api/objects/aaa/group/{ref}/usedby",
		"/api/objects/aaa/user/",
		"/api/objects/aaa/user/{ref}",
		"/api/objects/aaa/user/{ref}/usedby",
	}
}

// References returns the Aaa's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Aaa) References() []string {
	return []string{
		"REF_AaaGroup",
		"REF_AaaUser",
	}
}

// AaaGroups is an Sophos Endpoint subType and implements sophos.RestObject
type AaaGroups []AaaGroup

// AaaGroup is a generated Sophos object
type AaaGroup struct {
	Locked               string        `json:"_locked"`
	Reference            string        `json:"_ref"`
	ObjectType           string        `json:"_type"`
	AdirectoryGroups     []interface{} `json:"adirectory_groups"`
	AdirectoryGroupsSids struct{}      `json:"adirectory_groups_sids"`
	BackendMatch         string        `json:"backend_match"`
	Comment              string        `json:"comment"`
	Dynamic              string        `json:"dynamic"`
	EdirectoryGroups     []interface{} `json:"edirectory_groups"`
	IpsecDn              string        `json:"ipsec_dn"`
	LdapAttribute        string        `json:"ldap_attribute"`
	LdapAttributeValue   string        `json:"ldap_attribute_value"`
	Members              []string      `json:"members"`
	Name                 string        `json:"name"`
	Network              string        `json:"network"`
	RadiusGroups         []interface{} `json:"radius_groups"`
	TacacsGroups         []interface{} `json:"tacacs_groups"`
}

var _ sophos.RestGetter = &AaaGroup{}

// GetPath implements sophos.RestObject and returns the AaaGroups GET path
// Returns all available aaa/group objects
func (*AaaGroups) GetPath() string { return "/api/objects/aaa/group/" }

// RefRequired implements sophos.RestObject
func (*AaaGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AaaGroups GET path
// Returns all available group types
func (a *AaaGroup) GetPath() string { return fmt.Sprintf("/api/objects/aaa/group/%s", a.Reference) }

// RefRequired implements sophos.RestObject
func (a *AaaGroup) RefRequired() (string, bool) { return a.Reference, true }

// DeletePath implements sophos.RestObject and returns the AaaGroup DELETE path
// Creates or updates the complete object group
func (*AaaGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/aaa/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AaaGroup PATCH path
// Changes to parts of the object group types
func (*AaaGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/aaa/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AaaGroup POST path
// Create a new aaa/group object
func (*AaaGroup) PostPath() string {
	return "/api/objects/aaa/group/"
}

// PutPath implements sophos.RestObject and returns the AaaGroup PUT path
// Creates or updates the complete object group
func (*AaaGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/aaa/group/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AaaGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/aaa/group/%s/usedby", ref)
}

// GetType implements sophos.Object
func (a *AaaGroup) GetType() string { return a.ObjectType }

// AaaUsers is an Sophos Endpoint subType and implements sophos.RestObject
type AaaUsers []AaaUser

// AaaUser is a generated Sophos object
type AaaUser struct {
	Locked           string        `json:"_locked"`
	Reference        string        `json:"_ref"`
	ObjectType       string        `json:"_type"`
	AccManaged       bool          `json:"acc_managed"`
	AllowedNetworks  []string      `json:"allowed_networks"`
	Authentication   string        `json:"authentication"`
	BackendUpdate    bool          `json:"backend_update"`
	Clearpass        string        `json:"clearpass"`
	Comment          string        `json:"comment"`
	EmailPrimary     string        `json:"email_primary"`
	EmailSecondary   []interface{} `json:"email_secondary"`
	Enabled          bool          `json:"enabled"`
	LastauthBackend  string        `json:"lastauth_backend"`
	LastauthFacility string        `json:"lastauth_facility"`
	LastauthTime     int64         `json:"lastauth_time"`
	Loc              string        `json:"loc"`
	Md4hash          string        `json:"md4hash"`
	Name             string        `json:"name"`
	Network          string        `json:"network"`
	Pop3Accounts     []interface{} `json:"pop3_accounts"`
	RasIP            string        `json:"ras_ip"`
	RasOnline        bool          `json:"ras_online"`
	Realname         string        `json:"realname"`
	SenderBlacklist  []interface{} `json:"sender_blacklist"`
	SenderWhitelist  []interface{} `json:"sender_whitelist"`
	Status           bool          `json:"status"`
	UseRasIP         bool          `json:"use_ras_ip"`
	UserPreferences  string        `json:"user_preferences"`
	X509Cert         string        `json:"x509_cert"`
	X509CertGost     string        `json:"x509_cert_gost"`
}

var _ sophos.RestGetter = &AaaUser{}

// GetPath implements sophos.RestObject and returns the AaaUsers GET path
// Returns all available aaa/user objects
func (*AaaUsers) GetPath() string { return "/api/objects/aaa/user/" }

// RefRequired implements sophos.RestObject
func (*AaaUsers) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the AaaUsers GET path
// Returns all available user types
func (a *AaaUser) GetPath() string { return fmt.Sprintf("/api/objects/aaa/user/%s", a.Reference) }

// RefRequired implements sophos.RestObject
func (a *AaaUser) RefRequired() (string, bool) { return a.Reference, true }

// DeletePath implements sophos.RestObject and returns the AaaUser DELETE path
// Creates or updates the complete object user
func (*AaaUser) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/aaa/user/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the AaaUser PATCH path
// Changes to parts of the object user types
func (*AaaUser) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/aaa/user/%s", ref)
}

// PostPath implements sophos.RestObject and returns the AaaUser POST path
// Create a new aaa/user object
func (*AaaUser) PostPath() string {
	return "/api/objects/aaa/user/"
}

// PutPath implements sophos.RestObject and returns the AaaUser PUT path
// Creates or updates the complete object user
func (*AaaUser) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/aaa/user/%s", ref)
}

// UsedByPath implements sophos.RestObject
// Returns the objects and the nodes that use the object with the given ref
func (*AaaUser) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/aaa/user/%s/usedby", ref)
}

// GetType implements sophos.Object
func (a *AaaUser) GetType() string { return a.ObjectType }
