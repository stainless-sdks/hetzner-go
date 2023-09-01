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

// PrimaryIPService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPrimaryIPService] method instead.
type PrimaryIPService struct {
	Options []option.RequestOption
	Actions *PrimaryIPActionService
}

// NewPrimaryIPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrimaryIPService(opts ...option.RequestOption) (r *PrimaryIPService) {
	r = &PrimaryIPService{}
	r.Options = opts
	r.Actions = NewPrimaryIPActionService(opts...)
	return
}

// Creates a new Primary IP, optionally assigned to a Server.
//
// If you want to create a Primary IP that is not assigned to a Server, you need to
// provide the `datacenter` key instead of `assignee_id`. This can be either the ID
// or the name of the Datacenter this Primary IP shall be created in.
//
// Note that a Primary IP can only be assigned to a Server in the same Datacenter
// later on.
//
// #### Call specific error codes
//
// | Code                 | Description                                                  |
// | -------------------- | ------------------------------------------------------------ |
// | `server_not_stopped` | The specified server is running, but needs to be powered off |
// | `server_has_ipv4`    | The server already has an ipv4 address                       |
// | `server_has_ipv6`    | The server already has an ipv6 address                       |
func (r *PrimaryIPService) New(ctx context.Context, body PrimaryIPNewParams, opts ...option.RequestOption) (res *PrimaryIPNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "primary_ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a specific Primary IP object.
func (r *PrimaryIPService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *PrimaryIPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("primary_ips/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Primary IP.
//
// Note that when updating labels, the Primary IP's current set of labels will be
// replaced with the labels provided in the request body. So, for example, if you
// want to add a new label, you have to provide all existing labels plus the new
// label in the request body.
//
// If the Primary IP object changes during the request, the response will be a
// “conflict” error.
func (r *PrimaryIPService) Update(ctx context.Context, id int64, body PrimaryIPUpdateParams, opts ...option.RequestOption) (res *PrimaryIPUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("primary_ips/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all Primary IP objects.
func (r *PrimaryIPService) List(ctx context.Context, query PrimaryIPListParams, opts ...option.RequestOption) (res *PrimaryIPListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "primary_ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a Primary IP.
//
// The Primary IP may be assigned to a Server. In this case it is unassigned
// automatically. The Server must be powered off (status `off`) in order for this
// operation to succeed.
func (r *PrimaryIPService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("primary_ips/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type PrimaryIP struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// ID of the resource the Primary IP is assigned to, null if it is not assigned at
	// all
	AssigneeID int64 `json:"assignee_id,required,nullable"`
	// Resource type the Primary IP can be assigned to
	AssigneeType PrimaryIPAssigneeType `json:"assignee_type,required"`
	// Delete this Primary IP when the resource it is assigned to is deleted
	AutoDelete bool `json:"auto_delete,required"`
	// Whether the IP is blocked
	Blocked bool `json:"blocked,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Datacenter this Primary IP is located at | Datacenter this Resource is located
	// at
	Datacenter PrimaryIPDatacenter `json:"datacenter,required"`
	// Array of reverse DNS entries
	DNSPtr []PrimaryIPDNSPtr `json:"dns_ptr,required"`
	// IP address
	IP string `json:"ip,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Protection configuration for the Resource
	Protection PrimaryIPProtection `json:"protection,required"`
	// The type of the IP
	Type PrimaryIPType `json:"type,required"`
	JSON primaryIPJSON
}

// primaryIPJSON contains the JSON metadata for the struct [PrimaryIP]
type primaryIPJSON struct {
	ID           apijson.Field
	AssigneeID   apijson.Field
	AssigneeType apijson.Field
	AutoDelete   apijson.Field
	Blocked      apijson.Field
	Created      apijson.Field
	Datacenter   apijson.Field
	DNSPtr       apijson.Field
	IP           apijson.Field
	Labels       apijson.Field
	Name         apijson.Field
	Protection   apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrimaryIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type the Primary IP can be assigned to
type PrimaryIPAssigneeType string

const (
	PrimaryIPAssigneeTypeServer PrimaryIPAssigneeType = "server"
)

// Datacenter this Primary IP is located at | Datacenter this Resource is located
// at
type PrimaryIPDatacenter struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Description of the Datacenter
	Description string `json:"description,required"`
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	Location PrimaryIPDatacenterLocation `json:"location,required"`
	// Unique identifier of the Datacenter
	Name string `json:"name,required"`
	// The Server types the Datacenter can handle
	ServerTypes PrimaryIPDatacenterServerTypes `json:"server_types,required"`
	JSON        primaryIPDatacenterJSON
}

// primaryIPDatacenterJSON contains the JSON metadata for the struct
// [PrimaryIPDatacenter]
type primaryIPDatacenterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	ServerTypes apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIPDatacenter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type PrimaryIPDatacenterLocation struct {
	// ID of the Location
	ID int64 `json:"id,required"`
	// City the Location is closest to
	City string `json:"city,required"`
	// ISO 3166-1 alpha-2 code of the country the Location resides in
	Country string `json:"country,required"`
	// Description of the Location
	Description string `json:"description,required"`
	// Latitude of the city closest to the Location
	Latitude float64 `json:"latitude,required"`
	// Longitude of the city closest to the Location
	Longitude float64 `json:"longitude,required"`
	// Unique identifier of the Location
	Name string `json:"name,required"`
	// Name of network zone this Location resides in
	NetworkZone string `json:"network_zone,required"`
	JSON        primaryIPDatacenterLocationJSON
}

// primaryIPDatacenterLocationJSON contains the JSON metadata for the struct
// [PrimaryIPDatacenterLocation]
type primaryIPDatacenterLocationJSON struct {
	ID          apijson.Field
	City        apijson.Field
	Country     apijson.Field
	Description apijson.Field
	Latitude    apijson.Field
	Longitude   apijson.Field
	Name        apijson.Field
	NetworkZone apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIPDatacenterLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Server types the Datacenter can handle
type PrimaryIPDatacenterServerTypes struct {
	// IDs of Server types that are supported and for which the Datacenter has enough
	// resources left
	Available []int64 `json:"available,required"`
	// IDs of Server types that are supported and for which the Datacenter has enough
	// resources left
	AvailableForMigration []int64 `json:"available_for_migration,required"`
	// IDs of Server types that are supported in the Datacenter
	Supported []int64 `json:"supported,required"`
	JSON      primaryIPDatacenterServerTypesJSON
}

// primaryIPDatacenterServerTypesJSON contains the JSON metadata for the struct
// [PrimaryIPDatacenterServerTypes]
type primaryIPDatacenterServerTypesJSON struct {
	Available             apijson.Field
	AvailableForMigration apijson.Field
	Supported             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PrimaryIPDatacenterServerTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PrimaryIPDNSPtr struct {
	// DNS pointer for the specific IP address
	DNSPtr string `json:"dns_ptr,required"`
	// Single IPv4 or IPv6 address | Single IPv6 address of this Server for which the
	// reverse DNS entry has been set up
	IP   string `json:"ip,required"`
	JSON primaryIPDNSPtrJSON
}

// primaryIPDNSPtrJSON contains the JSON metadata for the struct [PrimaryIPDNSPtr]
type primaryIPDNSPtrJSON struct {
	DNSPtr      apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIPDNSPtr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type PrimaryIPProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   primaryIPProtectionJSON
}

// primaryIPProtectionJSON contains the JSON metadata for the struct
// [PrimaryIPProtection]
type primaryIPProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIPProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the IP
type PrimaryIPType string

const (
	PrimaryIPTypeIpv4 PrimaryIPType = "ipv4"
	PrimaryIPTypeIpv6 PrimaryIPType = "ipv6"
)

// Response to POST https://api.hetzner.cloud/v1/primary_ips
type PrimaryIPNewResponse struct {
	PrimaryIP PrimaryIP `json:"primary_ip,required"`
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action"`
	JSON   primaryIPNewResponseJSON
}

// primaryIPNewResponseJSON contains the JSON metadata for the struct
// [PrimaryIPNewResponse]
type primaryIPNewResponseJSON struct {
	PrimaryIP   apijson.Field
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIPNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/primary_ips/{id}
type PrimaryIPGetResponse struct {
	PrimaryIP PrimaryIP `json:"primary_ip,required"`
	JSON      primaryIPGetResponseJSON
}

// primaryIPGetResponseJSON contains the JSON metadata for the struct
// [PrimaryIPGetResponse]
type primaryIPGetResponseJSON struct {
	PrimaryIP   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIPGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to PUT https://api.hetzner.cloud/v1/primary_ips/{id}
type PrimaryIPUpdateResponse struct {
	PrimaryIP PrimaryIP `json:"primary_ip,required"`
	JSON      primaryIPUpdateResponseJSON
}

// primaryIPUpdateResponseJSON contains the JSON metadata for the struct
// [PrimaryIPUpdateResponse]
type primaryIPUpdateResponseJSON struct {
	PrimaryIP   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIPUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/primary_ips
type PrimaryIPListResponse struct {
	PrimaryIPs []PrimaryIP `json:"primary_ips,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON primaryIPListResponseJSON
}

// primaryIPListResponseJSON contains the JSON metadata for the struct
// [PrimaryIPListResponse]
type primaryIPListResponseJSON struct {
	PrimaryIPs  apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIPListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PrimaryIPNewParams struct {
	// Resource type the Primary IP can be assigned to
	AssigneeType param.Field[PrimaryIPNewParamsAssigneeType] `json:"assignee_type,required"`
	Name         param.Field[string]                         `json:"name,required"`
	// The type of the IP
	Type param.Field[PrimaryIPNewParamsType] `json:"type,required"`
	// ID of the resource the Primary IP should be assigned to. Omitted if it should
	// not be assigned.
	AssigneeID param.Field[int64] `json:"assignee_id"`
	// Delete the Primary IP when the Server it is assigned to is deleted. If omitted
	// defaults to `false`.
	AutoDelete param.Field[bool] `json:"auto_delete"`
	// ID or name of Datacenter the Primary IP will be bound to. Needs to be omitted if
	// `assignee_id` is passed.
	Datacenter param.Field[string] `json:"datacenter"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r PrimaryIPNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Resource type the Primary IP can be assigned to
type PrimaryIPNewParamsAssigneeType string

const (
	PrimaryIPNewParamsAssigneeTypeServer PrimaryIPNewParamsAssigneeType = "server"
)

// The type of the IP
type PrimaryIPNewParamsType string

const (
	PrimaryIPNewParamsTypeIpv4 PrimaryIPNewParamsType = "ipv4"
	PrimaryIPNewParamsTypeIpv6 PrimaryIPNewParamsType = "ipv6"
)

type PrimaryIPUpdateParams struct {
	// Delete this Primary IP when the resource it is assigned to is deleted
	AutoDelete param.Field[bool] `json:"auto_delete"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// New unique name to set
	Name param.Field[string] `json:"name"`
}

func (r PrimaryIPUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrimaryIPListParams struct {
	// Can be used to filter resources by their ip. The response will only contain the
	// resources matching the specified ip.
	IP param.Field[string] `query:"ip"`
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
	// Can be used multiple times. Choices id id:asc id:desc created created:asc
	// created:desc
	Sort param.Field[PrimaryIPListParamsSort] `query:"sort"`
}

// URLQuery serializes [PrimaryIPListParams]'s query parameters as `url.Values`.
func (r PrimaryIPListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times. Choices id id:asc id:desc created created:asc
// created:desc
type PrimaryIPListParamsSort string

const (
	PrimaryIPListParamsSortID          PrimaryIPListParamsSort = "id"
	PrimaryIPListParamsSortIDAsc       PrimaryIPListParamsSort = "id:asc"
	PrimaryIPListParamsSortIDDesc      PrimaryIPListParamsSort = "id:desc"
	PrimaryIPListParamsSortCreated     PrimaryIPListParamsSort = "created"
	PrimaryIPListParamsSortCreatedAsc  PrimaryIPListParamsSort = "created:asc"
	PrimaryIPListParamsSortCreatedDesc PrimaryIPListParamsSort = "created:desc"
)
