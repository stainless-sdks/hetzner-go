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

// PrimaryIpActionService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewPrimaryIpActionService] method
// instead.
type PrimaryIpActionService struct {
	Options []option.RequestOption
}

// NewPrimaryIpActionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrimaryIpActionService(opts ...option.RequestOption) (r *PrimaryIpActionService) {
	r = &PrimaryIpActionService{}
	r.Options = opts
	return
}

// Returns a specific Action object.
func (r *PrimaryIpActionService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *PrimaryIpActionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("primary_ips/actions/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all Action objects. You can `sort` the results by using the sort URI
// parameter, and filter them with the `status` and `id` parameter.
func (r *PrimaryIpActionService) List(ctx context.Context, query PrimaryIpActionListParams, opts ...option.RequestOption) (res *PrimaryIpActionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "primary_ips/actions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Assigns a Primary IP to a Server.
//
// A Server can only have one Primary IP of type `ipv4` and one of type `ipv6`
// assigned. If you need more IPs use Floating IPs.
//
// The Server must be powered off (status `off`) in order for this operation to
// succeed.
//
// #### Call specific error codes
//
// | Code                          | Description                                          |
// | ----------------------------- | ---------------------------------------------------- |
// | `server_not_stopped`          | The server is running, but needs to be powered off   |
// | `primary_ip_already_assigned` | Primary ip is already assigned to a different server |
// | `server_has_ipv4`             | The server already has an ipv4 address               |
// | `server_has_ipv6`             | The server already has an ipv6 address               |
func (r *PrimaryIpActionService) Assign(ctx context.Context, id int64, body PrimaryIpActionAssignParams, opts ...option.RequestOption) (res *PrimaryIpActionAssignResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("primary_ips/%v/actions/assign", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the hostname that will appear when getting the hostname belonging to
// this Primary IP.
func (r *PrimaryIpActionService) ChangeDnsPtr(ctx context.Context, id int64, body PrimaryIpActionChangeDnsPtrParams, opts ...option.RequestOption) (res *PrimaryIpActionChangeDnsPtrResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("primary_ips/%v/actions/change_dns_ptr", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the protection configuration of a Primary IP.
//
// A Primary IP can only be delete protected if its `auto_delete` property is set
// to `false`.
func (r *PrimaryIpActionService) ChangeProtection(ctx context.Context, id int64, body PrimaryIpActionChangeProtectionParams, opts ...option.RequestOption) (res *PrimaryIpActionChangeProtectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("primary_ips/%v/actions/change_protection", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unassigns a Primary IP from a Server.
//
// The Server must be powered off (status `off`) in order for this operation to
// succeed.
//
// Note that only Servers that have at least one network interface (public or
// private) attached can be powered on.
//
// #### Call specific error codes
//
// | Code                             | Description                                        |
// | -------------------------------- | -------------------------------------------------- |
// | `server_not_stopped`             | The server is running, but needs to be powered off |
// | `server_is_load_balancer_target` | The server ipv4 address is a loadbalancer target   |
func (r *PrimaryIpActionService) Unassign(ctx context.Context, id int64, opts ...option.RequestOption) (res *PrimaryIpActionUnassignResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("primary_ips/%v/actions/unassign", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/{resource}/actions
type PrimaryIpActionGetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   primaryIpActionGetResponseJSON
}

// primaryIpActionGetResponseJSON contains the JSON metadata for the struct
// [PrimaryIpActionGetResponse]
type primaryIpActionGetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/{resource}/actions/{id}
type PrimaryIpActionListResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON primaryIpActionListResponseJSON
}

// primaryIpActionListResponseJSON contains the JSON metadata for the struct
// [PrimaryIpActionListResponse]
type primaryIpActionListResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/primary_ips/{id}/actions/assign
type PrimaryIpActionAssignResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   primaryIpActionAssignResponseJSON
}

// primaryIpActionAssignResponseJSON contains the JSON metadata for the struct
// [PrimaryIpActionAssignResponse]
type primaryIpActionAssignResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpActionAssignResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/primary_ips/{id}/actions/change_dns_ptr
type PrimaryIpActionChangeDnsPtrResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   primaryIpActionChangeDnsPtrResponseJSON
}

// primaryIpActionChangeDnsPtrResponseJSON contains the JSON metadata for the
// struct [PrimaryIpActionChangeDnsPtrResponse]
type primaryIpActionChangeDnsPtrResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpActionChangeDnsPtrResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/primary_ips/{id}/actions/change_protection
type PrimaryIpActionChangeProtectionResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   primaryIpActionChangeProtectionResponseJSON
}

// primaryIpActionChangeProtectionResponseJSON contains the JSON metadata for the
// struct [PrimaryIpActionChangeProtectionResponse]
type primaryIpActionChangeProtectionResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpActionChangeProtectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/primary_ips/{id}/actions/unassign
type PrimaryIpActionUnassignResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   primaryIpActionUnassignResponseJSON
}

// primaryIpActionUnassignResponseJSON contains the JSON metadata for the struct
// [PrimaryIpActionUnassignResponse]
type primaryIpActionUnassignResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpActionUnassignResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PrimaryIpActionListParams struct {
	// Can be used multiple times, the response will contain only Actions with
	// specified IDs
	ID param.Field[int64] `query:"id"`
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
	// Can be used multiple times.
	Sort param.Field[shared.SortParam] `query:"sort"`
	// Can be used multiple times, the response will contain only Actions with
	// specified statuses
	Status param.Field[shared.StatusParam] `query:"status"`
}

// URLQuery serializes [PrimaryIpActionListParams]'s query parameters as
// `url.Values`.
func (r PrimaryIpActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PrimaryIpActionAssignParams struct {
	// ID of a resource of type `assignee_type`
	AssigneeID param.Field[int64] `json:"assignee_id,required"`
	// Type of resource assigning the Primary IP to
	AssigneeType param.Field[PrimaryIpActionAssignParamsAssigneeType] `json:"assignee_type,required"`
}

func (r PrimaryIpActionAssignParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of resource assigning the Primary IP to
type PrimaryIpActionAssignParamsAssigneeType string

const (
	PrimaryIpActionAssignParamsAssigneeTypeServer PrimaryIpActionAssignParamsAssigneeType = "server"
)

type PrimaryIpActionChangeDnsPtrParams struct {
	// Hostname to set as a reverse DNS PTR entry, will reset to original default value
	// if `null`
	DnsPtr param.Field[string] `json:"dns_ptr,required"`
	// IP address for which to set the reverse DNS entry
	Ip param.Field[string] `json:"ip,required"`
}

func (r PrimaryIpActionChangeDnsPtrParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrimaryIpActionChangeProtectionParams struct {
	// If true, prevents the Primary IP from being deleted
	Delete param.Field[bool] `json:"delete"`
}

func (r PrimaryIpActionChangeProtectionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
