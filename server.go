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

// ServerService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewServerService] method instead.
type ServerService struct {
	Options []option.RequestOption
	Actions *ServerActionService
	Metrics *ServerMetricService
}

// NewServerService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewServerService(opts ...option.RequestOption) (r *ServerService) {
	r = &ServerService{}
	r.Options = opts
	r.Actions = NewServerActionService(opts...)
	r.Metrics = NewServerMetricService(opts...)
	return
}

// Creates a new Server. Returns preliminary information about the Server as well
// as an Action that covers progress of creation.
func (r *ServerService) New(ctx context.Context, body ServerNewParams, opts ...option.RequestOption) (res *ServerNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "servers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a specific Server object. The Server must exist inside the Project
func (r *ServerService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a Server. You can update a Server’s name and a Server’s labels. Please
// note that Server names must be unique per Project and valid hostnames as per RFC
// 1123 (i.e. may only contain letters, digits, periods, and dashes). Also note
// that when updating labels, the Server’s current set of labels will be replaced
// with the labels provided in the request body. So, for example, if you want to
// add a new label, you have to provide all existing labels plus the new label in
// the request body.
func (r *ServerService) Update(ctx context.Context, id int64, body ServerUpdateParams, opts ...option.RequestOption) (res *ServerUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all existing Server objects
func (r *ServerService) List(ctx context.Context, query ServerListParams, opts ...option.RequestOption) (res *shared.ServersPage[Server], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "servers"
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

// Returns all existing Server objects
func (r *ServerService) ListAutoPaging(ctx context.Context, query ServerListParams, opts ...option.RequestOption) *shared.ServersPageAutoPager[Server] {
	return shared.NewServersPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a Server. This immediately removes the Server from your account, and it
// is no longer accessible.
func (r *ServerService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Servers are virtual machines that can be provisioned.
type Server struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Time window (UTC) in which the backup will run, or null if the backups are not
	// enabled
	BackupWindow string `json:"backup_window,required,nullable"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Datacenter this Primary IP is located at | Datacenter this Resource is located
	// at
	Datacenter ServerDatacenter `json:"datacenter,required"`
	Image      ServerImage      `json:"image,required,nullable"`
	// Free Traffic for the current billing period in bytes
	IncludedTraffic int64 `json:"included_traffic,required,nullable"`
	// Inbound Traffic for the current billing period in bytes
	IngoingTraffic int64 `json:"ingoing_traffic,required,nullable"`
	// ISO Image that is attached to this Server. Null if no ISO is attached.
	Iso ServerIso `json:"iso,required,nullable"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// True if Server has been locked and is not available to user
	Locked bool `json:"locked,required"`
	// Name of the Server (must be unique per Project and a valid hostname as per
	// RFC 1123)
	Name string `json:"name,required"`
	// Outbound Traffic for the current billing period in bytes
	OutgoingTraffic int64 `json:"outgoing_traffic,required,nullable"`
	// Size of the primary Disk
	PrimaryDiskSize int64 `json:"primary_disk_size,required"`
	// Private networks information
	PrivateNet []ServerPrivateNet `json:"private_net,required"`
	// Protection configuration for the Server
	Protection ServerProtection `json:"protection,required"`
	// Public network information. The Server's IPv4 address can be found in
	// `public_net->ipv4->ip`
	PublicNet ServerPublicNet `json:"public_net,required"`
	// True if rescue mode is enabled. Server will then boot into rescue system on next
	// reboot
	RescueEnabled bool `json:"rescue_enabled,required"`
	// Type of Server - determines how much ram, disk and cpu a Server has
	ServerType ServerServerType `json:"server_type,required"`
	// Status of the Server
	Status         ServerStatus         `json:"status,required"`
	LoadBalancers  []int64              `json:"load_balancers"`
	PlacementGroup ServerPlacementGroup `json:"placement_group,nullable"`
	// IDs of Volumes assigned to this Server
	Volumes []int64 `json:"volumes"`
	JSON    serverJSON
}

// serverJSON contains the JSON metadata for the struct [Server]
type serverJSON struct {
	ID              apijson.Field
	BackupWindow    apijson.Field
	Created         apijson.Field
	Datacenter      apijson.Field
	Image           apijson.Field
	IncludedTraffic apijson.Field
	IngoingTraffic  apijson.Field
	Iso             apijson.Field
	Labels          apijson.Field
	Locked          apijson.Field
	Name            apijson.Field
	OutgoingTraffic apijson.Field
	PrimaryDiskSize apijson.Field
	PrivateNet      apijson.Field
	Protection      apijson.Field
	PublicNet       apijson.Field
	RescueEnabled   apijson.Field
	ServerType      apijson.Field
	Status          apijson.Field
	LoadBalancers   apijson.Field
	PlacementGroup  apijson.Field
	Volumes         apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *Server) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Datacenter this Primary IP is located at | Datacenter this Resource is located
// at
type ServerDatacenter struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Description of the Datacenter
	Description string `json:"description,required"`
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	Location ServerDatacenterLocation `json:"location,required"`
	// Unique identifier of the Datacenter
	Name string `json:"name,required"`
	// The Server types the Datacenter can handle
	ServerTypes ServerDatacenterServerTypes `json:"server_types,required"`
	JSON        serverDatacenterJSON
}

// serverDatacenterJSON contains the JSON metadata for the struct
// [ServerDatacenter]
type serverDatacenterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	ServerTypes apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerDatacenter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type ServerDatacenterLocation struct {
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
	JSON        serverDatacenterLocationJSON
}

// serverDatacenterLocationJSON contains the JSON metadata for the struct
// [ServerDatacenterLocation]
type serverDatacenterLocationJSON struct {
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

func (r *ServerDatacenterLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Server types the Datacenter can handle
type ServerDatacenterServerTypes struct {
	// IDs of Server types that are supported and for which the Datacenter has enough
	// resources left
	Available []int64 `json:"available,required"`
	// IDs of Server types that are supported and for which the Datacenter has enough
	// resources left
	AvailableForMigration []int64 `json:"available_for_migration,required"`
	// IDs of Server types that are supported in the Datacenter
	Supported []int64 `json:"supported,required"`
	JSON      serverDatacenterServerTypesJSON
}

// serverDatacenterServerTypesJSON contains the JSON metadata for the struct
// [ServerDatacenterServerTypes]
type serverDatacenterServerTypesJSON struct {
	Available             apijson.Field
	AvailableForMigration apijson.Field
	Supported             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *ServerDatacenterServerTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServerImage struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Type of cpu architecture this image is compatible with. | Type of cpu
	// architecture
	Architecture ServerImageArchitecture `json:"architecture,required"`
	// ID of Server the Image is bound to. Only set for Images of type `backup`.
	BoundTo int64 `json:"bound_to,required,nullable"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Information about the Server the Image was created from
	CreatedFrom ServerImageCreatedFrom `json:"created_from,required,nullable"`
	// Point in time where the Image was deleted (in ISO-8601 format)
	Deleted string `json:"deleted,required,nullable"`
	// Point in time when the Image is considered to be deprecated (in ISO-8601 format)
	Deprecated string `json:"deprecated,required,nullable"`
	// Description of the Image
	Description string `json:"description,required"`
	// Size of the disk contained in the Image in GB
	DiskSize float64 `json:"disk_size,required"`
	// Size of the Image file in our storage in GB. For snapshot Images this is the
	// value relevant for calculating costs for the Image.
	ImageSize float64 `json:"image_size,required,nullable"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Unique identifier of the Image. This value is only set for system Images.
	Name string `json:"name,required,nullable"`
	// Flavor of operating system contained in the Image
	OsFlavor ServerImageOsFlavor `json:"os_flavor,required"`
	// Operating system version
	OsVersion string `json:"os_version,required,nullable"`
	// Protection configuration for the Resource
	Protection ServerImageProtection `json:"protection,required"`
	// Whether the Image can be used or if it's still being created or unavailable
	Status ServerImageStatus `json:"status,required"`
	// Type of the Image
	Type ServerImageType `json:"type,required"`
	// Indicates that rapid deploy of the Image is available
	RapidDeploy bool `json:"rapid_deploy"`
	JSON        serverImageJSON
}

// serverImageJSON contains the JSON metadata for the struct [ServerImage]
type serverImageJSON struct {
	ID           apijson.Field
	Architecture apijson.Field
	BoundTo      apijson.Field
	Created      apijson.Field
	CreatedFrom  apijson.Field
	Deleted      apijson.Field
	Deprecated   apijson.Field
	Description  apijson.Field
	DiskSize     apijson.Field
	ImageSize    apijson.Field
	Labels       apijson.Field
	Name         apijson.Field
	OsFlavor     apijson.Field
	OsVersion    apijson.Field
	Protection   apijson.Field
	Status       apijson.Field
	Type         apijson.Field
	RapidDeploy  apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServerImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cpu architecture this image is compatible with. | Type of cpu
// architecture
type ServerImageArchitecture string

const (
	ServerImageArchitectureArm ServerImageArchitecture = "arm"
	ServerImageArchitectureX86 ServerImageArchitecture = "x86"
)

// Information about the Server the Image was created from
type ServerImageCreatedFrom struct {
	// ID of the Server the Image was created from
	ID int64 `json:"id,required"`
	// Server name at the time the Image was created
	Name string `json:"name,required"`
	JSON serverImageCreatedFromJSON
}

// serverImageCreatedFromJSON contains the JSON metadata for the struct
// [ServerImageCreatedFrom]
type serverImageCreatedFromJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerImageCreatedFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor of operating system contained in the Image
type ServerImageOsFlavor string

const (
	ServerImageOsFlavorAlma    ServerImageOsFlavor = "alma"
	ServerImageOsFlavorCentos  ServerImageOsFlavor = "centos"
	ServerImageOsFlavorDebian  ServerImageOsFlavor = "debian"
	ServerImageOsFlavorFedora  ServerImageOsFlavor = "fedora"
	ServerImageOsFlavorRocky   ServerImageOsFlavor = "rocky"
	ServerImageOsFlavorUbuntu  ServerImageOsFlavor = "ubuntu"
	ServerImageOsFlavorUnknown ServerImageOsFlavor = "unknown"
)

// Protection configuration for the Resource
type ServerImageProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   serverImageProtectionJSON
}

// serverImageProtectionJSON contains the JSON metadata for the struct
// [ServerImageProtection]
type serverImageProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerImageProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the Image can be used or if it's still being created or unavailable
type ServerImageStatus string

const (
	ServerImageStatusAvailable   ServerImageStatus = "available"
	ServerImageStatusCreating    ServerImageStatus = "creating"
	ServerImageStatusUnavailable ServerImageStatus = "unavailable"
)

// Type of the Image
type ServerImageType string

const (
	ServerImageTypeApp       ServerImageType = "app"
	ServerImageTypeBackup    ServerImageType = "backup"
	ServerImageTypeSnapshot  ServerImageType = "snapshot"
	ServerImageTypeSystem    ServerImageType = "system"
	ServerImageTypeTemporary ServerImageType = "temporary"
)

// ISO Image that is attached to this Server. Null if no ISO is attached.
type ServerIso struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Type of cpu architecture this iso is compatible with. Null indicates no
	// restriction on the architecture (wildcard).
	Architecture ServerIsoArchitecture `json:"architecture,required,nullable"`
	// ISO 8601 timestamp of deprecation, null if ISO is still available. After the
	// deprecation time it will no longer be possible to attach the ISO to Servers.
	Deprecated string `json:"deprecated,required,nullable"`
	// Description of the ISO
	Description string `json:"description,required"`
	// Unique identifier of the ISO. Only set for public ISOs
	Name string `json:"name,required,nullable"`
	// Type of the ISO
	Type ServerIsoType `json:"type,required"`
	JSON serverIsoJSON
}

// serverIsoJSON contains the JSON metadata for the struct [ServerIso]
type serverIsoJSON struct {
	ID           apijson.Field
	Architecture apijson.Field
	Deprecated   apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServerIso) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cpu architecture this iso is compatible with. Null indicates no
// restriction on the architecture (wildcard).
type ServerIsoArchitecture string

const (
	ServerIsoArchitectureArm ServerIsoArchitecture = "arm"
	ServerIsoArchitectureX86 ServerIsoArchitecture = "x86"
)

// Type of the ISO
type ServerIsoType string

const (
	ServerIsoTypePrivate ServerIsoType = "private"
	ServerIsoTypePublic  ServerIsoType = "public"
)

type ServerPrivateNet struct {
	AliasIps   []string `json:"alias_ips"`
	Ip         string   `json:"ip"`
	MacAddress string   `json:"mac_address"`
	// ID of the Network
	Network int64 `json:"network"`
	JSON    serverPrivateNetJSON
}

// serverPrivateNetJSON contains the JSON metadata for the struct
// [ServerPrivateNet]
type serverPrivateNetJSON struct {
	AliasIps    apijson.Field
	Ip          apijson.Field
	MacAddress  apijson.Field
	Network     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerPrivateNet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Server
type ServerProtection struct {
	// If true, prevents the Server from being deleted
	Delete bool `json:"delete,required"`
	// If true, prevents the Server from being rebuilt
	Rebuild bool `json:"rebuild,required"`
	JSON    serverProtectionJSON
}

// serverProtectionJSON contains the JSON metadata for the struct
// [ServerProtection]
type serverProtectionJSON struct {
	Delete      apijson.Field
	Rebuild     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Public network information. The Server's IPv4 address can be found in
// `public_net->ipv4->ip`
type ServerPublicNet struct {
	// IDs of Floating IPs assigned to this Server
	FloatingIps []int64 `json:"floating_ips,required"`
	// IP address (v4) and its reverse DNS entry of this Server
	Ipv4 ServerPublicNetIpv4 `json:"ipv4,required,nullable"`
	// IPv6 network assigned to this Server and its reverse DNS entry
	Ipv6 ServerPublicNetIpv6 `json:"ipv6,required,nullable"`
	// Firewalls applied to the public network interface of this Server
	Firewalls []ServerPublicNetFirewall `json:"firewalls"`
	JSON      serverPublicNetJSON
}

// serverPublicNetJSON contains the JSON metadata for the struct [ServerPublicNet]
type serverPublicNetJSON struct {
	FloatingIps apijson.Field
	Ipv4        apijson.Field
	Ipv6        apijson.Field
	Firewalls   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerPublicNet) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// IP address (v4) and its reverse DNS entry of this Server
type ServerPublicNetIpv4 struct {
	// If the IP is blocked by our anti abuse dept
	Blocked bool `json:"blocked,required"`
	// Reverse DNS PTR entry for the IPv4 addresses of this Server
	DnsPtr string `json:"dns_ptr,required"`
	// IP address (v4) of this Server
	Ip string `json:"ip,required"`
	// ID of the Resource
	ID   int64 `json:"id"`
	JSON serverPublicNetIpv4JSON
}

// serverPublicNetIpv4JSON contains the JSON metadata for the struct
// [ServerPublicNetIpv4]
type serverPublicNetIpv4JSON struct {
	Blocked     apijson.Field
	DnsPtr      apijson.Field
	Ip          apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerPublicNetIpv4) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// IPv6 network assigned to this Server and its reverse DNS entry
type ServerPublicNetIpv6 struct {
	// If the IP is blocked by our anti abuse dept
	Blocked bool `json:"blocked,required"`
	// Reverse DNS PTR entries for the IPv6 addresses of this Server, `null` by default
	DnsPtr []ServerPublicNetIpv6DnsPtr `json:"dns_ptr,required,nullable"`
	// IP address (v6) of this Server
	Ip string `json:"ip,required"`
	// ID of the Resource
	ID   int64 `json:"id"`
	JSON serverPublicNetIpv6JSON
}

// serverPublicNetIpv6JSON contains the JSON metadata for the struct
// [ServerPublicNetIpv6]
type serverPublicNetIpv6JSON struct {
	Blocked     apijson.Field
	DnsPtr      apijson.Field
	Ip          apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerPublicNetIpv6) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServerPublicNetIpv6DnsPtr struct {
	// DNS pointer for the specific IP address
	DnsPtr string `json:"dns_ptr,required"`
	// Single IPv4 or IPv6 address | Single IPv6 address of this Server for which the
	// reverse DNS entry has been set up
	Ip   string `json:"ip,required"`
	JSON serverPublicNetIpv6DnsPtrJSON
}

// serverPublicNetIpv6DnsPtrJSON contains the JSON metadata for the struct
// [ServerPublicNetIpv6DnsPtr]
type serverPublicNetIpv6DnsPtrJSON struct {
	DnsPtr      apijson.Field
	Ip          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerPublicNetIpv6DnsPtr) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServerPublicNetFirewall struct {
	// ID of the Resource
	ID int64 `json:"id"`
	// Status of the Firewall on the Server
	Status ServerPublicNetFirewallsStatus `json:"status"`
	JSON   serverPublicNetFirewallJSON
}

// serverPublicNetFirewallJSON contains the JSON metadata for the struct
// [ServerPublicNetFirewall]
type serverPublicNetFirewallJSON struct {
	ID          apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerPublicNetFirewall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Firewall on the Server
type ServerPublicNetFirewallsStatus string

const (
	ServerPublicNetFirewallsStatusApplied ServerPublicNetFirewallsStatus = "applied"
	ServerPublicNetFirewallsStatusPending ServerPublicNetFirewallsStatus = "pending"
)

// Type of Server - determines how much ram, disk and cpu a Server has
type ServerServerType struct {
	// ID of the Server type
	ID int64 `json:"id,required"`
	// Number of cpu cores a Server of this type will have
	Cores int64 `json:"cores,required"`
	// Type of cpu
	CpuType CpuType `json:"cpu_type,required"`
	// True if Server type is deprecated
	Deprecated bool `json:"deprecated,required,nullable"`
	// Description of the Server type
	Description string `json:"description,required"`
	// Disk size a Server of this type will have in GB
	Disk float64 `json:"disk,required"`
	// Memory a Server of this type will have in GB
	Memory float64 `json:"memory,required"`
	// Unique identifier of the Server type
	Name string `json:"name,required"`
	// Prices in different Locations
	Prices []ServerServerTypePrice `json:"prices,required"`
	// Type of Server boot drive. Local has higher speed. Network has better
	// availability.
	StorageType ServerServerTypeStorageType `json:"storage_type,required"`
	JSON        serverServerTypeJSON
}

// serverServerTypeJSON contains the JSON metadata for the struct
// [ServerServerType]
type serverServerTypeJSON struct {
	ID          apijson.Field
	Cores       apijson.Field
	CpuType     apijson.Field
	Deprecated  apijson.Field
	Description apijson.Field
	Disk        apijson.Field
	Memory      apijson.Field
	Name        apijson.Field
	Prices      apijson.Field
	StorageType apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerServerType) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServerServerTypePrice struct {
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
	JSON         serverServerTypePriceJSON
}

// serverServerTypePriceJSON contains the JSON metadata for the struct
// [ServerServerTypePrice]
type serverServerTypePriceJSON struct {
	Location     apijson.Field
	PriceHourly  apijson.Field
	PriceMonthly apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServerServerTypePrice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of Server boot drive. Local has higher speed. Network has better
// availability.
type ServerServerTypeStorageType string

const (
	ServerServerTypeStorageTypeLocal   ServerServerTypeStorageType = "local"
	ServerServerTypeStorageTypeNetwork ServerServerTypeStorageType = "network"
)

// Status of the Server
type ServerStatus string

const (
	ServerStatusDeleting     ServerStatus = "deleting"
	ServerStatusInitializing ServerStatus = "initializing"
	ServerStatusMigrating    ServerStatus = "migrating"
	ServerStatusOff          ServerStatus = "off"
	ServerStatusRebuilding   ServerStatus = "rebuilding"
	ServerStatusRunning      ServerStatus = "running"
	ServerStatusStarting     ServerStatus = "starting"
	ServerStatusStopping     ServerStatus = "stopping"
	ServerStatusUnknown      ServerStatus = "unknown"
)

type ServerPlacementGroup struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Array of IDs of Servers that are part of this Placement Group
	Servers []int64 `json:"servers,required"`
	// Type of the Placement Group
	Type ServerPlacementGroupType `json:"type,required"`
	JSON serverPlacementGroupJSON
}

// serverPlacementGroupJSON contains the JSON metadata for the struct
// [ServerPlacementGroup]
type serverPlacementGroupJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Labels      apijson.Field
	Name        apijson.Field
	Servers     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerPlacementGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of the Placement Group
type ServerPlacementGroupType string

const (
	ServerPlacementGroupTypeSpread ServerPlacementGroupType = "spread"
)

// Response to POST https://api.hetzner.cloud/v1/servers
type ServerNewResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action      shared.Action   `json:"action,required"`
	NextActions []shared.Action `json:"next_actions,required"`
	// Root password when no SSH keys have been specified
	RootPassword string `json:"root_password,required,nullable"`
	// Servers are virtual machines that can be provisioned.
	Server Server `json:"server,required"`
	JSON   serverNewResponseJSON
}

// serverNewResponseJSON contains the JSON metadata for the struct
// [ServerNewResponse]
type serverNewResponseJSON struct {
	Action       apijson.Field
	NextActions  apijson.Field
	RootPassword apijson.Field
	Server       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServerNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/servers/{id}
type ServerGetResponse struct {
	// Servers are virtual machines that can be provisioned.
	Server Server `json:"server"`
	JSON   serverGetResponseJSON
}

// serverGetResponseJSON contains the JSON metadata for the struct
// [ServerGetResponse]
type serverGetResponseJSON struct {
	Server      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to PUT https://api.hetzner.cloud/v1/servers/{id}
type ServerUpdateResponse struct {
	// Servers are virtual machines that can be provisioned.
	Server Server `json:"server"`
	JSON   serverUpdateResponseJSON
}

// serverUpdateResponseJSON contains the JSON metadata for the struct
// [ServerUpdateResponse]
type serverUpdateResponseJSON struct {
	Server      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to DELETE https://api.hetzner.cloud/v1/servers/{id}
type ServerDeleteResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action"`
	JSON   serverDeleteResponseJSON
}

// serverDeleteResponseJSON contains the JSON metadata for the struct
// [ServerDeleteResponse]
type serverDeleteResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServerNewParams struct {
	// ID or name of the Image the Server is created from
	Image param.Field[string] `json:"image,required"`
	// Name of the Server to create (must be unique per Project and a valid hostname as
	// per RFC 1123)
	Name param.Field[string] `json:"name,required"`
	// ID or name of the Server type this Server should be created with
	ServerType param.Field[string] `json:"server_type,required"`
	// Auto-mount Volumes after attach
	Automount param.Field[bool] `json:"automount"`
	// ID or name of Datacenter to create Server in (must not be used together with
	// location)
	Datacenter param.Field[string] `json:"datacenter"`
	// Firewalls which should be applied on the Server's public network interface at
	// creation time
	Firewalls param.Field[[]ServerNewParamsFirewall] `json:"firewalls"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// ID or name of Location to create Server in (must not be used together with
	// datacenter)
	Location param.Field[string] `json:"location"`
	// Network IDs which should be attached to the Server private network interface at
	// the creation time
	Networks param.Field[[]int64] `json:"networks"`
	// ID of the Placement Group the server should be in
	PlacementGroup param.Field[int64] `json:"placement_group"`
	// Public Network options
	PublicNet param.Field[ServerNewParamsPublicNet] `json:"public_net"`
	// SSH key IDs (`integer`) or names (`string`) which should be injected into the
	// Server at creation time
	SshKeys param.Field[[]string] `json:"ssh_keys"`
	// Start Server right after creation. Defaults to true.
	StartAfterCreate param.Field[bool] `json:"start_after_create"`
	// Cloud-Init user data to use during Server creation. This field is limited to
	// 32KiB.
	UserData param.Field[string] `json:"user_data"`
	// Volume IDs which should be attached to the Server at the creation time. Volumes
	// must be in the same Location.
	Volumes param.Field[[]int64] `json:"volumes"`
}

func (r ServerNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServerNewParamsFirewall struct {
	// ID of the Firewall
	Firewall param.Field[int64] `json:"firewall,required"`
}

func (r ServerNewParamsFirewall) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Public Network options
type ServerNewParamsPublicNet struct {
	// Attach an IPv4 on the public NIC. If false, no IPv4 address will be attached.
	// Defaults to true.
	EnableIpv4 param.Field[bool] `json:"enable_ipv4"`
	// Attach an IPv6 on the public NIC. If false, no IPv6 address will be attached.
	// Defaults to true.
	EnableIpv6 param.Field[bool] `json:"enable_ipv6"`
	// ID of the ipv4 Primary IP to use. If omitted and enable_ipv4 is true, a new ipv4
	// Primary IP will automatically be created.
	Ipv4 param.Field[int64] `json:"ipv4"`
	// ID of the ipv6 Primary IP to use. If omitted and enable_ipv6 is true, a new ipv6
	// Primary IP will automatically be created.
	Ipv6 param.Field[int64] `json:"ipv6"`
}

func (r ServerNewParamsPublicNet) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServerUpdateParams struct {
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// New name to set
	Name param.Field[string] `json:"name"`
}

func (r ServerUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServerListParams struct {
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
	Sort param.Field[ServerListParamsSort] `query:"sort"`
	// Can be used multiple times. The response will only contain Server matching the
	// status
	Status param.Field[ServerListParamsStatus] `query:"status"`
}

// URLQuery serializes [ServerListParams]'s query parameters as `url.Values`.
func (r ServerListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times.
type ServerListParamsSort string

const (
	ServerListParamsSortID          ServerListParamsSort = "id"
	ServerListParamsSortIDAsc       ServerListParamsSort = "id:asc"
	ServerListParamsSortIDDesc      ServerListParamsSort = "id:desc"
	ServerListParamsSortName        ServerListParamsSort = "name"
	ServerListParamsSortNameAsc     ServerListParamsSort = "name:asc"
	ServerListParamsSortNameDesc    ServerListParamsSort = "name:desc"
	ServerListParamsSortCreated     ServerListParamsSort = "created"
	ServerListParamsSortCreatedAsc  ServerListParamsSort = "created:asc"
	ServerListParamsSortCreatedDesc ServerListParamsSort = "created:desc"
)

// Can be used multiple times. The response will only contain Server matching the
// status
type ServerListParamsStatus string

const (
	ServerListParamsStatusInitializing ServerListParamsStatus = "initializing"
	ServerListParamsStatusStarting     ServerListParamsStatus = "starting"
	ServerListParamsStatusRunning      ServerListParamsStatus = "running"
	ServerListParamsStatusStopping     ServerListParamsStatus = "stopping"
	ServerListParamsStatusOff          ServerListParamsStatus = "off"
	ServerListParamsStatusDeleting     ServerListParamsStatus = "deleting"
	ServerListParamsStatusRebuilding   ServerListParamsStatus = "rebuilding"
	ServerListParamsStatusMigrating    ServerListParamsStatus = "migrating"
	ServerListParamsStatusUnknown      ServerListParamsStatus = "unknown"
)
