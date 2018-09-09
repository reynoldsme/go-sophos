package objects

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Service is a generated struct representing the Sophos Service Endpoint
// GET /api/nodes/service
type Service struct {
	ServiceAh     ServiceAh     `json:"service_ah"`
	ServiceAny    ServiceAny    `json:"service_any"`
	ServiceEsp    ServiceEsp    `json:"service_esp"`
	ServiceGroup  ServiceGroup  `json:"service_group"`
	ServiceIcmp   ServiceIcmp   `json:"service_icmp"`
	ServiceIcmpv6 ServiceIcmpv6 `json:"service_icmpv6"`
	ServiceIp     ServiceIp     `json:"service_ip"`
	ServiceTcp    ServiceTcp    `json:"service_tcp"`
	ServiceTcpudp ServiceTcpudp `json:"service_tcpudp"`
	ServiceUdp    ServiceUdp    `json:"service_udp"`
}

var defsService = map[string]sophos.RestObject{
	"ServiceAh":     &ServiceAh{},
	"ServiceAny":    &ServiceAny{},
	"ServiceEsp":    &ServiceEsp{},
	"ServiceGroup":  &ServiceGroup{},
	"ServiceIcmp":   &ServiceIcmp{},
	"ServiceIcmpv6": &ServiceIcmpv6{},
	"ServiceIp":     &ServiceIp{},
	"ServiceTcp":    &ServiceTcp{},
	"ServiceTcpudp": &ServiceTcpudp{},
	"ServiceUdp":    &ServiceUdp{},
}

// RestObjects implements the sophos.Node interface and returns a map of Service's Objects
func (Service) RestObjects() map[string]sophos.RestObject { return defsService }

// GetPath implements sophos.RestGetter
func (*Service) GetPath() string { return "/api/nodes/service" }

// RefRequired implements sophos.RestGetter
func (*Service) RefRequired() (string, bool) { return "", false }

var defService = &sophos.Definition{Description: "service", Name: "service", Link: "/api/definitions/service"}

// Definition returns the /api/definitions struct of Service
func (Service) Definition() sophos.Definition { return *defService }

// ApiRoutes returns all known Service Paths
func (Service) ApiRoutes() []string {
	return []string{
		"/api/objects/service/ah/",
		"/api/objects/service/ah/{ref}",
		"/api/objects/service/ah/{ref}/usedby",
		"/api/objects/service/any/",
		"/api/objects/service/any/{ref}",
		"/api/objects/service/any/{ref}/usedby",
		"/api/objects/service/esp/",
		"/api/objects/service/esp/{ref}",
		"/api/objects/service/esp/{ref}/usedby",
		"/api/objects/service/group/",
		"/api/objects/service/group/{ref}",
		"/api/objects/service/group/{ref}/usedby",
		"/api/objects/service/icmp/",
		"/api/objects/service/icmp/{ref}",
		"/api/objects/service/icmp/{ref}/usedby",
		"/api/objects/service/icmpv6/",
		"/api/objects/service/icmpv6/{ref}",
		"/api/objects/service/icmpv6/{ref}/usedby",
		"/api/objects/service/ip/",
		"/api/objects/service/ip/{ref}",
		"/api/objects/service/ip/{ref}/usedby",
		"/api/objects/service/tcp/",
		"/api/objects/service/tcp/{ref}",
		"/api/objects/service/tcp/{ref}/usedby",
		"/api/objects/service/tcpudp/",
		"/api/objects/service/tcpudp/{ref}",
		"/api/objects/service/tcpudp/{ref}/usedby",
		"/api/objects/service/udp/",
		"/api/objects/service/udp/{ref}",
		"/api/objects/service/udp/{ref}/usedby",
	}
}

// References returns the Service's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Service) References() []string {
	return []string{
		"REF_ServiceAh",
		"REF_ServiceAny",
		"REF_ServiceEsp",
		"REF_ServiceGroup",
		"REF_ServiceIcmp",
		"REF_ServiceIcmpv6",
		"REF_ServiceIp",
		"REF_ServiceTcp",
		"REF_ServiceTcpudp",
		"REF_ServiceUdp",
	}
}

// ServiceAh is an Sophos Endpoint subType and implements sophos.RestObject
type ServiceAh []interface{}

// GetPath implements sophos.RestObject and returns the ServiceAh GET path
// Returns all available service/ah objects
func (*ServiceAh) GetPath() string { return "/api/objects/service/ah/" }

// RefRequired implements sophos.RestObject
func (*ServiceAh) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ServiceAh DELETE path
// Creates or updates the complete object ah
func (*ServiceAh) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/service/ah/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ServiceAh PATCH path
// Changes to parts of the object ah types
func (*ServiceAh) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/ah/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ServiceAh POST path
// Create a new service/ah object
func (*ServiceAh) PostPath() string {
	return "/api/objects/service/ah/"
}

// PutPath implements sophos.RestObject and returns the ServiceAh PUT path
// Creates or updates the complete object ah
func (*ServiceAh) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/ah/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ServiceAh) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/ah/%s/usedby", ref)
}

// ServiceAnys is an Sophos Endpoint subType and implements sophos.RestObject
type ServiceAnys []ServiceAny

// ServiceAny is a generated Sophos object
type ServiceAny struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
}

// GetPath implements sophos.RestObject and returns the ServiceAnys GET path
// Returns all available service/any objects
func (*ServiceAnys) GetPath() string { return "/api/objects/service/any/" }

// RefRequired implements sophos.RestObject
func (*ServiceAnys) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ServiceAnys GET path
// Returns all available any types
func (s *ServiceAny) GetPath() string { return fmt.Sprintf("/api/objects/service/any/%s", s.Reference) }

// RefRequired implements sophos.RestObject
func (s *ServiceAny) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the ServiceAny DELETE path
// Creates or updates the complete object any
func (*ServiceAny) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/service/any/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ServiceAny PATCH path
// Changes to parts of the object any types
func (*ServiceAny) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/any/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ServiceAny POST path
// Create a new service/any object
func (*ServiceAny) PostPath() string {
	return "/api/objects/service/any/"
}

// PutPath implements sophos.RestObject and returns the ServiceAny PUT path
// Creates or updates the complete object any
func (*ServiceAny) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/any/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ServiceAny) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/any/%s/usedby", ref)
}

// GetType implements sophos.Object
func (s *ServiceAny) GetType() string { return s._type }

// ServiceEsp is an Sophos Endpoint subType and implements sophos.RestObject
type ServiceEsp []interface{}

// GetPath implements sophos.RestObject and returns the ServiceEsp GET path
// Returns all available service/esp objects
func (*ServiceEsp) GetPath() string { return "/api/objects/service/esp/" }

// RefRequired implements sophos.RestObject
func (*ServiceEsp) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ServiceEsp DELETE path
// Creates or updates the complete object esp
func (*ServiceEsp) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/service/esp/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ServiceEsp PATCH path
// Changes to parts of the object esp types
func (*ServiceEsp) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/esp/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ServiceEsp POST path
// Create a new service/esp object
func (*ServiceEsp) PostPath() string {
	return "/api/objects/service/esp/"
}

// PutPath implements sophos.RestObject and returns the ServiceEsp PUT path
// Creates or updates the complete object esp
func (*ServiceEsp) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/esp/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ServiceEsp) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/esp/%s/usedby", ref)
}

// ServiceGroups is an Sophos Endpoint subType and implements sophos.RestObject
type ServiceGroups []ServiceGroup

// ServiceGroup is a generated Sophos object
type ServiceGroup struct {
	Locked    string   `json:"_locked"`
	Reference string   `json:"_ref"`
	_type     string   `json:"_type"`
	Comment   string   `json:"comment"`
	Members   []string `json:"members"`
	Name      string   `json:"name"`
	Types     []string `json:"types"`
}

// GetPath implements sophos.RestObject and returns the ServiceGroups GET path
// Returns all available service/group objects
func (*ServiceGroups) GetPath() string { return "/api/objects/service/group/" }

// RefRequired implements sophos.RestObject
func (*ServiceGroups) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ServiceGroups GET path
// Returns all available group types
func (s *ServiceGroup) GetPath() string {
	return fmt.Sprintf("/api/objects/service/group/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *ServiceGroup) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the ServiceGroup DELETE path
// Creates or updates the complete object group
func (*ServiceGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/service/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ServiceGroup PATCH path
// Changes to parts of the object group types
func (*ServiceGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ServiceGroup POST path
// Create a new service/group object
func (*ServiceGroup) PostPath() string {
	return "/api/objects/service/group/"
}

// PutPath implements sophos.RestObject and returns the ServiceGroup PUT path
// Creates or updates the complete object group
func (*ServiceGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/group/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ServiceGroup) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/group/%s/usedby", ref)
}

// GetType implements sophos.Object
func (s *ServiceGroup) GetType() string { return s._type }

// ServiceIcmps is an Sophos Endpoint subType and implements sophos.RestObject
type ServiceIcmps []ServiceIcmp

// ServiceIcmp is a generated Sophos object
type ServiceIcmp struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Code      int64  `json:"code"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
	Type      int64  `json:"type"`
}

// GetPath implements sophos.RestObject and returns the ServiceIcmps GET path
// Returns all available service/icmp objects
func (*ServiceIcmps) GetPath() string { return "/api/objects/service/icmp/" }

// RefRequired implements sophos.RestObject
func (*ServiceIcmps) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ServiceIcmps GET path
// Returns all available icmp types
func (s *ServiceIcmp) GetPath() string {
	return fmt.Sprintf("/api/objects/service/icmp/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *ServiceIcmp) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the ServiceIcmp DELETE path
// Creates or updates the complete object icmp
func (*ServiceIcmp) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/service/icmp/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ServiceIcmp PATCH path
// Changes to parts of the object icmp types
func (*ServiceIcmp) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/icmp/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ServiceIcmp POST path
// Create a new service/icmp object
func (*ServiceIcmp) PostPath() string {
	return "/api/objects/service/icmp/"
}

// PutPath implements sophos.RestObject and returns the ServiceIcmp PUT path
// Creates or updates the complete object icmp
func (*ServiceIcmp) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/icmp/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ServiceIcmp) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/icmp/%s/usedby", ref)
}

// GetType implements sophos.Object
func (s *ServiceIcmp) GetType() string { return s._type }

// ServiceIcmpv6 is an Sophos Endpoint subType and implements sophos.RestObject
type ServiceIcmpv6 []interface{}

// GetPath implements sophos.RestObject and returns the ServiceIcmpv6 GET path
// Returns all available service/icmpv6 objects
func (*ServiceIcmpv6) GetPath() string { return "/api/objects/service/icmpv6/" }

// RefRequired implements sophos.RestObject
func (*ServiceIcmpv6) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the ServiceIcmpv6 DELETE path
// Creates or updates the complete object icmpv6
func (*ServiceIcmpv6) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/service/icmpv6/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ServiceIcmpv6 PATCH path
// Changes to parts of the object icmpv6 types
func (*ServiceIcmpv6) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/icmpv6/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ServiceIcmpv6 POST path
// Create a new service/icmpv6 object
func (*ServiceIcmpv6) PostPath() string {
	return "/api/objects/service/icmpv6/"
}

// PutPath implements sophos.RestObject and returns the ServiceIcmpv6 PUT path
// Creates or updates the complete object icmpv6
func (*ServiceIcmpv6) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/icmpv6/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ServiceIcmpv6) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/icmpv6/%s/usedby", ref)
}

// ServiceIps is an Sophos Endpoint subType and implements sophos.RestObject
type ServiceIps []ServiceIp

// ServiceIp is a generated Sophos object
type ServiceIp struct {
	Locked    string `json:"_locked"`
	Reference string `json:"_ref"`
	_type     string `json:"_type"`
	Comment   string `json:"comment"`
	Name      string `json:"name"`
	Proto     int64  `json:"proto"`
}

// GetPath implements sophos.RestObject and returns the ServiceIps GET path
// Returns all available service/ip objects
func (*ServiceIps) GetPath() string { return "/api/objects/service/ip/" }

// RefRequired implements sophos.RestObject
func (*ServiceIps) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ServiceIps GET path
// Returns all available ip types
func (s *ServiceIp) GetPath() string { return fmt.Sprintf("/api/objects/service/ip/%s", s.Reference) }

// RefRequired implements sophos.RestObject
func (s *ServiceIp) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the ServiceIp DELETE path
// Creates or updates the complete object ip
func (*ServiceIp) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/service/ip/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ServiceIp PATCH path
// Changes to parts of the object ip types
func (*ServiceIp) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/ip/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ServiceIp POST path
// Create a new service/ip object
func (*ServiceIp) PostPath() string {
	return "/api/objects/service/ip/"
}

// PutPath implements sophos.RestObject and returns the ServiceIp PUT path
// Creates or updates the complete object ip
func (*ServiceIp) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/ip/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ServiceIp) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/ip/%s/usedby", ref)
}

// GetType implements sophos.Object
func (s *ServiceIp) GetType() string { return s._type }

// ServiceTcps is an Sophos Endpoint subType and implements sophos.RestObject
type ServiceTcps []ServiceTcp

// ServiceTcp is a generated Sophos object
type ServiceTcp struct {
	Locked       string `json:"_locked"`
	Reference    string `json:"_ref"`
	_type        string `json:"_type"`
	AutoPfSvcDst string `json:"auto_pf_svc_dst"`
	AutoPfSvcSrc string `json:"auto_pf_svc_src"`
	Comment      string `json:"comment"`
	DstHigh      int64  `json:"dst_high"`
	DstLow       int64  `json:"dst_low"`
	Name         string `json:"name"`
	SrcHigh      int64  `json:"src_high"`
	SrcLow       int64  `json:"src_low"`
}

// GetPath implements sophos.RestObject and returns the ServiceTcps GET path
// Returns all available service/tcp objects
func (*ServiceTcps) GetPath() string { return "/api/objects/service/tcp/" }

// RefRequired implements sophos.RestObject
func (*ServiceTcps) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ServiceTcps GET path
// Returns all available tcp types
func (s *ServiceTcp) GetPath() string { return fmt.Sprintf("/api/objects/service/tcp/%s", s.Reference) }

// RefRequired implements sophos.RestObject
func (s *ServiceTcp) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the ServiceTcp DELETE path
// Creates or updates the complete object tcp
func (*ServiceTcp) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/service/tcp/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ServiceTcp PATCH path
// Changes to parts of the object tcp types
func (*ServiceTcp) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/tcp/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ServiceTcp POST path
// Create a new service/tcp object
func (*ServiceTcp) PostPath() string {
	return "/api/objects/service/tcp/"
}

// PutPath implements sophos.RestObject and returns the ServiceTcp PUT path
// Creates or updates the complete object tcp
func (*ServiceTcp) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/tcp/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ServiceTcp) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/tcp/%s/usedby", ref)
}

// GetType implements sophos.Object
func (s *ServiceTcp) GetType() string { return s._type }

// ServiceTcpudps is an Sophos Endpoint subType and implements sophos.RestObject
type ServiceTcpudps []ServiceTcpudp

// ServiceTcpudp is a generated Sophos object
type ServiceTcpudp struct {
	Locked       string `json:"_locked"`
	Reference    string `json:"_ref"`
	_type        string `json:"_type"`
	AutoPfSvcDst string `json:"auto_pf_svc_dst"`
	AutoPfSvcSrc string `json:"auto_pf_svc_src"`
	Comment      string `json:"comment"`
	DstHigh      int64  `json:"dst_high"`
	DstLow       int64  `json:"dst_low"`
	Name         string `json:"name"`
	SrcHigh      int64  `json:"src_high"`
	SrcLow       int64  `json:"src_low"`
}

// GetPath implements sophos.RestObject and returns the ServiceTcpudps GET path
// Returns all available service/tcpudp objects
func (*ServiceTcpudps) GetPath() string { return "/api/objects/service/tcpudp/" }

// RefRequired implements sophos.RestObject
func (*ServiceTcpudps) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ServiceTcpudps GET path
// Returns all available tcpudp types
func (s *ServiceTcpudp) GetPath() string {
	return fmt.Sprintf("/api/objects/service/tcpudp/%s", s.Reference)
}

// RefRequired implements sophos.RestObject
func (s *ServiceTcpudp) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the ServiceTcpudp DELETE path
// Creates or updates the complete object tcpudp
func (*ServiceTcpudp) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/service/tcpudp/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ServiceTcpudp PATCH path
// Changes to parts of the object tcpudp types
func (*ServiceTcpudp) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/tcpudp/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ServiceTcpudp POST path
// Create a new service/tcpudp object
func (*ServiceTcpudp) PostPath() string {
	return "/api/objects/service/tcpudp/"
}

// PutPath implements sophos.RestObject and returns the ServiceTcpudp PUT path
// Creates or updates the complete object tcpudp
func (*ServiceTcpudp) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/tcpudp/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ServiceTcpudp) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/tcpudp/%s/usedby", ref)
}

// GetType implements sophos.Object
func (s *ServiceTcpudp) GetType() string { return s._type }

// ServiceUdps is an Sophos Endpoint subType and implements sophos.RestObject
type ServiceUdps []ServiceUdp

// ServiceUdp is a generated Sophos object
type ServiceUdp struct {
	Locked       string `json:"_locked"`
	Reference    string `json:"_ref"`
	_type        string `json:"_type"`
	AutoPfSvcDst string `json:"auto_pf_svc_dst"`
	AutoPfSvcSrc string `json:"auto_pf_svc_src"`
	Comment      string `json:"comment"`
	DstHigh      int64  `json:"dst_high"`
	DstLow       int64  `json:"dst_low"`
	Name         string `json:"name"`
	SrcHigh      int64  `json:"src_high"`
	SrcLow       int64  `json:"src_low"`
}

// GetPath implements sophos.RestObject and returns the ServiceUdps GET path
// Returns all available service/udp objects
func (*ServiceUdps) GetPath() string { return "/api/objects/service/udp/" }

// RefRequired implements sophos.RestObject
func (*ServiceUdps) RefRequired() (string, bool) { return "", false }

// GetPath implements sophos.RestObject and returns the ServiceUdps GET path
// Returns all available udp types
func (s *ServiceUdp) GetPath() string { return fmt.Sprintf("/api/objects/service/udp/%s", s.Reference) }

// RefRequired implements sophos.RestObject
func (s *ServiceUdp) RefRequired() (string, bool) { return s.Reference, true }

// DeletePath implements sophos.RestObject and returns the ServiceUdp DELETE path
// Creates or updates the complete object udp
func (*ServiceUdp) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/service/udp/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the ServiceUdp PATCH path
// Changes to parts of the object udp types
func (*ServiceUdp) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/udp/%s", ref)
}

// PostPath implements sophos.RestObject and returns the ServiceUdp POST path
// Create a new service/udp object
func (*ServiceUdp) PostPath() string {
	return "/api/objects/service/udp/"
}

// PutPath implements sophos.RestObject and returns the ServiceUdp PUT path
// Creates or updates the complete object udp
func (*ServiceUdp) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/udp/%s", ref)
}

// UsedByPath implements sophos.UsedObject
// Returns the objects and the nodes that use the object with the given ref
func (*ServiceUdp) UsedByPath(ref string) string {
	return fmt.Sprintf("/api/objects/service/udp/%s/usedby", ref)
}

// GetType implements sophos.Object
func (s *ServiceUdp) GetType() string { return s._type }
