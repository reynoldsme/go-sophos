package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Epp is a generated struct representing the Sophos Epp Endpoint
// GET /api/nodes/epp
type Epp struct {
	AllowedNetworks       []string      `json:"allowed_networks"`
	Certificate           string        `json:"certificate"`
	City                  string        `json:"city"`
	Country               string        `json:"country"`
	DefaultEndpointsGroup string        `json:"default_endpoints_group"`
	Devices               []interface{} `json:"devices"`
	Email                 string        `json:"email"`
	Endpoints             []interface{} `json:"endpoints"`
	EndpointsGroups       []string      `json:"endpoints_groups"`
	Exceptions            struct {
		Av []interface{} `json:"av"`
		Dc []interface{} `json:"dc"`
	} `json:"exceptions"`
	FallbackURL       string `json:"fallback_url"`
	MagnetPassword    string `json:"magnet_password"`
	MagnetUsername    string `json:"magnet_username"`
	Organization      string `json:"organization"`
	ParentProxyHost   string `json:"parent_proxy_host"`
	ParentProxyPort   int64  `json:"parent_proxy_port"`
	ParentProxyStatus int64  `json:"parent_proxy_status"`
	Policies          struct {
		Av []string `json:"av"`
		Dc []string `json:"dc"`
	} `json:"policies"`
	Port              int64  `json:"port"`
	PrivateKey        string `json:"private_key"`
	RegistrationToken string `json:"registration_token"`
	Status            struct {
		Av     int64 `json:"av"`
		Broker int64 `json:"broker"`
		Dc     int64 `json:"dc"`
		Epp    int64 `json:"epp"`
		Wc     int64 `json:"wc"`
	} `json:"status"`
	TamperPassword string `json:"tamper_password"`
	Version        string `json:"version"`
	WdxToken       string `json:"wdx_token"`
}

var defsEpp = map[string]sophos.RestObject{
	"EppAvException":    &EppAvException{},
	"EppAvPolicy":       &EppAvPolicy{},
	"EppDcException":    &EppDcException{},
	"EppDcPolicy":       &EppDcPolicy{},
	"EppDevice":         &EppDevice{},
	"EppEndpoint":       &EppEndpoint{},
	"EppEndpointsGroup": &EppEndpointsGroup{},
	"EppGroup":          &EppGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of Epp's Objects
func (Epp) RestObjects() map[string]sophos.RestObject { return defsEpp }

// GetPath implements sophos.RestGetter
func (*Epp) GetPath() string { return "/api/nodes/epp" }

// RefRequired implements sophos.RestGetter
func (*Epp) RefRequired() (string, bool) { return "", false }

var defEpp = &sophos.Definition{Description: "epp", Name: "epp", Link: "/api/definitions/epp"}

// Definition returns the /api/definitions struct of Epp
func (Epp) Definition() sophos.Definition { return *defEpp }

// ApiRoutes returns all known Epp Paths
func (Epp) ApiRoutes() []string {
	return []string{
		"/api/objects/epp/av_exception/",
		"/api/objects/epp/av_exception/{ref}",
		"/api/objects/epp/av_exception/{ref}/usedby",
		"/api/objects/epp/av_policy/",
		"/api/objects/epp/av_policy/{ref}",
		"/api/objects/epp/av_policy/{ref}/usedby",
		"/api/objects/epp/dc_exception/",
		"/api/objects/epp/dc_exception/{ref}",
		"/api/objects/epp/dc_exception/{ref}/usedby",
		"/api/objects/epp/dc_policy/",
		"/api/objects/epp/dc_policy/{ref}",
		"/api/objects/epp/dc_policy/{ref}/usedby",
		"/api/objects/epp/device/",
		"/api/objects/epp/device/{ref}",
		"/api/objects/epp/device/{ref}/usedby",
		"/api/objects/epp/endpoint/",
		"/api/objects/epp/endpoint/{ref}",
		"/api/objects/epp/endpoint/{ref}/usedby",
		"/api/objects/epp/endpoints_group/",
		"/api/objects/epp/endpoints_group/{ref}",
		"/api/objects/epp/endpoints_group/{ref}/usedby",
		"/api/objects/epp/group/",
		"/api/objects/epp/group/{ref}",
		"/api/objects/epp/group/{ref}/usedby",
	}
}

// References returns the Epp's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Epp) References() []string {
	return []string{
		"REF_EppAvException",
		"REF_EppAvPolicy",
		"REF_EppDcException",
		"REF_EppDcPolicy",
		"REF_EppDevice",
		"REF_EppEndpoint",
		"REF_EppEndpointsGroup",
		"REF_EppGroup",
	}
}

// EppAvException is an Sophos Endpoint subType and implements sophos.RestObject
type EppAvException []interface{}

// GetPath implements sophos.RestObject and returns the EppAvException GET path
// Returns all available epp/av_exception objects
func (*EppAvException) GetPath() string { return "/api/objects/epp/av_exception/" }

// RefRequired implements sophos.RestObject
func (*EppAvException) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the EppAvException DELETE path
// Creates or updates the complete object av_exception
func (*EppAvException) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_exception/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppAvException PATCH path
// Changes to parts of the object av_exception types
func (*EppAvException) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_exception/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppAvException POST path
// Create a new epp/av_exception object
func (*EppAvException) PostPath() string {
	return "/api/objects/epp/av_exception/"
}

// PutPath implements sophos.RestObject and returns the EppAvException PUT path
// Creates or updates the complete object av_exception
func (*EppAvException) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_exception/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*EppAvException) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_exception/%s/usedby", ref)
}

// EppAvPolicys is an Sophos Endpoint subType and implements sophos.RestObject
type EppAvPolicys []EppAvPolicy

// EppAvPolicy is a generated Sophos object
type EppAvPolicy struct {
	Locked                    string `json:"_locked"`
	Reference                 string `json:"_ref"`
	_type                     string `json:"_type"`
	AlertOnly                 bool   `json:"alert_only"`
	AutoCleanup               bool   `json:"auto_cleanup"`
	BlockMaliciousSites       bool   `json:"block_malicious_sites"`
	Comment                   string `json:"comment"`
	DetectBufferOverflow      bool   `json:"detect_buffer_overflow"`
	DetectMaliciousFiles      bool   `json:"detect_malicious_files"`
	DetectSuspiciousBehaviour bool   `json:"detect_suspicious_behaviour"`
	DownloadScanning          bool   `json:"download_scanning"`
	Hips                      bool   `json:"hips"`
	LowPriorityScan           bool   `json:"low_priority_scan"`
	Name                      string `json:"name"`
	OnAccessScanning          bool   `json:"on_access_scanning"`
	OnRead                    bool   `json:"on_read"`
	OnRename                  bool   `json:"on_rename"`
	OnWrite                   bool   `json:"on_write"`
	RootKitScan               bool   `json:"root_kit_scan"`
	ScanForPua                bool   `json:"scan_for_pua"`
	ScanForSuspiciousFiles    bool   `json:"scan_for_suspicious_files"`
	ScanInsideArchive         bool   `json:"scan_inside_archive"`
	ScanSystemMemory          bool   `json:"scan_system_memory"`
	ScheduledScanning         bool   `json:"scheduled_scanning"`
	SendSampleFile            bool   `json:"send_sample_file"`
	SophosLiveProtection      bool   `json:"sophos_live_protection"`
	TimeEvent                 string `json:"time_event"`
	WebProtection             bool   `json:"web_protection"`
}

// GetPath implements sophos.RestObject and returns the EppAvPolicys GET path
// Returns all available epp/av_policy objects
func (*EppAvPolicys) GetPath() string { return "/api/objects/epp/av_policy/" }

// RefRequired implements sophos.RestObject
func (*EppAvPolicys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EppAvPolicys GET path
// Returns all available av_policy types
func (e *EppAvPolicy) GetPath() string {
	return fmt.Sprintf("/api/objects/epp/av_policy/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EppAvPolicy) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EppAvPolicy DELETE path
// Creates or updates the complete object av_policy
func (*EppAvPolicy) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_policy/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppAvPolicy PATCH path
// Changes to parts of the object av_policy types
func (*EppAvPolicy) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_policy/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppAvPolicy POST path
// Create a new epp/av_policy object
func (*EppAvPolicy) PostPath() string {
	return "/api/objects/epp/av_policy/"
}

// PutPath implements sophos.RestObject and returns the EppAvPolicy PUT path
// Creates or updates the complete object av_policy
func (*EppAvPolicy) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_policy/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*EppAvPolicy) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/av_policy/%s/usedby", ref)
}

// GetType implements sophos.Object
func (e *EppAvPolicy) GetType() string { return e._type }

// EppDcException is an Sophos Endpoint subType and implements sophos.RestObject
type EppDcException []interface{}

// GetPath implements sophos.RestObject and returns the EppDcException GET path
// Returns all available epp/dc_exception objects
func (*EppDcException) GetPath() string { return "/api/objects/epp/dc_exception/" }

// RefRequired implements sophos.RestObject
func (*EppDcException) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the EppDcException DELETE path
// Creates or updates the complete object dc_exception
func (*EppDcException) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_exception/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppDcException PATCH path
// Changes to parts of the object dc_exception types
func (*EppDcException) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_exception/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppDcException POST path
// Create a new epp/dc_exception object
func (*EppDcException) PostPath() string {
	return "/api/objects/epp/dc_exception/"
}

// PutPath implements sophos.RestObject and returns the EppDcException PUT path
// Creates or updates the complete object dc_exception
func (*EppDcException) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_exception/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*EppDcException) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_exception/%s/usedby", ref)
}

// EppDcPolicys is an Sophos Endpoint subType and implements sophos.RestObject
type EppDcPolicys []EppDcPolicy

// EppDcPolicy is a generated Sophos object
type EppDcPolicy struct {
	Locked           string `json:"_locked"`
	Reference        string `json:"_ref"`
	_type            string `json:"_type"`
	Bluetooth        string `json:"bluetooth"`
	Comment          string `json:"comment"`
	EncryptedStorage string `json:"encrypted_storage"`
	FloppyDrive      string `json:"floppy_drive"`
	Infrared         string `json:"infrared"`
	Modem            string `json:"modem"`
	Name             string `json:"name"`
	OpticalDrive     string `json:"optical_drive"`
	RemovableStorage string `json:"removable_storage"`
	Wireless         string `json:"wireless"`
}

// GetPath implements sophos.RestObject and returns the EppDcPolicys GET path
// Returns all available epp/dc_policy objects
func (*EppDcPolicys) GetPath() string { return "/api/objects/epp/dc_policy/" }

// RefRequired implements sophos.RestObject
func (*EppDcPolicys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EppDcPolicys GET path
// Returns all available dc_policy types
func (e *EppDcPolicy) GetPath() string {
	return fmt.Sprintf("/api/objects/epp/dc_policy/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EppDcPolicy) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EppDcPolicy DELETE path
// Creates or updates the complete object dc_policy
func (*EppDcPolicy) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_policy/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppDcPolicy PATCH path
// Changes to parts of the object dc_policy types
func (*EppDcPolicy) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_policy/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppDcPolicy POST path
// Create a new epp/dc_policy object
func (*EppDcPolicy) PostPath() string {
	return "/api/objects/epp/dc_policy/"
}

// PutPath implements sophos.RestObject and returns the EppDcPolicy PUT path
// Creates or updates the complete object dc_policy
func (*EppDcPolicy) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_policy/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*EppDcPolicy) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/dc_policy/%s/usedby", ref)
}

// GetType implements sophos.Object
func (e *EppDcPolicy) GetType() string { return e._type }

// EppDevice is an Sophos Endpoint subType and implements sophos.RestObject
type EppDevice []interface{}

// GetPath implements sophos.RestObject and returns the EppDevice GET path
// Returns all available epp/device objects
func (*EppDevice) GetPath() string { return "/api/objects/epp/device/" }

// RefRequired implements sophos.RestObject
func (*EppDevice) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the EppDevice DELETE path
// Creates or updates the complete object device
func (*EppDevice) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/device/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppDevice PATCH path
// Changes to parts of the object device types
func (*EppDevice) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/device/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppDevice POST path
// Create a new epp/device object
func (*EppDevice) PostPath() string {
	return "/api/objects/epp/device/"
}

// PutPath implements sophos.RestObject and returns the EppDevice PUT path
// Creates or updates the complete object device
func (*EppDevice) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/device/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*EppDevice) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/device/%s/usedby", ref)
}

// EppEndpoint is an Sophos Endpoint subType and implements sophos.RestObject
type EppEndpoint []interface{}

// GetPath implements sophos.RestObject and returns the EppEndpoint GET path
// Returns all available epp/endpoint objects
func (*EppEndpoint) GetPath() string { return "/api/objects/epp/endpoint/" }

// RefRequired implements sophos.RestObject
func (*EppEndpoint) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the EppEndpoint DELETE path
// Creates or updates the complete object endpoint
func (*EppEndpoint) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoint/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppEndpoint PATCH path
// Changes to parts of the object endpoint types
func (*EppEndpoint) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoint/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppEndpoint POST path
// Create a new epp/endpoint object
func (*EppEndpoint) PostPath() string {
	return "/api/objects/epp/endpoint/"
}

// PutPath implements sophos.RestObject and returns the EppEndpoint PUT path
// Creates or updates the complete object endpoint
func (*EppEndpoint) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoint/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*EppEndpoint) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoint/%s/usedby", ref)
}

// EppEndpointsGroups is an Sophos Endpoint subType and implements sophos.RestObject
type EppEndpointsGroups []EppEndpointsGroup

// EppEndpointsGroup is a generated Sophos object
type EppEndpointsGroup struct {
	Locked           string        `json:"_locked"`
	Reference        string        `json:"_ref"`
	_type            string        `json:"_type"`
	AvPolicy         string        `json:"av_policy"`
	Comment          string        `json:"comment"`
	DcPolicy         string        `json:"dc_policy"`
	Endpoints        []interface{} `json:"endpoints"`
	Name             string        `json:"name"`
	ProxyAddress     string        `json:"proxy_address"`
	ProxyPassword    string        `json:"proxy_password"`
	ProxyPort        int64         `json:"proxy_port"`
	ProxySupport     bool          `json:"proxy_support"`
	ProxyUser        string        `json:"proxy_user"`
	TamperProtection bool          `json:"tamper_protection"`
	WebControl       bool          `json:"web_control"`
}

// GetPath implements sophos.RestObject and returns the EppEndpointsGroups GET path
// Returns all available epp/endpoints_group objects
func (*EppEndpointsGroups) GetPath() string { return "/api/objects/epp/endpoints_group/" }

// RefRequired implements sophos.RestObject
func (*EppEndpointsGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the EppEndpointsGroups GET path
// Returns all available endpoints_group types
func (e *EppEndpointsGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/epp/endpoints_group/%s", e.Reference)
}

// RefRequired implements sophos.RestObject
func (e *EppEndpointsGroup) RefRequired() (string, bool) { return e.Reference, true }

// DeletePath implements sophos.RestObject and returns the EppEndpointsGroup DELETE path
// Creates or updates the complete object endpoints_group
func (*EppEndpointsGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoints_group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppEndpointsGroup PATCH path
// Changes to parts of the object endpoints_group types
func (*EppEndpointsGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoints_group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppEndpointsGroup POST path
// Create a new epp/endpoints_group object
func (*EppEndpointsGroup) PostPath() string {
	return "/api/objects/epp/endpoints_group/"
}

// PutPath implements sophos.RestObject and returns the EppEndpointsGroup PUT path
// Creates or updates the complete object endpoints_group
func (*EppEndpointsGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoints_group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*EppEndpointsGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/endpoints_group/%s/usedby", ref)
}

// GetType implements sophos.Object
func (e *EppEndpointsGroup) GetType() string { return e._type }

// EppGroup is an Sophos Endpoint subType and implements sophos.RestObject
type EppGroup []interface{}

// GetPath implements sophos.RestObject and returns the EppGroup GET path
// Returns all available epp/group objects
func (*EppGroup) GetPath() string { return "/api/objects/epp/group/" }

// RefRequired implements sophos.RestObject
func (*EppGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the EppGroup DELETE path
// Creates or updates the complete object group
func (*EppGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the EppGroup PATCH path
// Changes to parts of the object group types
func (*EppGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the EppGroup POST path
// Create a new epp/group object
func (*EppGroup) PostPath() string {
	return "/api/objects/epp/group/"
}

// PutPath implements sophos.RestObject and returns the EppGroup PUT path
// Creates or updates the complete object group
func (*EppGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/group/%s", ref)
}

// UsedByPath implements sophos.Object
// Returns the objects and the nodes that use the object with the given ref
func (*EppGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/epp/group/%s/usedby", ref)
}