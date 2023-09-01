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

// VolumeService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewVolumeService] method instead.
type VolumeService struct {
	Options []option.RequestOption
	Actions *VolumeActionService
}

// NewVolumeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVolumeService(opts ...option.RequestOption) (r *VolumeService) {
	r = &VolumeService{}
	r.Options = opts
	r.Actions = NewVolumeActionService(opts...)
	return
}

// Creates a new Volume attached to a Server. If you want to create a Volume that
// is not attached to a Server, you need to provide the `location` key instead of
// `server`. This can be either the ID or the name of the Location this Volume will
// be created in. Note that a Volume can be attached to a Server only in the same
// Location as the Volume itself.
//
// Specifying the Server during Volume creation will automatically attach the
// Volume to that Server after it has been initialized. In that case, the
// `next_actions` key in the response is an array which contains a single
// `attach_volume` action.
//
// The minimum Volume size is 10GB and the maximum size is 10TB (10240GB).
//
// A volume’s name can consist of alphanumeric characters, dashes, underscores, and
// dots, but has to start and end with an alphanumeric character. The total length
// is limited to 64 characters. Volume names must be unique per Project.
//
// #### Call specific error codes
//
// | Code                        | Description                                         |
// | --------------------------- | --------------------------------------------------- |
// | `no_space_left_in_location` | There is no volume space left in the given location |
func (r *VolumeService) New(ctx context.Context, body VolumeNewParams, opts ...option.RequestOption) (res *VolumeNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a specific Volume object.
func (r *VolumeService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *VolumeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("volumes/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Volume properties.
//
// Note that when updating labels, the volume’s current set of labels will be
// replaced with the labels provided in the request body. So, for example, if you
// want to add a new label, you have to provide all existing labels plus the new
// label in the request body.
func (r *VolumeService) Update(ctx context.Context, id int64, body VolumeUpdateParams, opts ...option.RequestOption) (res *VolumeUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("volumes/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Gets all existing Volumes that you have available.
func (r *VolumeService) List(ctx context.Context, query VolumeListParams, opts ...option.RequestOption) (res *VolumeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a volume. All Volume data is irreversibly destroyed. The Volume must not
// be attached to a Server and it must not have delete protection enabled.
func (r *VolumeService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("volumes/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Response to POST https://api.hetzner.cloud/v1/volumes
type VolumeNewResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action      shared.Action   `json:"action,required"`
	NextActions []shared.Action `json:"next_actions,required"`
	// A Volume is a highly-available, scalable, and SSD-based block storage for
	// Servers. Pricing for Volumes depends on the Volume size and Location, not the
	// actual used storage. Please see
	// [Hetzner Wiki](https://wiki.hetzner.de/index.php/CloudServer/en#Volumes) for
	// more details about Volumes.
	Volume VolumeNewResponseVolume `json:"volume,required"`
	JSON   volumeNewResponseJSON
}

// volumeNewResponseJSON contains the JSON metadata for the struct
// [VolumeNewResponse]
type volumeNewResponseJSON struct {
	Action      apijson.Field
	NextActions apijson.Field
	Volume      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Volume is a highly-available, scalable, and SSD-based block storage for
// Servers. Pricing for Volumes depends on the Volume size and Location, not the
// actual used storage. Please see
// [Hetzner Wiki](https://wiki.hetzner.de/index.php/CloudServer/en#Volumes) for
// more details about Volumes.
type VolumeNewResponseVolume struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Filesystem of the Volume if formatted on creation, null if not formatted on
	// creation
	Format string `json:"format,required,nullable"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Device path on the file system for the Volume
	LinuxDevice string `json:"linux_device,required"`
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	Location VolumeNewResponseVolumeLocation `json:"location,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Protection configuration for the Resource
	Protection VolumeNewResponseVolumeProtection `json:"protection,required"`
	// ID of the Server the Volume is attached to, null if it is not attached at all
	Server int64 `json:"server,required,nullable"`
	// Size in GB of the Volume
	Size float64 `json:"size,required"`
	// Current status of the Volume
	Status VolumeNewResponseVolumeStatus `json:"status,required"`
	JSON   volumeNewResponseVolumeJSON
}

// volumeNewResponseVolumeJSON contains the JSON metadata for the struct
// [VolumeNewResponseVolume]
type volumeNewResponseVolumeJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Format      apijson.Field
	Labels      apijson.Field
	LinuxDevice apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	Protection  apijson.Field
	Server      apijson.Field
	Size        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeNewResponseVolume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type VolumeNewResponseVolumeLocation struct {
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
	JSON        volumeNewResponseVolumeLocationJSON
}

// volumeNewResponseVolumeLocationJSON contains the JSON metadata for the struct
// [VolumeNewResponseVolumeLocation]
type volumeNewResponseVolumeLocationJSON struct {
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

func (r *VolumeNewResponseVolumeLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type VolumeNewResponseVolumeProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   volumeNewResponseVolumeProtectionJSON
}

// volumeNewResponseVolumeProtectionJSON contains the JSON metadata for the struct
// [VolumeNewResponseVolumeProtection]
type volumeNewResponseVolumeProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeNewResponseVolumeProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the Volume
type VolumeNewResponseVolumeStatus string

const (
	VolumeNewResponseVolumeStatusAvailable VolumeNewResponseVolumeStatus = "available"
	VolumeNewResponseVolumeStatusCreating  VolumeNewResponseVolumeStatus = "creating"
)

// Response to GET https://api.hetzner.cloud/v1/volumes/{id}
type VolumeGetResponse struct {
	// A Volume is a highly-available, scalable, and SSD-based block storage for
	// Servers. Pricing for Volumes depends on the Volume size and Location, not the
	// actual used storage. Please see
	// [Hetzner Wiki](https://wiki.hetzner.de/index.php/CloudServer/en#Volumes) for
	// more details about Volumes.
	Volume VolumeGetResponseVolume `json:"volume,required"`
	JSON   volumeGetResponseJSON
}

// volumeGetResponseJSON contains the JSON metadata for the struct
// [VolumeGetResponse]
type volumeGetResponseJSON struct {
	Volume      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Volume is a highly-available, scalable, and SSD-based block storage for
// Servers. Pricing for Volumes depends on the Volume size and Location, not the
// actual used storage. Please see
// [Hetzner Wiki](https://wiki.hetzner.de/index.php/CloudServer/en#Volumes) for
// more details about Volumes.
type VolumeGetResponseVolume struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Filesystem of the Volume if formatted on creation, null if not formatted on
	// creation
	Format string `json:"format,required,nullable"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Device path on the file system for the Volume
	LinuxDevice string `json:"linux_device,required"`
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	Location VolumeGetResponseVolumeLocation `json:"location,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Protection configuration for the Resource
	Protection VolumeGetResponseVolumeProtection `json:"protection,required"`
	// ID of the Server the Volume is attached to, null if it is not attached at all
	Server int64 `json:"server,required,nullable"`
	// Size in GB of the Volume
	Size float64 `json:"size,required"`
	// Current status of the Volume
	Status VolumeGetResponseVolumeStatus `json:"status,required"`
	JSON   volumeGetResponseVolumeJSON
}

// volumeGetResponseVolumeJSON contains the JSON metadata for the struct
// [VolumeGetResponseVolume]
type volumeGetResponseVolumeJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Format      apijson.Field
	Labels      apijson.Field
	LinuxDevice apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	Protection  apijson.Field
	Server      apijson.Field
	Size        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeGetResponseVolume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type VolumeGetResponseVolumeLocation struct {
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
	JSON        volumeGetResponseVolumeLocationJSON
}

// volumeGetResponseVolumeLocationJSON contains the JSON metadata for the struct
// [VolumeGetResponseVolumeLocation]
type volumeGetResponseVolumeLocationJSON struct {
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

func (r *VolumeGetResponseVolumeLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type VolumeGetResponseVolumeProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   volumeGetResponseVolumeProtectionJSON
}

// volumeGetResponseVolumeProtectionJSON contains the JSON metadata for the struct
// [VolumeGetResponseVolumeProtection]
type volumeGetResponseVolumeProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeGetResponseVolumeProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the Volume
type VolumeGetResponseVolumeStatus string

const (
	VolumeGetResponseVolumeStatusAvailable VolumeGetResponseVolumeStatus = "available"
	VolumeGetResponseVolumeStatusCreating  VolumeGetResponseVolumeStatus = "creating"
)

// Response to PUT https://api.hetzner.cloud/v1/volumes/{id}
type VolumeUpdateResponse struct {
	// A Volume is a highly-available, scalable, and SSD-based block storage for
	// Servers. Pricing for Volumes depends on the Volume size and Location, not the
	// actual used storage. Please see
	// [Hetzner Wiki](https://wiki.hetzner.de/index.php/CloudServer/en#Volumes) for
	// more details about Volumes.
	Volume VolumeUpdateResponseVolume `json:"volume,required"`
	JSON   volumeUpdateResponseJSON
}

// volumeUpdateResponseJSON contains the JSON metadata for the struct
// [VolumeUpdateResponse]
type volumeUpdateResponseJSON struct {
	Volume      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Volume is a highly-available, scalable, and SSD-based block storage for
// Servers. Pricing for Volumes depends on the Volume size and Location, not the
// actual used storage. Please see
// [Hetzner Wiki](https://wiki.hetzner.de/index.php/CloudServer/en#Volumes) for
// more details about Volumes.
type VolumeUpdateResponseVolume struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Filesystem of the Volume if formatted on creation, null if not formatted on
	// creation
	Format string `json:"format,required,nullable"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Device path on the file system for the Volume
	LinuxDevice string `json:"linux_device,required"`
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	Location VolumeUpdateResponseVolumeLocation `json:"location,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Protection configuration for the Resource
	Protection VolumeUpdateResponseVolumeProtection `json:"protection,required"`
	// ID of the Server the Volume is attached to, null if it is not attached at all
	Server int64 `json:"server,required,nullable"`
	// Size in GB of the Volume
	Size float64 `json:"size,required"`
	// Current status of the Volume
	Status VolumeUpdateResponseVolumeStatus `json:"status,required"`
	JSON   volumeUpdateResponseVolumeJSON
}

// volumeUpdateResponseVolumeJSON contains the JSON metadata for the struct
// [VolumeUpdateResponseVolume]
type volumeUpdateResponseVolumeJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Format      apijson.Field
	Labels      apijson.Field
	LinuxDevice apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	Protection  apijson.Field
	Server      apijson.Field
	Size        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeUpdateResponseVolume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type VolumeUpdateResponseVolumeLocation struct {
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
	JSON        volumeUpdateResponseVolumeLocationJSON
}

// volumeUpdateResponseVolumeLocationJSON contains the JSON metadata for the struct
// [VolumeUpdateResponseVolumeLocation]
type volumeUpdateResponseVolumeLocationJSON struct {
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

func (r *VolumeUpdateResponseVolumeLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type VolumeUpdateResponseVolumeProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   volumeUpdateResponseVolumeProtectionJSON
}

// volumeUpdateResponseVolumeProtectionJSON contains the JSON metadata for the
// struct [VolumeUpdateResponseVolumeProtection]
type volumeUpdateResponseVolumeProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeUpdateResponseVolumeProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the Volume
type VolumeUpdateResponseVolumeStatus string

const (
	VolumeUpdateResponseVolumeStatusAvailable VolumeUpdateResponseVolumeStatus = "available"
	VolumeUpdateResponseVolumeStatusCreating  VolumeUpdateResponseVolumeStatus = "creating"
)

// Response to GET https://api.hetzner.cloud/v1/volumes
type VolumeListResponse struct {
	Volumes []VolumeListResponseVolume `json:"volumes,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON volumeListResponseJSON
}

// volumeListResponseJSON contains the JSON metadata for the struct
// [VolumeListResponse]
type volumeListResponseJSON struct {
	Volumes     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A Volume is a highly-available, scalable, and SSD-based block storage for
// Servers. Pricing for Volumes depends on the Volume size and Location, not the
// actual used storage. Please see
// [Hetzner Wiki](https://wiki.hetzner.de/index.php/CloudServer/en#Volumes) for
// more details about Volumes.
type VolumeListResponseVolume struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Filesystem of the Volume if formatted on creation, null if not formatted on
	// creation
	Format string `json:"format,required,nullable"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Device path on the file system for the Volume
	LinuxDevice string `json:"linux_device,required"`
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	Location VolumeListResponseVolumesLocation `json:"location,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Protection configuration for the Resource
	Protection VolumeListResponseVolumesProtection `json:"protection,required"`
	// ID of the Server the Volume is attached to, null if it is not attached at all
	Server int64 `json:"server,required,nullable"`
	// Size in GB of the Volume
	Size float64 `json:"size,required"`
	// Current status of the Volume
	Status VolumeListResponseVolumesStatus `json:"status,required"`
	JSON   volumeListResponseVolumeJSON
}

// volumeListResponseVolumeJSON contains the JSON metadata for the struct
// [VolumeListResponseVolume]
type volumeListResponseVolumeJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Format      apijson.Field
	Labels      apijson.Field
	LinuxDevice apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	Protection  apijson.Field
	Server      apijson.Field
	Size        apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeListResponseVolume) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type VolumeListResponseVolumesLocation struct {
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
	JSON        volumeListResponseVolumesLocationJSON
}

// volumeListResponseVolumesLocationJSON contains the JSON metadata for the struct
// [VolumeListResponseVolumesLocation]
type volumeListResponseVolumesLocationJSON struct {
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

func (r *VolumeListResponseVolumesLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Protection configuration for the Resource
type VolumeListResponseVolumesProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   volumeListResponseVolumesProtectionJSON
}

// volumeListResponseVolumesProtectionJSON contains the JSON metadata for the
// struct [VolumeListResponseVolumesProtection]
type volumeListResponseVolumesProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeListResponseVolumesProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the Volume
type VolumeListResponseVolumesStatus string

const (
	VolumeListResponseVolumesStatusAvailable VolumeListResponseVolumesStatus = "available"
	VolumeListResponseVolumesStatusCreating  VolumeListResponseVolumesStatus = "creating"
)

type VolumeNewParams struct {
	// Name of the volume
	Name param.Field[string] `json:"name,required"`
	// Size of the Volume in GB
	Size param.Field[int64] `json:"size,required"`
	// Auto-mount Volume after attach. `server` must be provided.
	Automount param.Field[bool] `json:"automount"`
	// Format Volume after creation. One of: `xfs`, `ext4`
	Format param.Field[string] `json:"format"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// Location to create the Volume in (can be omitted if Server is specified)
	Location param.Field[string] `json:"location"`
	// Server to which to attach the Volume once it's created (Volume will be created
	// in the same Location as the server)
	Server param.Field[int64] `json:"server"`
}

func (r VolumeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VolumeUpdateParams struct {
	// New Volume name
	Name param.Field[string] `json:"name,required"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r VolumeUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VolumeListParams struct {
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
	Sort param.Field[VolumeListParamsSort] `query:"sort"`
	// Can be used multiple times. The response will only contain Volumes matching the
	// status.
	Status param.Field[VolumeListParamsStatus] `query:"status"`
}

// URLQuery serializes [VolumeListParams]'s query parameters as `url.Values`.
func (r VolumeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times.
type VolumeListParamsSort string

const (
	VolumeListParamsSortID          VolumeListParamsSort = "id"
	VolumeListParamsSortIDAsc       VolumeListParamsSort = "id:asc"
	VolumeListParamsSortIDDesc      VolumeListParamsSort = "id:desc"
	VolumeListParamsSortName        VolumeListParamsSort = "name"
	VolumeListParamsSortNameAsc     VolumeListParamsSort = "name:asc"
	VolumeListParamsSortNameDesc    VolumeListParamsSort = "name:desc"
	VolumeListParamsSortCreated     VolumeListParamsSort = "created"
	VolumeListParamsSortCreatedAsc  VolumeListParamsSort = "created:asc"
	VolumeListParamsSortCreatedDesc VolumeListParamsSort = "created:desc"
)

// Can be used multiple times. The response will only contain Volumes matching the
// status.
type VolumeListParamsStatus string

const (
	VolumeListParamsStatusAvailable VolumeListParamsStatus = "available"
	VolumeListParamsStatusCreating  VolumeListParamsStatus = "creating"
)
