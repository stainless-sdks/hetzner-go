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

// ServerTypeService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewServerTypeService] method instead.
type ServerTypeService struct {
	Options []option.RequestOption
}

// NewServerTypeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewServerTypeService(opts ...option.RequestOption) (r *ServerTypeService) {
	r = &ServerTypeService{}
	r.Options = opts
	return
}

// Gets a specific Server type object.
func (r *ServerTypeService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerTypeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("server_types/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets all Server type objects.
func (r *ServerTypeService) List(ctx context.Context, query ServerTypeListParams, opts ...option.RequestOption) (res *ServerTypeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "server_types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/server_types/{id}
type ServerTypeGetResponse struct {
	ServerType ServerTypeGetResponseServerType `json:"server_type,required"`
	JSON       serverTypeGetResponseJSON
}

// serverTypeGetResponseJSON contains the JSON metadata for the struct
// [ServerTypeGetResponse]
type serverTypeGetResponseJSON struct {
	ServerType  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServerTypeGetResponseServerType struct {
	// ID of the Server type
	ID int64 `json:"id,required"`
	// Type of cpu architecture this image is compatible with. | Type of cpu
	// architecture
	Architecture ServerTypeGetResponseServerTypeArchitecture `json:"architecture,required"`
	// Number of cpu cores a Server of this type will have
	Cores int64 `json:"cores,required"`
	// Type of cpu
	CpuType ServerTypeGetResponseServerTypeCpuType `json:"cpu_type,required"`
	// This field is deprecated. Use the deprecation object instead
	Deprecated bool `json:"deprecated,required,nullable"`
	// Description of the Server type
	Description string `json:"description,required"`
	// Disk size a Server of this type will have in GB
	Disk float64 `json:"disk,required"`
	// Free traffic per month in bytes
	IncludedTraffic int64 `json:"included_traffic,required"`
	// Memory a Server of this type will have in GB
	Memory float64 `json:"memory,required"`
	// Unique identifier of the Server type
	Name string `json:"name,required"`
	// Prices in different Locations
	Prices []ServerTypeGetResponseServerTypePrice `json:"prices,required"`
	// Type of Server boot drive. Local has higher speed. Network has better
	// availability.
	StorageType ServerTypeGetResponseServerTypeStorageType `json:"storage_type,required"`
	// Describes if, when & how the resources was deprecated. If this field is set to
	// `null` the resource is not deprecated. If it has a value, it is considered
	// deprecated.
	Deprecation ServerTypeGetResponseServerTypeDeprecation `json:"deprecation,nullable"`
	JSON        serverTypeGetResponseServerTypeJSON
}

// serverTypeGetResponseServerTypeJSON contains the JSON metadata for the struct
// [ServerTypeGetResponseServerType]
type serverTypeGetResponseServerTypeJSON struct {
	ID              apijson.Field
	Architecture    apijson.Field
	Cores           apijson.Field
	CpuType         apijson.Field
	Deprecated      apijson.Field
	Description     apijson.Field
	Disk            apijson.Field
	IncludedTraffic apijson.Field
	Memory          apijson.Field
	Name            apijson.Field
	Prices          apijson.Field
	StorageType     apijson.Field
	Deprecation     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ServerTypeGetResponseServerType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cpu architecture this image is compatible with. | Type of cpu
// architecture
type ServerTypeGetResponseServerTypeArchitecture string

const (
	ServerTypeGetResponseServerTypeArchitectureArm ServerTypeGetResponseServerTypeArchitecture = "arm"
	ServerTypeGetResponseServerTypeArchitectureX86 ServerTypeGetResponseServerTypeArchitecture = "x86"
)

// Type of cpu
type ServerTypeGetResponseServerTypeCpuType string

const (
	ServerTypeGetResponseServerTypeCpuTypeDedicated ServerTypeGetResponseServerTypeCpuType = "dedicated"
	ServerTypeGetResponseServerTypeCpuTypeShared    ServerTypeGetResponseServerTypeCpuType = "shared"
)

type ServerTypeGetResponseServerTypePrice struct {
	// Name of the Location the price is for
	Location string `json:"location,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceHourly ServerTypeGetResponseServerTypePricesPriceHourly `json:"price_hourly,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceMonthly ServerTypeGetResponseServerTypePricesPriceMonthly `json:"price_monthly,required"`
	JSON         serverTypeGetResponseServerTypePriceJSON
}

// serverTypeGetResponseServerTypePriceJSON contains the JSON metadata for the
// struct [ServerTypeGetResponseServerTypePrice]
type serverTypeGetResponseServerTypePriceJSON struct {
	Location     apijson.Field
	PriceHourly  apijson.Field
	PriceMonthly apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServerTypeGetResponseServerTypePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
// this Location | Monthly costs for a Floating IP type in this Location | Hourly
// costs for a Load Balancer type in this network zone | Monthly costs for a Load
// Balancer type in this network zone | Hourly costs for a Primary IP type in this
// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
// for a Server type in this Location | Monthly costs for a Server type in this
// Location
type ServerTypeGetResponseServerTypePricesPriceHourly struct {
	// Price with VAT added
	Gross string `json:"gross,required" format:"decimal"`
	// Price without VAT
	Net  string `json:"net,required" format:"decimal"`
	JSON serverTypeGetResponseServerTypePricesPriceHourlyJSON
}

// serverTypeGetResponseServerTypePricesPriceHourlyJSON contains the JSON metadata
// for the struct [ServerTypeGetResponseServerTypePricesPriceHourly]
type serverTypeGetResponseServerTypePricesPriceHourlyJSON struct {
	Gross       apijson.Field
	Net         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerTypeGetResponseServerTypePricesPriceHourly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
// this Location | Monthly costs for a Floating IP type in this Location | Hourly
// costs for a Load Balancer type in this network zone | Monthly costs for a Load
// Balancer type in this network zone | Hourly costs for a Primary IP type in this
// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
// for a Server type in this Location | Monthly costs for a Server type in this
// Location
type ServerTypeGetResponseServerTypePricesPriceMonthly struct {
	// Price with VAT added
	Gross string `json:"gross,required" format:"decimal"`
	// Price without VAT
	Net  string `json:"net,required" format:"decimal"`
	JSON serverTypeGetResponseServerTypePricesPriceMonthlyJSON
}

// serverTypeGetResponseServerTypePricesPriceMonthlyJSON contains the JSON metadata
// for the struct [ServerTypeGetResponseServerTypePricesPriceMonthly]
type serverTypeGetResponseServerTypePricesPriceMonthlyJSON struct {
	Gross       apijson.Field
	Net         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerTypeGetResponseServerTypePricesPriceMonthly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of Server boot drive. Local has higher speed. Network has better
// availability.
type ServerTypeGetResponseServerTypeStorageType string

const (
	ServerTypeGetResponseServerTypeStorageTypeLocal   ServerTypeGetResponseServerTypeStorageType = "local"
	ServerTypeGetResponseServerTypeStorageTypeNetwork ServerTypeGetResponseServerTypeStorageType = "network"
)

// Describes if, when & how the resources was deprecated. If this field is set to
// `null` the resource is not deprecated. If it has a value, it is considered
// deprecated.
type ServerTypeGetResponseServerTypeDeprecation struct {
	// Date of when the deprecation was announced.
	Announced string `json:"announced,required" format:"iso-8601"`
	// After the time in this field, the resource will not be available from the
	// general listing endpoint of the resource type, and it can not be used in new
	// resources. For example, if this is an image, you can not create new servers with
	// this image after the mentioned date.
	UnavailableAfter string `json:"unavailable_after,required" format:"iso-8601"`
	JSON             serverTypeGetResponseServerTypeDeprecationJSON
}

// serverTypeGetResponseServerTypeDeprecationJSON contains the JSON metadata for
// the struct [ServerTypeGetResponseServerTypeDeprecation]
type serverTypeGetResponseServerTypeDeprecationJSON struct {
	Announced        apijson.Field
	UnavailableAfter apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ServerTypeGetResponseServerTypeDeprecation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/server_types
type ServerTypeListResponse struct {
	ServerTypes []ServerTypeListResponseServerType `json:"server_types,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON serverTypeListResponseJSON
}

// serverTypeListResponseJSON contains the JSON metadata for the struct
// [ServerTypeListResponse]
type serverTypeListResponseJSON struct {
	ServerTypes apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServerTypeListResponseServerType struct {
	// ID of the Server type
	ID int64 `json:"id,required"`
	// Type of cpu architecture this image is compatible with. | Type of cpu
	// architecture
	Architecture ServerTypeListResponseServerTypesArchitecture `json:"architecture,required"`
	// Number of cpu cores a Server of this type will have
	Cores int64 `json:"cores,required"`
	// Type of cpu
	CpuType ServerTypeListResponseServerTypesCpuType `json:"cpu_type,required"`
	// This field is deprecated. Use the deprecation object instead
	Deprecated bool `json:"deprecated,required,nullable"`
	// Description of the Server type
	Description string `json:"description,required"`
	// Disk size a Server of this type will have in GB
	Disk float64 `json:"disk,required"`
	// Free traffic per month in bytes
	IncludedTraffic int64 `json:"included_traffic,required"`
	// Memory a Server of this type will have in GB
	Memory float64 `json:"memory,required"`
	// Unique identifier of the Server type
	Name string `json:"name,required"`
	// Prices in different Locations
	Prices []ServerTypeListResponseServerTypesPrice `json:"prices,required"`
	// Type of Server boot drive. Local has higher speed. Network has better
	// availability.
	StorageType ServerTypeListResponseServerTypesStorageType `json:"storage_type,required"`
	// Describes if, when & how the resources was deprecated. If this field is set to
	// `null` the resource is not deprecated. If it has a value, it is considered
	// deprecated.
	Deprecation ServerTypeListResponseServerTypesDeprecation `json:"deprecation,nullable"`
	JSON        serverTypeListResponseServerTypeJSON
}

// serverTypeListResponseServerTypeJSON contains the JSON metadata for the struct
// [ServerTypeListResponseServerType]
type serverTypeListResponseServerTypeJSON struct {
	ID              apijson.Field
	Architecture    apijson.Field
	Cores           apijson.Field
	CpuType         apijson.Field
	Deprecated      apijson.Field
	Description     apijson.Field
	Disk            apijson.Field
	IncludedTraffic apijson.Field
	Memory          apijson.Field
	Name            apijson.Field
	Prices          apijson.Field
	StorageType     apijson.Field
	Deprecation     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ServerTypeListResponseServerType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cpu architecture this image is compatible with. | Type of cpu
// architecture
type ServerTypeListResponseServerTypesArchitecture string

const (
	ServerTypeListResponseServerTypesArchitectureArm ServerTypeListResponseServerTypesArchitecture = "arm"
	ServerTypeListResponseServerTypesArchitectureX86 ServerTypeListResponseServerTypesArchitecture = "x86"
)

// Type of cpu
type ServerTypeListResponseServerTypesCpuType string

const (
	ServerTypeListResponseServerTypesCpuTypeDedicated ServerTypeListResponseServerTypesCpuType = "dedicated"
	ServerTypeListResponseServerTypesCpuTypeShared    ServerTypeListResponseServerTypesCpuType = "shared"
)

type ServerTypeListResponseServerTypesPrice struct {
	// Name of the Location the price is for
	Location string `json:"location,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceHourly ServerTypeListResponseServerTypesPricesPriceHourly `json:"price_hourly,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceMonthly ServerTypeListResponseServerTypesPricesPriceMonthly `json:"price_monthly,required"`
	JSON         serverTypeListResponseServerTypesPriceJSON
}

// serverTypeListResponseServerTypesPriceJSON contains the JSON metadata for the
// struct [ServerTypeListResponseServerTypesPrice]
type serverTypeListResponseServerTypesPriceJSON struct {
	Location     apijson.Field
	PriceHourly  apijson.Field
	PriceMonthly apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServerTypeListResponseServerTypesPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
// this Location | Monthly costs for a Floating IP type in this Location | Hourly
// costs for a Load Balancer type in this network zone | Monthly costs for a Load
// Balancer type in this network zone | Hourly costs for a Primary IP type in this
// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
// for a Server type in this Location | Monthly costs for a Server type in this
// Location
type ServerTypeListResponseServerTypesPricesPriceHourly struct {
	// Price with VAT added
	Gross string `json:"gross,required" format:"decimal"`
	// Price without VAT
	Net  string `json:"net,required" format:"decimal"`
	JSON serverTypeListResponseServerTypesPricesPriceHourlyJSON
}

// serverTypeListResponseServerTypesPricesPriceHourlyJSON contains the JSON
// metadata for the struct [ServerTypeListResponseServerTypesPricesPriceHourly]
type serverTypeListResponseServerTypesPricesPriceHourlyJSON struct {
	Gross       apijson.Field
	Net         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerTypeListResponseServerTypesPricesPriceHourly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
// this Location | Monthly costs for a Floating IP type in this Location | Hourly
// costs for a Load Balancer type in this network zone | Monthly costs for a Load
// Balancer type in this network zone | Hourly costs for a Primary IP type in this
// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
// for a Server type in this Location | Monthly costs for a Server type in this
// Location
type ServerTypeListResponseServerTypesPricesPriceMonthly struct {
	// Price with VAT added
	Gross string `json:"gross,required" format:"decimal"`
	// Price without VAT
	Net  string `json:"net,required" format:"decimal"`
	JSON serverTypeListResponseServerTypesPricesPriceMonthlyJSON
}

// serverTypeListResponseServerTypesPricesPriceMonthlyJSON contains the JSON
// metadata for the struct [ServerTypeListResponseServerTypesPricesPriceMonthly]
type serverTypeListResponseServerTypesPricesPriceMonthlyJSON struct {
	Gross       apijson.Field
	Net         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerTypeListResponseServerTypesPricesPriceMonthly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of Server boot drive. Local has higher speed. Network has better
// availability.
type ServerTypeListResponseServerTypesStorageType string

const (
	ServerTypeListResponseServerTypesStorageTypeLocal   ServerTypeListResponseServerTypesStorageType = "local"
	ServerTypeListResponseServerTypesStorageTypeNetwork ServerTypeListResponseServerTypesStorageType = "network"
)

// Describes if, when & how the resources was deprecated. If this field is set to
// `null` the resource is not deprecated. If it has a value, it is considered
// deprecated.
type ServerTypeListResponseServerTypesDeprecation struct {
	// Date of when the deprecation was announced.
	Announced string `json:"announced,required" format:"iso-8601"`
	// After the time in this field, the resource will not be available from the
	// general listing endpoint of the resource type, and it can not be used in new
	// resources. For example, if this is an image, you can not create new servers with
	// this image after the mentioned date.
	UnavailableAfter string `json:"unavailable_after,required" format:"iso-8601"`
	JSON             serverTypeListResponseServerTypesDeprecationJSON
}

// serverTypeListResponseServerTypesDeprecationJSON contains the JSON metadata for
// the struct [ServerTypeListResponseServerTypesDeprecation]
type serverTypeListResponseServerTypesDeprecationJSON struct {
	Announced        apijson.Field
	UnavailableAfter apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ServerTypeListResponseServerTypesDeprecation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServerTypeListParams struct {
	// Can be used to filter Server types by their name. The response will only contain
	// the Server type matching the specified name.
	Name param.Field[string] `query:"name"`
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ServerTypeListParams]'s query parameters as `url.Values`.
func (r ServerTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
