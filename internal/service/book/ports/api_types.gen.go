// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.4 DO NOT EDIT.
package ports

import (
	"time"
)

const (
	BearerAuthScopes = "bearerAuth.Scopes"
)

// Defines values for ContactGender.
const (
	ContactGenderFEMALE ContactGender = "FEMALE"
	ContactGenderMALE   ContactGender = "MALE"
	ContactGenderOTHER  ContactGender = "OTHER"
)

// Defines values for FlowFormat.
const (
	FlowFormatAPIJSON    FlowFormat = "API_JSON"
	FlowFormatAPIXML     FlowFormat = "API_XML"
	FlowFormatPOLIRISCSV FlowFormat = "POLIRIS_CSV"
)

// Defines values for FlowStatus.
const (
	FlowStatusCREATED   FlowStatus = "CREATED"
	FlowStatusDISABLED  FlowStatus = "DISABLED"
	FlowStatusENABLED   FlowStatus = "ENABLED"
	FlowStatusSUSPENDED FlowStatus = "SUSPENDED"
)

// Defines values for FlowTypeOffers.
const (
	FlowTypeOffersANCIEN FlowTypeOffers = "ANCIEN"
	FlowTypeOffersNEUF   FlowTypeOffers = "NEUF"
)

// Defines values for NewContactGender.
const (
	NewContactGenderFEMALE NewContactGender = "FEMALE"
	NewContactGenderMALE   NewContactGender = "MALE"
	NewContactGenderOTHER  NewContactGender = "OTHER"
)

// Defines values for NewFlowFormat.
const (
	NewFlowFormatAPIJSON    NewFlowFormat = "API_JSON"
	NewFlowFormatAPIXML     NewFlowFormat = "API_XML"
	NewFlowFormatPOLIRISCSV NewFlowFormat = "POLIRIS_CSV"
)

// Defines values for NewFlowStatus.
const (
	NewFlowStatusCREATED   NewFlowStatus = "CREATED"
	NewFlowStatusDISABLED  NewFlowStatus = "DISABLED"
	NewFlowStatusENABLED   NewFlowStatus = "ENABLED"
	NewFlowStatusSUSPENDED NewFlowStatus = "SUSPENDED"
)

// Defines values for NewFlowTypeOffers.
const (
	NewFlowTypeOffersANCIEN NewFlowTypeOffers = "ANCIEN"
	NewFlowTypeOffersNEUF   NewFlowTypeOffers = "NEUF"
)

// Defines values for QueryStatus.
const (
	QueryStatusCREATED   QueryStatus = "CREATED"
	QueryStatusDISABLED  QueryStatus = "DISABLED"
	QueryStatusENABLED   QueryStatus = "ENABLED"
	QueryStatusSUSPENDED QueryStatus = "SUSPENDED"
)

// Defines values for FindFlowsOriginsParamsStatus.
const (
	FindFlowsOriginsParamsStatusCREATED   FindFlowsOriginsParamsStatus = "CREATED"
	FindFlowsOriginsParamsStatusDISABLED  FindFlowsOriginsParamsStatus = "DISABLED"
	FindFlowsOriginsParamsStatusENABLED   FindFlowsOriginsParamsStatus = "ENABLED"
	FindFlowsOriginsParamsStatusSUSPENDED FindFlowsOriginsParamsStatus = "SUSPENDED"
)

// ModelError A representation of an errorMessage
type ModelError struct {
	// ErrorMessage error Message
	ErrorMessage *string `json:"errorMessage,omitempty"`
}

// Contact defines model for contact.
type Contact struct {
	// ID Unique id (format UUID) of the Contact
	ID *string `json:"ID,omitempty"`

	// CreatedAt Created date (Format ISO 8601)
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Email Email of the contact
	Email *string `json:"email,omitempty" validate:"omitempty,email"`

	// FirstName Firstname of the contact
	FirstName string `json:"first_name" validate:"required"`

	// FlowID Id of the flow linked with the contact (required if project_ID null)
	FlowID *string `json:"flow_ID,omitempty" validate:"omitempty,required_without=ProjectID,uuid4"`

	// Gender Gender of the contact
	Gender ContactGender `json:"gender" validate:"required,oneof=OTHER MALE FEMALE"`

	// Job Type of the contact
	Job *string `json:"job,omitempty" validate:"omitempty"`

	// LastName LastName of the contact
	LastName string `json:"last_name" validate:"required"`

	// Mobile Mobile of the contact
	Mobile *string `json:"mobile,omitempty" validate:"omitempty"`

	// ProjectID Id of the project linked with the contact (required if flow_id null)
	ProjectID *string `json:"project_ID,omitempty" validate:"omitempty,required_without=FlowID,uuid4"`

	// UpdatedAt Last update date (Format ISO 8601)
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// ContactGender Gender of the contact
type ContactGender string

// Flow defines model for flow.
type Flow struct {
	// ID Unique id (format UUID) of the flow
	ID *string `json:"ID,omitempty"`

	// CreatedAt Created flow date (Format ISO 8601)
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// Description Description of the flow
	Description string `json:"description" validate:"required"`

	// FlowOrigin Code origin of the flow
	FlowOrigin string `json:"flow_origin" validate:"required"`

	// Format Format of the data flow
	Format FlowFormat `json:"format" validate:"required,oneof=POLIRIS_CSV API_XML API_JSON"`

	// JobName Job name of the flow
	JobName string `json:"job_name" validate:"required"`

	// Logo Link to the logo of the flow provider
	Logo string `json:"logo" validate:"required"`

	// Name Name of the flow
	Name string `json:"name" validate:"required"`

	// Status Status of the flow
	Status FlowStatus `json:"status" validate:"required,oneof=CREATED ENABLED DISABLED SUSPENDED"`

	// TypeOffers Type of the offers in the flow
	TypeOffers FlowTypeOffers `json:"type_offers" validate:"required,oneof=NEUF ANCIEN"`

	// UpdatedAt Last update flow date (Format ISO 8601)
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// FlowFormat Format of the data flow
type FlowFormat string

// FlowStatus Status of the flow
type FlowStatus string

// FlowTypeOffers Type of the offers in the flow
type FlowTypeOffers string

// FlowJobs defines model for flowJobs.
type FlowJobs struct {
	// Count number of objects in the response
	Count *int64 `json:"count,omitempty"`

	// FlowJobs Apikeys fetched
	FlowJobs *[]Job `json:"flow_jobs,omitempty"`
}

// FlowsStatistics defines model for flowsStatistics.
type FlowsStatistics struct {
	JobsError []string `json:"jobs_error"`
	JobsOk    []string `json:"jobs_ok"`

	// NbAds number of active property
	NbAds int64 `json:"nb_ads"`

	// NbFlows number of active flows
	NbFlows int64 `json:"nb_flows"`

	// NbFlowsTotal number of total flows
	NbFlowsTotal *int64 `json:"nb_flows_total,omitempty"`
}

// Job defines model for job.
type Job struct {
	// Context Job run context
	Context string `json:"context"`

	// Duration Job run duration
	Duration int64     `json:"duration"`
	Error    *JobError `json:"error,omitempty"`

	// Job Job name (same as flow code)
	Job *string `json:"job,omitempty"`

	// Message Job run status
	Message *string `json:"message,omitempty"`

	// Moment Job run date
	Moment *time.Time `json:"moment,omitempty"`

	// Pid Job pid
	Pid string `json:"pid"`
}

// JobError defines model for jobError.
type JobError struct {
	// Message Error message
	Message string `json:"message"`

	// Origin Error component origin
	Origin string `json:"origin"`

	// Pid Job pid
	Pid string `json:"pid"`

	// Type Error type
	Type string `json:"type"`
}

// NewContact defines model for newContact.
type NewContact struct {
	// Email Email of the contact
	Email *string `json:"email,omitempty" validate:"omitempty,email"`

	// FirstName Firstname of the contact
	FirstName string `json:"first_name" validate:"required"`

	// FlowID Id of the flow linked with the contact (required if project_ID null)
	FlowID *string `json:"flow_ID,omitempty" validate:"omitempty,required_without=ProjectID,uuid4"`

	// Gender Gender of the contact
	Gender NewContactGender `json:"gender" validate:"required,oneof=OTHER MALE FEMALE"`

	// Job Type of the contact
	Job *string `json:"job,omitempty" validate:"omitempty"`

	// LastName LastName of the contact
	LastName string `json:"last_name" validate:"required"`

	// Mobile Mobile of the contact
	Mobile *string `json:"mobile,omitempty" validate:"omitempty"`

	// ProjectID Id of the project linked with the contact (required if flow_id null)
	ProjectID *string `json:"project_ID,omitempty" validate:"omitempty,required_without=FlowID,uuid4"`
}

// NewContactGender Gender of the contact
type NewContactGender string

// NewFlow defines model for newFlow.
type NewFlow struct {
	// Description Description of the flow
	Description string `json:"description" validate:"required"`

	// FlowOrigin Code origin of the flow
	FlowOrigin string `json:"flow_origin" validate:"required"`

	// Format Format of the data flow
	Format NewFlowFormat `json:"format" validate:"required,oneof=POLIRIS_CSV API_XML API_JSON"`

	// JobName Job name of the flow
	JobName string `json:"job_name" validate:"required"`

	// Logo Link to the logo of the flow provider
	Logo string `json:"logo" validate:"required"`

	// Name Name of the flow
	Name string `json:"name" validate:"required"`

	// Status Status of the flow
	Status NewFlowStatus `json:"status" validate:"required,oneof=CREATED ENABLED DISABLED SUSPENDED"`

	// TypeOffers Type of the offers in the flow
	TypeOffers NewFlowTypeOffers `json:"type_offers" validate:"required,oneof=NEUF ANCIEN"`
}

// NewFlowFormat Format of the data flow
type NewFlowFormat string

// NewFlowStatus Status of the flow
type NewFlowStatus string

// NewFlowTypeOffers Type of the offers in the flow
type NewFlowTypeOffers string

// ResponseFetchContacts defines model for responseFetchContacts.
type ResponseFetchContacts struct {
	// Contacts Invoice contacts fetched
	Contacts *[]Contact `json:"contacts,omitempty"`

	// Count number of contacts in the response
	Count *int64 `json:"count,omitempty"`
}

// ResponseFetchFlows defines model for responseFetchFlows.
type ResponseFetchFlows struct {
	// Count number of objects in the response
	Count *int64 `json:"count,omitempty"`

	// Flows Flows fetched
	Flows *[]Flow `json:"flows,omitempty"`
}

// ResponseFetchFlowsOrigins defines model for responseFetchFlowsOrigins.
type ResponseFetchFlowsOrigins struct {
	// FlowsOrigins Flows origins fetched
	FlowsOrigins *[]string `json:"flows_origins,omitempty"`
}

// StatusIn defines model for statusIn.
type StatusIn struct {
	Enable *bool `json:"enable,omitempty" validate:"required"`

	// SubscriptionID Id of the subscription
	SubscriptionID string `json:"subscription_ID" validate:"required,uuid4"`
}

// QueryOrderBy defines model for queryOrderBy.
type QueryOrderBy = string

// QueryPageSize defines model for queryPageSize.
type QueryPageSize = int64

// QueryPageToken defines model for queryPageToken.
type QueryPageToken = int64

// QuerySortDirDesc defines model for querySortDirDesc.
type QuerySortDirDesc = bool

// QueryStatus defines model for queryStatus.
type QueryStatus string

// QueryTypeOffers defines model for queryTypeOffers.
type QueryTypeOffers = string

// N400Error A representation of an errorMessage
type N400Error = ModelError

// N401Error A representation of an errorMessage
type N401Error = ModelError

// N403Error A representation of an errorMessage
type N403Error = ModelError

// N404Error A representation of an errorMessage
type N404Error = ModelError

// N500Error A representation of an errorMessage
type N500Error = ModelError

// N503Error A representation of an errorMessage
type N503Error = ModelError

// FindFlowsParams defines parameters for FindFlows.
type FindFlowsParams struct {
	// OrderBy order by field
	OrderBy *QueryOrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// PageSize pagination size
	PageSize *QueryPageSize `form:"page_size,omitempty" json:"page_size,omitempty"`

	// PageToken Page index token (starting from 0)
	PageToken *QueryPageToken `form:"page_token,omitempty" json:"page_token,omitempty"`

	// SortDirDesc is sorting desc
	SortDirDesc *QuerySortDirDesc `form:"sort_dir_desc,omitempty" json:"sort_dir_desc,omitempty"`
}

// FindFlowsOriginsParams defines parameters for FindFlowsOrigins.
type FindFlowsOriginsParams struct {
	// TypeOffers Type offers sorting
	TypeOffers *QueryTypeOffers `form:"type_offers,omitempty" json:"type_offers,omitempty"`

	// Status Status.
	Status *FindFlowsOriginsParamsStatus `form:"status,omitempty" json:"status,omitempty" validate:"omitempty,oneof=CREATED ENABLED DISABLED SUSPENDED"`
}

// FindFlowsOriginsParamsStatus defines parameters for FindFlowsOrigins.
type FindFlowsOriginsParamsStatus string

// FindContactsByFlowIDParams defines parameters for FindContactsByFlowID.
type FindContactsByFlowIDParams struct {
	// OrderBy order by field
	OrderBy *QueryOrderBy `form:"order_by,omitempty" json:"order_by,omitempty"`

	// PageSize pagination size
	PageSize *QueryPageSize `form:"page_size,omitempty" json:"page_size,omitempty"`

	// PageToken Page index token (starting from 0)
	PageToken *QueryPageToken `form:"page_token,omitempty" json:"page_token,omitempty"`

	// SortDirDesc is sorting desc
	SortDirDesc *QuerySortDirDesc `form:"sort_dir_desc,omitempty" json:"sort_dir_desc,omitempty"`
}

// FindJobsByFlowJobNameParams defines parameters for FindJobsByFlowJobName.
type FindJobsByFlowJobNameParams struct {
	// PageSize pagination size
	PageSize *QueryPageSize `form:"page_size,omitempty" json:"page_size,omitempty"`

	// PageToken Page index token (starting from 0)
	PageToken *QueryPageToken `form:"page_token,omitempty" json:"page_token,omitempty"`
}

// CreateFlowJSONRequestBody defines body for CreateFlow for application/json ContentType.
type CreateFlowJSONRequestBody = NewFlow

// UpdateFlowJSONRequestBody defines body for UpdateFlow for application/json ContentType.
type UpdateFlowJSONRequestBody = NewFlow

// UpdateFlowVisibilityJSONRequestBody defines body for UpdateFlowVisibility for application/json ContentType.
type UpdateFlowVisibilityJSONRequestBody = StatusIn
