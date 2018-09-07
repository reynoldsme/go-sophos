// Package types contains the generated Sophos types
//
// This file was generated by bin/gen.go! DO NOT EDIT!
package types

import (
	"fmt"

	"github.com/esurdam/go-sophos"
)

// Notification is a generated struct representing the Sophos Notification Endpoint
// GET /api/nodes/notification
type Notification struct {
	NotificationGroup        NotificationGroup        `json:"notification_group"`
	NotificationNotification NotificationNotification `json:"notification_notification"`
}

var defsNotification = map[string]sophos.RestObject{
	"NotificationGroup":        &NotificationGroup{},
	"NotificationNotification": &NotificationNotification{},
}

// RestObjects implements the sophos.Node interface and returns a map of Notification's Objects
func (Notification) RestObjects() map[string]sophos.RestObject {
	return defsNotification
}

// GetPath implements sophos.RestGetter
func (*Notification) GetPath() string { return "/api/nodes/notification" }

// RefRequired implements sophos.RestGetter
func (*Notification) RefRequired() (string, bool) { return "", false }

var defNotification = &sophos.Definition{Description: "notification", Name: "notification", Link: "/api/definitions/notification", Swag: map[string]sophos.MethodMap{"/objects/notification/notification/{ref}": {"delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object notification", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"notification/notification"}, Responses: map[int]struct{ Description string }{204: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available notification types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"notification/notification"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}, "patch": sophos.MethodDescriptions{Description: "Changes to parts of the object notification types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "notification/notification that will be changes", Type: "", Required: true}}, Tags: []string{"notification/notification"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object notification", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "notification/notification that will be updated", Type: "", Required: true}}, Tags: []string{"notification/notification"}, Responses: map[int]struct{ Description string }{404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/notification/notification/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"notification/notification"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/notification/group/": {"get": sophos.MethodDescriptions{Description: "Returns all available notification/group objects", Parameters: []sophos.Parameter(nil), Tags: []string{"notification/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "post": sophos.MethodDescriptions{Description: "Create a new notification/group object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "notification/group that will be created", Type: "", Required: true}}, Tags: []string{"notification/group"}, Responses: map[int]struct{ Description string }{201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/notification/group/{ref}": {"patch": sophos.MethodDescriptions{Description: "Changes to parts of the object group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "body", In: "body", Description: "notification/group that will be changes", Type: "", Required: true}}, Tags: []string{"notification/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "put": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "notification/group that will be updated", Type: "", Required: true}}, Tags: []string{"notification/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 404: {Description: "NotFound"}}}, "delete": sophos.MethodDescriptions{Description: "Creates or updates the complete object group", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}, {Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}}, Tags: []string{"notification/group"}, Responses: map[int]struct{ Description string }{401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}, 204: {Description: "OK"}, 400: {Description: "BadRequest"}}}, "get": sophos.MethodDescriptions{Description: "Returns all available group types", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"notification/group"}, Responses: map[int]struct{ Description string }{404: {Description: "NotFound"}, 200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/notification/group/{ref}/usedby": {"get": sophos.MethodDescriptions{Description: "Returns the objects and the nodes that use the object with the given ref", Parameters: []sophos.Parameter{{Name: "ref", In: "path", Description: "id of the object", Type: "string", Required: true}}, Tags: []string{"notification/group"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}}, "/objects/notification/notification/": {"get": sophos.MethodDescriptions{Description: "Returns all available notification/notification objects", Parameters: []sophos.Parameter(nil), Tags: []string{"notification/notification"}, Responses: map[int]struct{ Description string }{200: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}, 403: {Description: "Forbidden"}}}, "post": sophos.MethodDescriptions{Description: "Create a new notification/notification object", Parameters: []sophos.Parameter{{Name: "X-Restd-Err-Ack", In: "header", Description: "Acknowledge confd errors (required for DELETE calls).", Type: "string", Required: false}, {Name: "X-Restd-Lock-Override", In: "header", Description: "Override confd lock (required to perform action on {'_locked': 'user'} objects).", Type: "string", Required: false}, {Name: "X-Restd-Insert", In: "header", Description: "Path and position (optional for arrays, required for hashes, not used for strings) of a node, where to insert the newly created object, format 'node-path [index]', e.g. array: 'packetfilter.rules 2', string: 'ha.aws.cloudwatch.profile', hash: 'auth.api_tokens myToken123'", Type: "string", Required: false}, {Name: "body", In: "body", Description: "notification/notification that will be created", Type: "", Required: true}}, Tags: []string{"notification/notification"}, Responses: map[int]struct{ Description string }{403: {Description: "Forbidden"}, 201: {Description: "OK"}, 400: {Description: "BadRequest"}, 401: {Description: "Unauthorized"}}}}}}

// Definition returns the /api/definitions struct of Notification
func (Notification) Definition() sophos.Definition { return *defNotification }

// ApiRoutes returns all known Notification Paths
func (Notification) ApiRoutes() []string {
	return []string{
		"/api/objects/notification/group/",
		"/api/objects/notification/group/{ref}",
		"/api/objects/notification/group/{ref}/usedby",
		"/api/objects/notification/notification/",
		"/api/objects/notification/notification/{ref}",
		"/api/objects/notification/notification/{ref}/usedby",
	}
}

// References returns the Notification's references. These strings serve no purpose other than to demonstrate which
// Reference keys are used for this Endpoint
func (Notification) References() []string {
	return []string{
		"REF_NotificationGroup",
		"REF_NotificationNotification",
	}
}

// NotificationGroup is an Sophos Endpoint subType and implements sophos.RestObject
type NotificationGroup []interface{}

// GetPath implements sophos.RestObject and returns the NotificationGroup GET path
// Returns all available notification/group objects
func (*NotificationGroup) GetPath() string { return "/api/objects/notification/group/" }

// RefRequired implements sophos.RestObject
func (*NotificationGroup) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the NotificationGroup DELETE path
// Creates or updates the complete object group
func (*NotificationGroup) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/group/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NotificationGroup PATCH path
// Changes to parts of the object group types
func (*NotificationGroup) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/group/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NotificationGroup POST path
// Create a new notification/group object
func (*NotificationGroup) PostPath() string {
	return "/api/objects/notification/group/"
}

// PutPath implements sophos.RestObject and returns the NotificationGroup PUT path
// Creates or updates the complete object group
func (*NotificationGroup) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/group/%s", ref)
}

// NotificationNotification is an Sophos Endpoint subType and implements sophos.RestObject
type NotificationNotification []interface{}

// GetPath implements sophos.RestObject and returns the NotificationNotification GET path
// Returns all available notification/notification objects
func (*NotificationNotification) GetPath() string { return "/api/objects/notification/notification/" }

// RefRequired implements sophos.RestObject
func (*NotificationNotification) RefRequired() (string, bool) { return "", false }

// DeletePath implements sophos.RestObject and returns the NotificationNotification DELETE path
// Creates or updates the complete object notification
func (*NotificationNotification) DeletePath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/notification/%s", ref)
}

// PatchPath implements sophos.RestObject and returns the NotificationNotification PATCH path
// Changes to parts of the object notification types
func (*NotificationNotification) PatchPath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/notification/%s", ref)
}

// PostPath implements sophos.RestObject and returns the NotificationNotification POST path
// Create a new notification/notification object
func (*NotificationNotification) PostPath() string {
	return "/api/objects/notification/notification/"
}

// PutPath implements sophos.RestObject and returns the NotificationNotification PUT path
// Creates or updates the complete object notification
func (*NotificationNotification) PutPath(ref string) string {
	return fmt.Sprintf("/api/objects/notification/notification/%s", ref)
}