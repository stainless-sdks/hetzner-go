// File generated from our OpenAPI spec by Stainless.

package hetzner

import (
	"os"

	"github.com/hetzner/hetzner-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the hetzner API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options           []option.RequestOption
	Actions           *ActionService
	Certificates      *CertificateService
	Datacenters       *DatacenterService
	Firewalls         *FirewallService
	FloatingIps       *FloatingIpService
	Images            *ImageService
	Isos              *IsoService
	LoadBalancerTypes *LoadBalancerTypeService
	LoadBalancers     *LoadBalancerService
	Locations         *LocationService
	Networks          *NetworkService
	PlacementGroups   *PlacementGroupService
	Pricing           *PricingService
	PrimaryIps        *PrimaryIpService
	ServerTypes       *ServerTypeService
	Servers           *ServerService
	SshKeys           *SshKeyService
	Volumes           *VolumeService
}

// NewClient generates a new client with the default option read from the
// environment (HETZNER_API_TOKEN). The option passed in as arguments are applied
// after these default arguments, and all option will be passed down to the
// services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r *Client) {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("HETZNER_API_TOKEN"); ok {
		defaults = append(defaults, option.WithAPIToken(o))
	}
	opts = append(defaults, opts...)

	r = &Client{Options: opts}

	r.Actions = NewActionService(opts...)
	r.Certificates = NewCertificateService(opts...)
	r.Datacenters = NewDatacenterService(opts...)
	r.Firewalls = NewFirewallService(opts...)
	r.FloatingIps = NewFloatingIpService(opts...)
	r.Images = NewImageService(opts...)
	r.Isos = NewIsoService(opts...)
	r.LoadBalancerTypes = NewLoadBalancerTypeService(opts...)
	r.LoadBalancers = NewLoadBalancerService(opts...)
	r.Locations = NewLocationService(opts...)
	r.Networks = NewNetworkService(opts...)
	r.PlacementGroups = NewPlacementGroupService(opts...)
	r.Pricing = NewPricingService(opts...)
	r.PrimaryIps = NewPrimaryIpService(opts...)
	r.ServerTypes = NewServerTypeService(opts...)
	r.Servers = NewServerService(opts...)
	r.SshKeys = NewSshKeyService(opts...)
	r.Volumes = NewVolumeService(opts...)

	return
}
