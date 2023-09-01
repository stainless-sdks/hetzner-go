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

// PrimaryIpService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPrimaryIpService] method instead.
type PrimaryIpService struct {
	Options []option.RequestOption
	Actions *PrimaryIpActionService
}

// NewPrimaryIpService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPrimaryIpService(opts ...option.RequestOption) (r *PrimaryIpService) {
	r = &PrimaryIpService{}
	r.Options = opts
	r.Actions = NewPrimaryIpActionService(opts...)
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
func (r *PrimaryIpService) New(ctx context.Context, body PrimaryIpNewParams, opts ...option.RequestOption) (res *PrimaryIpNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "primary_ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a specific Primary IP object.
func (r *PrimaryIpService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *PrimaryIpGetResponse, err error) {
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
func (r *PrimaryIpService) Update(ctx context.Context, id int64, body PrimaryIpUpdateParams, opts ...option.RequestOption) (res *PrimaryIpUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("primary_ips/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all Primary IP objects.
func (r *PrimaryIpService) List(ctx context.Context, query PrimaryIpListParams, opts ...option.RequestOption) (res *PrimaryIpListResponse, err error) {
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
func (r *PrimaryIpService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("primary_ips/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type PrimaryIp struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// ID of the resource the Primary IP is assigned to, null if it is not assigned at
	// all
	AssigneeID int64 `json:"assignee_id,required,nullable"`
	// Resource type the Primary IP can be assigned to
	AssigneeType PrimaryIpAssigneeType `json:"assignee_type,required"`
	// Delete this Primary IP when the resource it is assigned to is deleted
	AutoDelete bool `json:"auto_delete,required"`
	// Whether the IP is blocked
	Blocked bool `json:"blocked,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Datacenter this Primary IP is located at | Datacenter this Resource is located
	// at
	Datacenter PrimaryIpDatacenter `json:"datacenter,required"`
	// Array of reverse DNS entries
	DnsPtr []PrimaryIpDnsPtr `json:"dns_ptr,required"`
	// IP address
	Ip string `json:"ip,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Protection configuration for the Resource
	Protection PrimaryIpProtection `json:"protection,required"`
	// The type of the IP
	Type PrimaryIpType `json:"type,required"`
	JSON primaryIpJSON
}

// primaryIpJSON contains the JSON metadata for the struct [PrimaryIp]
type primaryIpJSON struct {
	ID           apijson.Field
	AssigneeID   apijson.Field
	AssigneeType apijson.Field
	AutoDelete   apijson.Field
	Blocked      apijson.Field
	Created      apijson.Field
	Datacenter   apijson.Field
	DnsPtr       apijson.Field
	Ip           apijson.Field
	Labels       apijson.Field
	Name         apijson.Field
	Protection   apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PrimaryIp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type the Primary IP can be assigned to
type PrimaryIpAssigneeType string

const (
	PrimaryIpAssigneeTypeServer PrimaryIpAssigneeType = "server"
)

// Datacenter this Primary IP is located at | Datacenter this Resource is located
// at
type PrimaryIpDatacenter struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Description of the Datacenter
	Description string `json:"description,required"`
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	Location PrimaryIpDatacenterLocation `json:"location,required"`
	// Unique identifier of the Datacenter
	Name string `json:"name,required"`
	// The Server types the Datacenter can handle
	ServerTypes PrimaryIpDatacenterServerTypes `json:"server_types,required"`
	JSON        primaryIpDatacenterJSON
}

// primaryIpDatacenterJSON contains the JSON metadata for the struct
// [PrimaryIpDatacenter]
type primaryIpDatacenterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	ServerTypes apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpDatacenter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type PrimaryIpDatacenterLocation struct {
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
	JSON        primaryIpDatacenterLocationJSON
}

// primaryIpDatacenterLocationJSON contains the JSON metadata for the struct
// [PrimaryIpDatacenterLocation]
type primaryIpDatacenterLocationJSON struct {
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

func (r *PrimaryIpDatacenterLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Server types the Datacenter can handle
type PrimaryIpDatacenterServerTypes struct {
	// IDs of Server types that are supported and for which the Datacenter has enough
	// resources left
	Available []int64 `json:"available,required"`
	// IDs of Server types that are supported and for which the Datacenter has enough
	// resources left
	AvailableForMigration []int64 `json:"available_for_migration,required"`
	// IDs of Server types that are supported in the Datacenter
	Supported []int64 `json:"supported,required"`
	JSON      primaryIpDatacenterServerTypesJSON
}

// primaryIpDatacenterServerTypesJSON contains the JSON metadata for the struct
// [PrimaryIpDatacenterServerTypes]
type primaryIpDatacenterServerTypesJSON struct {
	Available             apijson.Field
	AvailableForMigration apijson.Field
	Supported             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *PrimaryIpDatacenterServerTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PrimaryIpDnsPtr struct {
	// DNS pointer for the specific IP address
	DnsPtr string `json:"dns_ptr,required"`
	// Single IPv4 or IPv6 address | Single IPv6 address of this Server for which the
	// reverse DNS entry has been set up
	Ip   string `json:"ip,required"`
	JSON primaryIpDnsPtrJSON
}

// primaryIpDnsPtrJSON contains the JSON metadata for the struct [PrimaryIpDnsPtr]
type primaryIpDnsPtrJSON struct {
	DnsPtr      apijson.Field
	Ip          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpDnsPtr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type PrimaryIpProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   primaryIpProtectionJSON
}

// primaryIpProtectionJSON contains the JSON metadata for the struct
// [PrimaryIpProtection]
type primaryIpProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the IP
type PrimaryIpType string

const (
	PrimaryIpTypeIpv4 PrimaryIpType = "ipv4"
	PrimaryIpTypeIpv6 PrimaryIpType = "ipv6"
)

// Response to POST https://api.hetzner.cloud/v1/primary_ips
type PrimaryIpNewResponse struct {
	PrimaryIp PrimaryIp `json:"primary_ip,required"`
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action"`
	JSON   primaryIpNewResponseJSON
}

// primaryIpNewResponseJSON contains the JSON metadata for the struct
// [PrimaryIpNewResponse]
type primaryIpNewResponseJSON struct {
	PrimaryIp   apijson.Field
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/primary_ips/{id}
type PrimaryIpGetResponse struct {
	PrimaryIp PrimaryIp `json:"primary_ip,required"`
	JSON      primaryIpGetResponseJSON
}

// primaryIpGetResponseJSON contains the JSON metadata for the struct
// [PrimaryIpGetResponse]
type primaryIpGetResponseJSON struct {
	PrimaryIp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to PUT https://api.hetzner.cloud/v1/primary_ips/{id}
type PrimaryIpUpdateResponse struct {
	PrimaryIp PrimaryIp `json:"primary_ip,required"`
	JSON      primaryIpUpdateResponseJSON
}

// primaryIpUpdateResponseJSON contains the JSON metadata for the struct
// [PrimaryIpUpdateResponse]
type primaryIpUpdateResponseJSON struct {
	PrimaryIp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/primary_ips
type PrimaryIpListResponse struct {
	PrimaryIps []PrimaryIp `json:"primary_ips,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON primaryIpListResponseJSON
}

// primaryIpListResponseJSON contains the JSON metadata for the struct
// [PrimaryIpListResponse]
type primaryIpListResponseJSON struct {
	PrimaryIps  apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PrimaryIpListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PrimaryIpNewParams struct {
	// Resource type the Primary IP can be assigned to
	AssigneeType param.Field[PrimaryIpNewParamsAssigneeType] `json:"assignee_type,required"`
	Name         param.Field[string]                         `json:"name,required"`
	// The type of the IP
	Type param.Field[PrimaryIpNewParamsType] `json:"type,required"`
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

func (r PrimaryIpNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Resource type the Primary IP can be assigned to
type PrimaryIpNewParamsAssigneeType string

const (
	PrimaryIpNewParamsAssigneeTypeServer PrimaryIpNewParamsAssigneeType = "server"
)

// The type of the IP
type PrimaryIpNewParamsType string

const (
	PrimaryIpNewParamsTypeIpv4 PrimaryIpNewParamsType = "ipv4"
	PrimaryIpNewParamsTypeIpv6 PrimaryIpNewParamsType = "ipv6"
)

type PrimaryIpUpdateParams struct {
	// Delete this Primary IP when the resource it is assigned to is deleted
	AutoDelete param.Field[bool] `json:"auto_delete"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// New unique name to set
	Name param.Field[string] `json:"name"`
}

func (r PrimaryIpUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PrimaryIpListParams struct {
	// Can be used to filter resources by their ip. The response will only contain the
	// resources matching the specified ip.
	Ip param.Field[string] `query:"ip"`
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
	Sort param.Field[PrimaryIpListParamsSort] `query:"sort"`
}

// URLQuery serializes [PrimaryIpListParams]'s query parameters as `url.Values`.
func (r PrimaryIpListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times. Choices id id:asc id:desc created created:asc
// created:desc
type PrimaryIpListParamsSort string

const (
	PrimaryIpListParamsSortID          PrimaryIpListParamsSort = "id"
	PrimaryIpListParamsSortIDAsc       PrimaryIpListParamsSort = "id:asc"
	PrimaryIpListParamsSortIDDesc      PrimaryIpListParamsSort = "id:desc"
	PrimaryIpListParamsSortCreated     PrimaryIpListParamsSort = "created"
	PrimaryIpListParamsSortCreatedAsc  PrimaryIpListParamsSort = "created:asc"
	PrimaryIpListParamsSortCreatedDesc PrimaryIpListParamsSort = "created:desc"
)
