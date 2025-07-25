// Package capability provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.5.0 DO NOT EDIT.
package capability

// Defines values for CapabilityEntityState.
const (
	Declaration CapabilityEntityState = "declaration"
	Instance    CapabilityEntityState = "instance"
)

// Defines values for CapabilityStatus.
const (
	Disabled CapabilityStatus = "disabled"
	Enabled  CapabilityStatus = "enabled"
)

// Capability Meshery manages entities in accordance with their specific capabilities. This field explicitly identifies those capabilities largely by what actions a given component supports; e.g. metric-scrape, sub-interface, and so on. This field is extensible. Entities may define a broad array of capabilities, which are in-turn dynamically interpretted by Meshery for full lifecycle management.
type Capability struct {
	// Description A written representation of the purpose and characteristics of the capability.
	Description string `json:"description" yaml:"description"`

	// DisplayName Name of the capability in human-readible format.
	DisplayName string `json:"displayName" yaml:"displayName"`

	// EntityState State of the entity in which the capability is applicable.
	EntityState []CapabilityEntityState `json:"entityState" yaml:"entityState"`

	// Key Key that backs the capability.
	Key string `json:"key" yaml:"key"`

	// Kind Top-level categorization of the capability
	Kind string `json:"kind" yaml:"kind"`

	// Metadata Metadata contains additional information associated with the capability. Extension point.
	Metadata *map[string]interface{} `json:"metadata,omitempty" yaml:"metadata,omitempty"`

	// SchemaVersion Specifies the version of the schema to which the capability definition conforms.
	SchemaVersion string `json:"schemaVersion" yaml:"schemaVersion"`

	// Status Status of the capability
	Status CapabilityStatus `json:"status" yaml:"status"`

	// SubType Most granular unit of capability classification. The combination of Kind, Type and SubType together uniquely identify a Capability.
	SubType string `json:"subType" yaml:"subType"`

	// Type Classification of capabilities. Used to group capabilities similar in nature.
	Type string `json:"type" yaml:"type"`

	// Version Version of the capability definition.
	Version string `json:"version" yaml:"version"`
}

// CapabilityEntityState A string starting with an alphanumeric character. Spaces and hyphens allowed.
type CapabilityEntityState string

// CapabilityStatus Status of the capability
type CapabilityStatus string
