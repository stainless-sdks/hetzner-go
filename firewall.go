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

// FirewallService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewFirewallService] method instead.
type FirewallService struct {
	Options []option.RequestOption
	Actions *FirewallActionService
}

// NewFirewallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFirewallService(opts ...option.RequestOption) (r *FirewallService) {
	r = &FirewallService{}
	r.Options = opts
	r.Actions = NewFirewallActionService(opts...)
	return
}

// Creates a new Firewall.
//
// #### Call specific error codes
//
// | Code                          | Description                                                   |
// | ----------------------------- | ------------------------------------------------------------- |
// | `server_already_added`        | Server added more than one time to resource                   |
// | `incompatible_network_type`   | The Network type is incompatible for the given resource       |
// | `firewall_resource_not_found` | The resource the Firewall should be attached to was not found |
func (r *FirewallService) New(ctx context.Context, body FirewallNewParams, opts ...option.RequestOption) (res *FirewallNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "firewalls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a specific Firewall object.
func (r *FirewallService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *FirewallGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("firewalls/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Firewall.
//
// Note that when updating labels, the Firewall's current set of labels will be
// replaced with the labels provided in the request body. So, for example, if you
// want to add a new label, you have to provide all existing labels plus the new
// label in the request body.
//
// Note: if the Firewall object changes during the request, the response will be a
// “conflict” error.
func (r *FirewallService) Update(ctx context.Context, id int64, body FirewallUpdateParams, opts ...option.RequestOption) (res *FirewallUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("firewalls/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all Firewall objects.
func (r *FirewallService) List(ctx context.Context, query FirewallListParams, opts ...option.RequestOption) (res *FirewallListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "firewalls"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a Firewall.
//
// #### Call specific error codes
//
// | Code              | Description                               |
// | ----------------- | ----------------------------------------- |
// | `resource_in_use` | Firewall must not be in use to be deleted |
func (r *FirewallService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("firewalls/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Firewalls can limit the network access to or from your resources.
type Firewall struct {
	// ID of the Resource
	ID        int64               `json:"id,required"`
	AppliedTo []FirewallAppliedTo `json:"applied_to,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Name of the Resource. Must be unique per Project.
	Name  string `json:"name,required"`
	Rules []Rule `json:"rules,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels"`
	JSON   firewallJSON
}

// firewallJSON contains the JSON metadata for the struct [Firewall]
type firewallJSON struct {
	ID          apijson.Field
	AppliedTo   apijson.Field
	Created     apijson.Field
	Name        apijson.Field
	Rules       apijson.Field
	Labels      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Firewall) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Resource a Firewall should be applied to.
type FirewallAppliedTo struct {
	// Type of resource referenced
	Type               FirewallAppliedToType                `json:"type,required"`
	AppliedToResources []FirewallAppliedToAppliedToResource `json:"applied_to_resources"`
	// Configuration for type LabelSelector, required if type is `label_selector`
	LabelSelector FirewallAppliedToLabelSelector `json:"label_selector"`
	// ID of the Resource
	Server FirewallAppliedToServer `json:"server"`
	JSON   firewallAppliedToJSON
}

// firewallAppliedToJSON contains the JSON metadata for the struct
// [FirewallAppliedTo]
type firewallAppliedToJSON struct {
	Type               apijson.Field
	AppliedToResources apijson.Field
	LabelSelector      apijson.Field
	Server             apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *FirewallAppliedTo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of resource referenced
type FirewallAppliedToType string

const (
	FirewallAppliedToTypeLabelSelector FirewallAppliedToType = "label_selector"
	FirewallAppliedToTypeServer        FirewallAppliedToType = "server"
)

type FirewallAppliedToAppliedToResource struct {
	// ID of the Resource
	Server FirewallAppliedToAppliedToResourcesServer `json:"server"`
	// Type of resource referenced
	Type FirewallAppliedToAppliedToResourcesType `json:"type"`
	JSON firewallAppliedToAppliedToResourceJSON
}

// firewallAppliedToAppliedToResourceJSON contains the JSON metadata for the struct
// [FirewallAppliedToAppliedToResource]
type firewallAppliedToAppliedToResourceJSON struct {
	Server      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAppliedToAppliedToResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the Resource
type FirewallAppliedToAppliedToResourcesServer struct {
	// ID of the Resource | ID of the Server
	ID   int64 `json:"id,required"`
	JSON firewallAppliedToAppliedToResourcesServerJSON
}

// firewallAppliedToAppliedToResourcesServerJSON contains the JSON metadata for the
// struct [FirewallAppliedToAppliedToResourcesServer]
type firewallAppliedToAppliedToResourcesServerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAppliedToAppliedToResourcesServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of resource referenced
type FirewallAppliedToAppliedToResourcesType string

const (
	FirewallAppliedToAppliedToResourcesTypeServer FirewallAppliedToAppliedToResourcesType = "server"
)

// Configuration for type LabelSelector, required if type is `label_selector`
type FirewallAppliedToLabelSelector struct {
	// Label selector
	Selector string `json:"selector,required"`
	JSON     firewallAppliedToLabelSelectorJSON
}

// firewallAppliedToLabelSelectorJSON contains the JSON metadata for the struct
// [FirewallAppliedToLabelSelector]
type firewallAppliedToLabelSelectorJSON struct {
	Selector    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAppliedToLabelSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// ID of the Resource
type FirewallAppliedToServer struct {
	// ID of the Resource | ID of the Server
	ID   int64 `json:"id,required"`
	JSON firewallAppliedToServerJSON
}

// firewallAppliedToServerJSON contains the JSON metadata for the struct
// [FirewallAppliedToServer]
type firewallAppliedToServerJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAppliedToServer) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Rule of a firewall.
type Rule struct {
	// Select traffic direction on which rule should be applied. Use `source_ips` for
	// direction `in` and `destination_ips` for direction `out`.
	Direction RuleDirection `json:"direction,required"`
	// Type of traffic to allow
	Protocol RuleProtocol `json:"protocol,required"`
	// Description of the Rule
	Description string `json:"description,nullable"`
	// List of permitted IPv4/IPv6 addresses in CIDR notation. Use `0.0.0.0/0` to allow
	// all IPv4 addresses and `::/0` to allow all IPv6 addresses. You can specify 100
	// CIDRs at most.
	DestinationIPs []string `json:"destination_ips"`
	// Port or port range to which traffic will be allowed, only applicable for
	// protocols TCP and UDP. A port range can be specified by separating two ports
	// with a dash, e.g `1024-5000`.
	Port string `json:"port"`
	// List of permitted IPv4/IPv6 addresses in CIDR notation. Use `0.0.0.0/0` to allow
	// all IPv4 addresses and `::/0` to allow all IPv6 addresses. You can specify 100
	// CIDRs at most.
	SourceIPs []string `json:"source_ips"`
	JSON      ruleJSON
}

// ruleJSON contains the JSON metadata for the struct [Rule]
type ruleJSON struct {
	Direction      apijson.Field
	Protocol       apijson.Field
	Description    apijson.Field
	DestinationIPs apijson.Field
	Port           apijson.Field
	SourceIPs      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Rule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Select traffic direction on which rule should be applied. Use `source_ips` for
// direction `in` and `destination_ips` for direction `out`.
type RuleDirection string

const (
	RuleDirectionIn  RuleDirection = "in"
	RuleDirectionOut RuleDirection = "out"
)

// Type of traffic to allow
type RuleProtocol string

const (
	RuleProtocolEsp  RuleProtocol = "esp"
	RuleProtocolGre  RuleProtocol = "gre"
	RuleProtocolIcmp RuleProtocol = "icmp"
	RuleProtocolTcp  RuleProtocol = "tcp"
	RuleProtocolUdp  RuleProtocol = "udp"
)

// Rule of a firewall.
type RuleParam struct {
	// Select traffic direction on which rule should be applied. Use `source_ips` for
	// direction `in` and `destination_ips` for direction `out`.
	Direction param.Field[RuleDirection] `json:"direction,required"`
	// Type of traffic to allow
	Protocol param.Field[RuleProtocol] `json:"protocol,required"`
	// Description of the Rule
	Description param.Field[string] `json:"description"`
	// List of permitted IPv4/IPv6 addresses in CIDR notation. Use `0.0.0.0/0` to allow
	// all IPv4 addresses and `::/0` to allow all IPv6 addresses. You can specify 100
	// CIDRs at most.
	DestinationIPs param.Field[[]string] `json:"destination_ips"`
	// Port or port range to which traffic will be allowed, only applicable for
	// protocols TCP and UDP. A port range can be specified by separating two ports
	// with a dash, e.g `1024-5000`.
	Port param.Field[string] `json:"port"`
	// List of permitted IPv4/IPv6 addresses in CIDR notation. Use `0.0.0.0/0` to allow
	// all IPv4 addresses and `::/0` to allow all IPv6 addresses. You can specify 100
	// CIDRs at most.
	SourceIPs param.Field[[]string] `json:"source_ips"`
}

func (r RuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Response to POST https://api.hetzner.cloud/v1/firewalls
type FirewallNewResponse struct {
	Actions []shared.Action `json:"actions"`
	// Firewalls can limit the network access to or from your resources.
	Firewall Firewall `json:"firewall"`
	JSON     firewallNewResponseJSON
}

// firewallNewResponseJSON contains the JSON metadata for the struct
// [FirewallNewResponse]
type firewallNewResponseJSON struct {
	Actions     apijson.Field
	Firewall    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/firewalls/{id}
type FirewallGetResponse struct {
	// Firewalls can limit the network access to or from your resources.
	Firewall Firewall `json:"firewall,required"`
	JSON     firewallGetResponseJSON
}

// firewallGetResponseJSON contains the JSON metadata for the struct
// [FirewallGetResponse]
type firewallGetResponseJSON struct {
	Firewall    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to PUT https://api.hetzner.cloud/v1/firewalls/{id}
type FirewallUpdateResponse struct {
	// Firewalls can limit the network access to or from your resources.
	Firewall Firewall `json:"firewall,required"`
	JSON     firewallUpdateResponseJSON
}

// firewallUpdateResponseJSON contains the JSON metadata for the struct
// [FirewallUpdateResponse]
type firewallUpdateResponseJSON struct {
	Firewall    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/firewalls
type FirewallListResponse struct {
	Firewalls []Firewall `json:"firewalls,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON firewallListResponseJSON
}

// firewallListResponseJSON contains the JSON metadata for the struct
// [FirewallListResponse]
type firewallListResponseJSON struct {
	Firewalls   apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallNewParams struct {
	// Name of the Firewall
	Name param.Field[string] `json:"name,required"`
	// Resources the Firewall should be applied to after creation
	ApplyTo param.Field[[]FirewallNewParamsApplyTo] `json:"apply_to"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// Array of rules
	Rules param.Field[[]RuleParam] `json:"rules"`
}

func (r FirewallNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Resource a Firewall should be applied to.
type FirewallNewParamsApplyTo struct {
	// Type of the resource
	Type param.Field[FirewallNewParamsApplyToType] `json:"type,required"`
	// Configuration for type LabelSelector, required if type is `label_selector`
	LabelSelector param.Field[FirewallNewParamsApplyToLabelSelector] `json:"label_selector"`
	// ID of the Resource
	Server param.Field[FirewallNewParamsApplyToServer] `json:"server"`
}

func (r FirewallNewParamsApplyTo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the resource
type FirewallNewParamsApplyToType string

const (
	FirewallNewParamsApplyToTypeLabelSelector FirewallNewParamsApplyToType = "label_selector"
	FirewallNewParamsApplyToTypeServer        FirewallNewParamsApplyToType = "server"
)

// Configuration for type LabelSelector, required if type is `label_selector`
type FirewallNewParamsApplyToLabelSelector struct {
	// Label selector
	Selector param.Field[string] `json:"selector,required"`
}

func (r FirewallNewParamsApplyToLabelSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the Resource
type FirewallNewParamsApplyToServer struct {
	// ID of the Resource | ID of the Server
	ID param.Field[int64] `json:"id,required"`
}

func (r FirewallNewParamsApplyToServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallUpdateParams struct {
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// New Firewall name
	Name param.Field[string] `json:"name"`
}

func (r FirewallUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FirewallListParams struct {
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
	Sort param.Field[FirewallListParamsSort] `query:"sort"`
}

// URLQuery serializes [FirewallListParams]'s query parameters as `url.Values`.
func (r FirewallListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times.
type FirewallListParamsSort string

const (
	FirewallListParamsSortID          FirewallListParamsSort = "id"
	FirewallListParamsSortIDAsc       FirewallListParamsSort = "id:asc"
	FirewallListParamsSortIDDesc      FirewallListParamsSort = "id:desc"
	FirewallListParamsSortName        FirewallListParamsSort = "name"
	FirewallListParamsSortNameAsc     FirewallListParamsSort = "name:asc"
	FirewallListParamsSortNameDesc    FirewallListParamsSort = "name:desc"
	FirewallListParamsSortCreated     FirewallListParamsSort = "created"
	FirewallListParamsSortCreatedAsc  FirewallListParamsSort = "created:asc"
	FirewallListParamsSortCreatedDesc FirewallListParamsSort = "created:desc"
)
