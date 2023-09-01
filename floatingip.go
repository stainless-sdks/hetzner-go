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

// FloatingIPService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewFloatingIPService] method instead.
type FloatingIPService struct {
	Options []option.RequestOption
	Actions *FloatingIPActionService
}

// NewFloatingIPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFloatingIPService(opts ...option.RequestOption) (r *FloatingIPService) {
	r = &FloatingIPService{}
	r.Options = opts
	r.Actions = NewFloatingIPActionService(opts...)
	return
}

// Creates a new Floating IP assigned to a Server. If you want to create a Floating
// IP that is not bound to a Server, you need to provide the `home_location` key
// instead of `server`. This can be either the ID or the name of the Location this
// IP shall be created in. Note that a Floating IP can be assigned to a Server in
// any Location later on. For optimal routing it is advised to use the Floating IP
// in the same Location it was created in.
func (r *FloatingIPService) New(ctx context.Context, body FloatingIPNewParams, opts ...option.RequestOption) (res *FloatingIPNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "floating_ips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a specific Floating IP object.
func (r *FloatingIPService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *FloatingIPGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the description or labels of a Floating IP. Also note that when updating
// labels, the Floating IP’s current set of labels will be replaced with the labels
// provided in the request body. So, for example, if you want to add a new label,
// you have to provide all existing labels plus the new label in the request body.
func (r *FloatingIPService) Update(ctx context.Context, id int64, body FloatingIPUpdateParams, opts ...option.RequestOption) (res *FloatingIPUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("floating_ips/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all Floating IP objects.
func (r *FloatingIPService) List(ctx context.Context, query FloatingIPListParams, opts ...option.RequestOption) (res *shared.FloatingIPsPage[FloatingIP], err error) {
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
func (r *FloatingIPService) ListAutoPaging(ctx context.Context, query FloatingIPListParams, opts ...option.RequestOption) *shared.FloatingIPsPageAutoPager[FloatingIP] {
	return shared.NewFloatingIPsPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a Floating IP. If it is currently assigned to a Server it will
// automatically get unassigned.
func (r *FloatingIPService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("floating_ips/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type FloatingIP struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Whether the IP is blocked
	Blocked bool `json:"blocked,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Description of the Resource
	Description string `json:"description,required,nullable"`
	// Array of reverse DNS entries
	DNSPtr []FloatingIPDNSPtr `json:"dns_ptr,required"`
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	HomeLocation FloatingIPHomeLocation `json:"home_location,required"`
	// IP address
	IP string `json:"ip,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Protection configuration for the Resource
	Protection FloatingIPProtection `json:"protection,required"`
	// ID of the Server the Floating IP is assigned to, null if it is not assigned at
	// all
	Server int64 `json:"server,required,nullable"`
	// The type of the IP
	Type FloatingIPType `json:"type,required"`
	JSON floatingIPJSON
}

// floatingIPJSON contains the JSON metadata for the struct [FloatingIP]
type floatingIPJSON struct {
	ID           apijson.Field
	Blocked      apijson.Field
	Created      apijson.Field
	Description  apijson.Field
	DNSPtr       apijson.Field
	HomeLocation apijson.Field
	IP           apijson.Field
	Labels       apijson.Field
	Name         apijson.Field
	Protection   apijson.Field
	Server       apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *FloatingIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPDNSPtr struct {
	// DNS pointer for the specific IP address
	DNSPtr string `json:"dns_ptr,required"`
	// Single IPv4 or IPv6 address | Single IPv6 address of this Server for which the
	// reverse DNS entry has been set up
	IP   string `json:"ip,required"`
	JSON floatingIPDNSPtrJSON
}

// floatingIPDNSPtrJSON contains the JSON metadata for the struct
// [FloatingIPDNSPtr]
type floatingIPDNSPtrJSON struct {
	DNSPtr      apijson.Field
	IP          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPDNSPtr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type FloatingIPHomeLocation struct {
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
	JSON        floatingIPHomeLocationJSON
}

// floatingIPHomeLocationJSON contains the JSON metadata for the struct
// [FloatingIPHomeLocation]
type floatingIPHomeLocationJSON struct {
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

func (r *FloatingIPHomeLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type FloatingIPProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   floatingIPProtectionJSON
}

// floatingIPProtectionJSON contains the JSON metadata for the struct
// [FloatingIPProtection]
type floatingIPProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the IP
type FloatingIPType string

const (
	FloatingIPTypeIpv4 FloatingIPType = "ipv4"
	FloatingIPTypeIpv6 FloatingIPType = "ipv6"
)

// Response to POST https://api.hetzner.cloud/v1/floating_ips
type FloatingIPNewResponse struct {
	FloatingIP FloatingIP `json:"floating_ip,required"`
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action"`
	JSON   floatingIPNewResponseJSON
}

// floatingIPNewResponseJSON contains the JSON metadata for the struct
// [FloatingIPNewResponse]
type floatingIPNewResponseJSON struct {
	FloatingIP  apijson.Field
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/floating_ips/{id}
type FloatingIPGetResponse struct {
	FloatingIP FloatingIP `json:"floating_ip,required"`
	JSON       floatingIPGetResponseJSON
}

// floatingIPGetResponseJSON contains the JSON metadata for the struct
// [FloatingIPGetResponse]
type floatingIPGetResponseJSON struct {
	FloatingIP  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to PUT https://api.hetzner.cloud/v1/floating_ips/{id}
type FloatingIPUpdateResponse struct {
	FloatingIP FloatingIP `json:"floating_ip,required"`
	JSON       floatingIPUpdateResponseJSON
}

// floatingIPUpdateResponseJSON contains the JSON metadata for the struct
// [FloatingIPUpdateResponse]
type floatingIPUpdateResponseJSON struct {
	FloatingIP  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FloatingIPNewParams struct {
	// The type of the IP
	Type        param.Field[FloatingIPNewParamsType] `json:"type,required"`
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

func (r FloatingIPNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of the IP
type FloatingIPNewParamsType string

const (
	FloatingIPNewParamsTypeIpv4 FloatingIPNewParamsType = "ipv4"
	FloatingIPNewParamsTypeIpv6 FloatingIPNewParamsType = "ipv6"
)

type FloatingIPUpdateParams struct {
	// New Description to set
	Description param.Field[string] `json:"description"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// New unique name to set
	Name param.Field[string] `json:"name"`
}

func (r FloatingIPUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FloatingIPListParams struct {
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
	Sort param.Field[FloatingIPListParamsSort] `query:"sort"`
}

// URLQuery serializes [FloatingIPListParams]'s query parameters as `url.Values`.
func (r FloatingIPListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times. Choices id id:asc id:desc created created:asc
// created:desc
type FloatingIPListParamsSort string

const (
	FloatingIPListParamsSortID          FloatingIPListParamsSort = "id"
	FloatingIPListParamsSortIDAsc       FloatingIPListParamsSort = "id:asc"
	FloatingIPListParamsSortIDDesc      FloatingIPListParamsSort = "id:desc"
	FloatingIPListParamsSortCreated     FloatingIPListParamsSort = "created"
	FloatingIPListParamsSortCreatedAsc  FloatingIPListParamsSort = "created:asc"
	FloatingIPListParamsSortCreatedDesc FloatingIPListParamsSort = "created:desc"
)
