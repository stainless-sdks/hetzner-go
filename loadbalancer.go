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

// LoadBalancerService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerService] method
// instead.
type LoadBalancerService struct {
	Options []option.RequestOption
	Actions *LoadBalancerActionService
	Metrics *LoadBalancerMetricService
}

// NewLoadBalancerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLoadBalancerService(opts ...option.RequestOption) (r *LoadBalancerService) {
	r = &LoadBalancerService{}
	r.Options = opts
	r.Actions = NewLoadBalancerActionService(opts...)
	r.Metrics = NewLoadBalancerMetricService(opts...)
	return
}

// Creates a Load Balancer.
//
// #### Call specific error codes
//
// | Code                                    | Description                                                                                           |
// | --------------------------------------- | ----------------------------------------------------------------------------------------------------- |
// | `cloud_resource_ip_not_allowed`         | The IP you are trying to add as a target belongs to a Hetzner Cloud resource                          |
// | `ip_not_owned`                          | The IP is not owned by the owner of the project of the Load Balancer                                  |
// | `load_balancer_not_attached_to_network` | The Load Balancer is not attached to a network                                                        |
// | `robot_unavailable`                     | Robot was not available. The caller may retry the operation after a short delay.                      |
// | `server_not_attached_to_network`        | The server you are trying to add as a target is not attached to the same network as the Load Balancer |
// | `source_port_already_used`              | The source port you are trying to add is already in use                                               |
// | `target_already_defined`                | The Load Balancer target you are trying to define is already defined                                  |
func (r *LoadBalancerService) New(ctx context.Context, body LoadBalancerNewParams, opts ...option.RequestOption) (res *LoadBalancerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "load_balancers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a specific Load Balancer object.
func (r *LoadBalancerService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *LoadBalancerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a Load Balancer. You can update a Load Balancer’s name and a Load
// Balancer’s labels.
//
// Note that when updating labels, the Load Balancer’s current set of labels will
// be replaced with the labels provided in the request body. So, for example, if
// you want to add a new label, you have to provide all existing labels plus the
// new label in the request body.
//
// Note: if the Load Balancer object changes during the request, the response will
// be a “conflict” error.
func (r *LoadBalancerService) Update(ctx context.Context, id int64, body LoadBalancerUpdateParams, opts ...option.RequestOption) (res *LoadBalancerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Gets all existing Load Balancers that you have available.
func (r *LoadBalancerService) List(ctx context.Context, query LoadBalancerListParams, opts ...option.RequestOption) (res *LoadBalancerListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "load_balancers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a Load Balancer.
func (r *LoadBalancerService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("load_balancers/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type LoadBalancer struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Algorithm of the Load Balancer | Request for POST
	// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/change_algorithm
	Algorithm LoadBalancerAlgorithm `json:"algorithm,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Free Traffic for the current billing period in bytes
	IncludedTraffic int64 `json:"included_traffic,required"`
	// Inbound Traffic for the current billing period in bytes
	IngoingTraffic int64 `json:"ingoing_traffic,required,nullable"`
	// User-defined labels (key-value pairs)
	Labels           map[string]string            `json:"labels,required"`
	LoadBalancerType LoadBalancerLoadBalancerType `json:"load_balancer_type,required"`
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	Location LoadBalancerLocation `json:"location,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Outbound Traffic for the current billing period in bytes
	OutgoingTraffic int64 `json:"outgoing_traffic,required,nullable"`
	// Private networks information
	PrivateNet []LoadBalancerPrivateNet `json:"private_net,required"`
	// Protection configuration for the Resource
	Protection LoadBalancerProtection `json:"protection,required"`
	// Public network information
	PublicNet LoadBalancerPublicNet `json:"public_net,required"`
	// List of services that belong to this Load Balancer
	Services []LoadBalancerServiceModel `json:"services,required"`
	// List of targets that belong to this Load Balancer
	Targets []LoadBalancerTarget `json:"targets,required"`
	JSON    loadBalancerJSON
}

// loadBalancerJSON contains the JSON metadata for the struct [LoadBalancer]
type loadBalancerJSON struct {
	ID               apijson.Field
	Algorithm        apijson.Field
	Created          apijson.Field
	IncludedTraffic  apijson.Field
	IngoingTraffic   apijson.Field
	Labels           apijson.Field
	LoadBalancerType apijson.Field
	Location         apijson.Field
	Name             apijson.Field
	OutgoingTraffic  apijson.Field
	PrivateNet       apijson.Field
	Protection       apijson.Field
	PublicNet        apijson.Field
	Services         apijson.Field
	Targets          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *LoadBalancer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Algorithm of the Load Balancer | Request for POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/change_algorithm
type LoadBalancerAlgorithm struct {
	// Type of the algorithm | Algorithm of the Load Balancer
	Type LoadBalancerAlgorithmType `json:"type,required"`
	JSON loadBalancerAlgorithmJSON
}

// loadBalancerAlgorithmJSON contains the JSON metadata for the struct
// [LoadBalancerAlgorithm]
type loadBalancerAlgorithmJSON struct {
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerAlgorithm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the algorithm | Algorithm of the Load Balancer
type LoadBalancerAlgorithmType string

const (
	LoadBalancerAlgorithmTypeLeastConnections LoadBalancerAlgorithmType = "least_connections"
	LoadBalancerAlgorithmTypeRoundRobin       LoadBalancerAlgorithmType = "round_robin"
)

type LoadBalancerLoadBalancerType struct {
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
	Prices []LoadBalancerLoadBalancerTypePrice `json:"prices,required"`
	JSON   loadBalancerLoadBalancerTypeJSON
}

// loadBalancerLoadBalancerTypeJSON contains the JSON metadata for the struct
// [LoadBalancerLoadBalancerType]
type loadBalancerLoadBalancerTypeJSON struct {
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

func (r *LoadBalancerLoadBalancerType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerLoadBalancerTypePrice struct {
	// Name of the Location the price is for
	Location string `json:"location,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceHourly LoadBalancerLoadBalancerTypePricesPriceHourly `json:"price_hourly,required"`
	// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
	// this Location | Monthly costs for a Floating IP type in this Location | Hourly
	// costs for a Load Balancer type in this network zone | Monthly costs for a Load
	// Balancer type in this network zone | Hourly costs for a Primary IP type in this
	// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
	// for a Server type in this Location | Monthly costs for a Server type in this
	// Location
	PriceMonthly LoadBalancerLoadBalancerTypePricesPriceMonthly `json:"price_monthly,required"`
	JSON         loadBalancerLoadBalancerTypePriceJSON
}

// loadBalancerLoadBalancerTypePriceJSON contains the JSON metadata for the struct
// [LoadBalancerLoadBalancerTypePrice]
type loadBalancerLoadBalancerTypePriceJSON struct {
	Location     apijson.Field
	PriceHourly  apijson.Field
	PriceMonthly apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LoadBalancerLoadBalancerTypePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
// this Location | Monthly costs for a Floating IP type in this Location | Hourly
// costs for a Load Balancer type in this network zone | Monthly costs for a Load
// Balancer type in this network zone | Hourly costs for a Primary IP type in this
// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
// for a Server type in this Location | Monthly costs for a Server type in this
// Location
type LoadBalancerLoadBalancerTypePricesPriceHourly struct {
	// Price with VAT added
	Gross string `json:"gross,required" format:"decimal"`
	// Price without VAT
	Net  string `json:"net,required" format:"decimal"`
	JSON loadBalancerLoadBalancerTypePricesPriceHourlyJSON
}

// loadBalancerLoadBalancerTypePricesPriceHourlyJSON contains the JSON metadata for
// the struct [LoadBalancerLoadBalancerTypePricesPriceHourly]
type loadBalancerLoadBalancerTypePricesPriceHourlyJSON struct {
	Gross       apijson.Field
	Net         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerLoadBalancerTypePricesPriceHourly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Hourly costs for a Resource in this Location | Monthly costs for a Resource in
// this Location | Monthly costs for a Floating IP type in this Location | Hourly
// costs for a Load Balancer type in this network zone | Monthly costs for a Load
// Balancer type in this network zone | Hourly costs for a Primary IP type in this
// Location | Monthly costs for a Primary IP type in this Location | Hourly costs
// for a Server type in this Location | Monthly costs for a Server type in this
// Location
type LoadBalancerLoadBalancerTypePricesPriceMonthly struct {
	// Price with VAT added
	Gross string `json:"gross,required" format:"decimal"`
	// Price without VAT
	Net  string `json:"net,required" format:"decimal"`
	JSON loadBalancerLoadBalancerTypePricesPriceMonthlyJSON
}

// loadBalancerLoadBalancerTypePricesPriceMonthlyJSON contains the JSON metadata
// for the struct [LoadBalancerLoadBalancerTypePricesPriceMonthly]
type loadBalancerLoadBalancerTypePricesPriceMonthlyJSON struct {
	Gross       apijson.Field
	Net         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerLoadBalancerTypePricesPriceMonthly) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type LoadBalancerLocation struct {
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
	JSON        loadBalancerLocationJSON
}

// loadBalancerLocationJSON contains the JSON metadata for the struct
// [LoadBalancerLocation]
type loadBalancerLocationJSON struct {
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

func (r *LoadBalancerLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPrivateNet struct {
	// IP address (v4) of this Load Balancer in this Network
	Ip string `json:"ip"`
	// ID of the Network
	Network int64 `json:"network"`
	JSON    loadBalancerPrivateNetJSON
}

// loadBalancerPrivateNetJSON contains the JSON metadata for the struct
// [LoadBalancerPrivateNet]
type loadBalancerPrivateNetJSON struct {
	Ip          apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPrivateNet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type LoadBalancerProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   loadBalancerProtectionJSON
}

// loadBalancerProtectionJSON contains the JSON metadata for the struct
// [LoadBalancerProtection]
type loadBalancerProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Public network information
type LoadBalancerPublicNet struct {
	// Public Interface enabled or not
	Enabled bool `json:"enabled,required"`
	// IP address (v4)
	Ipv4 LoadBalancerPublicNetIpv4 `json:"ipv4,required"`
	// IP address (v6)
	Ipv6 LoadBalancerPublicNetIpv6 `json:"ipv6,required"`
	JSON loadBalancerPublicNetJSON
}

// loadBalancerPublicNetJSON contains the JSON metadata for the struct
// [LoadBalancerPublicNet]
type loadBalancerPublicNetJSON struct {
	Enabled     apijson.Field
	Ipv4        apijson.Field
	Ipv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPublicNet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// IP address (v4)
type LoadBalancerPublicNetIpv4 struct {
	// Reverse DNS PTR entry for the IPv4 address of this Load Balancer
	DnsPtr string `json:"dns_ptr,nullable"`
	// IP address (v4) of this Load Balancer
	Ip   string `json:"ip,nullable"`
	JSON loadBalancerPublicNetIpv4JSON
}

// loadBalancerPublicNetIpv4JSON contains the JSON metadata for the struct
// [LoadBalancerPublicNetIpv4]
type loadBalancerPublicNetIpv4JSON struct {
	DnsPtr      apijson.Field
	Ip          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPublicNetIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// IP address (v6)
type LoadBalancerPublicNetIpv6 struct {
	// Reverse DNS PTR entry for the IPv6 address of this Load Balancer
	DnsPtr string `json:"dns_ptr,nullable"`
	// IP address (v6) of this Load Balancer
	Ip   string `json:"ip,nullable"`
	JSON loadBalancerPublicNetIpv6JSON
}

// loadBalancerPublicNetIpv6JSON contains the JSON metadata for the struct
// [LoadBalancerPublicNetIpv6]
type loadBalancerPublicNetIpv6JSON struct {
	DnsPtr      apijson.Field
	Ip          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPublicNetIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A service for a Load Balancer.
type LoadBalancerServiceModel struct {
	// Port the Load Balancer will balance to
	DestinationPort int64 `json:"destination_port,required"`
	// Service health check
	HealthCheck LoadBalancerServiceModelHealthCheck `json:"health_check,required"`
	// Port the Load Balancer listens on
	ListenPort int64 `json:"listen_port,required"`
	// Protocol of the Load Balancer
	Protocol LoadBalancerServiceModelProtocol `json:"protocol,required"`
	// Is Proxyprotocol enabled or not
	Proxyprotocol bool `json:"proxyprotocol,required"`
	// Configuration option for protocols http and https
	HTTP LoadBalancerServiceModelHTTP `json:"http"`
	JSON loadBalancerServiceModelJSON
}

// loadBalancerServiceModelJSON contains the JSON metadata for the struct
// [LoadBalancerServiceModel]
type loadBalancerServiceModelJSON struct {
	DestinationPort apijson.Field
	HealthCheck     apijson.Field
	ListenPort      apijson.Field
	Protocol        apijson.Field
	Proxyprotocol   apijson.Field
	HTTP            apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LoadBalancerServiceModel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Service health check
type LoadBalancerServiceModelHealthCheck struct {
	// Time interval in seconds health checks are performed
	Interval int64 `json:"interval,required"`
	// Port the health check will be performed on
	Port int64 `json:"port,required"`
	// Type of the health check
	Protocol LoadBalancerServiceModelHealthCheckProtocol `json:"protocol,required"`
	// Unsuccessful retries needed until a target is considered unhealthy; an unhealthy
	// target needs the same number of successful retries to become healthy again
	Retries int64 `json:"retries,required"`
	// Time in seconds after an attempt is considered a timeout
	Timeout int64 `json:"timeout,required"`
	// Additional configuration for protocol http
	HTTP LoadBalancerServiceModelHealthCheckHTTP `json:"http"`
	JSON loadBalancerServiceModelHealthCheckJSON
}

// loadBalancerServiceModelHealthCheckJSON contains the JSON metadata for the
// struct [LoadBalancerServiceModelHealthCheck]
type loadBalancerServiceModelHealthCheckJSON struct {
	Interval    apijson.Field
	Port        apijson.Field
	Protocol    apijson.Field
	Retries     apijson.Field
	Timeout     apijson.Field
	HTTP        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerServiceModelHealthCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the health check
type LoadBalancerServiceModelHealthCheckProtocol string

const (
	LoadBalancerServiceModelHealthCheckProtocolHTTP LoadBalancerServiceModelHealthCheckProtocol = "http"
	LoadBalancerServiceModelHealthCheckProtocolTcp  LoadBalancerServiceModelHealthCheckProtocol = "tcp"
)

// Additional configuration for protocol http
type LoadBalancerServiceModelHealthCheckHTTP struct {
	// Host header to send in the HTTP request. May not contain spaces, percent or
	// backslash symbols. Can be null, in that case no host header is sent.
	Domain string `json:"domain,required,nullable"`
	// HTTP path to use for health checks. May not contain literal spaces, use
	// percent-encoding instead.
	Path string `json:"path,required"`
	// String that must be contained in HTTP response in order to pass the health check
	Response string `json:"response"`
	// List of returned HTTP status codes in order to pass the health check. Supports
	// the wildcards `?` for exactly one character and `*` for multiple ones. The
	// default is to pass the health check for any status code between 2?? and 3??.
	StatusCodes []string `json:"status_codes"`
	// Use HTTPS for health check
	Tls  bool `json:"tls"`
	JSON loadBalancerServiceModelHealthCheckHTTPJSON
}

// loadBalancerServiceModelHealthCheckHTTPJSON contains the JSON metadata for the
// struct [LoadBalancerServiceModelHealthCheckHTTP]
type loadBalancerServiceModelHealthCheckHTTPJSON struct {
	Domain      apijson.Field
	Path        apijson.Field
	Response    apijson.Field
	StatusCodes apijson.Field
	Tls         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerServiceModelHealthCheckHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protocol of the Load Balancer
type LoadBalancerServiceModelProtocol string

const (
	LoadBalancerServiceModelProtocolHTTP  LoadBalancerServiceModelProtocol = "http"
	LoadBalancerServiceModelProtocolHTTPs LoadBalancerServiceModelProtocol = "https"
	LoadBalancerServiceModelProtocolTcp   LoadBalancerServiceModelProtocol = "tcp"
)

// Configuration option for protocols http and https
type LoadBalancerServiceModelHTTP struct {
	// IDs of the Certificates to use for TLS/SSL termination by the Load Balancer;
	// empty for TLS/SSL passthrough or if `protocol` is "http"
	Certificates []int64 `json:"certificates"`
	// Lifetime of the cookie used for sticky sessions
	CookieLifetime int64 `json:"cookie_lifetime"`
	// Name of the cookie used for sticky sessions
	CookieName string `json:"cookie_name"`
	// Redirect HTTP requests to HTTPS. Only available if protocol is "https". Default
	// `false`
	RedirectHTTP bool `json:"redirect_http"`
	// Use sticky sessions. Only available if protocol is "http" or "https". Default
	// `false`
	StickySessions bool `json:"sticky_sessions"`
	JSON           loadBalancerServiceModelHTTPJSON
}

// loadBalancerServiceModelHTTPJSON contains the JSON metadata for the struct
// [LoadBalancerServiceModelHTTP]
type loadBalancerServiceModelHTTPJSON struct {
	Certificates   apijson.Field
	CookieLifetime apijson.Field
	CookieName     apijson.Field
	RedirectHTTP   apijson.Field
	StickySessions apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *LoadBalancerServiceModelHTTP) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A target for a load balancer
type LoadBalancerTarget struct {
	// Type of the resource
	Type LoadBalancerTargetsType `json:"type,required"`
	// List of health statuses of the services on this target. Only present for target
	// types "server" and "ip".
	HealthStatus []shared.HealthStatus `json:"health_status"`
	// IP targets where the traffic should be routed to. It is only possible to use the
	// (Public or vSwitch) IPs of Hetzner Online Root Servers belonging to the project
	// owner. IPs belonging to other users are blocked. Additionally IPs belonging to
	// services provided by Hetzner Cloud (Servers, Load Balancers, ...) are blocked as
	// well. Only present for target type "ip".
	Ip LoadBalancerTargetIp `json:"ip"`
	// Configuration for type LabelSelector, required if type is `label_selector`
	LabelSelector LoadBalancerTargetsLabelSelector `json:"label_selector"`
	// ID of the Resource
	Server LoadBalancerTargetsServer `json:"server"`
	// List of resolved label selector target Servers. Only present for type
	// "label_selector".
	Targets []LoadBalancerTargetsTarget `json:"targets"`
	// Use the private network IP instead of the public IP. Default value is false.
	// Only present for target types "server" and "label_selector".
	UsePrivateIp bool `json:"use_private_ip"`
	JSON         loadBalancerTargetJSON
}

// loadBalancerTargetJSON contains the JSON metadata for the struct
// [LoadBalancerTarget]
type loadBalancerTargetJSON struct {
	Type          apijson.Field
	HealthStatus  apijson.Field
	Ip            apijson.Field
	LabelSelector apijson.Field
	Server        apijson.Field
	Targets       apijson.Field
	UsePrivateIp  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the resource
type LoadBalancerTargetsType string

const (
	LoadBalancerTargetsTypeIp            LoadBalancerTargetsType = "ip"
	LoadBalancerTargetsTypeLabelSelector LoadBalancerTargetsType = "label_selector"
	LoadBalancerTargetsTypeServer        LoadBalancerTargetsType = "server"
)

// IP targets where the traffic should be routed to. It is only possible to use the
// (Public or vSwitch) IPs of Hetzner Online Root Servers belonging to the project
// owner. IPs belonging to other users are blocked. Additionally IPs belonging to
// services provided by Hetzner Cloud (Servers, Load Balancers, ...) are blocked as
// well. Only present for target type "ip".
type LoadBalancerTargetIp struct {
	// IP of a server that belongs to the same customer (public IPv4/IPv6) or private
	// IP in a Subnetwork type vswitch.
	Ip   string `json:"ip,required"`
	JSON loadBalancerTargetIpJSON
}

// loadBalancerTargetIpJSON contains the JSON metadata for the struct
// [LoadBalancerTargetIp]
type loadBalancerTargetIpJSON struct {
	Ip          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerTargetIp) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for type LabelSelector, required if type is `label_selector`
type LoadBalancerTargetsLabelSelector struct {
	// Label selector
	Selector string `json:"selector,required"`
	JSON     loadBalancerTargetsLabelSelectorJSON
}

// loadBalancerTargetsLabelSelectorJSON contains the JSON metadata for the struct
// [LoadBalancerTargetsLabelSelector]
type loadBalancerTargetsLabelSelectorJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerTargetsLabelSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the Resource
type LoadBalancerTargetsServer struct {
	// ID of the Resource | ID of the Server
	ID   int64 `json:"id,required"`
	JSON loadBalancerTargetsServerJSON
}

// loadBalancerTargetsServerJSON contains the JSON metadata for the struct
// [LoadBalancerTargetsServer]
type loadBalancerTargetsServerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerTargetsServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerTargetsTarget struct {
	// List of health statuses of the services on this target. Only present for target
	// types "server" and "ip".
	HealthStatus []shared.HealthStatus `json:"health_status"`
	// ID of the Resource
	Server LoadBalancerTargetsTargetsServer `json:"server"`
	// Type of the resource. Here always "server".
	Type string `json:"type"`
	// Use the private network IP instead of the public IP. Default value is false.
	// Only present for target types "server" and "label_selector".
	UsePrivateIp bool `json:"use_private_ip"`
	JSON         loadBalancerTargetsTargetJSON
}

// loadBalancerTargetsTargetJSON contains the JSON metadata for the struct
// [LoadBalancerTargetsTarget]
type loadBalancerTargetsTargetJSON struct {
	HealthStatus apijson.Field
	Server       apijson.Field
	Type         apijson.Field
	UsePrivateIp apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LoadBalancerTargetsTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the Resource
type LoadBalancerTargetsTargetsServer struct {
	// ID of the Resource | ID of the Server
	ID   int64 `json:"id,required"`
	JSON loadBalancerTargetsTargetsServerJSON
}

// loadBalancerTargetsTargetsServerJSON contains the JSON metadata for the struct
// [LoadBalancerTargetsTargetsServer]
type loadBalancerTargetsTargetsServerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerTargetsTargetsServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A service for a Load Balancer.
type LoadBalancerServiceModelParam struct {
	// Port the Load Balancer will balance to
	DestinationPort param.Field[int64] `json:"destination_port,required"`
	// Service health check
	HealthCheck param.Field[LoadBalancerServiceModelHealthCheckParam] `json:"health_check,required"`
	// Port the Load Balancer listens on
	ListenPort param.Field[int64] `json:"listen_port,required"`
	// Protocol of the Load Balancer
	Protocol param.Field[LoadBalancerServiceModelProtocol] `json:"protocol,required"`
	// Is Proxyprotocol enabled or not
	Proxyprotocol param.Field[bool] `json:"proxyprotocol,required"`
	// Configuration option for protocols http and https
	HTTP param.Field[LoadBalancerServiceModelHTTPParam] `json:"http"`
}

func (r LoadBalancerServiceModelParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service health check
type LoadBalancerServiceModelHealthCheckParam struct {
	// Time interval in seconds health checks are performed
	Interval param.Field[int64] `json:"interval,required"`
	// Port the health check will be performed on
	Port param.Field[int64] `json:"port,required"`
	// Type of the health check
	Protocol param.Field[LoadBalancerServiceModelHealthCheckProtocol] `json:"protocol,required"`
	// Unsuccessful retries needed until a target is considered unhealthy; an unhealthy
	// target needs the same number of successful retries to become healthy again
	Retries param.Field[int64] `json:"retries,required"`
	// Time in seconds after an attempt is considered a timeout
	Timeout param.Field[int64] `json:"timeout,required"`
	// Additional configuration for protocol http
	HTTP param.Field[LoadBalancerServiceModelHealthCheckHTTPParam] `json:"http"`
}

func (r LoadBalancerServiceModelHealthCheckParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Additional configuration for protocol http
type LoadBalancerServiceModelHealthCheckHTTPParam struct {
	// Host header to send in the HTTP request. May not contain spaces, percent or
	// backslash symbols. Can be null, in that case no host header is sent.
	Domain param.Field[string] `json:"domain,required"`
	// HTTP path to use for health checks. May not contain literal spaces, use
	// percent-encoding instead.
	Path param.Field[string] `json:"path,required"`
	// String that must be contained in HTTP response in order to pass the health check
	Response param.Field[string] `json:"response"`
	// List of returned HTTP status codes in order to pass the health check. Supports
	// the wildcards `?` for exactly one character and `*` for multiple ones. The
	// default is to pass the health check for any status code between 2?? and 3??.
	StatusCodes param.Field[[]string] `json:"status_codes"`
	// Use HTTPS for health check
	Tls param.Field[bool] `json:"tls"`
}

func (r LoadBalancerServiceModelHealthCheckHTTPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration option for protocols http and https
type LoadBalancerServiceModelHTTPParam struct {
	// IDs of the Certificates to use for TLS/SSL termination by the Load Balancer;
	// empty for TLS/SSL passthrough or if `protocol` is "http"
	Certificates param.Field[[]int64] `json:"certificates"`
	// Lifetime of the cookie used for sticky sessions
	CookieLifetime param.Field[int64] `json:"cookie_lifetime"`
	// Name of the cookie used for sticky sessions
	CookieName param.Field[string] `json:"cookie_name"`
	// Redirect HTTP requests to HTTPS. Only available if protocol is "https". Default
	// `false`
	RedirectHTTP param.Field[bool] `json:"redirect_http"`
	// Use sticky sessions. Only available if protocol is "http" or "https". Default
	// `false`
	StickySessions param.Field[bool] `json:"sticky_sessions"`
}

func (r LoadBalancerServiceModelHTTPParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// IP targets where the traffic should be routed to. It is only possible to use the
// (Public or vSwitch) IPs of Hetzner Online Root Servers belonging to the project
// owner. IPs belonging to other users are blocked. Additionally IPs belonging to
// services provided by Hetzner Cloud (Servers, Load Balancers, ...) are blocked as
// well. Only present for target type "ip".
type LoadBalancerTargetIpParam struct {
	// IP of a server that belongs to the same customer (public IPv4/IPv6) or private
	// IP in a Subnetwork type vswitch.
	Ip param.Field[string] `json:"ip,required"`
}

func (r LoadBalancerTargetIpParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response to POST https://api.hetzner.cloud/v1/load_balancers
type LoadBalancerNewResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action       shared.Action `json:"action,required"`
	LoadBalancer LoadBalancer  `json:"load_balancer,required"`
	JSON         loadBalancerNewResponseJSON
}

// loadBalancerNewResponseJSON contains the JSON metadata for the struct
// [LoadBalancerNewResponse]
type loadBalancerNewResponseJSON struct {
	Action       apijson.Field
	LoadBalancer apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LoadBalancerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/load_balancers/{id}
type LoadBalancerGetResponse struct {
	LoadBalancer LoadBalancer `json:"load_balancer,required"`
	JSON         loadBalancerGetResponseJSON
}

// loadBalancerGetResponseJSON contains the JSON metadata for the struct
// [LoadBalancerGetResponse]
type loadBalancerGetResponseJSON struct {
	LoadBalancer apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LoadBalancerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to PUT https://api.hetzner.cloud/v1/load_balancers/{id}
type LoadBalancerUpdateResponse struct {
	LoadBalancer LoadBalancer `json:"load_balancer,required"`
	JSON         loadBalancerUpdateResponseJSON
}

// loadBalancerUpdateResponseJSON contains the JSON metadata for the struct
// [LoadBalancerUpdateResponse]
type loadBalancerUpdateResponseJSON struct {
	LoadBalancer apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *LoadBalancerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/load_balancers
type LoadBalancerListResponse struct {
	LoadBalancers []LoadBalancer `json:"load_balancers,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON loadBalancerListResponseJSON
}

// loadBalancerListResponseJSON contains the JSON metadata for the struct
// [LoadBalancerListResponse]
type loadBalancerListResponseJSON struct {
	LoadBalancers apijson.Field
	Meta          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancerListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerNewParams struct {
	// Algorithm of the Load Balancer | Request for POST
	// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/change_algorithm
	Algorithm param.Field[LoadBalancerNewParamsAlgorithm] `json:"algorithm,required"`
	// ID or name of the Load Balancer type this Load Balancer should be created with
	LoadBalancerType param.Field[string] `json:"load_balancer_type,required"`
	// Name of the Load Balancer
	Name param.Field[string] `json:"name,required"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// ID or name of Location to create Load Balancer in
	Location param.Field[string] `json:"location"`
	// ID of the network the Load Balancer should be attached to on creation
	Network param.Field[int64] `json:"network"`
	// Name of network zone
	NetworkZone param.Field[string] `json:"network_zone"`
	// Enable or disable the public interface of the Load Balancer
	PublicInterface param.Field[bool] `json:"public_interface"`
	// Array of services
	Services param.Field[[]LoadBalancerServiceModelParam] `json:"services"`
	// Array of targets
	Targets param.Field[[]LoadBalancerNewParamsTarget] `json:"targets"`
}

func (r LoadBalancerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Algorithm of the Load Balancer | Request for POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/change_algorithm
type LoadBalancerNewParamsAlgorithm struct {
	// Type of the algorithm | Algorithm of the Load Balancer
	Type param.Field[LoadBalancerNewParamsAlgorithmType] `json:"type,required"`
}

func (r LoadBalancerNewParamsAlgorithm) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the algorithm | Algorithm of the Load Balancer
type LoadBalancerNewParamsAlgorithmType string

const (
	LoadBalancerNewParamsAlgorithmTypeLeastConnections LoadBalancerNewParamsAlgorithmType = "least_connections"
	LoadBalancerNewParamsAlgorithmTypeRoundRobin       LoadBalancerNewParamsAlgorithmType = "round_robin"
)

// A target for a load balancer
type LoadBalancerNewParamsTarget struct {
	// Type of the resource
	Type param.Field[LoadBalancerNewParamsTargetsType] `json:"type,required"`
	// List of health statuses of the services on this target. Only present for target
	// types "server" and "ip".
	HealthStatus param.Field[[]shared.HealthStatusParam] `json:"health_status"`
	// IP targets where the traffic should be routed to. It is only possible to use the
	// (Public or vSwitch) IPs of Hetzner Online Root Servers belonging to the project
	// owner. IPs belonging to other users are blocked. Additionally IPs belonging to
	// services provided by Hetzner Cloud (Servers, Load Balancers, ...) are blocked as
	// well. Only present for target type "ip".
	Ip param.Field[LoadBalancerTargetIpParam] `json:"ip"`
	// Configuration for type LabelSelector, required if type is `label_selector`
	LabelSelector param.Field[LoadBalancerNewParamsTargetsLabelSelector] `json:"label_selector"`
	// ID of the Resource
	Server param.Field[LoadBalancerNewParamsTargetsServer] `json:"server"`
	// List of resolved label selector target Servers. Only present for type
	// "label_selector".
	Targets param.Field[[]LoadBalancerNewParamsTargetsTarget] `json:"targets"`
	// Use the private network IP instead of the public IP. Default value is false.
	// Only present for target types "server" and "label_selector".
	UsePrivateIp param.Field[bool] `json:"use_private_ip"`
}

func (r LoadBalancerNewParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the resource
type LoadBalancerNewParamsTargetsType string

const (
	LoadBalancerNewParamsTargetsTypeIp            LoadBalancerNewParamsTargetsType = "ip"
	LoadBalancerNewParamsTargetsTypeLabelSelector LoadBalancerNewParamsTargetsType = "label_selector"
	LoadBalancerNewParamsTargetsTypeServer        LoadBalancerNewParamsTargetsType = "server"
)

// Configuration for type LabelSelector, required if type is `label_selector`
type LoadBalancerNewParamsTargetsLabelSelector struct {
	// Label selector
	Selector param.Field[string] `json:"selector,required"`
}

func (r LoadBalancerNewParamsTargetsLabelSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the Resource
type LoadBalancerNewParamsTargetsServer struct {
	// ID of the Resource | ID of the Server
	ID param.Field[int64] `json:"id,required"`
}

func (r LoadBalancerNewParamsTargetsServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerNewParamsTargetsTarget struct {
	// List of health statuses of the services on this target. Only present for target
	// types "server" and "ip".
	HealthStatus param.Field[[]shared.HealthStatusParam] `json:"health_status"`
	// ID of the Resource
	Server param.Field[LoadBalancerNewParamsTargetsTargetsServer] `json:"server"`
	// Type of the resource. Here always "server".
	Type param.Field[string] `json:"type"`
	// Use the private network IP instead of the public IP. Default value is false.
	// Only present for target types "server" and "label_selector".
	UsePrivateIp param.Field[bool] `json:"use_private_ip"`
}

func (r LoadBalancerNewParamsTargetsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the Resource
type LoadBalancerNewParamsTargetsTargetsServer struct {
	// ID of the Resource | ID of the Server
	ID param.Field[int64] `json:"id,required"`
}

func (r LoadBalancerNewParamsTargetsTargetsServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerUpdateParams struct {
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// New Load Balancer name
	Name param.Field[string] `json:"name"`
}

func (r LoadBalancerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerListParams struct {
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
	// Can be used multiple times.
	Sort param.Field[LoadBalancerListParamsSort] `query:"sort"`
}

// URLQuery serializes [LoadBalancerListParams]'s query parameters as `url.Values`.
func (r LoadBalancerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times.
type LoadBalancerListParamsSort string

const (
	LoadBalancerListParamsSortID          LoadBalancerListParamsSort = "id"
	LoadBalancerListParamsSortIDAsc       LoadBalancerListParamsSort = "id:asc"
	LoadBalancerListParamsSortIDDesc      LoadBalancerListParamsSort = "id:desc"
	LoadBalancerListParamsSortName        LoadBalancerListParamsSort = "name"
	LoadBalancerListParamsSortNameAsc     LoadBalancerListParamsSort = "name:asc"
	LoadBalancerListParamsSortNameDesc    LoadBalancerListParamsSort = "name:desc"
	LoadBalancerListParamsSortCreated     LoadBalancerListParamsSort = "created"
	LoadBalancerListParamsSortCreatedAsc  LoadBalancerListParamsSort = "created:asc"
	LoadBalancerListParamsSortCreatedDesc LoadBalancerListParamsSort = "created:desc"
)
