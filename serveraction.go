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

// ServerActionService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewServerActionService] method
// instead.
type ServerActionService struct {
	Options []option.RequestOption
}

// NewServerActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewServerActionService(opts ...option.RequestOption) (r *ServerActionService) {
	r = &ServerActionService{}
	r.Options = opts
	return
}

// Returns a specific Action object for a Server.
func (r *ServerActionService) Get(ctx context.Context, id int64, actionID int64, opts ...option.RequestOption) (res *ServerActionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/%v", id, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all Action objects for a Server. You can `sort` the results by using the
// sort URI parameter, and filter them with the `status` parameter.
func (r *ServerActionService) List(ctx context.Context, id int64, query ServerActionListParams, opts ...option.RequestOption) (res *ServerActionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Adds a Server to a Placement Group.
//
// Server must be powered off for this command to succeed.
//
// #### Call specific error codes
//
// | Code                 | Description                          |
// | -------------------- | ------------------------------------ |
// | `server_not_stopped` | The action requires a stopped server |
func (r *ServerActionService) AddToPlacementGroup(ctx context.Context, id int64, body ServerActionAddToPlacementGroupParams, opts ...option.RequestOption) (res *ServerActionAddToPlacementGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/add_to_placement_group", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Attaches an ISO to a Server. The Server will immediately see it as a new disk.
// An already attached ISO will automatically be detached before the new ISO is
// attached.
//
// Servers with attached ISOs have a modified boot order: They will try to boot
// from the ISO first before falling back to hard disk.
func (r *ServerActionService) AttachIso(ctx context.Context, id int64, body ServerActionAttachIsoParams, opts ...option.RequestOption) (res *ServerActionAttachIsoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/attach_iso", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Attaches a Server to a network. This will complement the fixed public Server
// interface by adding an additional ethernet interface to the Server which is
// connected to the specified network.
//
// The Server will get an IP auto assigned from a subnet of type `server` in the
// same `network_zone`.
//
// Using the `alias_ips` attribute you can also define one or more additional IPs
// to the Servers. Please note that you will have to configure these IPs by hand on
// your Server since only the primary IP will be given out by DHCP.
//
// **Call specific error codes**
//
// | Code                      | Description                                                    |
// | ------------------------- | -------------------------------------------------------------- |
// | `server_already_attached` | The server is already attached to the network                  |
// | `ip_not_available`        | The provided Network IP is not available                       |
// | `no_subnet_available`     | No Subnet or IP is available for the Server within the network |
// | `networks_overlap`        | The network IP range overlaps with one of the server networks  |
func (r *ServerActionService) AttachToNetwork(ctx context.Context, id int64, body ServerActionAttachToNetworkParams, opts ...option.RequestOption) (res *ServerActionAttachToNetworkResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/attach_to_network", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the alias IPs of an already attached Network. Note that the existing
// aliases for the specified Network will be replaced with these provided in the
// request body. So if you want to add an alias IP, you have to provide the
// existing ones from the Network plus the new alias IP in the request body.
func (r *ServerActionService) ChangeAliasIps(ctx context.Context, id int64, body ServerActionChangeAliasIpsParams, opts ...option.RequestOption) (res *ServerActionChangeAliasIpsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/change_alias_ips", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the hostname that will appear when getting the hostname belonging to the
// primary IPs (IPv4 and IPv6) of this Server.
//
// Floating IPs assigned to the Server are not affected by this.
func (r *ServerActionService) ChangeDnsPtr(ctx context.Context, id int64, body ServerActionChangeDnsPtrParams, opts ...option.RequestOption) (res *ServerActionChangeDnsPtrResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/change_dns_ptr", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the protection configuration of the Server.
func (r *ServerActionService) ChangeProtection(ctx context.Context, id int64, body ServerActionChangeProtectionParams, opts ...option.RequestOption) (res *ServerActionChangeProtectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/change_protection", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the type (Cores, RAM and disk sizes) of a Server.
//
// Server must be powered off for this command to succeed.
//
// This copies the content of its disk, and starts it again.
//
// You can only migrate to Server types with the same `storage_type` and equal or
// bigger disks. Shrinking disks is not possible as it might destroy data.
//
// If the disk gets upgraded, the Server type can not be downgraded any more. If
// you plan to downgrade the Server type, set `upgrade_disk` to `false`.
//
// #### Call specific error codes
//
// | Code                  | Description                                                        |
// | --------------------- | ------------------------------------------------------------------ |
// | `invalid_server_type` | The server type does not fit for the given server or is deprecated |
// | `server_not_stopped`  | The action requires a stopped server                               |
func (r *ServerActionService) ChangeType(ctx context.Context, id int64, body ServerActionChangeTypeParams, opts ...option.RequestOption) (res *ServerActionChangeTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/change_type", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Creates an Image (snapshot) from a Server by copying the contents of its disks.
// This creates a snapshot of the current state of the disk and copies it into an
// Image. If the Server is currently running you must make sure that its disk
// content is consistent. Otherwise, the created Image may not be readable.
//
// To make sure disk content is consistent, we recommend to shut down the Server
// prior to creating an Image.
//
// You can either create a `backup` Image that is bound to the Server and therefore
// will be deleted when the Server is deleted, or you can create an `snapshot`
// Image which is completely independent of the Server it was created from and will
// survive Server deletion. Backup Images are only available when the backup option
// is enabled for the Server. Snapshot Images are billed on a per GB basis.
func (r *ServerActionService) NewImage(ctx context.Context, id int64, body ServerActionNewImageParams, opts ...option.RequestOption) (res *ServerActionNewImageResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/create_image", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detaches a Server from a network. The interface for this network will vanish.
func (r *ServerActionService) DetachFromNetwork(ctx context.Context, id int64, body ServerActionDetachFromNetworkParams, opts ...option.RequestOption) (res *ServerActionDetachFromNetworkResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/detach_from_network", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detaches an ISO from a Server. In case no ISO Image is attached to the Server,
// the status of the returned Action is immediately set to `success`
func (r *ServerActionService) DetachIso(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerActionDetachIsoResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/detach_iso", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Disables the automatic backup option and deletes all existing Backups for a
// Server. No more additional charges for backups will be made.
//
// Caution: This immediately removes all existing backups for the Server!
func (r *ServerActionService) DisableBackup(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerActionDisableBackupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/disable_backup", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Disables the Hetzner Rescue System for a Server. This makes a Server start from
// its disks on next reboot.
//
// Rescue Mode is automatically disabled when you first boot into it or if you do
// not use it for 60 minutes.
//
// Disabling rescue mode will not reboot your Server — you will have to do this
// yourself.
func (r *ServerActionService) DisableRescue(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerActionDisableRescueResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/disable_rescue", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Enables and configures the automatic daily backup option for the Server.
// Enabling automatic backups will increase the price of the Server by 20%. In
// return, you will get seven slots where Images of type backup can be stored.
//
// Backups are automatically created daily.
func (r *ServerActionService) EnableBackup(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerActionEnableBackupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/enable_backup", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Enable the Hetzner Rescue System for this Server. The next time a Server with
// enabled rescue mode boots it will start a special minimal Linux distribution
// designed for repair and reinstall.
//
// In case a Server cannot boot on its own you can use this to access a Server’s
// disks.
//
// Rescue Mode is automatically disabled when you first boot into it or if you do
// not use it for 60 minutes.
//
// Enabling rescue mode will not
// [reboot](https://docs.hetzner.cloud/#server-actions-soft-reboot-a-server) your
// Server — you will have to do this yourself.
func (r *ServerActionService) EnableRescue(ctx context.Context, id int64, body ServerActionEnableRescueParams, opts ...option.RequestOption) (res *ServerActionEnableRescueResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/enable_rescue", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Cuts power to the Server. This forcefully stops it without giving the Server
// operating system time to gracefully stop. May lead to data loss, equivalent to
// pulling the power cord. Power off should only be used when shutdown does not
// work.
func (r *ServerActionService) Poweroff(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerActionPoweroffResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/poweroff", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Starts a Server by turning its power on.
func (r *ServerActionService) Poweron(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerActionPoweronResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/poweron", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Reboots a Server gracefully by sending an ACPI request. The Server operating
// system must support ACPI and react to the request, otherwise the Server will not
// reboot.
func (r *ServerActionService) Reboot(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerActionRebootResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/reboot", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Rebuilds a Server overwriting its disk with the content of an Image, thereby
// **destroying all data** on the target Server
//
// The Image can either be one you have created earlier (`backup` or `snapshot`
// Image) or it can be a completely fresh `system` Image provided by us. You can
// get a list of all available Images with `GET /images`.
//
// Your Server will automatically be powered off before the rebuild command
// executes.
func (r *ServerActionService) Rebuild(ctx context.Context, id int64, body ServerActionRebuildParams, opts ...option.RequestOption) (res *ServerActionRebuildResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/rebuild", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes a Server from a Placement Group.
func (r *ServerActionService) RemoveFromPlacementGroup(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerActionRemoveFromPlacementGroupResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/remove_from_placement_group", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Requests credentials for remote access via VNC over websocket to keyboard,
// monitor, and mouse for a Server. The provided URL is valid for 1 minute, after
// this period a new url needs to be created to connect to the Server. How long the
// connection is open after the initial connect is not subject to this timeout.
func (r *ServerActionService) RequestConsole(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerActionRequestConsoleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/request_console", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cuts power to a Server and starts it again. This forcefully stops it without
// giving the Server operating system time to gracefully stop. This may lead to
// data loss, it’s equivalent to pulling the power cord and plugging it in again.
// Reset should only be used when reboot does not work.
func (r *ServerActionService) Reset(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerActionResetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/reset", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Resets the root password. Only works for Linux systems that are running the qemu
// guest agent. Server must be powered on (status `running`) in order for this
// operation to succeed.
//
// This will generate a new password for this Server and return it.
//
// If this does not succeed you can use the rescue system to netboot the Server and
// manually change your Server password by hand.
func (r *ServerActionService) ResetPassword(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerActionResetPasswordResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/reset_password", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Shuts down a Server gracefully by sending an ACPI shutdown request. The Server
// operating system must support ACPI and react to the request, otherwise the
// Server will not shut down. Please note that the `action` status in this case
// only reflects whether the action was sent to the server. It does not mean that
// the server actually shut down successfully. If you need to ensure that the
// server is off, use the `poweroff` action
func (r *ServerActionService) Shutdown(ctx context.Context, id int64, opts ...option.RequestOption) (res *ServerActionShutdownResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/actions/shutdown", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/servers/{id}/actions/{action_id}
type ServerActionGetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionGetResponseJSON
}

// serverActionGetResponseJSON contains the JSON metadata for the struct
// [ServerActionGetResponse]
type serverActionGetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/servers/{id}/actions
type ServerActionListResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON serverActionListResponseJSON
}

// serverActionListResponseJSON contains the JSON metadata for the struct
// [ServerActionListResponse]
type serverActionListResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/servers/{id}/actions/add_to_placement_group
type ServerActionAddToPlacementGroupResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionAddToPlacementGroupResponseJSON
}

// serverActionAddToPlacementGroupResponseJSON contains the JSON metadata for the
// struct [ServerActionAddToPlacementGroupResponse]
type serverActionAddToPlacementGroupResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionAddToPlacementGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/servers/{id}/actions/attach_iso
type ServerActionAttachIsoResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionAttachIsoResponseJSON
}

// serverActionAttachIsoResponseJSON contains the JSON metadata for the struct
// [ServerActionAttachIsoResponse]
type serverActionAttachIsoResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionAttachIsoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/servers/{id}/actions/attach_to_network
type ServerActionAttachToNetworkResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionAttachToNetworkResponseJSON
}

// serverActionAttachToNetworkResponseJSON contains the JSON metadata for the
// struct [ServerActionAttachToNetworkResponse]
type serverActionAttachToNetworkResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionAttachToNetworkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/servers/{id}/actions/change_alias_ips
type ServerActionChangeAliasIpsResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionChangeAliasIpsResponseJSON
}

// serverActionChangeAliasIpsResponseJSON contains the JSON metadata for the struct
// [ServerActionChangeAliasIpsResponse]
type serverActionChangeAliasIpsResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionChangeAliasIpsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/servers/{id}/actions/change_dns_ptr
type ServerActionChangeDnsPtrResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionChangeDnsPtrResponseJSON
}

// serverActionChangeDnsPtrResponseJSON contains the JSON metadata for the struct
// [ServerActionChangeDnsPtrResponse]
type serverActionChangeDnsPtrResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionChangeDnsPtrResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/servers/{id}/actions/change_protection
type ServerActionChangeProtectionResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionChangeProtectionResponseJSON
}

// serverActionChangeProtectionResponseJSON contains the JSON metadata for the
// struct [ServerActionChangeProtectionResponse]
type serverActionChangeProtectionResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionChangeProtectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/servers/{id}/actions/change_type
type ServerActionChangeTypeResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionChangeTypeResponseJSON
}

// serverActionChangeTypeResponseJSON contains the JSON metadata for the struct
// [ServerActionChangeTypeResponse]
type serverActionChangeTypeResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionChangeTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/servers/{id}/actions/create_image
type ServerActionNewImageResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action                     `json:"action"`
	Image  ServerActionNewImageResponseImage `json:"image"`
	JSON   serverActionNewImageResponseJSON
}

// serverActionNewImageResponseJSON contains the JSON metadata for the struct
// [ServerActionNewImageResponse]
type serverActionNewImageResponseJSON struct {
	Action      apijson.Field
	Image       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionNewImageResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServerActionNewImageResponseImage struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Type of cpu architecture this image is compatible with. | Type of cpu
	// architecture
	Architecture ServerActionNewImageResponseImageArchitecture `json:"architecture,required"`
	// ID of Server the Image is bound to. Only set for Images of type `backup`.
	BoundTo int64 `json:"bound_to,required,nullable"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Information about the Server the Image was created from
	CreatedFrom ServerActionNewImageResponseImageCreatedFrom `json:"created_from,required,nullable"`
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
	OsFlavor ServerActionNewImageResponseImageOsFlavor `json:"os_flavor,required"`
	// Operating system version
	OsVersion string `json:"os_version,required,nullable"`
	// Protection configuration for the Resource
	Protection ServerActionNewImageResponseImageProtection `json:"protection,required"`
	// Whether the Image can be used or if it's still being created or unavailable
	Status ServerActionNewImageResponseImageStatus `json:"status,required"`
	// Type of the Image
	Type ServerActionNewImageResponseImageType `json:"type,required"`
	// Indicates that rapid deploy of the Image is available
	RapidDeploy bool `json:"rapid_deploy"`
	JSON        serverActionNewImageResponseImageJSON
}

// serverActionNewImageResponseImageJSON contains the JSON metadata for the struct
// [ServerActionNewImageResponseImage]
type serverActionNewImageResponseImageJSON struct {
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

func (r *ServerActionNewImageResponseImage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cpu architecture this image is compatible with. | Type of cpu
// architecture
type ServerActionNewImageResponseImageArchitecture string

const (
	ServerActionNewImageResponseImageArchitectureArm ServerActionNewImageResponseImageArchitecture = "arm"
	ServerActionNewImageResponseImageArchitectureX86 ServerActionNewImageResponseImageArchitecture = "x86"
)

// Information about the Server the Image was created from
type ServerActionNewImageResponseImageCreatedFrom struct {
	// ID of the Server the Image was created from
	ID int64 `json:"id,required"`
	// Server name at the time the Image was created
	Name string `json:"name,required"`
	JSON serverActionNewImageResponseImageCreatedFromJSON
}

// serverActionNewImageResponseImageCreatedFromJSON contains the JSON metadata for
// the struct [ServerActionNewImageResponseImageCreatedFrom]
type serverActionNewImageResponseImageCreatedFromJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionNewImageResponseImageCreatedFrom) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Flavor of operating system contained in the Image
type ServerActionNewImageResponseImageOsFlavor string

const (
	ServerActionNewImageResponseImageOsFlavorAlma    ServerActionNewImageResponseImageOsFlavor = "alma"
	ServerActionNewImageResponseImageOsFlavorCentos  ServerActionNewImageResponseImageOsFlavor = "centos"
	ServerActionNewImageResponseImageOsFlavorDebian  ServerActionNewImageResponseImageOsFlavor = "debian"
	ServerActionNewImageResponseImageOsFlavorFedora  ServerActionNewImageResponseImageOsFlavor = "fedora"
	ServerActionNewImageResponseImageOsFlavorRocky   ServerActionNewImageResponseImageOsFlavor = "rocky"
	ServerActionNewImageResponseImageOsFlavorUbuntu  ServerActionNewImageResponseImageOsFlavor = "ubuntu"
	ServerActionNewImageResponseImageOsFlavorUnknown ServerActionNewImageResponseImageOsFlavor = "unknown"
)

// Protection configuration for the Resource
type ServerActionNewImageResponseImageProtection struct {
	// If true, prevents the Resource from being deleted | If true, prevents the
	// Network from being deleted
	Delete bool `json:"delete,required"`
	JSON   serverActionNewImageResponseImageProtectionJSON
}

// serverActionNewImageResponseImageProtectionJSON contains the JSON metadata for
// the struct [ServerActionNewImageResponseImageProtection]
type serverActionNewImageResponseImageProtectionJSON struct {
	Delete      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionNewImageResponseImageProtection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the Image can be used or if it's still being created or unavailable
type ServerActionNewImageResponseImageStatus string

const (
	ServerActionNewImageResponseImageStatusAvailable   ServerActionNewImageResponseImageStatus = "available"
	ServerActionNewImageResponseImageStatusCreating    ServerActionNewImageResponseImageStatus = "creating"
	ServerActionNewImageResponseImageStatusUnavailable ServerActionNewImageResponseImageStatus = "unavailable"
)

// Type of the Image
type ServerActionNewImageResponseImageType string

const (
	ServerActionNewImageResponseImageTypeApp       ServerActionNewImageResponseImageType = "app"
	ServerActionNewImageResponseImageTypeBackup    ServerActionNewImageResponseImageType = "backup"
	ServerActionNewImageResponseImageTypeSnapshot  ServerActionNewImageResponseImageType = "snapshot"
	ServerActionNewImageResponseImageTypeSystem    ServerActionNewImageResponseImageType = "system"
	ServerActionNewImageResponseImageTypeTemporary ServerActionNewImageResponseImageType = "temporary"
)

// Response to POST
// https://api.hetzner.cloud/v1/servers/{id}/actions/detach_from_network
type ServerActionDetachFromNetworkResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionDetachFromNetworkResponseJSON
}

// serverActionDetachFromNetworkResponseJSON contains the JSON metadata for the
// struct [ServerActionDetachFromNetworkResponse]
type serverActionDetachFromNetworkResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionDetachFromNetworkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/servers/{id}/actions/detach_iso
type ServerActionDetachIsoResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionDetachIsoResponseJSON
}

// serverActionDetachIsoResponseJSON contains the JSON metadata for the struct
// [ServerActionDetachIsoResponse]
type serverActionDetachIsoResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionDetachIsoResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/servers/{id}/actions/disable_backup
type ServerActionDisableBackupResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionDisableBackupResponseJSON
}

// serverActionDisableBackupResponseJSON contains the JSON metadata for the struct
// [ServerActionDisableBackupResponse]
type serverActionDisableBackupResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionDisableBackupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/servers/{id}/actions/disable_rescue
type ServerActionDisableRescueResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionDisableRescueResponseJSON
}

// serverActionDisableRescueResponseJSON contains the JSON metadata for the struct
// [ServerActionDisableRescueResponse]
type serverActionDisableRescueResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionDisableRescueResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/servers/{id}/actions/enable_backup
type ServerActionEnableBackupResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionEnableBackupResponseJSON
}

// serverActionEnableBackupResponseJSON contains the JSON metadata for the struct
// [ServerActionEnableBackupResponse]
type serverActionEnableBackupResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionEnableBackupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/servers/{id}/actions/enable_rescue
type ServerActionEnableRescueResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action"`
	// Password that will be set for this Server once the Action succeeds
	RootPassword string `json:"root_password"`
	JSON         serverActionEnableRescueResponseJSON
}

// serverActionEnableRescueResponseJSON contains the JSON metadata for the struct
// [ServerActionEnableRescueResponse]
type serverActionEnableRescueResponseJSON struct {
	Action       apijson.Field
	RootPassword apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServerActionEnableRescueResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/servers/{id}/actions/poweroff
type ServerActionPoweroffResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionPoweroffResponseJSON
}

// serverActionPoweroffResponseJSON contains the JSON metadata for the struct
// [ServerActionPoweroffResponse]
type serverActionPoweroffResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionPoweroffResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/servers/{id}/actions/poweron
type ServerActionPoweronResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionPoweronResponseJSON
}

// serverActionPoweronResponseJSON contains the JSON metadata for the struct
// [ServerActionPoweronResponse]
type serverActionPoweronResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionPoweronResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/servers/{id}/actions/reboot
type ServerActionRebootResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionRebootResponseJSON
}

// serverActionRebootResponseJSON contains the JSON metadata for the struct
// [ServerActionRebootResponse]
type serverActionRebootResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionRebootResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/servers/{id}/actions/rebuild
type ServerActionRebuildResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action"`
	// New root password when not using SSH keys
	RootPassword string `json:"root_password,nullable"`
	JSON         serverActionRebuildResponseJSON
}

// serverActionRebuildResponseJSON contains the JSON metadata for the struct
// [ServerActionRebuildResponse]
type serverActionRebuildResponseJSON struct {
	Action       apijson.Field
	RootPassword apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServerActionRebuildResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/servers/{id}/actions/remove_from_placement_group
type ServerActionRemoveFromPlacementGroupResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionRemoveFromPlacementGroupResponseJSON
}

// serverActionRemoveFromPlacementGroupResponseJSON contains the JSON metadata for
// the struct [ServerActionRemoveFromPlacementGroupResponse]
type serverActionRemoveFromPlacementGroupResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionRemoveFromPlacementGroupResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/servers/{id}/actions/request_console
type ServerActionRequestConsoleResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	// VNC password to use for this connection (this password only works in combination
	// with a wss_url with valid token)
	Password string `json:"password,required"`
	// URL of websocket proxy to use; this includes a token which is valid for a
	// limited time only
	WssURL string `json:"wss_url,required"`
	JSON   serverActionRequestConsoleResponseJSON
}

// serverActionRequestConsoleResponseJSON contains the JSON metadata for the struct
// [ServerActionRequestConsoleResponse]
type serverActionRequestConsoleResponseJSON struct {
	Action      apijson.Field
	Password    apijson.Field
	WssURL      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionRequestConsoleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/servers/{id}/actions/reset
type ServerActionResetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionResetResponseJSON
}

// serverActionResetResponseJSON contains the JSON metadata for the struct
// [ServerActionResetResponse]
type serverActionResetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionResetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/servers/{id}/actions/reset_password
type ServerActionResetPasswordResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action"`
	// Password that will be set for this Server once the Action succeeds
	RootPassword string `json:"root_password"`
	JSON         serverActionResetPasswordResponseJSON
}

// serverActionResetPasswordResponseJSON contains the JSON metadata for the struct
// [ServerActionResetPasswordResponse]
type serverActionResetPasswordResponseJSON struct {
	Action       apijson.Field
	RootPassword apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ServerActionResetPasswordResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/servers/{id}/actions/shutdown
type ServerActionShutdownResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   serverActionShutdownResponseJSON
}

// serverActionShutdownResponseJSON contains the JSON metadata for the struct
// [ServerActionShutdownResponse]
type serverActionShutdownResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerActionShutdownResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServerActionListParams struct {
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
	// Can be used multiple times.
	Sort param.Field[shared.SortParam] `query:"sort"`
	// Can be used multiple times, the response will contain only Actions with
	// specified statuses
	Status param.Field[shared.StatusParam] `query:"status"`
}

// URLQuery serializes [ServerActionListParams]'s query parameters as `url.Values`.
func (r ServerActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ServerActionAddToPlacementGroupParams struct {
	// ID of Placement Group the Server should be added to
	PlacementGroup param.Field[int64] `json:"placement_group,required"`
}

func (r ServerActionAddToPlacementGroupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServerActionAttachIsoParams struct {
	// ID or name of ISO to attach to the Server as listed in GET `/isos`
	Iso param.Field[string] `json:"iso,required"`
}

func (r ServerActionAttachIsoParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServerActionAttachToNetworkParams struct {
	// ID of an existing network to attach the Server to
	Network param.Field[int64] `json:"network,required"`
	// Additional IPs to be assigned to this Server
	AliasIps param.Field[[]string] `json:"alias_ips"`
	// IP to request to be assigned to this Server; if you do not provide this then you
	// will be auto assigned an IP address
	Ip param.Field[string] `json:"ip"`
}

func (r ServerActionAttachToNetworkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServerActionChangeAliasIpsParams struct {
	// New alias IPs to set for this Server
	AliasIps param.Field[[]string] `json:"alias_ips,required"`
	// ID of an existing Network already attached to the Server
	Network param.Field[int64] `json:"network,required"`
}

func (r ServerActionChangeAliasIpsParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServerActionChangeDnsPtrParams struct {
	// Hostname to set as a reverse DNS PTR entry, reset to original value if `null`
	DnsPtr param.Field[string] `json:"dns_ptr,required"`
	// Primary IP address for which the reverse DNS entry should be set
	Ip param.Field[string] `json:"ip,required"`
}

func (r ServerActionChangeDnsPtrParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServerActionChangeProtectionParams struct {
	// If true, prevents the Server from being deleted (currently delete and rebuild
	// attribute needs to have the same value)
	Delete param.Field[bool] `json:"delete"`
	// If true, prevents the Server from being rebuilt (currently delete and rebuild
	// attribute needs to have the same value)
	Rebuild param.Field[bool] `json:"rebuild"`
}

func (r ServerActionChangeProtectionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServerActionChangeTypeParams struct {
	// ID or name of Server type the Server should migrate to
	ServerType param.Field[string] `json:"server_type,required"`
	// If false, do not upgrade the disk (this allows downgrading the Server type
	// later)
	UpgradeDisk param.Field[bool] `json:"upgrade_disk,required"`
}

func (r ServerActionChangeTypeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServerActionNewImageParams struct {
	// Description of the Image, will be auto-generated if not set
	Description param.Field[string] `json:"description"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// Type of Image to create (default: `snapshot`)
	Type param.Field[ServerActionNewImageParamsType] `json:"type"`
}

func (r ServerActionNewImageParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of Image to create (default: `snapshot`)
type ServerActionNewImageParamsType string

const (
	ServerActionNewImageParamsTypeBackup   ServerActionNewImageParamsType = "backup"
	ServerActionNewImageParamsTypeSnapshot ServerActionNewImageParamsType = "snapshot"
)

type ServerActionDetachFromNetworkParams struct {
	// ID of an existing network to detach the Server from
	Network param.Field[int64] `json:"network,required"`
}

func (r ServerActionDetachFromNetworkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServerActionEnableRescueParams struct {
	// Array of SSH key IDs which should be injected into the rescue system.
	SshKeys param.Field[[]int64] `json:"ssh_keys"`
	// Type of rescue system to boot (default: `linux64`)
	Type param.Field[ServerActionEnableRescueParamsType] `json:"type"`
}

func (r ServerActionEnableRescueParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of rescue system to boot (default: `linux64`)
type ServerActionEnableRescueParamsType string

const (
	ServerActionEnableRescueParamsTypeLinux32 ServerActionEnableRescueParamsType = "linux32"
	ServerActionEnableRescueParamsTypeLinux64 ServerActionEnableRescueParamsType = "linux64"
)

type ServerActionRebuildParams struct {
	// ID or name of Image to rebuilt from.
	Image param.Field[string] `json:"image,required"`
}

func (r ServerActionRebuildParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
