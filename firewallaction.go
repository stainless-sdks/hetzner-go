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

// FirewallActionService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewFirewallActionService] method
// instead.
type FirewallActionService struct {
	Options []option.RequestOption
}

// NewFirewallActionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFirewallActionService(opts ...option.RequestOption) (r *FirewallActionService) {
	r = &FirewallActionService{}
	r.Options = opts
	return
}

// Returns a specific Action for a Firewall.
func (r *FirewallActionService) Get(ctx context.Context, id int64, actionID int64, opts ...option.RequestOption) (res *FirewallActionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("firewalls/%v/actions/%v", id, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all Action objects for a Firewall. You can sort the results by using the
// `sort` URI parameter, and filter them with the `status` parameter.
func (r *FirewallActionService) List(ctx context.Context, id int64, query FirewallActionListParams, opts ...option.RequestOption) (res *FirewallActionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("firewalls/%v/actions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Applies one Firewall to multiple resources.
//
// Currently servers (public network interface) and label selectors are supported.
//
// #### Call specific error codes
//
// | Code                          | Description                                                   |
// | ----------------------------- | ------------------------------------------------------------- |
// | `firewall_already_applied`    | Firewall was already applied on resource                      |
// | `incompatible_network_type`   | The Network type is incompatible for the given resource       |
// | `firewall_resource_not_found` | The resource the Firewall should be attached to was not found |
func (r *FirewallActionService) ApplyToResources(ctx context.Context, id int64, body FirewallActionApplyToResourcesParams, opts ...option.RequestOption) (res *FirewallActionApplyToResourcesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("firewalls/%v/actions/apply_to_resources", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes one Firewall from multiple resources.
//
// Currently only Servers (and their public network interfaces) are supported.
//
// #### Call specific error codes
//
// | Code                                 | Description                                                            |
// | ------------------------------------ | ---------------------------------------------------------------------- |
// | `firewall_already_removed`           | Firewall was already removed from the resource                         |
// | `firewall_resource_not_found`        | The resource the Firewall should be attached to was not found          |
// | `firewall_managed_by_label_selector` | Firewall was applied via label selector and cannot be removed manually |
func (r *FirewallActionService) RemoveFromResources(ctx context.Context, id int64, body FirewallActionRemoveFromResourcesParams, opts ...option.RequestOption) (res *FirewallActionRemoveFromResourcesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("firewalls/%v/actions/remove_from_resources", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Sets the rules of a Firewall.
//
// All existing rules will be overwritten. Pass an empty `rules` array to remove
// all rules. The maximum amount of rules that can be defined is 50.
//
// #### Call specific error codes
//
// | Code                          | Description                                                   |
// | ----------------------------- | ------------------------------------------------------------- |
// | `firewall_resource_not_found` | The resource the Firewall should be attached to was not found |
func (r *FirewallActionService) SetRules(ctx context.Context, id int64, body FirewallActionSetRulesParams, opts ...option.RequestOption) (res *FirewallActionSetRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("firewalls/%v/actions/set_rules", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/firewalls/{id}/actions/{action_id}
type FirewallActionGetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   firewallActionGetResponseJSON
}

// firewallActionGetResponseJSON contains the JSON metadata for the struct
// [FirewallActionGetResponse]
type firewallActionGetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/firewalls/{id}/actions
type FirewallActionListResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON firewallActionListResponseJSON
}

// firewallActionListResponseJSON contains the JSON metadata for the struct
// [FirewallActionListResponse]
type firewallActionListResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/firewalls/{id}/actions/apply_to_resources
type FirewallActionApplyToResourcesResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON firewallActionApplyToResourcesResponseJSON
}

// firewallActionApplyToResourcesResponseJSON contains the JSON metadata for the
// struct [FirewallActionApplyToResourcesResponse]
type firewallActionApplyToResourcesResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallActionApplyToResourcesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/firewalls/{id}/actions/remove_from_resources
type FirewallActionRemoveFromResourcesResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON firewallActionRemoveFromResourcesResponseJSON
}

// firewallActionRemoveFromResourcesResponseJSON contains the JSON metadata for the
// struct [FirewallActionRemoveFromResourcesResponse]
type firewallActionRemoveFromResourcesResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallActionRemoveFromResourcesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/firewalls/{id}/actions/set_rules
type FirewallActionSetRulesResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON firewallActionSetRulesResponseJSON
}

// firewallActionSetRulesResponseJSON contains the JSON metadata for the struct
// [FirewallActionSetRulesResponse]
type firewallActionSetRulesResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallActionSetRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type FirewallActionListParams struct {
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

// URLQuery serializes [FirewallActionListParams]'s query parameters as
// `url.Values`.
func (r FirewallActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FirewallActionApplyToResourcesParams struct {
	// Resources the Firewall should be applied to
	ApplyTo param.Field[[]FirewallActionApplyToResourcesParamsApplyTo] `json:"apply_to,required"`
}

func (r FirewallActionApplyToResourcesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Resource a Firewall should be applied to.
type FirewallActionApplyToResourcesParamsApplyTo struct {
	// Configuration for type LabelSelector, required if type is `label_selector`
	LabelSelector param.Field[FirewallActionApplyToResourcesParamsApplyToLabelSelector] `json:"label_selector"`
	// ID of the Resource
	Server param.Field[FirewallActionApplyToResourcesParamsApplyToServer] `json:"server"`
	// Type of the resource
	Type param.Field[FirewallActionApplyToResourcesParamsApplyToType] `json:"type"`
}

func (r FirewallActionApplyToResourcesParamsApplyTo) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for type LabelSelector, required if type is `label_selector`
type FirewallActionApplyToResourcesParamsApplyToLabelSelector struct {
	// Label selector
	Selector param.Field[string] `json:"selector,required"`
}

func (r FirewallActionApplyToResourcesParamsApplyToLabelSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the Resource
type FirewallActionApplyToResourcesParamsApplyToServer struct {
	// ID of the Resource | ID of the Server
	ID param.Field[int64] `json:"id,required"`
}

func (r FirewallActionApplyToResourcesParamsApplyToServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the resource
type FirewallActionApplyToResourcesParamsApplyToType string

const (
	FirewallActionApplyToResourcesParamsApplyToTypeLabelSelector FirewallActionApplyToResourcesParamsApplyToType = "label_selector"
	FirewallActionApplyToResourcesParamsApplyToTypeServer        FirewallActionApplyToResourcesParamsApplyToType = "server"
)

type FirewallActionRemoveFromResourcesParams struct {
	// Resources the Firewall should be removed from
	RemoveFrom param.Field[[]FirewallActionRemoveFromResourcesParamsRemoveFrom] `json:"remove_from,required"`
}

func (r FirewallActionRemoveFromResourcesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Resource a Firewall should be applied to.
type FirewallActionRemoveFromResourcesParamsRemoveFrom struct {
	// Configuration for type LabelSelector, required if type is `label_selector`
	LabelSelector param.Field[FirewallActionRemoveFromResourcesParamsRemoveFromLabelSelector] `json:"label_selector"`
	// ID of the Resource
	Server param.Field[FirewallActionRemoveFromResourcesParamsRemoveFromServer] `json:"server"`
	// Type of the resource
	Type param.Field[FirewallActionRemoveFromResourcesParamsRemoveFromType] `json:"type"`
}

func (r FirewallActionRemoveFromResourcesParamsRemoveFrom) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for type LabelSelector, required if type is `label_selector`
type FirewallActionRemoveFromResourcesParamsRemoveFromLabelSelector struct {
	// Label selector
	Selector param.Field[string] `json:"selector,required"`
}

func (r FirewallActionRemoveFromResourcesParamsRemoveFromLabelSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// ID of the Resource
type FirewallActionRemoveFromResourcesParamsRemoveFromServer struct {
	// ID of the Resource | ID of the Server
	ID param.Field[int64] `json:"id,required"`
}

func (r FirewallActionRemoveFromResourcesParamsRemoveFromServer) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Type of the resource
type FirewallActionRemoveFromResourcesParamsRemoveFromType string

const (
	FirewallActionRemoveFromResourcesParamsRemoveFromTypeLabelSelector FirewallActionRemoveFromResourcesParamsRemoveFromType = "label_selector"
	FirewallActionRemoveFromResourcesParamsRemoveFromTypeServer        FirewallActionRemoveFromResourcesParamsRemoveFromType = "server"
)

type FirewallActionSetRulesParams struct {
	// Array of rules
	Rules param.Field[[]RuleParam] `json:"rules,required"`
}

func (r FirewallActionSetRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
