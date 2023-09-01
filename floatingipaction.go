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

// FloatingIpActionService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFloatingIpActionService] method
// instead.
type FloatingIpActionService struct {
	Options []option.RequestOption
}

// NewFloatingIpActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFloatingIpActionService(opts ...option.RequestOption) (r *FloatingIpActionService) {
	r = &FloatingIpActionService{}
	r.Options = opts
	return
}

// Returns a specific Action object for a Floating IP.
func (r *FloatingIpActionService) Get(ctx context.Context, id int64, actionID int64, opts ...option.RequestOption) (res *FloatingIpActionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v/actions/%v", id, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all Action objects for a Floating IP. You can sort the results by using
// the `sort` URI parameter, and filter them with the `status` parameter.
func (r *FloatingIpActionService) List(ctx context.Context, id int64, query FloatingIpActionListParams, opts ...option.RequestOption) (res *FloatingIpActionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v/actions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Assigns a Floating IP to a Server.
func (r *FloatingIpActionService) Assign(ctx context.Context, id int64, body FloatingIpActionAssignParams, opts ...option.RequestOption) (res *FloatingIpActionAssignResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v/actions/assign", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the hostname that will appear when getting the hostname belonging to
// this Floating IP.
func (r *FloatingIpActionService) ChangeDnsPtr(ctx context.Context, id int64, body FloatingIpActionChangeDnsPtrParams, opts ...option.RequestOption) (res *FloatingIpActionChangeDnsPtrResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v/actions/change_dns_ptr", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the protection configuration of the Floating IP.
func (r *FloatingIpActionService) ChangeProtection(ctx context.Context, id int64, body FloatingIpActionChangeProtectionParams, opts ...option.RequestOption) (res *FloatingIpActionChangeProtectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v/actions/change_protection", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unassigns a Floating IP, resulting in it being unreachable. You may assign it to
// a Server again at a later time.
func (r *FloatingIpActionService) Unassign(ctx context.Context, id int64, opts ...option.RequestOption) (res *FloatingIpActionUnassignResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v/actions/unassign", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Response to GET
// https://api.hetzner.cloud/v1/floating_ips/{id}/actions/{action_id}
type FloatingIpActionGetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   floatingIpActionGetResponseJSON
}

// floatingIpActionGetResponseJSON contains the JSON metadata for the struct
// [FloatingIpActionGetResponse]
type floatingIpActionGetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIpActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/floating_ips/{id}/actions
type FloatingIpActionListResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON floatingIpActionListResponseJSON
}

// floatingIpActionListResponseJSON contains the JSON metadata for the struct
// [FloatingIpActionListResponse]
type floatingIpActionListResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIpActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/floating_ips/{id}/actions/assign
type FloatingIpActionAssignResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   floatingIpActionAssignResponseJSON
}

// floatingIpActionAssignResponseJSON contains the JSON metadata for the struct
// [FloatingIpActionAssignResponse]
type floatingIpActionAssignResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIpActionAssignResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/floating_ips/{id}/actions/change_dns_ptr
type FloatingIpActionChangeDnsPtrResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   floatingIpActionChangeDnsPtrResponseJSON
}

// floatingIpActionChangeDnsPtrResponseJSON contains the JSON metadata for the
// struct [FloatingIpActionChangeDnsPtrResponse]
type floatingIpActionChangeDnsPtrResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIpActionChangeDnsPtrResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/floating_ips/{id}/actions/change_protection
type FloatingIpActionChangeProtectionResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   floatingIpActionChangeProtectionResponseJSON
}

// floatingIpActionChangeProtectionResponseJSON contains the JSON metadata for the
// struct [FloatingIpActionChangeProtectionResponse]
type floatingIpActionChangeProtectionResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIpActionChangeProtectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/floating_ips/{id}/actions/unassign
type FloatingIpActionUnassignResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   floatingIpActionUnassignResponseJSON
}

// floatingIpActionUnassignResponseJSON contains the JSON metadata for the struct
// [FloatingIpActionUnassignResponse]
type floatingIpActionUnassignResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIpActionUnassignResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIpActionListParams struct {
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

// URLQuery serializes [FloatingIpActionListParams]'s query parameters as
// `url.Values`.
func (r FloatingIpActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FloatingIpActionAssignParams struct {
	// ID of the Server the Floating IP shall be assigned to
	Server param.Field[int64] `json:"server,required"`
}

func (r FloatingIpActionAssignParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FloatingIpActionChangeDnsPtrParams struct {
	// Hostname to set as a reverse DNS PTR entry, will reset to original default value
	// if `null`
	DnsPtr param.Field[string] `json:"dns_ptr,required"`
	// IP address for which to set the reverse DNS entry
	Ip param.Field[string] `json:"ip,required"`
}

func (r FloatingIpActionChangeDnsPtrParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FloatingIpActionChangeProtectionParams struct {
	// If true, prevents the Floating IP from being deleted
	Delete param.Field[bool] `json:"delete"`
}

func (r FloatingIpActionChangeProtectionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
