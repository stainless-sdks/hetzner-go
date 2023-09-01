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

// LoadBalancerActionService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerActionService] method
// instead.
type LoadBalancerActionService struct {
	Options []option.RequestOption
}

// NewLoadBalancerActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerActionService(opts ...option.RequestOption) (r *LoadBalancerActionService) {
	r = &LoadBalancerActionService{}
	r.Options = opts
	return
}

// Returns a specific Action for a Load Balancer.
func (r *LoadBalancerActionService) Get(ctx context.Context, id int64, actionID int64, opts ...option.RequestOption) (res *LoadBalancerActionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/%v", id, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all Action objects for a Load Balancer. You can sort the results by
// using the `sort` URI parameter, and filter them with the `status` parameter.
func (r *LoadBalancerActionService) List(ctx context.Context, id int64, query LoadBalancerActionListParams, opts ...option.RequestOption) (res *LoadBalancerActionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Adds a service to a Load Balancer.
//
// #### Call specific error codes
//
// | Code                       | Description                                             |
// | -------------------------- | ------------------------------------------------------- |
// | `source_port_already_used` | The source port you are trying to add is already in use |
func (r *LoadBalancerActionService) AddService(ctx context.Context, id int64, body LoadBalancerActionAddServiceParams, opts ...option.RequestOption) (res *LoadBalancerActionAddServiceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/add_service", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Adds a target to a Load Balancer.
//
// #### Call specific error codes
//
// | Code                                    | Description                                                                                           |
// | --------------------------------------- | ----------------------------------------------------------------------------------------------------- |
// | `cloud_resource_ip_not_allowed`         | The IP you are trying to add as a target belongs to a Hetzner Cloud resource                          |
// | `ip_not_owned`                          | The IP you are trying to add as a target is not owned by the Project owner                            |
// | `load_balancer_not_attached_to_network` | The Load Balancer is not attached to a network                                                        |
// | `robot_unavailable`                     | Robot was not available. The caller may retry the operation after a short delay.                      |
// | `server_not_attached_to_network`        | The server you are trying to add as a target is not attached to the same network as the Load Balancer |
// | `target_already_defined`                | The Load Balancer target you are trying to define is already defined                                  |
func (r *LoadBalancerActionService) AssTarget(ctx context.Context, id int64, body LoadBalancerActionAssTargetParams, opts ...option.RequestOption) (res *LoadBalancerActionAssTargetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/add_target", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Attach a Load Balancer to a Network.
//
// **Call specific error codes**
//
// | Code                             | Description                                                           |
// | -------------------------------- | --------------------------------------------------------------------- |
// | `load_balancer_already_attached` | The Load Balancer is already attached to a network                    |
// | `ip_not_available`               | The provided Network IP is not available                              |
// | `no_subnet_available`            | No Subnet or IP is available for the Load Balancer within the network |
func (r *LoadBalancerActionService) AttachToNetwork(ctx context.Context, id int64, body LoadBalancerActionAttachToNetworkParams, opts ...option.RequestOption) (res *LoadBalancerActionAttachToNetworkResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/attach_to_network", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Change the algorithm that determines to which target new requests are sent.
func (r *LoadBalancerActionService) ChangeAlgorithm(ctx context.Context, id int64, body LoadBalancerActionChangeAlgorithmParams, opts ...option.RequestOption) (res *LoadBalancerActionChangeAlgorithmResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/change_algorithm", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the hostname that will appear when getting the hostname belonging to the
// public IPs (IPv4 and IPv6) of this Load Balancer.
//
// Floating IPs assigned to the Server are not affected by this.
func (r *LoadBalancerActionService) ChangeDnsPtr(ctx context.Context, id int64, body LoadBalancerActionChangeDnsPtrParams, opts ...option.RequestOption) (res *LoadBalancerActionChangeDnsPtrResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/change_dns_ptr", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the protection configuration of a Load Balancer.
func (r *LoadBalancerActionService) ChangeProtection(ctx context.Context, id int64, body LoadBalancerActionChangeProtectionParams, opts ...option.RequestOption) (res *LoadBalancerActionChangeProtectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/change_protection", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the type (Max Services, Max Targets and Max Connections) of a Load
// Balancer.
//
// **Call specific error codes**
//
// | Code                         | Description                                                     |
// | ---------------------------- | --------------------------------------------------------------- |
// | `invalid_load_balancer_type` | The Load Balancer type does not fit for the given Load Balancer |
func (r *LoadBalancerActionService) ChangeType(ctx context.Context, id int64, body LoadBalancerActionChangeTypeParams, opts ...option.RequestOption) (res *LoadBalancerActionChangeTypeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/change_type", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Delete a service of a Load Balancer.
func (r *LoadBalancerActionService) DeleteService(ctx context.Context, id int64, body LoadBalancerActionDeleteServiceParams, opts ...option.RequestOption) (res *LoadBalancerActionDeleteServiceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/delete_service", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detaches a Load Balancer from a network.
func (r *LoadBalancerActionService) DetachFromNetwork(ctx context.Context, id int64, body LoadBalancerActionDetachFromNetworkParams, opts ...option.RequestOption) (res *LoadBalancerActionDetachFromNetworkResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/detach_from_network", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Disable the public interface of a Load Balancer. The Load Balancer will be not
// accessible from the internet via its public IPs.
//
// #### Call specific error codes
//
// | Code                                    | Description                                                                    |
// | --------------------------------------- | ------------------------------------------------------------------------------ |
// | `load_balancer_not_attached_to_network` | The Load Balancer is not attached to a network                                 |
// | `targets_without_use_private_ip`        | The Load Balancer has targets that use the public IP instead of the private IP |
func (r *LoadBalancerActionService) DisablePublicInterface(ctx context.Context, id int64, opts ...option.RequestOption) (res *LoadBalancerActionDisablePublicInterfaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/disable_public_interface", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Enable the public interface of a Load Balancer. The Load Balancer will be
// accessible from the internet via its public IPs.
func (r *LoadBalancerActionService) EnablePublicInterface(ctx context.Context, id int64, opts ...option.RequestOption) (res *LoadBalancerActionEnablePublicInterfaceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/enable_public_interface", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Removes a target from a Load Balancer.
func (r *LoadBalancerActionService) RemoveTarget(ctx context.Context, id int64, body LoadBalancerActionRemoveTargetParams, opts ...option.RequestOption) (res *LoadBalancerActionRemoveTargetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/remove_target", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates a Load Balancer Service.
//
// #### Call specific error codes
//
// | Code                       | Description                                             |
// | -------------------------- | ------------------------------------------------------- |
// | `source_port_already_used` | The source port you are trying to add is already in use |
func (r *LoadBalancerActionService) UpdateService(ctx context.Context, id int64, body LoadBalancerActionUpdateServiceParams, opts ...option.RequestOption) (res *LoadBalancerActionUpdateServiceResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/actions/update_service", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response to GET
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/{action_id}
type LoadBalancerActionGetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionGetResponseJSON
}

// loadBalancerActionGetResponseJSON contains the JSON metadata for the struct
// [LoadBalancerActionGetResponse]
type loadBalancerActionGetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/load_balancers/{id}/actions
type LoadBalancerActionListResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON loadBalancerActionListResponseJSON
}

// loadBalancerActionListResponseJSON contains the JSON metadata for the struct
// [LoadBalancerActionListResponse]
type loadBalancerActionListResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/add_service
type LoadBalancerActionAddServiceResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionAddServiceResponseJSON
}

// loadBalancerActionAddServiceResponseJSON contains the JSON metadata for the
// struct [LoadBalancerActionAddServiceResponse]
type loadBalancerActionAddServiceResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionAddServiceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/add_target
type LoadBalancerActionAssTargetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionAssTargetResponseJSON
}

// loadBalancerActionAssTargetResponseJSON contains the JSON metadata for the
// struct [LoadBalancerActionAssTargetResponse]
type loadBalancerActionAssTargetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionAssTargetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/attach_to_network
type LoadBalancerActionAttachToNetworkResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionAttachToNetworkResponseJSON
}

// loadBalancerActionAttachToNetworkResponseJSON contains the JSON metadata for the
// struct [LoadBalancerActionAttachToNetworkResponse]
type loadBalancerActionAttachToNetworkResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionAttachToNetworkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/change_algorithm
type LoadBalancerActionChangeAlgorithmResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionChangeAlgorithmResponseJSON
}

// loadBalancerActionChangeAlgorithmResponseJSON contains the JSON metadata for the
// struct [LoadBalancerActionChangeAlgorithmResponse]
type loadBalancerActionChangeAlgorithmResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionChangeAlgorithmResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/change_dns_ptr
type LoadBalancerActionChangeDnsPtrResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionChangeDnsPtrResponseJSON
}

// loadBalancerActionChangeDnsPtrResponseJSON contains the JSON metadata for the
// struct [LoadBalancerActionChangeDnsPtrResponse]
type loadBalancerActionChangeDnsPtrResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionChangeDnsPtrResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/change_protection
type LoadBalancerActionChangeProtectionResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionChangeProtectionResponseJSON
}

// loadBalancerActionChangeProtectionResponseJSON contains the JSON metadata for
// the struct [LoadBalancerActionChangeProtectionResponse]
type loadBalancerActionChangeProtectionResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionChangeProtectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/change_type
type LoadBalancerActionChangeTypeResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionChangeTypeResponseJSON
}

// loadBalancerActionChangeTypeResponseJSON contains the JSON metadata for the
// struct [LoadBalancerActionChangeTypeResponse]
type loadBalancerActionChangeTypeResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionChangeTypeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/delete_service
type LoadBalancerActionDeleteServiceResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionDeleteServiceResponseJSON
}

// loadBalancerActionDeleteServiceResponseJSON contains the JSON metadata for the
// struct [LoadBalancerActionDeleteServiceResponse]
type loadBalancerActionDeleteServiceResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionDeleteServiceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/detach_from_network
type LoadBalancerActionDetachFromNetworkResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionDetachFromNetworkResponseJSON
}

// loadBalancerActionDetachFromNetworkResponseJSON contains the JSON metadata for
// the struct [LoadBalancerActionDetachFromNetworkResponse]
type loadBalancerActionDetachFromNetworkResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionDetachFromNetworkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/disable_public_interface
type LoadBalancerActionDisablePublicInterfaceResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionDisablePublicInterfaceResponseJSON
}

// loadBalancerActionDisablePublicInterfaceResponseJSON contains the JSON metadata
// for the struct [LoadBalancerActionDisablePublicInterfaceResponse]
type loadBalancerActionDisablePublicInterfaceResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionDisablePublicInterfaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/enable_public_interface
type LoadBalancerActionEnablePublicInterfaceResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionEnablePublicInterfaceResponseJSON
}

// loadBalancerActionEnablePublicInterfaceResponseJSON contains the JSON metadata
// for the struct [LoadBalancerActionEnablePublicInterfaceResponse]
type loadBalancerActionEnablePublicInterfaceResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionEnablePublicInterfaceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/remove_target
type LoadBalancerActionRemoveTargetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionRemoveTargetResponseJSON
}

// loadBalancerActionRemoveTargetResponseJSON contains the JSON metadata for the
// struct [LoadBalancerActionRemoveTargetResponse]
type loadBalancerActionRemoveTargetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionRemoveTargetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/load_balancers/{id}/actions/update_service
type LoadBalancerActionUpdateServiceResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   loadBalancerActionUpdateServiceResponseJSON
}

// loadBalancerActionUpdateServiceResponseJSON contains the JSON metadata for the
// struct [LoadBalancerActionUpdateServiceResponse]
type loadBalancerActionUpdateServiceResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerActionUpdateServiceResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerActionListParams struct {
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

// URLQuery serializes [LoadBalancerActionListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LoadBalancerActionAddServiceParams struct {
	// Port the Load Balancer will balance to
	DestinationPort param.Field[int64] `json:"destination_port,required"`
	// Service health check
	HealthCheck param.Field[LoadBalancerActionAddServiceParamsHealthCheck] `json:"health_check,required"`
	// Port the Load Balancer listens on
	ListenPort param.Field[int64] `json:"listen_port,required"`
	// Protocol of the Load Balancer
	Protocol param.Field[LoadBalancerActionAddServiceParamsProtocol] `json:"protocol,required"`
	// Is Proxyprotocol enabled or not
	Proxyprotocol param.Field[bool] `json:"proxyprotocol,required"`
	// Configuration option for protocols http and https
	HTTP param.Field[LoadBalancerActionAddServiceParamsHTTP] `json:"http"`
}

func (r LoadBalancerActionAddServiceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service health check
type LoadBalancerActionAddServiceParamsHealthCheck struct {
	// Time interval in seconds health checks are performed
	Interval param.Field[int64] `json:"interval,required"`
	// Port the health check will be performed on
	Port param.Field[int64] `json:"port,required"`
	// Type of the health check
	Protocol param.Field[LoadBalancerActionAddServiceParamsHealthCheckProtocol] `json:"protocol,required"`
	// Unsuccessful retries needed until a target is considered unhealthy; an unhealthy
	// target needs the same number of successful retries to become healthy again
	Retries param.Field[int64] `json:"retries,required"`
	// Time in seconds after an attempt is considered a timeout
	Timeout param.Field[int64] `json:"timeout,required"`
	// Additional configuration for protocol http
	HTTP param.Field[LoadBalancerActionAddServiceParamsHealthCheckHTTP] `json:"http"`
}

func (r LoadBalancerActionAddServiceParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the health check
type LoadBalancerActionAddServiceParamsHealthCheckProtocol string

const (
	LoadBalancerActionAddServiceParamsHealthCheckProtocolHTTP LoadBalancerActionAddServiceParamsHealthCheckProtocol = "http"
	LoadBalancerActionAddServiceParamsHealthCheckProtocolTcp  LoadBalancerActionAddServiceParamsHealthCheckProtocol = "tcp"
)

// Additional configuration for protocol http
type LoadBalancerActionAddServiceParamsHealthCheckHTTP struct {
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

func (r LoadBalancerActionAddServiceParamsHealthCheckHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol of the Load Balancer
type LoadBalancerActionAddServiceParamsProtocol string

const (
	LoadBalancerActionAddServiceParamsProtocolHTTP  LoadBalancerActionAddServiceParamsProtocol = "http"
	LoadBalancerActionAddServiceParamsProtocolHTTPs LoadBalancerActionAddServiceParamsProtocol = "https"
	LoadBalancerActionAddServiceParamsProtocolTcp   LoadBalancerActionAddServiceParamsProtocol = "tcp"
)

// Configuration option for protocols http and https
type LoadBalancerActionAddServiceParamsHTTP struct {
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

func (r LoadBalancerActionAddServiceParamsHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerActionAssTargetParams struct {
	// Type of the resource
	Type param.Field[LoadBalancerActionAssTargetParamsType] `json:"type,required"`
	// IP targets where the traffic should be routed to. It is only possible to use the
	// (Public or vSwitch) IPs of Hetzner Online Root Servers belonging to the project
	// owner. IPs belonging to other users are blocked. Additionally IPs belonging to
	// services provided by Hetzner Cloud (Servers, Load Balancers, ...) are blocked as
	// well. Only present for target type "ip".
	Ip param.Field[LoadBalancerTargetIpParam] `json:"ip"`
	// Configuration for type LabelSelector, required if type is `label_selector`
	LabelSelector param.Field[LoadBalancerActionAssTargetParamsLabelSelector] `json:"label_selector"`
	// Configuration for type Server, required if type is `server`
	Server param.Field[LoadBalancerActionAssTargetParamsServer] `json:"server"`
	// Use the private network IP instead of the public IP of the Server, requires the
	// Server and Load Balancer to be in the same network. Default value is false.
	UsePrivateIp param.Field[bool] `json:"use_private_ip"`
}

func (r LoadBalancerActionAssTargetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the resource
type LoadBalancerActionAssTargetParamsType string

const (
	LoadBalancerActionAssTargetParamsTypeIp            LoadBalancerActionAssTargetParamsType = "ip"
	LoadBalancerActionAssTargetParamsTypeLabelSelector LoadBalancerActionAssTargetParamsType = "label_selector"
	LoadBalancerActionAssTargetParamsTypeServer        LoadBalancerActionAssTargetParamsType = "server"
)

// Configuration for type LabelSelector, required if type is `label_selector`
type LoadBalancerActionAssTargetParamsLabelSelector struct {
	// Label selector
	Selector param.Field[string] `json:"selector,required"`
}

func (r LoadBalancerActionAssTargetParamsLabelSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for type Server, required if type is `server`
type LoadBalancerActionAssTargetParamsServer struct {
	// ID of the Server
	ID param.Field[int64] `json:"id,required"`
}

func (r LoadBalancerActionAssTargetParamsServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerActionAttachToNetworkParams struct {
	// ID of an existing network to attach the Load Balancer to
	Network param.Field[int64] `json:"network,required"`
	// IP to request to be assigned to this Load Balancer; if you do not provide this
	// then you will be auto assigned an IP address
	Ip param.Field[string] `json:"ip"`
}

func (r LoadBalancerActionAttachToNetworkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerActionChangeAlgorithmParams struct {
	// Type of the algorithm | Algorithm of the Load Balancer
	Type param.Field[LoadBalancerActionChangeAlgorithmParamsType] `json:"type,required"`
}

func (r LoadBalancerActionChangeAlgorithmParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the algorithm | Algorithm of the Load Balancer
type LoadBalancerActionChangeAlgorithmParamsType string

const (
	LoadBalancerActionChangeAlgorithmParamsTypeLeastConnections LoadBalancerActionChangeAlgorithmParamsType = "least_connections"
	LoadBalancerActionChangeAlgorithmParamsTypeRoundRobin       LoadBalancerActionChangeAlgorithmParamsType = "round_robin"
)

type LoadBalancerActionChangeDnsPtrParams struct {
	// Hostname to set as a reverse DNS PTR entry
	DnsPtr param.Field[string] `json:"dns_ptr,required"`
	// Public IP address for which the reverse DNS entry should be set
	Ip param.Field[string] `json:"ip,required"`
}

func (r LoadBalancerActionChangeDnsPtrParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerActionChangeProtectionParams struct {
	// If true, prevents the Load Balancer from being deleted
	Delete param.Field[bool] `json:"delete"`
}

func (r LoadBalancerActionChangeProtectionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerActionChangeTypeParams struct {
	// ID or name of Load Balancer type the Load Balancer should migrate to
	LoadBalancerType param.Field[string] `json:"load_balancer_type,required"`
}

func (r LoadBalancerActionChangeTypeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerActionDeleteServiceParams struct {
	// The listen port of the service you want to delete
	ListenPort param.Field[int64] `json:"listen_port,required"`
}

func (r LoadBalancerActionDeleteServiceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerActionDetachFromNetworkParams struct {
	// ID of an existing network to detach the Load Balancer from
	Network param.Field[int64] `json:"network,required"`
}

func (r LoadBalancerActionDetachFromNetworkParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerActionRemoveTargetParams struct {
	// Type of the resource
	Type param.Field[LoadBalancerActionRemoveTargetParamsType] `json:"type,required"`
	// IP targets where the traffic should be routed to. It is only possible to use the
	// (Public or vSwitch) IPs of Hetzner Online Root Servers belonging to the project
	// owner. IPs belonging to other users are blocked. Additionally IPs belonging to
	// services provided by Hetzner Cloud (Servers, Load Balancers, ...) are blocked as
	// well. Only present for target type "ip".
	Ip param.Field[LoadBalancerTargetIpParam] `json:"ip"`
	// Configuration for type LabelSelector, required if type is `label_selector`
	LabelSelector param.Field[LoadBalancerActionRemoveTargetParamsLabelSelector] `json:"label_selector"`
	// Configuration for type Server, required if type is `server`
	Server param.Field[LoadBalancerActionRemoveTargetParamsServer] `json:"server"`
}

func (r LoadBalancerActionRemoveTargetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the resource
type LoadBalancerActionRemoveTargetParamsType string

const (
	LoadBalancerActionRemoveTargetParamsTypeIp            LoadBalancerActionRemoveTargetParamsType = "ip"
	LoadBalancerActionRemoveTargetParamsTypeLabelSelector LoadBalancerActionRemoveTargetParamsType = "label_selector"
	LoadBalancerActionRemoveTargetParamsTypeServer        LoadBalancerActionRemoveTargetParamsType = "server"
)

// Configuration for type LabelSelector, required if type is `label_selector`
type LoadBalancerActionRemoveTargetParamsLabelSelector struct {
	// Label selector
	Selector param.Field[string] `json:"selector,required"`
}

func (r LoadBalancerActionRemoveTargetParamsLabelSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for type Server, required if type is `server`
type LoadBalancerActionRemoveTargetParamsServer struct {
	// ID of the Server
	ID param.Field[int64] `json:"id,required"`
}

func (r LoadBalancerActionRemoveTargetParamsServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LoadBalancerActionUpdateServiceParams struct {
	// Port the Load Balancer will balance to
	DestinationPort param.Field[int64] `json:"destination_port,required"`
	// Service health check
	HealthCheck param.Field[LoadBalancerActionUpdateServiceParamsHealthCheck] `json:"health_check,required"`
	// Port the Load Balancer listens on
	ListenPort param.Field[int64] `json:"listen_port,required"`
	// Protocol of the Load Balancer
	Protocol param.Field[LoadBalancerActionUpdateServiceParamsProtocol] `json:"protocol,required"`
	// Is Proxyprotocol enabled or not
	Proxyprotocol param.Field[bool] `json:"proxyprotocol,required"`
	// Configuration option for protocols http and https
	HTTP param.Field[LoadBalancerActionUpdateServiceParamsHTTP] `json:"http"`
}

func (r LoadBalancerActionUpdateServiceParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Service health check
type LoadBalancerActionUpdateServiceParamsHealthCheck struct {
	// Time interval in seconds health checks are performed
	Interval param.Field[int64] `json:"interval,required"`
	// Port the health check will be performed on
	Port param.Field[int64] `json:"port,required"`
	// Type of the health check
	Protocol param.Field[LoadBalancerActionUpdateServiceParamsHealthCheckProtocol] `json:"protocol,required"`
	// Unsuccessful retries needed until a target is considered unhealthy; an unhealthy
	// target needs the same number of successful retries to become healthy again
	Retries param.Field[int64] `json:"retries,required"`
	// Time in seconds after an attempt is considered a timeout
	Timeout param.Field[int64] `json:"timeout,required"`
	// Additional configuration for protocol http
	HTTP param.Field[LoadBalancerActionUpdateServiceParamsHealthCheckHTTP] `json:"http"`
}

func (r LoadBalancerActionUpdateServiceParamsHealthCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the health check
type LoadBalancerActionUpdateServiceParamsHealthCheckProtocol string

const (
	LoadBalancerActionUpdateServiceParamsHealthCheckProtocolHTTP LoadBalancerActionUpdateServiceParamsHealthCheckProtocol = "http"
	LoadBalancerActionUpdateServiceParamsHealthCheckProtocolTcp  LoadBalancerActionUpdateServiceParamsHealthCheckProtocol = "tcp"
)

// Additional configuration for protocol http
type LoadBalancerActionUpdateServiceParamsHealthCheckHTTP struct {
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

func (r LoadBalancerActionUpdateServiceParamsHealthCheckHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Protocol of the Load Balancer
type LoadBalancerActionUpdateServiceParamsProtocol string

const (
	LoadBalancerActionUpdateServiceParamsProtocolHTTP  LoadBalancerActionUpdateServiceParamsProtocol = "http"
	LoadBalancerActionUpdateServiceParamsProtocolHTTPs LoadBalancerActionUpdateServiceParamsProtocol = "https"
	LoadBalancerActionUpdateServiceParamsProtocolTcp   LoadBalancerActionUpdateServiceParamsProtocol = "tcp"
)

// Configuration option for protocols http and https
type LoadBalancerActionUpdateServiceParamsHTTP struct {
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

func (r LoadBalancerActionUpdateServiceParamsHTTP) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
