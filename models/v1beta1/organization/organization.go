// Package organization provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.5.0 DO NOT EDIT.
package organization

import (
	"time"

	"database/sql"

	"github.com/gofrs/uuid"
)

// DashboardPrefs Preferences specific to dashboard behavior
type DashboardPrefs map[string]interface{}

// Location defines model for Location.
type Location struct {
	Location string `json:"location" yaml:"location"`
	Svg      string `json:"svg" yaml:"svg"`
}

// Logo defines model for Logo.
type Logo struct {
	DarkDesktopView Location `json:"dark_desktop_view" yaml:"dark_desktop_view"`
	DarkMobileView  Location `json:"dark_mobile_view" yaml:"dark_mobile_view"`
	DesktopView     Location `json:"desktop_view" yaml:"desktop_view"`
	MobileView      Location `json:"mobile_view" yaml:"mobile_view"`
}

// NullableTime defines model for NullableTime.
type NullableTime = sql.NullTime

// OrgMetadata defines model for OrgMetadata.
type OrgMetadata struct {
	Preferences Preferences `json:"preferences" yaml:"preferences"`
}

// Organization defines model for Organization.
type Organization struct {
	Country     string       `db:"country" json:"country" yaml:"country"`
	CreatedAt   time.Time    `db:"created_at" json:"created_at" yaml:"created_at"`
	DeletedAt   sql.NullTime `db:"deleted_at" json:"deleted_at,omitempty" yaml:"deleted_at,omitempty"`
	Description string       `db:"description" json:"description" yaml:"description"`
	Domain      *string      `db:"domain" json:"domain" yaml:"domain"`

	// Id A Universally Unique Identifier used to uniquely identify entities in Meshery. The UUID core definition is used across different schemas.
	Id       uuid.UUID   `db:"id" json:"id" yaml:"id"`
	Metadata OrgMetadata `db:"metadata" json:"metadata" yaml:"metadata"`
	Name     string      `db:"name" json:"name" yaml:"name"`

	// Owner A Universally Unique Identifier used to uniquely identify entities in Meshery. The UUID core definition is used across different schemas.
	Owner     uuid.UUID `db:"owner" json:"owner" yaml:"owner"`
	Region    string    `db:"region" json:"region" yaml:"region"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at" yaml:"updated_at"`
}

// Preferences defines model for Preferences.
type Preferences struct {
	// Dashboard Preferences specific to dashboard behavior
	Dashboard DashboardPrefs `json:"dashboard" yaml:"dashboard"`
	Theme     Theme          `json:"theme" yaml:"theme"`
}

// Theme defines model for Theme.
type Theme struct {
	Id   string                  `json:"id" yaml:"id"`
	Logo Logo                    `json:"logo" yaml:"logo"`
	Vars *map[string]interface{} `json:"vars,omitempty" yaml:"vars,omitempty"`
}

// Time defines model for Time.
type Time = time.Time

// UUID A Universally Unique Identifier used to uniquely identify entities in Meshery. The UUID core definition is used across different schemas.
type UUID = uuid.UUID
