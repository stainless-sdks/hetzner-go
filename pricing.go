// File generated from our OpenAPI spec by Stainless.

package hetzner

import (
	"context"
	"net/http"

	"github.com/hetzner/hetzner-go/internal/apijson"
	"github.com/hetzner/hetzner-go/internal/requestconfig"
	"github.com/hetzner/hetzner-go/option"
)

// PricingService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPricingService] method instead.
type PricingService struct {
	Options []option.RequestOption
}

// NewPricingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPricingService(opts ...option.RequestOption) (r *PricingService) {
	r = &PricingService{}
	r.Options = opts
	return
}

// Returns prices for all resources available on the platform. VAT and currency of
// the Project owner are used for calculations.
//
// Both net and gross prices are included in the response.
func (r *PricingService) Get(ctx context.Context, opts ...option.RequestOption) (res *PricingGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "pricing"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FloatingIPPriceDetails struct {
	// Floating IP type costs per Location
	Prices []PricePerTimeMonthly `json:"prices,required"`
	// The type of the IP
	Type FloatingIPPriceDetailsType `json:"type,required"`
	JSON floatingIPPriceDetailsJSON
}

// floatingIPPriceDetailsJSON contains the JSON metadata for the struct
// [FloatingIPPriceDetails]
type floatingIPPriceDetailsJSON struct {
	Prices      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPPriceDetails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PricePerTimeMonthly struct {
	// Name of the Location the price is for
	Location string `json:"location,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceMonthly Price `json:"price_monthly,required"`
	JSON         pricePerTimeMonthlyJSON
}

// pricePerTimeMonthlyJSON contains the JSON metadata for the struct
// [PricePerTimeMonthly]
type pricePerTimeMonthlyJSON struct {
	Location     apijson.Field
	PriceMonthly apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PricePerTimeMonthly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
// this Location | Monthly costs for a Floating IP type in this Location | Hourly
// costs for a Load Balancer type in this network zone | Monthly costs for a Load
// Balancer type in this network zone | Hourly costs for a Primary IP type in this
// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
// for a Server type in this Location | Monthly costs for a Server type in this
// Location
type Price struct {
	// Price with VAT added
	Gross string `json:"gross,required" format:"decimal"`
	// Price without VAT
	Net  string `json:"net,required" format:"decimal"`
	JSON priceJSON
}

// priceJSON contains the JSON metadata for the struct [Price]
type priceJSON struct {
	Gross       apijson.Field
	Net         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Price) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the IP
type FloatingIPPriceDetailsType string

const (
	FloatingIPPriceDetailsTypeIpv4 FloatingIPPriceDetailsType = "ipv4"
	FloatingIPPriceDetailsTypeIpv6 FloatingIPPriceDetailsType = "ipv6"
)

// Response to GET https://api.hetzner.cloud/v1/pricing
type PricingGetResponse struct {
	Pricing PricingGetResponsePricing `json:"pricing,required"`
	JSON    pricingGetResponseJSON
}

// pricingGetResponseJSON contains the JSON metadata for the struct
// [PricingGetResponse]
type pricingGetResponseJSON struct {
	Pricing     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PricingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PricingGetResponsePricing struct {
	// Currency the returned prices are expressed in, coded according to ISO 4217
	Currency string `json:"currency,required"`
	// The cost of one Floating IP per month
	FloatingIP PricingGetResponsePricingFloatingIP `json:"floating_ip,required"`
	// Costs of Floating IPs types per Location and type
	FloatingIPs []FloatingIPPriceDetails `json:"floating_ips,required"`
	// The cost of Image per GB/month
	Image PricingGetResponsePricingImage `json:"image,required"`
	// Costs of Load Balancer types per Location and type
	LoadBalancerTypes []PricingGetResponsePricingLoadBalancerType `json:"load_balancer_types,required"`
	// Costs of Primary IPs types per Location
	PrimaryIPs []PricingGetResponsePricingPrimaryIP `json:"primary_ips,required"`
	// Will increase base Server costs by specific percentage
	ServerBackup PricingGetResponsePricingServerBackup `json:"server_backup,required"`
	// Costs of Server types per Location and type
	ServerTypes []PricingGetResponsePricingServerType `json:"server_types,required"`
	// The cost of additional traffic per TB
	Traffic PricingGetResponsePricingTraffic `json:"traffic,required"`
	// The VAT rate used for calculating prices with VAT
	VatRate string `json:"vat_rate,required"`
	// The cost of Volume per GB/month
	Volume PricingGetResponsePricingVolume `json:"volume,required"`
	JSON   pricingGetResponsePricingJSON
}

// pricingGetResponsePricingJSON contains the JSON metadata for the struct
// [PricingGetResponsePricing]
type pricingGetResponsePricingJSON struct {
	Currency          apijson.Field
	FloatingIP        apijson.Field
	FloatingIPs       apijson.Field
	Image             apijson.Field
	LoadBalancerTypes apijson.Field
	PrimaryIPs        apijson.Field
	ServerBackup      apijson.Field
	ServerTypes       apijson.Field
	Traffic           apijson.Field
	VatRate           apijson.Field
	Volume            apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *PricingGetResponsePricing) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The cost of one Floating IP per month
type PricingGetResponsePricingFloatingIP struct {
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceMonthly Price `json:"price_monthly,required"`
	JSON         pricingGetResponsePricingFloatingIPJSON
}

// pricingGetResponsePricingFloatingIPJSON contains the JSON metadata for the
// struct [PricingGetResponsePricingFloatingIP]
type pricingGetResponsePricingFloatingIPJSON struct {
	PriceMonthly apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PricingGetResponsePricingFloatingIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The cost of Image per GB/month
type PricingGetResponsePricingImage struct {
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PricePerGBMonth Price `json:"price_per_gb_month,required"`
	JSON            pricingGetResponsePricingImageJSON
}

// pricingGetResponsePricingImageJSON contains the JSON metadata for the struct
// [PricingGetResponsePricingImage]
type pricingGetResponsePricingImageJSON struct {
	PricePerGBMonth apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PricingGetResponsePricingImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PricingGetResponsePricingLoadBalancerType struct {
	// ID of the Load Balancer type the price is for
	ID int64 `json:"id,required"`
	// Name of the Load Balancer type the price is for
	Name string `json:"name,required"`
	// Load Balancer type costs per Location
	Prices []PricingGetResponsePricingLoadBalancerTypesPrice `json:"prices,required"`
	JSON   pricingGetResponsePricingLoadBalancerTypeJSON
}

// pricingGetResponsePricingLoadBalancerTypeJSON contains the JSON metadata for the
// struct [PricingGetResponsePricingLoadBalancerType]
type pricingGetResponsePricingLoadBalancerTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Prices      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PricingGetResponsePricingLoadBalancerType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PricingGetResponsePricingLoadBalancerTypesPrice struct {
	// Name of the Location the price is for
	Location string `json:"location,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceHourly Price `json:"price_hourly,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceMonthly Price `json:"price_monthly,required"`
	JSON         pricingGetResponsePricingLoadBalancerTypesPriceJSON
}

// pricingGetResponsePricingLoadBalancerTypesPriceJSON contains the JSON metadata
// for the struct [PricingGetResponsePricingLoadBalancerTypesPrice]
type pricingGetResponsePricingLoadBalancerTypesPriceJSON struct {
	Location     apijson.Field
	PriceHourly  apijson.Field
	PriceMonthly apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PricingGetResponsePricingLoadBalancerTypesPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PricingGetResponsePricingPrimaryIP struct {
	// Primary IP type costs per Location
	Prices []PricingGetResponsePricingPrimaryIPsPrice `json:"prices,required"`
	// The type of the IP
	Type PricingGetResponsePricingPrimaryIPsType `json:"type,required"`
	JSON pricingGetResponsePricingPrimaryIPJSON
}

// pricingGetResponsePricingPrimaryIPJSON contains the JSON metadata for the struct
// [PricingGetResponsePricingPrimaryIP]
type pricingGetResponsePricingPrimaryIPJSON struct {
	Prices      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PricingGetResponsePricingPrimaryIP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PricingGetResponsePricingPrimaryIPsPrice struct {
	// Name of the Location the price is for
	Location string `json:"location,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceHourly Price `json:"price_hourly,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceMonthly Price `json:"price_monthly,required"`
	JSON         pricingGetResponsePricingPrimaryIPsPriceJSON
}

// pricingGetResponsePricingPrimaryIPsPriceJSON contains the JSON metadata for the
// struct [PricingGetResponsePricingPrimaryIPsPrice]
type pricingGetResponsePricingPrimaryIPsPriceJSON struct {
	Location     apijson.Field
	PriceHourly  apijson.Field
	PriceMonthly apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PricingGetResponsePricingPrimaryIPsPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the IP
type PricingGetResponsePricingPrimaryIPsType string

const (
	PricingGetResponsePricingPrimaryIPsTypeIpv4 PricingGetResponsePricingPrimaryIPsType = "ipv4"
	PricingGetResponsePricingPrimaryIPsTypeIpv6 PricingGetResponsePricingPrimaryIPsType = "ipv6"
)

// Will increase base Server costs by specific percentage
type PricingGetResponsePricingServerBackup struct {
	// Percentage by how much the base price will increase
	Percentage string `json:"percentage,required"`
	JSON       pricingGetResponsePricingServerBackupJSON
}

// pricingGetResponsePricingServerBackupJSON contains the JSON metadata for the
// struct [PricingGetResponsePricingServerBackup]
type pricingGetResponsePricingServerBackupJSON struct {
	Percentage  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PricingGetResponsePricingServerBackup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PricingGetResponsePricingServerType struct {
	// ID of the Server type the price is for
	ID int64 `json:"id,required"`
	// Name of the Server type the price is for
	Name string `json:"name,required"`
	// Server type costs per Location
	Prices []PricingGetResponsePricingServerTypesPrice `json:"prices,required"`
	JSON   pricingGetResponsePricingServerTypeJSON
}

// pricingGetResponsePricingServerTypeJSON contains the JSON metadata for the
// struct [PricingGetResponsePricingServerType]
type pricingGetResponsePricingServerTypeJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Prices      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PricingGetResponsePricingServerType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PricingGetResponsePricingServerTypesPrice struct {
	// Name of the Location the price is for
	Location string `json:"location,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceHourly Price `json:"price_hourly,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceMonthly Price `json:"price_monthly,required"`
	JSON         pricingGetResponsePricingServerTypesPriceJSON
}

// pricingGetResponsePricingServerTypesPriceJSON contains the JSON metadata for the
// struct [PricingGetResponsePricingServerTypesPrice]
type pricingGetResponsePricingServerTypesPriceJSON struct {
	Location     apijson.Field
	PriceHourly  apijson.Field
	PriceMonthly apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *PricingGetResponsePricingServerTypesPrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The cost of additional traffic per TB
type PricingGetResponsePricingTraffic struct {
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PricePerTb Price `json:"price_per_tb,required"`
	JSON       pricingGetResponsePricingTrafficJSON
}

// pricingGetResponsePricingTrafficJSON contains the JSON metadata for the struct
// [PricingGetResponsePricingTraffic]
type pricingGetResponsePricingTrafficJSON struct {
	PricePerTb  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PricingGetResponsePricingTraffic) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The cost of Volume per GB/month
type PricingGetResponsePricingVolume struct {
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PricePerGBMonth Price `json:"price_per_gb_month,required"`
	JSON            pricingGetResponsePricingVolumeJSON
}

// pricingGetResponsePricingVolumeJSON contains the JSON metadata for the struct
// [PricingGetResponsePricingVolume]
type pricingGetResponsePricingVolumeJSON struct {
	PricePerGBMonth apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *PricingGetResponsePricingVolume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
