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

// FloatingIpService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewFloatingIpService] method instead.
type FloatingIpService struct {
	Options []option.RequestOption
	Actions *FloatingIpActionService
}

// NewFloatingIpService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFloatingIpService(opts ...option.RequestOption) (r *FloatingIpService) {
	r = &FloatingIpService{}
	r.Options = opts
	r.Actions = NewFloatingIpActionService(opts...)
	return
}

// Creates a new Floating IP assigned to a Server. If you want to create a Floating
// IP that is not bound to a Server, you need to provide the `home_location` key
// instead of `server`. This can be either the ID or the name of the Location this
// IP shall be created in. Note that a Floating IP can be assigned to a Server in
// any Location later on. For optimal routing it is advised to use the Floating IP
// in the same Location it was created in.
func (r *FloatingIpService) New(ctx context.Context, body FloatingIpNewParams, opts ...option.RequestOption) (res *FloatingIpNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "floating_ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a specific Floating IP object.
func (r *FloatingIpService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *FloatingIpGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the description or labels of a Floating IP. Also note that when updating
// labels, the Floating IP’s current set of labels will be replaced with the labels
// provided in the request body. So, for example, if you want to add a new label,
// you have to provide all existing labels plus the new label in the request body.
func (r *FloatingIpService) Update(ctx context.Context, id int64, body FloatingIpUpdateParams, opts ...option.RequestOption) (res *FloatingIpUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all Floating IP objects.
func (r *FloatingIpService) List(ctx context.Context, query FloatingIpListParams, opts ...option.RequestOption) (res *shared.FloatingIpsPage[FloatingIp], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "floating_ips"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns all Floating IP objects.
func (r *FloatingIpService) ListAutoPaging(ctx context.Context, query FloatingIpListParams, opts ...option.RequestOption) *shared.FloatingIpsPageAutoPager[FloatingIp] {
	return shared.NewFloatingIpsPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a Floating IP. If it is currently assigned to a Server it will
// automatically get unassigned.
func (r *FloatingIpService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("floating_ips/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type FloatingIp struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Whether the IP is blocked
	Blocked bool `json:"blocked,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Description of the Resource
	Description string `json:"description,required,nullable"`
	// Array of reverse DNS entries
	DnsPtr []FloatingIpDnsPtr `json:"dns_ptr,required"`
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	HomeLocation FloatingIpHomeLocation `json:"home_location,required"`
	// IP address
	Ip string `json:"ip,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Protection configuration for the Resource
	Protection FloatingIpProtection `json:"protection,required"`
	// ID of the Server the Floating IP is assigned to, null if it is not assigned at
	// all
	Server int64 `json:"server,required,nullable"`
	// The type of the IP
	Type FloatingIpType `json:"type,required"`
	JSON floatingIpJSON
}

// floatingIpJSON contains the JSON metadata for the struct [FloatingIp]
type floatingIpJSON struct {
	ID           apijson.Field
	Blocked      apijson.Field
	Created      apijson.Field
	Description  apijson.Field
	DnsPtr       apijson.Field
	HomeLocation apijson.Field
	Ip           apijson.Field
	Labels       apijson.Field
	Name         apijson.Field
	Protection   apijson.Field
	Server       apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FloatingIp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIpDnsPtr struct {
	// DNS pointer for the specific IP address
	DnsPtr string `json:"dns_ptr,required"`
	// Single IPv4 or IPv6 address | Single IPv6 address of this Server for which the
	// reverse DNS entry has been set up
	Ip   string `json:"ip,required"`
	JSON floatingIpDnsPtrJSON
}

// floatingIpDnsPtrJSON contains the JSON metadata for the struct
// [FloatingIpDnsPtr]
type floatingIpDnsPtrJSON struct {
	DnsPtr      apijson.Field
	Ip          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIpDnsPtr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type FloatingIpHomeLocation struct {
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
	JSON        floatingIpHomeLocationJSON
}

// floatingIpHomeLocationJSON contains the JSON metadata for the struct
// [FloatingIpHomeLocation]
type floatingIpHomeLocationJSON struct {
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

func (r *FloatingIpHomeLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type FloatingIpProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   floatingIpProtectionJSON
}

// floatingIpProtectionJSON contains the JSON metadata for the struct
// [FloatingIpProtection]
type floatingIpProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIpProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the IP
type FloatingIpType string

const (
	FloatingIpTypeIpv4 FloatingIpType = "ipv4"
	FloatingIpTypeIpv6 FloatingIpType = "ipv6"
)

// Response to POST https://api.hetzner.cloud/v1/floating_ips
type FloatingIpNewResponse struct {
	FloatingIp FloatingIp `json:"floating_ip,required"`
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action"`
	JSON   floatingIpNewResponseJSON
}

// floatingIpNewResponseJSON contains the JSON metadata for the struct
// [FloatingIpNewResponse]
type floatingIpNewResponseJSON struct {
	FloatingIp  apijson.Field
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIpNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/floating_ips/{id}
type FloatingIpGetResponse struct {
	FloatingIp FloatingIp `json:"floating_ip,required"`
	JSON       floatingIpGetResponseJSON
}

// floatingIpGetResponseJSON contains the JSON metadata for the struct
// [FloatingIpGetResponse]
type floatingIpGetResponseJSON struct {
	FloatingIp  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIpGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to PUT https://api.hetzner.cloud/v1/floating_ips/{id}
type FloatingIpUpdateResponse struct {
	FloatingIp FloatingIp `json:"floating_ip,required"`
	JSON       floatingIpUpdateResponseJSON
}

// floatingIpUpdateResponseJSON contains the JSON metadata for the struct
// [FloatingIpUpdateResponse]
type floatingIpUpdateResponseJSON struct {
	FloatingIp  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIpUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIpNewParams struct {
	// The type of the IP
	Type        param.Field[FloatingIpNewParamsType] `json:"type,required"`
	Description param.Field[string]                  `json:"description"`
	// Home Location (routing is optimized for that Location). Only optional if Server
	// argument is passed.
	HomeLocation param.Field[string] `json:"home_location"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	Name   param.Field[string]            `json:"name"`
	// ID of the Server to assign the Floating IP to
	Server param.Field[int64] `json:"server"`
}

func (r FloatingIpNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the IP
type FloatingIpNewParamsType string

const (
	FloatingIpNewParamsTypeIpv4 FloatingIpNewParamsType = "ipv4"
	FloatingIpNewParamsTypeIpv6 FloatingIpNewParamsType = "ipv6"
)

type FloatingIpUpdateParams struct {
	// New Description to set
	Description param.Field[string] `json:"description"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// New unique name to set
	Name param.Field[string] `json:"name"`
}

func (r FloatingIpUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FloatingIpListParams struct {
	// Can be used to filter Floating IPs by labels. The response will only contain
	// Floating IPs matching the label selector.
	LabelSelector param.Field[string] `query:"label_selector"`
	// Can be used to filter Floating IPs by their name. The response will only contain
	// the Floating IP matching the specified name.
	Name param.Field[string] `query:"name"`
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
	// Can be used multiple times. Choices id id:asc id:desc created created:asc
	// created:desc
	Sort param.Field[FloatingIpListParamsSort] `query:"sort"`
}

// URLQuery serializes [FloatingIpListParams]'s query parameters as `url.Values`.
func (r FloatingIpListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times. Choices id id:asc id:desc created created:asc
// created:desc
type FloatingIpListParamsSort string

const (
	FloatingIpListParamsSortID          FloatingIpListParamsSort = "id"
	FloatingIpListParamsSortIDAsc       FloatingIpListParamsSort = "id:asc"
	FloatingIpListParamsSortIDDesc      FloatingIpListParamsSort = "id:desc"
	FloatingIpListParamsSortCreated     FloatingIpListParamsSort = "created"
	FloatingIpListParamsSortCreatedAsc  FloatingIpListParamsSort = "created:asc"
	FloatingIpListParamsSortCreatedDesc FloatingIpListParamsSort = "created:desc"
)
