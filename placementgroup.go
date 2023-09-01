// File generated from our OpenAPI spec by Stainless.

package hetzner

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hetzner/hetzner-go/internal/apijson"
	"github.com/hetzner/hetzner-go/internal/apiquery"
	"github.com/hetzner/hetzner-go/internal/param"
	"github.com/hetzner/hetzner-go/internal/requestconfig"
	"github.com/hetzner/hetzner-go/internal/shared"
	"github.com/hetzner/hetzner-go/option"
)

// PlacementGroupService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPlacementGroupService] method
// instead.
type PlacementGroupService struct {
	Options []option.RequestOption
}

// NewPlacementGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPlacementGroupService(opts ...option.RequestOption) (r *PlacementGroupService) {
	r = &PlacementGroupService{}
	r.Options = opts
	return
}

// Creates a new PlacementGroup.
func (r *PlacementGroupService) New(ctx context.Context, body PlacementGroupNewParams, opts ...option.RequestOption) (res *PlacementGroupNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "placement_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a specific PlacementGroup object.
func (r *PlacementGroupService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *PlacementGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("placement_groups/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the PlacementGroup properties.
//
// Note that when updating labels, the PlacementGroup’s current set of labels will
// be replaced with the labels provided in the request body. So, for example, if
// you want to add a new label, you have to provide all existing labels plus the
// new label in the request body.
//
// Note: if the PlacementGroup object changes during the request, the response will
// be a “conflict” error.
func (r *PlacementGroupService) Update(ctx context.Context, id int64, body PlacementGroupUpdateParams, opts ...option.RequestOption) (res *PlacementGroupUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("placement_groups/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all PlacementGroup objects.
func (r *PlacementGroupService) List(ctx context.Context, query PlacementGroupListParams, opts ...option.RequestOption) (res *PlacementGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "placement_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a PlacementGroup.
func (r *PlacementGroupService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("placement_groups/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type PlacementGroup struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Array of IDs of Servers that are part of this Placement Group
	Servers []int64 `json:"servers,required"`
	// Type of the Placement Group
	Type PlacementGroupType `json:"type,required"`
	JSON placementGroupJSON
}

// placementGroupJSON contains the JSON metadata for the struct [PlacementGroup]
type placementGroupJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Labels      apijson.Field
	Name        apijson.Field
	Servers     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the Placement Group
type PlacementGroupType string

const (
	PlacementGroupTypeSpread PlacementGroupType = "spread"
)

// Response to POST https://api.hetzner.cloud/v1/placement_groups
type PlacementGroupNewResponse struct {
	PlacementGroup PlacementGroup `json:"placement_group,required"`
	// Actions show the results and progress of asynchronous requests to the API.
	Action PlacementGroupNewResponseAction `json:"action,nullable"`
	JSON   placementGroupNewResponseJSON
}

// placementGroupNewResponseJSON contains the JSON metadata for the struct
// [PlacementGroupNewResponse]
type placementGroupNewResponseJSON struct {
	PlacementGroup apijson.Field
	Action         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PlacementGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions show the results and progress of asynchronous requests to the API.
type PlacementGroupNewResponseAction struct {
	// ID of the Action
	ID int64 `json:"id,required"`
	// Command executed in the Action
	Command string `json:"command,required"`
	// Error message for the Action if error occurred, otherwise null
	Error PlacementGroupNewResponseActionError `json:"error,required,nullable"`
	// Point in time when the Action was finished (in ISO-8601 format). Only set if the
	// Action is finished otherwise null.
	Finished string `json:"finished,required,nullable"`
	// Progress of Action in percent
	Progress float64 `json:"progress,required"`
	// Resources the Action relates to
	Resources []PlacementGroupNewResponseActionResource `json:"resources,required"`
	// Point in time when the Action was started (in ISO-8601 format)
	Started string `json:"started,required"`
	// Status of the Action
	Status PlacementGroupNewResponseActionStatus `json:"status,required"`
	JSON   placementGroupNewResponseActionJSON
}

// placementGroupNewResponseActionJSON contains the JSON metadata for the struct
// [PlacementGroupNewResponseAction]
type placementGroupNewResponseActionJSON struct {
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

func (r *PlacementGroupNewResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Error message for the Action if error occurred, otherwise null
type PlacementGroupNewResponseActionError struct {
	// Fixed machine readable code
	Code string `json:"code,required"`
	// Humanized error message
	Message string `json:"message,required"`
	JSON    placementGroupNewResponseActionErrorJSON
}

// placementGroupNewResponseActionErrorJSON contains the JSON metadata for the
// struct [PlacementGroupNewResponseActionError]
type placementGroupNewResponseActionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementGroupNewResponseActionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementGroupNewResponseActionResource struct {
	// ID of the Resource | ID of resource referenced
	ID int64 `json:"id,required"`
	// Type of resource referenced
	Type string `json:"type,required"`
	JSON placementGroupNewResponseActionResourceJSON
}

// placementGroupNewResponseActionResourceJSON contains the JSON metadata for the
// struct [PlacementGroupNewResponseActionResource]
type placementGroupNewResponseActionResourceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PlacementGroupNewResponseActionResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Action
type PlacementGroupNewResponseActionStatus string

const (
	PlacementGroupNewResponseActionStatusError   PlacementGroupNewResponseActionStatus = "error"
	PlacementGroupNewResponseActionStatusRunning PlacementGroupNewResponseActionStatus = "running"
	PlacementGroupNewResponseActionStatusSuccess PlacementGroupNewResponseActionStatus = "success"
)

// Response to GET https://api.hetzner.cloud/v1/placement_groups/{id}
type PlacementGroupGetResponse struct {
	PlacementGroup PlacementGroup `json:"placement_group,required"`
	JSON           placementGroupGetResponseJSON
}

// placementGroupGetResponseJSON contains the JSON metadata for the struct
// [PlacementGroupGetResponse]
type placementGroupGetResponseJSON struct {
	PlacementGroup apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PlacementGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to PUT https://api.hetzner.cloud/v1/placement_groups/{id}
type PlacementGroupUpdateResponse struct {
	PlacementGroup PlacementGroup `json:"placement_group,required"`
	JSON           placementGroupUpdateResponseJSON
}

// placementGroupUpdateResponseJSON contains the JSON metadata for the struct
// [PlacementGroupUpdateResponse]
type placementGroupUpdateResponseJSON struct {
	PlacementGroup apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PlacementGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/placement_groups
type PlacementGroupListResponse struct {
	PlacementGroups []PlacementGroup `json:"placement_groups,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON placementGroupListResponseJSON
}

// placementGroupListResponseJSON contains the JSON metadata for the struct
// [PlacementGroupListResponse]
type placementGroupListResponseJSON struct {
	PlacementGroups apijson.Field
	Meta            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PlacementGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PlacementGroupNewParams struct {
	// Name of the PlacementGroup
	Name param.Field[string] `json:"name,required"`
	// Define the Placement Group Type.
	Type param.Field[PlacementGroupNewParamsType] `json:"type,required"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r PlacementGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define the Placement Group Type.
type PlacementGroupNewParamsType string

const (
	PlacementGroupNewParamsTypeSpread PlacementGroupNewParamsType = "spread"
)

type PlacementGroupUpdateParams struct {
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// New PlacementGroup name
	Name param.Field[string] `json:"name"`
}

func (r PlacementGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PlacementGroupListParams struct {
	// Can be used to filter resources by labels. The response will only contain
	// resources matching the label selector.
	LabelSelector param.Field[string] `query:"label_selector"`
	// Can be used to filter resources by their name. The response will only contain
	// the resources matching the specified name
	Name param.Field[string] `query:"name"`
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
	// Can be used multiple times.
	Sort param.Field[PlacementGroupListParamsSort] `query:"sort"`
	// Can be used multiple times. The response will only contain PlacementGroups
	// matching the type.
	Type param.Field[PlacementGroupListParamsType] `query:"type"`
}

// URLQuery serializes [PlacementGroupListParams]'s query parameters as
// `url.Values`.
func (r PlacementGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times.
type PlacementGroupListParamsSort string

const (
	PlacementGroupListParamsSortID          PlacementGroupListParamsSort = "id"
	PlacementGroupListParamsSortIDAsc       PlacementGroupListParamsSort = "id:asc"
	PlacementGroupListParamsSortIDDesc      PlacementGroupListParamsSort = "id:desc"
	PlacementGroupListParamsSortName        PlacementGroupListParamsSort = "name"
	PlacementGroupListParamsSortNameAsc     PlacementGroupListParamsSort = "name:asc"
	PlacementGroupListParamsSortNameDesc    PlacementGroupListParamsSort = "name:desc"
	PlacementGroupListParamsSortCreated     PlacementGroupListParamsSort = "created"
	PlacementGroupListParamsSortCreatedAsc  PlacementGroupListParamsSort = "created:asc"
	PlacementGroupListParamsSortCreatedDesc PlacementGroupListParamsSort = "created:desc"
)

// Can be used multiple times. The response will only contain PlacementGroups
// matching the type.
type PlacementGroupListParamsType string

const (
	PlacementGroupListParamsTypeSpread PlacementGroupListParamsType = "spread"
)
