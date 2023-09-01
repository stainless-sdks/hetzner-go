// File generated from our OpenAPI spec by Stainless.

package shared

import (
	"github.com/hetzner/hetzner-go/internal/apijson"
	"github.com/hetzner/hetzner-go/internal/param"
)

// Actions show the results and progress of asynchronous requests to the API.
type Action struct {
	// ID of the Action
	ID int64 `json:"id,required"`
	// Command executed in the Action
	Command string `json:"command,required"`
	// Error message for the Action if error occurred, otherwise null
	Error ActionError `json:"error,required,nullable"`
	// Point in time when the Action was finished (in ISO-8601 format). Only set if the
	// Action is finished otherwise null.
	Finished string `json:"finished,required,nullable"`
	// Progress of Action in percent
	Progress float64 `json:"progress,required"`
	// Resources the Action relates to
	Resources []ActionResource `json:"resources,required"`
	// Point in time when the Action was started (in ISO-8601 format)
	Started string `json:"started,required"`
	// Status of the Action
	Status ActionStatus `json:"status,required"`
	JSON   actionJSON
}

// actionJSON contains the JSON metadata for the struct [Action]
type actionJSON struct {
	ID          apijson.Field
	Command     apijson.Field
	Error       apijson.Field
	Finished    apijson.Field
	Progress    apijson.Field
	Resources   apijson.Field
	Started     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Action) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Error message for the Action if error occurred, otherwise null
type ActionError struct {
	// Fixed machine readable code
	Code string `json:"code,required"`
	// Humanized error message
	Message string `json:"message,required"`
	JSON    actionErrorJSON
}

// actionErrorJSON contains the JSON metadata for the struct [ActionError]
type actionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ActionResource struct {
	// ID of the Resource | ID of resource referenced
	ID int64 `json:"id,required"`
	// Type of resource referenced
	Type string `json:"type,required"`
	JSON actionResourceJSON
}

// actionResourceJSON contains the JSON metadata for the struct [ActionResource]
type actionResourceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActionResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Action
type ActionStatus string

const (
	ActionStatusError   ActionStatus = "error"
	ActionStatusRunning ActionStatus = "running"
	ActionStatusSuccess ActionStatus = "success"
)

type HealthStatus struct {
	ListenPort int64              `json:"listen_port"`
	Status     HealthStatusStatus `json:"status"`
	JSON       healthStatusJSON
}

// healthStatusJSON contains the JSON metadata for the struct [HealthStatus]
type healthStatusJSON struct {
	ListenPort  apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthStatusStatus string

const (
	HealthStatusStatusHealthy   HealthStatusStatus = "healthy"
	HealthStatusStatusUnhealthy HealthStatusStatus = "unhealthy"
	HealthStatusStatusUnknown   HealthStatusStatus = "unknown"
)

type HealthStatusParam struct {
	ListenPort param.Field[int64]              `json:"listen_port"`
	Status     param.Field[HealthStatusStatus] `json:"status"`
}

func (r HealthStatusParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metadata contained in the response
type ResponseMeta struct {
	// Information about the current pagination. The keys previous_page, next_page,
	// last_page, and total_entries may be null when on the first page, last page, or
	// when the total number of entries is unknown
	Pagination ResponseMetaPagination `json:"pagination,required"`
	JSON       responseMetaJSON
}

// responseMetaJSON contains the JSON metadata for the struct [ResponseMeta]
type responseMetaJSON struct {
	Pagination  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Information about the current pagination. The keys previous_page, next_page,
// last_page, and total_entries may be null when on the first page, last page, or
// when the total number of entries is unknown
type ResponseMetaPagination struct {
	// ID of the last page available. Can be null if the current page is the last one.
	// | The last page number
	LastPage int64 `json:"last_page,required,nullable"`
	// ID of the next page. Can be null if the current page is the last one. | The next
	// page number
	NextPage int64 `json:"next_page,required,nullable"`
	// Current page number | The current page number
	Page int64 `json:"page,required"`
	// Maximum number of items shown per page in the response | The number of entries
	// per page
	PerPage int64 `json:"per_page,required"`
	// ID of the previous page. Can be null if the current page is the first one. | The
	// previous page number
	PreviousPage int64 `json:"previous_page,required,nullable"`
	// The total number of entries that exist in the database for this query. Nullable
	// if unknown. | The total number of entries
	TotalEntries int64 `json:"total_entries,required,nullable"`
	JSON         responseMetaPaginationJSON
}

// responseMetaPaginationJSON contains the JSON metadata for the struct
// [ResponseMetaPagination]
type responseMetaPaginationJSON struct {
	LastPage     apijson.Field
	NextPage     apijson.Field
	Page         apijson.Field
	PerPage      apijson.Field
	PreviousPage apijson.Field
	TotalEntries apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ResponseMetaPagination) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SortParam string

const (
	SortParamID           SortParam = "id"
	SortParamIDAsc        SortParam = "id:asc"
	SortParamIDDesc       SortParam = "id:desc"
	SortParamCommand      SortParam = "command"
	SortParamCommandAsc   SortParam = "command:asc"
	SortParamCommandDesc  SortParam = "command:desc"
	SortParamStatus       SortParam = "status"
	SortParamStatusAsc    SortParam = "status:asc"
	SortParamStatusDesc   SortParam = "status:desc"
	SortParamProgress     SortParam = "progress"
	SortParamProgressAsc  SortParam = "progress:asc"
	SortParamProgressDesc SortParam = "progress:desc"
	SortParamStarted      SortParam = "started"
	SortParamStartedAsc   SortParam = "started:asc"
	SortParamStartedDesc  SortParam = "started:desc"
	SortParamFinished     SortParam = "finished"
	SortParamFinishedAsc  SortParam = "finished:asc"
	SortParamFinishedDesc SortParam = "finished:desc"
)

type StatusParam string

const (
	StatusParamRunning StatusParam = "running"
	StatusParamSuccess StatusParam = "success"
	StatusParamError   StatusParam = "error"
)
