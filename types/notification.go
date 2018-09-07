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
	NotificationNotification NotificationNotification `json:"notification_notification"`
	NotificationGroup        NotificationGroup        `json:"notification_group"`
}

var defsNotification = map[string]sophos.RestObject{
	"NotificationNotification": &NotificationNotification{},
	"NotificationGroup":        &NotificationGroup{},
}

// RestObjects implements the sophos.Node interface and returns a map of Notification's Objects
func (Notification) RestObjects() map[string]sophos.RestObject {
	return defsNotification
}

// GetPath implements sophos.RestGetter
func (*Notification) GetPath() string { return "/api/nodes/notification" }

// RefRequired implements sophos.RestGetter
func (*Notification) RefRequired() (string, bool) { return "", false }

var defNotification = &sophos.Definition{Description: "notification", Name: "notification", Link: "/api/definitions/notification"}

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
