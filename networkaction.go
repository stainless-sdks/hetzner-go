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

// NetworkActionService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewNetworkActionService] method
// instead.
type NetworkActionService struct {
	Options []option.RequestOption
}

// NewNetworkActionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkActionService(opts ...option.RequestOption) (r *NetworkActionService) {
	r = &NetworkActionService{}
	r.Options = opts
	return
}

// Returns a specific Action for a Network.
func (r *NetworkActionService) Get(ctx context.Context, id int64, actionID int64, opts ...option.RequestOption) (res *NetworkActionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("networks/%v/actions/%v", id, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all Action objects for a Network. You can sort the results by using the
// `sort` URI parameter, and filter them with the `status` parameter.
func (r *NetworkActionService) List(ctx context.Context, id int64, query NetworkActionListParams, opts ...option.RequestOption) (res *NetworkActionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("networks/%v/actions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Adds a route entry to a Network.
//
// Note: if the Network object changes during the request, the response will be a
// “conflict” error.
func (r *NetworkActionService) AddRoute(ctx context.Context, id int64, body NetworkActionAddRouteParams, opts ...option.RequestOption) (res *NetworkActionAddRouteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("networks/%v/actions/add_route", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Adds a new subnet object to the Network. If you do not specify an `ip_range` for
// the subnet we will automatically pick the first available /24 range for you if
// possible.
//
// Note: if the parent Network object changes during the request, the response will
// be a “conflict” error.
func (r *NetworkActionService) AddSubnet(ctx context.Context, id int64, body NetworkActionAddSubnetParams, opts ...option.RequestOption) (res *NetworkActionAddSubnetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("networks/%v/actions/add_subnet", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the IP range of a Network. IP ranges can only be extended and never
// shrunk. You can only add IPs at the end of an existing IP range. This means that
// the IP part of your existing range must stay the same and you can only change
// its netmask.
//
// For example if you have a range `10.0.0.0/16` you want to extend then your new
// range must also start with the IP `10.0.0.0`. Your CIDR netmask `/16` may change
// to a number that is smaller than `16` thereby increasing the IP range. So valid
// entries would be `10.0.0.0/15`, `10.0.0.0/14`, `10.0.0.0/13` and so on.
//
// After changing the IP range you will have to adjust the routes on your connected
// Servers by either rebooting them or manually changing the routes to your private
// Network interface.
//
// Note: if the Network object changes during the request, the response will be a
// “conflict” error.
func (r *NetworkActionService) ChangeIpRange(ctx context.Context, id int64, body NetworkActionChangeIpRangeParams, opts ...option.RequestOption) (res *NetworkActionChangeIpRangeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("networks/%v/actions/change_ip_range", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the protection configuration of a Network.
//
// Note: if the Network object changes during the request, the response will be a
// “conflict” error.
func (r *NetworkActionService) ChangeProtection(ctx context.Context, id int64, body NetworkActionChangeProtectionParams, opts ...option.RequestOption) (res *NetworkActionChangeProtectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("networks/%v/actions/change_protection", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a route entry from a Network.
//
// Note: if the Network object changes during the request, the response will be a
// “conflict” error.
func (r *NetworkActionService) DeleteRoute(ctx context.Context, id int64, body NetworkActionDeleteRouteParams, opts ...option.RequestOption) (res *NetworkActionDeleteRouteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("networks/%v/actions/delete_route", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Deletes a single subnet entry from a Network. You cannot delete subnets which
// still have Servers attached. If you have Servers attached you first need to
// detach all Servers that use IPs from this subnet before you can delete the
// subnet.
//
// Note: if the Network object changes during the request, the response will be a
// “conflict” error.
func (r *NetworkActionService) DeleteSubnet(ctx context.Context, id int64, body NetworkActionDeleteSubnetParams, opts ...option.RequestOption) (res *NetworkActionDeleteSubnetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("networks/%v/actions/delete_subnet", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/networks/{id}/actions/{action_id}
type NetworkActionGetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   networkActionGetResponseJSON
}

// networkActionGetResponseJSON contains the JSON metadata for the struct
// [NetworkActionGetResponse]
type networkActionGetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/networks/{id}/actions
type NetworkActionListResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON networkActionListResponseJSON
}

// networkActionListResponseJSON contains the JSON metadata for the struct
// [NetworkActionListResponse]
type networkActionListResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/networks/{id}/actions/add_route
type NetworkActionAddRouteResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   networkActionAddRouteResponseJSON
}

// networkActionAddRouteResponseJSON contains the JSON metadata for the struct
// [NetworkActionAddRouteResponse]
type networkActionAddRouteResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkActionAddRouteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/networks/{id}/actions/add_subnet
type NetworkActionAddSubnetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   networkActionAddSubnetResponseJSON
}

// networkActionAddSubnetResponseJSON contains the JSON metadata for the struct
// [NetworkActionAddSubnetResponse]
type networkActionAddSubnetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkActionAddSubnetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/networks/{id}/actions/change_ip_range
type NetworkActionChangeIpRangeResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   networkActionChangeIpRangeResponseJSON
}

// networkActionChangeIpRangeResponseJSON contains the JSON metadata for the struct
// [NetworkActionChangeIpRangeResponse]
type networkActionChangeIpRangeResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkActionChangeIpRangeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/networks/{id}/actions/change_protection
type NetworkActionChangeProtectionResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   networkActionChangeProtectionResponseJSON
}

// networkActionChangeProtectionResponseJSON contains the JSON metadata for the
// struct [NetworkActionChangeProtectionResponse]
type networkActionChangeProtectionResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkActionChangeProtectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/networks/{id}/actions/delete_route
type NetworkActionDeleteRouteResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   networkActionDeleteRouteResponseJSON
}

// networkActionDeleteRouteResponseJSON contains the JSON metadata for the struct
// [NetworkActionDeleteRouteResponse]
type networkActionDeleteRouteResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkActionDeleteRouteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/networks/{id}/actions/delete_subnet
type NetworkActionDeleteSubnetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   networkActionDeleteSubnetResponseJSON
}

// networkActionDeleteSubnetResponseJSON contains the JSON metadata for the struct
// [NetworkActionDeleteSubnetResponse]
type networkActionDeleteSubnetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkActionDeleteSubnetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkActionListParams struct {
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

// URLQuery serializes [NetworkActionListParams]'s query parameters as
// `url.Values`.
func (r NetworkActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type NetworkActionAddRouteParams struct {
	// Destination network or host of this route. Must not overlap with an existing
	// ip_range in any subnets or with any destinations in other routes or with the
	// first IP of the networks ip_range or with 172.31.1.1. Must be one of the private
	// IPv4 ranges of RFC1918.
	Destination param.Field[string] `json:"destination,required"`
	// Gateway for the route. Cannot be the first IP of the networks ip_range and also
	// cannot be 172.31.1.1 as this IP is being used as a gateway for the public
	// network interface of Servers. | Gateway for the route. Cannot be the first IP of
	// the networks ip_range, an IP behind a vSwitch or 172.31.1.1, as this IP is being
	// used as a gateway for the public network interface of Servers.
	Gateway param.Field[string] `json:"gateway,required"`
}

func (r NetworkActionAddRouteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkActionAddSubnetParams struct {
	// Name of Network zone. The Location object contains the `network_zone` property
	// each Location belongs to.
	NetworkZone param.Field[string] `json:"network_zone,required"`
	// Type of Subnetwork
	Type param.Field[NetworkActionAddSubnetParamsType] `json:"type,required"`
	// Range to allocate IPs from. Must be a Subnet of the ip_range of the parent
	// network object and must not overlap with any other subnets or with any
	// destinations in routes. Minimum Network size is /30. We suggest that you pick a
	// bigger Network with a /24 netmask. | Range to allocate IPs from. Must be a
	// Subnet of the ip_range of the parent network object and must not overlap with
	// any other subnets or with any destinations in routes. If the Subnet is of type
	// vSwitch, it also can not overlap with any gateway in routes. Minimum Network
	// size is /30. We suggest that you pick a bigger Network with a /24 netmask.
	IpRange param.Field[string] `json:"ip_range"`
	// ID of the robot vSwitch. Must be supplied if the subnet is of type vswitch.
	VswitchID param.Field[int64] `json:"vswitch_id"`
}

func (r NetworkActionAddSubnetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of Subnetwork
type NetworkActionAddSubnetParamsType string

const (
	NetworkActionAddSubnetParamsTypeCloud   NetworkActionAddSubnetParamsType = "cloud"
	NetworkActionAddSubnetParamsTypeServer  NetworkActionAddSubnetParamsType = "server"
	NetworkActionAddSubnetParamsTypeVswitch NetworkActionAddSubnetParamsType = "vswitch"
)

type NetworkActionChangeIpRangeParams struct {
	// The new prefix for the whole Network
	IpRange param.Field[string] `json:"ip_range,required"`
}

func (r NetworkActionChangeIpRangeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkActionChangeProtectionParams struct {
	// If true, prevents the Network from being deleted
	Delete param.Field[bool] `json:"delete"`
}

func (r NetworkActionChangeProtectionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkActionDeleteRouteParams struct {
	// Destination network or host of this route. Must not overlap with an existing
	// ip_range in any subnets or with any destinations in other routes or with the
	// first IP of the networks ip_range or with 172.31.1.1. Must be one of the private
	// IPv4 ranges of RFC1918.
	Destination param.Field[string] `json:"destination,required"`
	// Gateway for the route. Cannot be the first IP of the networks ip_range and also
	// cannot be 172.31.1.1 as this IP is being used as a gateway for the public
	// network interface of Servers. | Gateway for the route. Cannot be the first IP of
	// the networks ip_range, an IP behind a vSwitch or 172.31.1.1, as this IP is being
	// used as a gateway for the public network interface of Servers.
	Gateway param.Field[string] `json:"gateway,required"`
}

func (r NetworkActionDeleteRouteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkActionDeleteSubnetParams struct {
	// IP range of subnet to delete
	IpRange param.Field[string] `json:"ip_range,required"`
}

func (r NetworkActionDeleteSubnetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
