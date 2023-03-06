// Code generated - EDITING IS FUTILE. DO NOT EDIT.
//
// Generated by:
//     public/app/plugins/gen.go
// Using jennies:
//     PluginGoTypesJenny
//
// Run 'make gen-cue' from repository root to regenerate.

package dataquery

// Defines values for PhlareQueryType.
const (
	PhlareQueryTypeBoth    PhlareQueryType = "both"
	PhlareQueryTypeMetrics PhlareQueryType = "metrics"
	PhlareQueryTypeProfile PhlareQueryType = "profile"
)

// PhlareDataQuery defines model for PhlareDataQuery.
type PhlareDataQuery struct {
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource *interface{} `json:"datasource,omitempty"`

	// Allows to group the results.
	GroupBy []string `json:"groupBy"`

	// Hide true if query is disabled (ie should not be returned to the dashboard)
	Hide *bool `json:"hide,omitempty"`

	// Unique, guid like, string used in explore mode
	Key *string `json:"key,omitempty"`

	// Specifies the query label selectors.
	LabelSelector string `json:"labelSelector"`

	// Specifies the type of profile to query.
	ProfileTypeId string `json:"profileTypeId"`

	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`

	// A - Z
	RefId string `json:"refId"`
}

// PhlareQueryType defines model for PhlareQueryType.
type PhlareQueryType string
