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

// FloatingIPActionService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFloatingIPActionService] method
// instead.
type FloatingIPActionService struct {
	Options []option.RequestOption
}

// NewFloatingIPActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewFloatingIPActionService(opts ...option.RequestOption) (r *FloatingIPActionService) {
	r = &FloatingIPActionService{}
	r.Options = opts
	return
}

// Returns a specific Action object for a Floating IP.
func (r *FloatingIPActionService) Get(ctx context.Context, id int64, actionID int64, opts ...option.RequestOption) (res *FloatingIPActionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v/actions/%v", id, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all Action objects for a Floating IP. You can sort the results by using
// the `sort` URI parameter, and filter them with the `status` parameter.
func (r *FloatingIPActionService) List(ctx context.Context, id int64, query FloatingIPActionListParams, opts ...option.RequestOption) (res *FloatingIPActionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v/actions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Assigns a Floating IP to a Server.
func (r *FloatingIPActionService) Assign(ctx context.Context, id int64, body FloatingIPActionAssignParams, opts ...option.RequestOption) (res *FloatingIPActionAssignResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v/actions/assign", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the hostname that will appear when getting the hostname belonging to
// this Floating IP.
func (r *FloatingIPActionService) ChangeDNSPtr(ctx context.Context, id int64, body FloatingIPActionChangeDNSPtrParams, opts ...option.RequestOption) (res *FloatingIPActionChangeDNSPtrResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v/actions/change_dns_ptr", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the protection configuration of the Floating IP.
func (r *FloatingIPActionService) ChangeProtection(ctx context.Context, id int64, body FloatingIPActionChangeProtectionParams, opts ...option.RequestOption) (res *FloatingIPActionChangeProtectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v/actions/change_protection", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Unassigns a Floating IP, resulting in it being unreachable. You may assign it to
// a Server again at a later time.
func (r *FloatingIPActionService) Unassign(ctx context.Context, id int64, opts ...option.RequestOption) (res *FloatingIPActionUnassignResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v/actions/unassign", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Response to GET
// https://api.hetzner.cloud/v1/floating_ips/{id}/actions/{action_id}
type FloatingIPActionGetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   floatingIPActionGetResponseJSON
}

// floatingIPActionGetResponseJSON contains the JSON metadata for the struct
// [FloatingIPActionGetResponse]
type floatingIPActionGetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/floating_ips/{id}/actions
type FloatingIPActionListResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON floatingIPActionListResponseJSON
}

// floatingIPActionListResponseJSON contains the JSON metadata for the struct
// [FloatingIPActionListResponse]
type floatingIPActionListResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/floating_ips/{id}/actions/assign
type FloatingIPActionAssignResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   floatingIPActionAssignResponseJSON
}

// floatingIPActionAssignResponseJSON contains the JSON metadata for the struct
// [FloatingIPActionAssignResponse]
type floatingIPActionAssignResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPActionAssignResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/floating_ips/{id}/actions/change_dns_ptr
type FloatingIPActionChangeDNSPtrResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   floatingIPActionChangeDNSPtrResponseJSON
}

// floatingIPActionChangeDNSPtrResponseJSON contains the JSON metadata for the
// struct [FloatingIPActionChangeDNSPtrResponse]
type floatingIPActionChangeDNSPtrResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPActionChangeDNSPtrResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/floating_ips/{id}/actions/change_protection
type FloatingIPActionChangeProtectionResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   floatingIPActionChangeProtectionResponseJSON
}

// floatingIPActionChangeProtectionResponseJSON contains the JSON metadata for the
// struct [FloatingIPActionChangeProtectionResponse]
type floatingIPActionChangeProtectionResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPActionChangeProtectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/floating_ips/{id}/actions/unassign
type FloatingIPActionUnassignResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   floatingIPActionUnassignResponseJSON
}

// floatingIPActionUnassignResponseJSON contains the JSON metadata for the struct
// [FloatingIPActionUnassignResponse]
type floatingIPActionUnassignResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPActionUnassignResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPActionListParams struct {
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

// URLQuery serializes [FloatingIPActionListParams]'s query parameters as
// `url.Values`.
func (r FloatingIPActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FloatingIPActionAssignParams struct {
	// ID of the Server the Floating IP shall be assigned to
	Server param.Field[int64] `json:"server,required"`
}

func (r FloatingIPActionAssignParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FloatingIPActionChangeDNSPtrParams struct {
	// Hostname to set as a reverse DNS PTR entry, will reset to original default value
	// if `null`
	DNSPtr param.Field[string] `json:"dns_ptr,required"`
	// IP address for which to set the reverse DNS entry
	IP param.Field[string] `json:"ip,required"`
}

func (r FloatingIPActionChangeDNSPtrParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FloatingIPActionChangeProtectionParams struct {
	// If true, prevents the Floating IP from being deleted
	Delete param.Field[bool] `json:"delete"`
}

func (r FloatingIPActionChangeProtectionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
