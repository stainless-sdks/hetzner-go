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

// NetworkService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewNetworkService] method instead.
type NetworkService struct {
	Options []option.RequestOption
	Actions *NetworkActionService
}

// NewNetworkService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewNetworkService(opts ...option.RequestOption) (r *NetworkService) {
	r = &NetworkService{}
	r.Options = opts
	r.Actions = NewNetworkActionService(opts...)
	return
}

// Creates a network with the specified `ip_range`.
//
// You may specify one or more `subnets`. You can also add more Subnets later by
// using the
// [add subnet action](https://docs.hetzner.cloud/#network-actions-add-a-subnet-to-a-network).
// If you do not specify an `ip_range` in the subnet we will automatically pick the
// first available /24 range for you.
//
// You may specify one or more routes in `routes`. You can also add more routes
// later by using the
// [add route action](https://docs.hetzner.cloud/#network-actions-add-a-route-to-a-network).
func (r *NetworkService) New(ctx context.Context, body NetworkNewParams, opts ...option.RequestOption) (res *NetworkNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "networks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a specific network object.
func (r *NetworkService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *NetworkGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("networks/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the network properties.
//
// Note that when updating labels, the network’s current set of labels will be
// replaced with the labels provided in the request body. So, for example, if you
// want to add a new label, you have to provide all existing labels plus the new
// label in the request body.
//
// Note: if the network object changes during the request, the response will be a
// “conflict” error.
func (r *NetworkService) Update(ctx context.Context, id int64, body NetworkUpdateParams, opts ...option.RequestOption) (res *NetworkUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("networks/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Gets all existing networks that you have available.
func (r *NetworkService) List(ctx context.Context, query NetworkListParams, opts ...option.RequestOption) (res *NetworkListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "networks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a network. If there are Servers attached they will be detached in the
// background.
//
// Note: if the network object changes during the request, the response will be a
// “conflict” error.
func (r *NetworkService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("networks/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Response to POST https://api.hetzner.cloud/v1/networks
type NetworkNewResponse struct {
	Network NetworkNewResponseNetwork `json:"network"`
	JSON    networkNewResponseJSON
}

// networkNewResponseJSON contains the JSON metadata for the struct
// [NetworkNewResponse]
type networkNewResponseJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkNewResponseNetwork struct {
	// ID of the Network
	ID int64 `json:"id,required"`
	// Point in time when the Network was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Indicates if the routes from this network should be exposed to the vSwitch
	// connection. The exposing only takes effect if a vSwitch connection is active.
	//
	// Currently the default route 0.0.0.0/0 is not exposed to the vSwitch connection.
	// We are aware of the issue and are working on a solution.
	ExposeRoutesToVswitch bool `json:"expose_routes_to_vswitch,required"`
	// IPv4 prefix of the whole Network
	IPRange string `json:"ip_range,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Network
	Name string `json:"name,required"`
	// Protection configuration for the Resource
	Protection NetworkNewResponseNetworkProtection `json:"protection,required"`
	// Array of routes set in this Network
	Routes []NetworkNewResponseNetworkRoute `json:"routes,required"`
	// Array of IDs of Servers attached to this Network
	Servers []int64 `json:"servers,required"`
	// Array subnets allocated in this Network
	Subnets []NetworkNewResponseNetworkSubnet `json:"subnets,required"`
	// Array of IDs of Load Balancers attached to this Network
	LoadBalancers []int64 `json:"load_balancers"`
	JSON          networkNewResponseNetworkJSON
}

// networkNewResponseNetworkJSON contains the JSON metadata for the struct
// [NetworkNewResponseNetwork]
type networkNewResponseNetworkJSON struct {
	ID                    apijson.Field
	Created               apijson.Field
	ExposeRoutesToVswitch apijson.Field
	IPRange               apijson.Field
	Labels                apijson.Field
	Name                  apijson.Field
	Protection            apijson.Field
	Routes                apijson.Field
	Servers               apijson.Field
	Subnets               apijson.Field
	LoadBalancers         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *NetworkNewResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type NetworkNewResponseNetworkProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   networkNewResponseNetworkProtectionJSON
}

// networkNewResponseNetworkProtectionJSON contains the JSON metadata for the
// struct [NetworkNewResponseNetworkProtection]
type networkNewResponseNetworkProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkNewResponseNetworkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Request for POST https://api.hetzner.cloud/v1/networks/{id}/actions/add_route |
// Request for POST https://api.hetzner.cloud/v1/networks/{id}/actions/delete_route
type NetworkNewResponseNetworkRoute struct {
	// Destination network or host of this route. Must not overlap with an existing
	// ip_range in any subnets or with any destinations in other routes or with the
	// first IP of the networks ip_range or with 172.31.1.1. Must be one of the private
	// IPv4 ranges of RFC1918.
	Destination string `json:"destination,required"`
	// Gateway for the route. Cannot be the first IP of the networks ip_range and also
	// cannot be 172.31.1.1 as this IP is being used as a gateway for the public
	// network interface of Servers. | Gateway for the route. Cannot be the first IP of
	// the networks ip_range, an IP behind a vSwitch or 172.31.1.1, as this IP is being
	// used as a gateway for the public network interface of Servers.
	Gateway string `json:"gateway,required"`
	JSON    networkNewResponseNetworkRouteJSON
}

// networkNewResponseNetworkRouteJSON contains the JSON metadata for the struct
// [NetworkNewResponseNetworkRoute]
type networkNewResponseNetworkRouteJSON struct {
	Destination apijson.Field
	Gateway     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkNewResponseNetworkRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkNewResponseNetworkSubnet struct {
	// Gateway for Servers attached to this subnet. For subnets of type Server this is
	// always the first IP of the network IP range.
	Gateway string `json:"gateway,required"`
	// Name of Network zone. The Location object contains the `network_zone` property
	// each Location belongs to.
	NetworkZone string `json:"network_zone,required"`
	// Type of Subnetwork
	Type NetworkNewResponseNetworkSubnetsType `json:"type,required"`
	// Range to allocate IPs from. Must be a Subnet of the ip_range of the parent
	// network object and must not overlap with any other subnets or with any
	// destinations in routes. Minimum Network size is /30. We suggest that you pick a
	// bigger Network with a /24 netmask.
	IPRange string `json:"ip_range"`
	// ID of the robot vSwitch if the subnet is of type vswitch.
	VswitchID int64 `json:"vswitch_id,nullable"`
	JSON      networkNewResponseNetworkSubnetJSON
}

// networkNewResponseNetworkSubnetJSON contains the JSON metadata for the struct
// [NetworkNewResponseNetworkSubnet]
type networkNewResponseNetworkSubnetJSON struct {
	Gateway     apijson.Field
	NetworkZone apijson.Field
	Type        apijson.Field
	IPRange     apijson.Field
	VswitchID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkNewResponseNetworkSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of Subnetwork
type NetworkNewResponseNetworkSubnetsType string

const (
	NetworkNewResponseNetworkSubnetsTypeCloud   NetworkNewResponseNetworkSubnetsType = "cloud"
	NetworkNewResponseNetworkSubnetsTypeServer  NetworkNewResponseNetworkSubnetsType = "server"
	NetworkNewResponseNetworkSubnetsTypeVswitch NetworkNewResponseNetworkSubnetsType = "vswitch"
)

// Response to GET https://api.hetzner.cloud/v1/networks/{id}
type NetworkGetResponse struct {
	Network NetworkGetResponseNetwork `json:"network"`
	JSON    networkGetResponseJSON
}

// networkGetResponseJSON contains the JSON metadata for the struct
// [NetworkGetResponse]
type networkGetResponseJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkGetResponseNetwork struct {
	// ID of the Network
	ID int64 `json:"id,required"`
	// Point in time when the Network was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Indicates if the routes from this network should be exposed to the vSwitch
	// connection. The exposing only takes effect if a vSwitch connection is active.
	//
	// Currently the default route 0.0.0.0/0 is not exposed to the vSwitch connection.
	// We are aware of the issue and are working on a solution.
	ExposeRoutesToVswitch bool `json:"expose_routes_to_vswitch,required"`
	// IPv4 prefix of the whole Network
	IPRange string `json:"ip_range,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Network
	Name string `json:"name,required"`
	// Protection configuration for the Resource
	Protection NetworkGetResponseNetworkProtection `json:"protection,required"`
	// Array of routes set in this Network
	Routes []NetworkGetResponseNetworkRoute `json:"routes,required"`
	// Array of IDs of Servers attached to this Network
	Servers []int64 `json:"servers,required"`
	// Array subnets allocated in this Network
	Subnets []NetworkGetResponseNetworkSubnet `json:"subnets,required"`
	// Array of IDs of Load Balancers attached to this Network
	LoadBalancers []int64 `json:"load_balancers"`
	JSON          networkGetResponseNetworkJSON
}

// networkGetResponseNetworkJSON contains the JSON metadata for the struct
// [NetworkGetResponseNetwork]
type networkGetResponseNetworkJSON struct {
	ID                    apijson.Field
	Created               apijson.Field
	ExposeRoutesToVswitch apijson.Field
	IPRange               apijson.Field
	Labels                apijson.Field
	Name                  apijson.Field
	Protection            apijson.Field
	Routes                apijson.Field
	Servers               apijson.Field
	Subnets               apijson.Field
	LoadBalancers         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *NetworkGetResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type NetworkGetResponseNetworkProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   networkGetResponseNetworkProtectionJSON
}

// networkGetResponseNetworkProtectionJSON contains the JSON metadata for the
// struct [NetworkGetResponseNetworkProtection]
type networkGetResponseNetworkProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkGetResponseNetworkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Request for POST https://api.hetzner.cloud/v1/networks/{id}/actions/add_route |
// Request for POST https://api.hetzner.cloud/v1/networks/{id}/actions/delete_route
type NetworkGetResponseNetworkRoute struct {
	// Destination network or host of this route. Must not overlap with an existing
	// ip_range in any subnets or with any destinations in other routes or with the
	// first IP of the networks ip_range or with 172.31.1.1. Must be one of the private
	// IPv4 ranges of RFC1918.
	Destination string `json:"destination,required"`
	// Gateway for the route. Cannot be the first IP of the networks ip_range and also
	// cannot be 172.31.1.1 as this IP is being used as a gateway for the public
	// network interface of Servers. | Gateway for the route. Cannot be the first IP of
	// the networks ip_range, an IP behind a vSwitch or 172.31.1.1, as this IP is being
	// used as a gateway for the public network interface of Servers.
	Gateway string `json:"gateway,required"`
	JSON    networkGetResponseNetworkRouteJSON
}

// networkGetResponseNetworkRouteJSON contains the JSON metadata for the struct
// [NetworkGetResponseNetworkRoute]
type networkGetResponseNetworkRouteJSON struct {
	Destination apijson.Field
	Gateway     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkGetResponseNetworkRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkGetResponseNetworkSubnet struct {
	// Gateway for Servers attached to this subnet. For subnets of type Server this is
	// always the first IP of the network IP range.
	Gateway string `json:"gateway,required"`
	// Name of Network zone. The Location object contains the `network_zone` property
	// each Location belongs to.
	NetworkZone string `json:"network_zone,required"`
	// Type of Subnetwork
	Type NetworkGetResponseNetworkSubnetsType `json:"type,required"`
	// Range to allocate IPs from. Must be a Subnet of the ip_range of the parent
	// network object and must not overlap with any other subnets or with any
	// destinations in routes. Minimum Network size is /30. We suggest that you pick a
	// bigger Network with a /24 netmask.
	IPRange string `json:"ip_range"`
	// ID of the robot vSwitch if the subnet is of type vswitch.
	VswitchID int64 `json:"vswitch_id,nullable"`
	JSON      networkGetResponseNetworkSubnetJSON
}

// networkGetResponseNetworkSubnetJSON contains the JSON metadata for the struct
// [NetworkGetResponseNetworkSubnet]
type networkGetResponseNetworkSubnetJSON struct {
	Gateway     apijson.Field
	NetworkZone apijson.Field
	Type        apijson.Field
	IPRange     apijson.Field
	VswitchID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkGetResponseNetworkSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of Subnetwork
type NetworkGetResponseNetworkSubnetsType string

const (
	NetworkGetResponseNetworkSubnetsTypeCloud   NetworkGetResponseNetworkSubnetsType = "cloud"
	NetworkGetResponseNetworkSubnetsTypeServer  NetworkGetResponseNetworkSubnetsType = "server"
	NetworkGetResponseNetworkSubnetsTypeVswitch NetworkGetResponseNetworkSubnetsType = "vswitch"
)

// Response to PUT https://api.hetzner.cloud/v1/networks/{id}
type NetworkUpdateResponse struct {
	Network NetworkUpdateResponseNetwork `json:"network"`
	JSON    networkUpdateResponseJSON
}

// networkUpdateResponseJSON contains the JSON metadata for the struct
// [NetworkUpdateResponse]
type networkUpdateResponseJSON struct {
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkUpdateResponseNetwork struct {
	// ID of the Network
	ID int64 `json:"id,required"`
	// Point in time when the Network was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Indicates if the routes from this network should be exposed to the vSwitch
	// connection. The exposing only takes effect if a vSwitch connection is active.
	//
	// Currently the default route 0.0.0.0/0 is not exposed to the vSwitch connection.
	// We are aware of the issue and are working on a solution.
	ExposeRoutesToVswitch bool `json:"expose_routes_to_vswitch,required"`
	// IPv4 prefix of the whole Network
	IPRange string `json:"ip_range,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Network
	Name string `json:"name,required"`
	// Protection configuration for the Resource
	Protection NetworkUpdateResponseNetworkProtection `json:"protection,required"`
	// Array of routes set in this Network
	Routes []NetworkUpdateResponseNetworkRoute `json:"routes,required"`
	// Array of IDs of Servers attached to this Network
	Servers []int64 `json:"servers,required"`
	// Array subnets allocated in this Network
	Subnets []NetworkUpdateResponseNetworkSubnet `json:"subnets,required"`
	// Array of IDs of Load Balancers attached to this Network
	LoadBalancers []int64 `json:"load_balancers"`
	JSON          networkUpdateResponseNetworkJSON
}

// networkUpdateResponseNetworkJSON contains the JSON metadata for the struct
// [NetworkUpdateResponseNetwork]
type networkUpdateResponseNetworkJSON struct {
	ID                    apijson.Field
	Created               apijson.Field
	ExposeRoutesToVswitch apijson.Field
	IPRange               apijson.Field
	Labels                apijson.Field
	Name                  apijson.Field
	Protection            apijson.Field
	Routes                apijson.Field
	Servers               apijson.Field
	Subnets               apijson.Field
	LoadBalancers         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *NetworkUpdateResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type NetworkUpdateResponseNetworkProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   networkUpdateResponseNetworkProtectionJSON
}

// networkUpdateResponseNetworkProtectionJSON contains the JSON metadata for the
// struct [NetworkUpdateResponseNetworkProtection]
type networkUpdateResponseNetworkProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkUpdateResponseNetworkProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Request for POST https://api.hetzner.cloud/v1/networks/{id}/actions/add_route |
// Request for POST https://api.hetzner.cloud/v1/networks/{id}/actions/delete_route
type NetworkUpdateResponseNetworkRoute struct {
	// Destination network or host of this route. Must not overlap with an existing
	// ip_range in any subnets or with any destinations in other routes or with the
	// first IP of the networks ip_range or with 172.31.1.1. Must be one of the private
	// IPv4 ranges of RFC1918.
	Destination string `json:"destination,required"`
	// Gateway for the route. Cannot be the first IP of the networks ip_range and also
	// cannot be 172.31.1.1 as this IP is being used as a gateway for the public
	// network interface of Servers. | Gateway for the route. Cannot be the first IP of
	// the networks ip_range, an IP behind a vSwitch or 172.31.1.1, as this IP is being
	// used as a gateway for the public network interface of Servers.
	Gateway string `json:"gateway,required"`
	JSON    networkUpdateResponseNetworkRouteJSON
}

// networkUpdateResponseNetworkRouteJSON contains the JSON metadata for the struct
// [NetworkUpdateResponseNetworkRoute]
type networkUpdateResponseNetworkRouteJSON struct {
	Destination apijson.Field
	Gateway     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkUpdateResponseNetworkRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkUpdateResponseNetworkSubnet struct {
	// Gateway for Servers attached to this subnet. For subnets of type Server this is
	// always the first IP of the network IP range.
	Gateway string `json:"gateway,required"`
	// Name of Network zone. The Location object contains the `network_zone` property
	// each Location belongs to.
	NetworkZone string `json:"network_zone,required"`
	// Type of Subnetwork
	Type NetworkUpdateResponseNetworkSubnetsType `json:"type,required"`
	// Range to allocate IPs from. Must be a Subnet of the ip_range of the parent
	// network object and must not overlap with any other subnets or with any
	// destinations in routes. Minimum Network size is /30. We suggest that you pick a
	// bigger Network with a /24 netmask.
	IPRange string `json:"ip_range"`
	// ID of the robot vSwitch if the subnet is of type vswitch.
	VswitchID int64 `json:"vswitch_id,nullable"`
	JSON      networkUpdateResponseNetworkSubnetJSON
}

// networkUpdateResponseNetworkSubnetJSON contains the JSON metadata for the struct
// [NetworkUpdateResponseNetworkSubnet]
type networkUpdateResponseNetworkSubnetJSON struct {
	Gateway     apijson.Field
	NetworkZone apijson.Field
	Type        apijson.Field
	IPRange     apijson.Field
	VswitchID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkUpdateResponseNetworkSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of Subnetwork
type NetworkUpdateResponseNetworkSubnetsType string

const (
	NetworkUpdateResponseNetworkSubnetsTypeCloud   NetworkUpdateResponseNetworkSubnetsType = "cloud"
	NetworkUpdateResponseNetworkSubnetsTypeServer  NetworkUpdateResponseNetworkSubnetsType = "server"
	NetworkUpdateResponseNetworkSubnetsTypeVswitch NetworkUpdateResponseNetworkSubnetsType = "vswitch"
)

// Response to GET https://api.hetzner.cloud/v1/networks
type NetworkListResponse struct {
	Networks []NetworkListResponseNetwork `json:"networks,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON networkListResponseJSON
}

// networkListResponseJSON contains the JSON metadata for the struct
// [NetworkListResponse]
type networkListResponseJSON struct {
	Networks    apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkListResponseNetwork struct {
	// ID of the Network
	ID int64 `json:"id,required"`
	// Point in time when the Network was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Indicates if the routes from this network should be exposed to the vSwitch
	// connection. The exposing only takes effect if a vSwitch connection is active.
	//
	// Currently the default route 0.0.0.0/0 is not exposed to the vSwitch connection.
	// We are aware of the issue and are working on a solution.
	ExposeRoutesToVswitch bool `json:"expose_routes_to_vswitch,required"`
	// IPv4 prefix of the whole Network
	IPRange string `json:"ip_range,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Network
	Name string `json:"name,required"`
	// Protection configuration for the Resource
	Protection NetworkListResponseNetworksProtection `json:"protection,required"`
	// Array of routes set in this Network
	Routes []NetworkListResponseNetworksRoute `json:"routes,required"`
	// Array of IDs of Servers attached to this Network
	Servers []int64 `json:"servers,required"`
	// Array subnets allocated in this Network
	Subnets []NetworkListResponseNetworksSubnet `json:"subnets,required"`
	// Array of IDs of Load Balancers attached to this Network
	LoadBalancers []int64 `json:"load_balancers"`
	JSON          networkListResponseNetworkJSON
}

// networkListResponseNetworkJSON contains the JSON metadata for the struct
// [NetworkListResponseNetwork]
type networkListResponseNetworkJSON struct {
	ID                    apijson.Field
	Created               apijson.Field
	ExposeRoutesToVswitch apijson.Field
	IPRange               apijson.Field
	Labels                apijson.Field
	Name                  apijson.Field
	Protection            apijson.Field
	Routes                apijson.Field
	Servers               apijson.Field
	Subnets               apijson.Field
	LoadBalancers         apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *NetworkListResponseNetwork) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type NetworkListResponseNetworksProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   networkListResponseNetworksProtectionJSON
}

// networkListResponseNetworksProtectionJSON contains the JSON metadata for the
// struct [NetworkListResponseNetworksProtection]
type networkListResponseNetworksProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkListResponseNetworksProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Request for POST https://api.hetzner.cloud/v1/networks/{id}/actions/add_route |
// Request for POST https://api.hetzner.cloud/v1/networks/{id}/actions/delete_route
type NetworkListResponseNetworksRoute struct {
	// Destination network or host of this route. Must not overlap with an existing
	// ip_range in any subnets or with any destinations in other routes or with the
	// first IP of the networks ip_range or with 172.31.1.1. Must be one of the private
	// IPv4 ranges of RFC1918.
	Destination string `json:"destination,required"`
	// Gateway for the route. Cannot be the first IP of the networks ip_range and also
	// cannot be 172.31.1.1 as this IP is being used as a gateway for the public
	// network interface of Servers. | Gateway for the route. Cannot be the first IP of
	// the networks ip_range, an IP behind a vSwitch or 172.31.1.1, as this IP is being
	// used as a gateway for the public network interface of Servers.
	Gateway string `json:"gateway,required"`
	JSON    networkListResponseNetworksRouteJSON
}

// networkListResponseNetworksRouteJSON contains the JSON metadata for the struct
// [NetworkListResponseNetworksRoute]
type networkListResponseNetworksRouteJSON struct {
	Destination apijson.Field
	Gateway     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkListResponseNetworksRoute) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkListResponseNetworksSubnet struct {
	// Gateway for Servers attached to this subnet. For subnets of type Server this is
	// always the first IP of the network IP range.
	Gateway string `json:"gateway,required"`
	// Name of Network zone. The Location object contains the `network_zone` property
	// each Location belongs to.
	NetworkZone string `json:"network_zone,required"`
	// Type of Subnetwork
	Type NetworkListResponseNetworksSubnetsType `json:"type,required"`
	// Range to allocate IPs from. Must be a Subnet of the ip_range of the parent
	// network object and must not overlap with any other subnets or with any
	// destinations in routes. Minimum Network size is /30. We suggest that you pick a
	// bigger Network with a /24 netmask.
	IPRange string `json:"ip_range"`
	// ID of the robot vSwitch if the subnet is of type vswitch.
	VswitchID int64 `json:"vswitch_id,nullable"`
	JSON      networkListResponseNetworksSubnetJSON
}

// networkListResponseNetworksSubnetJSON contains the JSON metadata for the struct
// [NetworkListResponseNetworksSubnet]
type networkListResponseNetworksSubnetJSON struct {
	Gateway     apijson.Field
	NetworkZone apijson.Field
	Type        apijson.Field
	IPRange     apijson.Field
	VswitchID   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NetworkListResponseNetworksSubnet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of Subnetwork
type NetworkListResponseNetworksSubnetsType string

const (
	NetworkListResponseNetworksSubnetsTypeCloud   NetworkListResponseNetworksSubnetsType = "cloud"
	NetworkListResponseNetworksSubnetsTypeServer  NetworkListResponseNetworksSubnetsType = "server"
	NetworkListResponseNetworksSubnetsTypeVswitch NetworkListResponseNetworksSubnetsType = "vswitch"
)

type NetworkNewParams struct {
	// IP range of the whole network which must span all included subnets. Must be one
	// of the private IPv4 ranges of RFC1918. Minimum network size is /24. We highly
	// recommend that you pick a larger network with a /16 netmask.
	IPRange param.Field[string] `json:"ip_range,required"`
	// Name of the network
	Name param.Field[string] `json:"name,required"`
	// Indicates if the routes from this network should be exposed to the vSwitch
	// connection. The exposing only takes effect if a vSwitch connection is active.
	//
	// Currently the default route 0.0.0.0/0 is not exposed to the vSwitch connection.
	// We are aware of the issue and are working on a solution.
	ExposeRoutesToVswitch param.Field[bool] `json:"expose_routes_to_vswitch"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// Array of routes set in this network. The destination of the route must be one of
	// the private IPv4 ranges of RFC1918. The gateway must be a subnet/IP of the
	// ip_range of the network object. The destination must not overlap with an
	// existing ip_range in any subnets or with any destinations in other routes or
	// with the first IP of the networks ip_range or with 172.31.1.1. The gateway
	// cannot be the first IP of the networks ip_range and also cannot be 172.31.1.1.
	Routes param.Field[[]NetworkNewParamsRoute] `json:"routes"`
	// Array of subnets allocated.
	Subnets param.Field[[]NetworkNewParamsSubnet] `json:"subnets"`
}

func (r NetworkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Request for POST https://api.hetzner.cloud/v1/networks/{id}/actions/add_route |
// Request for POST https://api.hetzner.cloud/v1/networks/{id}/actions/delete_route
type NetworkNewParamsRoute struct {
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

func (r NetworkNewParamsRoute) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Subnets divide the ip_range from the parent Network object into multiple
// Subnetworks that you can use for different specific purposes.
type NetworkNewParamsSubnet struct {
	// Name of Network zone. The Location object contains the `network_zone` property
	// each Location belongs to.
	NetworkZone param.Field[string] `json:"network_zone,required"`
	// Type of Subnetwork
	Type param.Field[NetworkNewParamsSubnetsType] `json:"type,required"`
	// Range to allocate IPs from. Must be a Subnet of the ip_range of the parent
	// network object and must not overlap with any other subnets or with any
	// destinations in routes. Minimum Network size is /30. We suggest that you pick a
	// bigger Network with a /24 netmask. | Range to allocate IPs from. Must be a
	// Subnet of the ip_range of the parent network object and must not overlap with
	// any other subnets or with any destinations in routes. If the Subnet is of type
	// vSwitch, it also can not overlap with any gateway in routes. Minimum Network
	// size is /30. We suggest that you pick a bigger Network with a /24 netmask.
	IPRange param.Field[string] `json:"ip_range"`
	// ID of the robot vSwitch. Must be supplied if the subnet is of type vswitch.
	VswitchID param.Field[int64] `json:"vswitch_id"`
}

func (r NetworkNewParamsSubnet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of Subnetwork
type NetworkNewParamsSubnetsType string

const (
	NetworkNewParamsSubnetsTypeCloud   NetworkNewParamsSubnetsType = "cloud"
	NetworkNewParamsSubnetsTypeServer  NetworkNewParamsSubnetsType = "server"
	NetworkNewParamsSubnetsTypeVswitch NetworkNewParamsSubnetsType = "vswitch"
)

type NetworkUpdateParams struct {
	// Indicates if the routes from this network should be exposed to the vSwitch
	// connection. The exposing only takes effect if a vSwitch connection is active.
	//
	// Currently the default route 0.0.0.0/0 is not exposed to the vSwitch connection.
	// We are aware of the issue and are working on a solution.
	ExposeRoutesToVswitch param.Field[bool] `json:"expose_routes_to_vswitch"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// New network name
	Name param.Field[string] `json:"name"`
}

func (r NetworkUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type NetworkListParams struct {
	// Can be used to filter networks by labels. The response will only contain
	// networks with a matching label selector pattern.
	LabelSelector param.Field[string] `query:"label_selector"`
	// Can be used to filter networks by their name. The response will only contain the
	// networks matching the specified name.
	Name param.Field[string] `query:"name"`
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [NetworkListParams]'s query parameters as `url.Values`.
func (r NetworkListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
