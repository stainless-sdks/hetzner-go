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

// LoadBalancerTypeService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerTypeService] method
// instead.
type LoadBalancerTypeService struct {
	Options []option.RequestOption
}

// NewLoadBalancerTypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerTypeService(opts ...option.RequestOption) (r *LoadBalancerTypeService) {
	r = &LoadBalancerTypeService{}
	r.Options = opts
	return
}

// Gets a specific Load Balancer type object.
func (r *LoadBalancerTypeService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *LoadBalancerTypeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancer_types/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets all Load Balancer type objects.
func (r *LoadBalancerTypeService) List(ctx context.Context, query LoadBalancerTypeListParams, opts ...option.RequestOption) (res *LoadBalancerTypeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "load_balancer_types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/load_balancer_types/{id}
type LoadBalancerTypeGetResponse struct {
	LoadBalancerType LoadBalancerTypeGetResponseLoadBalancerType `json:"load_balancer_type"`
	JSON             loadBalancerTypeGetResponseJSON
}

// loadBalancerTypeGetResponseJSON contains the JSON metadata for the struct
// [LoadBalancerTypeGetResponse]
type loadBalancerTypeGetResponseJSON struct {
	LoadBalancerType apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LoadBalancerTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerTypeGetResponseLoadBalancerType struct {
	// ID of the Load Balancer type
	ID int64 `json:"id,required"`
	// Point in time when the Load Balancer type is deprecated (in ISO-8601 format)
	Deprecated string `json:"deprecated,required,nullable"`
	// Description of the Load Balancer type
	Description string `json:"description,required"`
	// Number of SSL Certificates that can be assigned to a single Load Balancer
	MaxAssignedCertificates int64 `json:"max_assigned_certificates,required"`
	// Number of maximum simultaneous open connections
	MaxConnections int64 `json:"max_connections,required"`
	// Number of services a Load Balancer of this type can have
	MaxServices int64 `json:"max_services,required"`
	// Number of targets a single Load Balancer can have
	MaxTargets int64 `json:"max_targets,required"`
	// Unique identifier of the Load Balancer type
	Name string `json:"name,required"`
	// Prices in different network zones
	Prices []LoadBalancerTypeGetResponseLoadBalancerTypePrice `json:"prices,required"`
	JSON   loadBalancerTypeGetResponseLoadBalancerTypeJSON
}

// loadBalancerTypeGetResponseLoadBalancerTypeJSON contains the JSON metadata for
// the struct [LoadBalancerTypeGetResponseLoadBalancerType]
type loadBalancerTypeGetResponseLoadBalancerTypeJSON struct {
	ID                      apijson.Field
	Deprecated              apijson.Field
	Description             apijson.Field
	MaxAssignedCertificates apijson.Field
	MaxConnections          apijson.Field
	MaxServices             apijson.Field
	MaxTargets              apijson.Field
	Name                    apijson.Field
	Prices                  apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *LoadBalancerTypeGetResponseLoadBalancerType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerTypeGetResponseLoadBalancerTypePrice struct {
	// Name of the Location the price is for
	Location string `json:"location,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceHourly LoadBalancerTypeGetResponseLoadBalancerTypePricesPriceHourly `json:"price_hourly,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceMonthly LoadBalancerTypeGetResponseLoadBalancerTypePricesPriceMonthly `json:"price_monthly,required"`
	JSON         loadBalancerTypeGetResponseLoadBalancerTypePriceJSON
}

// loadBalancerTypeGetResponseLoadBalancerTypePriceJSON contains the JSON metadata
// for the struct [LoadBalancerTypeGetResponseLoadBalancerTypePrice]
type loadBalancerTypeGetResponseLoadBalancerTypePriceJSON struct {
	Location     apijson.Field
	PriceHourly  apijson.Field
	PriceMonthly apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LoadBalancerTypeGetResponseLoadBalancerTypePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
// this Location | Monthly costs for a Floating IP type in this Location | Hourly
// costs for a Load Balancer type in this network zone | Monthly costs for a Load
// Balancer type in this network zone | Hourly costs for a Primary IP type in this
// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
// for a Server type in this Location | Monthly costs for a Server type in this
// Location
type LoadBalancerTypeGetResponseLoadBalancerTypePricesPriceHourly struct {
	// Price with VAT added
	Gross string `json:"gross,required" format:"decimal"`
	// Price without VAT
	Net  string `json:"net,required" format:"decimal"`
	JSON loadBalancerTypeGetResponseLoadBalancerTypePricesPriceHourlyJSON
}

// loadBalancerTypeGetResponseLoadBalancerTypePricesPriceHourlyJSON contains the
// JSON metadata for the struct
// [LoadBalancerTypeGetResponseLoadBalancerTypePricesPriceHourly]
type loadBalancerTypeGetResponseLoadBalancerTypePricesPriceHourlyJSON struct {
	Gross       apijson.Field
	Net         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerTypeGetResponseLoadBalancerTypePricesPriceHourly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
// this Location | Monthly costs for a Floating IP type in this Location | Hourly
// costs for a Load Balancer type in this network zone | Monthly costs for a Load
// Balancer type in this network zone | Hourly costs for a Primary IP type in this
// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
// for a Server type in this Location | Monthly costs for a Server type in this
// Location
type LoadBalancerTypeGetResponseLoadBalancerTypePricesPriceMonthly struct {
	// Price with VAT added
	Gross string `json:"gross,required" format:"decimal"`
	// Price without VAT
	Net  string `json:"net,required" format:"decimal"`
	JSON loadBalancerTypeGetResponseLoadBalancerTypePricesPriceMonthlyJSON
}

// loadBalancerTypeGetResponseLoadBalancerTypePricesPriceMonthlyJSON contains the
// JSON metadata for the struct
// [LoadBalancerTypeGetResponseLoadBalancerTypePricesPriceMonthly]
type loadBalancerTypeGetResponseLoadBalancerTypePricesPriceMonthlyJSON struct {
	Gross       apijson.Field
	Net         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerTypeGetResponseLoadBalancerTypePricesPriceMonthly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/load_balancer_types
type LoadBalancerTypeListResponse struct {
	LoadBalancerTypes []LoadBalancerTypeListResponseLoadBalancerType `json:"load_balancer_types,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON loadBalancerTypeListResponseJSON
}

// loadBalancerTypeListResponseJSON contains the JSON metadata for the struct
// [LoadBalancerTypeListResponse]
type loadBalancerTypeListResponseJSON struct {
	LoadBalancerTypes apijson.Field
	Meta              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *LoadBalancerTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerTypeListResponseLoadBalancerType struct {
	// ID of the Load Balancer type
	ID int64 `json:"id,required"`
	// Point in time when the Load Balancer type is deprecated (in ISO-8601 format)
	Deprecated string `json:"deprecated,required,nullable"`
	// Description of the Load Balancer type
	Description string `json:"description,required"`
	// Number of SSL Certificates that can be assigned to a single Load Balancer
	MaxAssignedCertificates int64 `json:"max_assigned_certificates,required"`
	// Number of maximum simultaneous open connections
	MaxConnections int64 `json:"max_connections,required"`
	// Number of services a Load Balancer of this type can have
	MaxServices int64 `json:"max_services,required"`
	// Number of targets a single Load Balancer can have
	MaxTargets int64 `json:"max_targets,required"`
	// Unique identifier of the Load Balancer type
	Name string `json:"name,required"`
	// Prices in different network zones
	Prices []LoadBalancerTypeListResponseLoadBalancerTypesPrice `json:"prices,required"`
	JSON   loadBalancerTypeListResponseLoadBalancerTypeJSON
}

// loadBalancerTypeListResponseLoadBalancerTypeJSON contains the JSON metadata for
// the struct [LoadBalancerTypeListResponseLoadBalancerType]
type loadBalancerTypeListResponseLoadBalancerTypeJSON struct {
	ID                      apijson.Field
	Deprecated              apijson.Field
	Description             apijson.Field
	MaxAssignedCertificates apijson.Field
	MaxConnections          apijson.Field
	MaxServices             apijson.Field
	MaxTargets              apijson.Field
	Name                    apijson.Field
	Prices                  apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *LoadBalancerTypeListResponseLoadBalancerType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerTypeListResponseLoadBalancerTypesPrice struct {
	// Name of the Location the price is for
	Location string `json:"location,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceHourly LoadBalancerTypeListResponseLoadBalancerTypesPricesPriceHourly `json:"price_hourly,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceMonthly LoadBalancerTypeListResponseLoadBalancerTypesPricesPriceMonthly `json:"price_monthly,required"`
	JSON         loadBalancerTypeListResponseLoadBalancerTypesPriceJSON
}

// loadBalancerTypeListResponseLoadBalancerTypesPriceJSON contains the JSON
// metadata for the struct [LoadBalancerTypeListResponseLoadBalancerTypesPrice]
type loadBalancerTypeListResponseLoadBalancerTypesPriceJSON struct {
	Location     apijson.Field
	PriceHourly  apijson.Field
	PriceMonthly apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LoadBalancerTypeListResponseLoadBalancerTypesPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
// this Location | Monthly costs for a Floating IP type in this Location | Hourly
// costs for a Load Balancer type in this network zone | Monthly costs for a Load
// Balancer type in this network zone | Hourly costs for a Primary IP type in this
// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
// for a Server type in this Location | Monthly costs for a Server type in this
// Location
type LoadBalancerTypeListResponseLoadBalancerTypesPricesPriceHourly struct {
	// Price with VAT added
	Gross string `json:"gross,required" format:"decimal"`
	// Price without VAT
	Net  string `json:"net,required" format:"decimal"`
	JSON loadBalancerTypeListResponseLoadBalancerTypesPricesPriceHourlyJSON
}

// loadBalancerTypeListResponseLoadBalancerTypesPricesPriceHourlyJSON contains the
// JSON metadata for the struct
// [LoadBalancerTypeListResponseLoadBalancerTypesPricesPriceHourly]
type loadBalancerTypeListResponseLoadBalancerTypesPricesPriceHourlyJSON struct {
	Gross       apijson.Field
	Net         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerTypeListResponseLoadBalancerTypesPricesPriceHourly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
// this Location | Monthly costs for a Floating IP type in this Location | Hourly
// costs for a Load Balancer type in this network zone | Monthly costs for a Load
// Balancer type in this network zone | Hourly costs for a Primary IP type in this
// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
// for a Server type in this Location | Monthly costs for a Server type in this
// Location
type LoadBalancerTypeListResponseLoadBalancerTypesPricesPriceMonthly struct {
	// Price with VAT added
	Gross string `json:"gross,required" format:"decimal"`
	// Price without VAT
	Net  string `json:"net,required" format:"decimal"`
	JSON loadBalancerTypeListResponseLoadBalancerTypesPricesPriceMonthlyJSON
}

// loadBalancerTypeListResponseLoadBalancerTypesPricesPriceMonthlyJSON contains the
// JSON metadata for the struct
// [LoadBalancerTypeListResponseLoadBalancerTypesPricesPriceMonthly]
type loadBalancerTypeListResponseLoadBalancerTypesPricesPriceMonthlyJSON struct {
	Gross       apijson.Field
	Net         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerTypeListResponseLoadBalancerTypesPricesPriceMonthly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerTypeListParams struct {
	// Can be used to filter Load Balancer types by their name. The response will only
	// contain the Load Balancer type matching the specified name.
	Name param.Field[string] `query:"name"`
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [LoadBalancerTypeListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
