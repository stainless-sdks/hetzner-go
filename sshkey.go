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

// SshKeyService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewSshKeyService] method instead.
type SshKeyService struct {
	Options []option.RequestOption
}

// NewSshKeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSshKeyService(opts ...option.RequestOption) (r *SshKeyService) {
	r = &SshKeyService{}
	r.Options = opts
	return
}

// Creates a new SSH key with the given `name` and `public_key`. Once an SSH key is
// created, it can be used in other calls such as creating Servers.
func (r *SshKeyService) New(ctx context.Context, body SshKeyNewParams, opts ...option.RequestOption) (res *SshKeyNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ssh_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a specific SSH key object.
func (r *SshKeyService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *SshKeyGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ssh_keys/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an SSH key. You can update an SSH key name and an SSH key labels.
//
// Please note that when updating labels, the SSH key current set of labels will be
// replaced with the labels provided in the request body. So, for example, if you
// want to add a new label, you have to provide all existing labels plus the new
// label in the request body.
func (r *SshKeyService) Update(ctx context.Context, id int64, body SshKeyUpdateParams, opts ...option.RequestOption) (res *SshKeyUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("ssh_keys/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all SSH key objects.
func (r *SshKeyService) List(ctx context.Context, query SshKeyListParams, opts ...option.RequestOption) (res *SshKeyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "ssh_keys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes an SSH key. It cannot be used anymore.
func (r *SshKeyService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("ssh_keys/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Response to POST https://api.hetzner.cloud/v1/ssh_keys
type SshKeyNewResponse struct {
	// SSH keys are public keys you provide to the cloud system. They can be injected
	// into Servers at creation time. We highly recommend that you use keys instead of
	// passwords to manage your Servers.
	SshKey SshKeyNewResponseSshKey `json:"ssh_key,required"`
	JSON   sshKeyNewResponseJSON
}

// sshKeyNewResponseJSON contains the JSON metadata for the struct
// [SshKeyNewResponse]
type sshKeyNewResponseJSON struct {
	SshKey      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SshKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSH keys are public keys you provide to the cloud system. They can be injected
// into Servers at creation time. We highly recommend that you use keys instead of
// passwords to manage your Servers.
type SshKeyNewResponseSshKey struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Fingerprint of public key
	Fingerprint string `json:"fingerprint,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Public key
	PublicKey string `json:"public_key,required"`
	JSON      sshKeyNewResponseSshKeyJSON
}

// sshKeyNewResponseSshKeyJSON contains the JSON metadata for the struct
// [SshKeyNewResponseSshKey]
type sshKeyNewResponseSshKeyJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Fingerprint apijson.Field
	Labels      apijson.Field
	Name        apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SshKeyNewResponseSshKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/ssh_keys/{id}
type SshKeyGetResponse struct {
	// SSH keys are public keys you provide to the cloud system. They can be injected
	// into Servers at creation time. We highly recommend that you use keys instead of
	// passwords to manage your Servers.
	SshKey SshKeyGetResponseSshKey `json:"ssh_key,required"`
	JSON   sshKeyGetResponseJSON
}

// sshKeyGetResponseJSON contains the JSON metadata for the struct
// [SshKeyGetResponse]
type sshKeyGetResponseJSON struct {
	SshKey      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SshKeyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSH keys are public keys you provide to the cloud system. They can be injected
// into Servers at creation time. We highly recommend that you use keys instead of
// passwords to manage your Servers.
type SshKeyGetResponseSshKey struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Fingerprint of public key
	Fingerprint string `json:"fingerprint,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Public key
	PublicKey string `json:"public_key,required"`
	JSON      sshKeyGetResponseSshKeyJSON
}

// sshKeyGetResponseSshKeyJSON contains the JSON metadata for the struct
// [SshKeyGetResponseSshKey]
type sshKeyGetResponseSshKeyJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Fingerprint apijson.Field
	Labels      apijson.Field
	Name        apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SshKeyGetResponseSshKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to PUT https://api.hetzner.cloud/v1/ssh_keys/{id}
type SshKeyUpdateResponse struct {
	// SSH keys are public keys you provide to the cloud system. They can be injected
	// into Servers at creation time. We highly recommend that you use keys instead of
	// passwords to manage your Servers.
	SshKey SshKeyUpdateResponseSshKey `json:"ssh_key,required"`
	JSON   sshKeyUpdateResponseJSON
}

// sshKeyUpdateResponseJSON contains the JSON metadata for the struct
// [SshKeyUpdateResponse]
type sshKeyUpdateResponseJSON struct {
	SshKey      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SshKeyUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSH keys are public keys you provide to the cloud system. They can be injected
// into Servers at creation time. We highly recommend that you use keys instead of
// passwords to manage your Servers.
type SshKeyUpdateResponseSshKey struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Fingerprint of public key
	Fingerprint string `json:"fingerprint,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Public key
	PublicKey string `json:"public_key,required"`
	JSON      sshKeyUpdateResponseSshKeyJSON
}

// sshKeyUpdateResponseSshKeyJSON contains the JSON metadata for the struct
// [SshKeyUpdateResponseSshKey]
type sshKeyUpdateResponseSshKeyJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Fingerprint apijson.Field
	Labels      apijson.Field
	Name        apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SshKeyUpdateResponseSshKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/ssh_keys
type SshKeyListResponse struct {
	SshKeys []SshKeyListResponseSshKey `json:"ssh_keys,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON sshKeyListResponseJSON
}

// sshKeyListResponseJSON contains the JSON metadata for the struct
// [SshKeyListResponse]
type sshKeyListResponseJSON struct {
	SshKeys     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SshKeyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// SSH keys are public keys you provide to the cloud system. They can be injected
// into Servers at creation time. We highly recommend that you use keys instead of
// passwords to manage your Servers.
type SshKeyListResponseSshKey struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Fingerprint of public key
	Fingerprint string `json:"fingerprint,required"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Public key
	PublicKey string `json:"public_key,required"`
	JSON      sshKeyListResponseSshKeyJSON
}

// sshKeyListResponseSshKeyJSON contains the JSON metadata for the struct
// [SshKeyListResponseSshKey]
type sshKeyListResponseSshKeyJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Fingerprint apijson.Field
	Labels      apijson.Field
	Name        apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SshKeyListResponseSshKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SshKeyNewParams struct {
	// Name of the SSH key
	Name param.Field[string] `json:"name,required"`
	// Public key
	PublicKey param.Field[string] `json:"public_key,required"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
}

func (r SshKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SshKeyUpdateParams struct {
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// New name Name to set
	Name param.Field[string] `json:"name"`
}

func (r SshKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SshKeyListParams struct {
	// Can be used to filter SSH keys by their fingerprint. The response will only
	// contain the SSH key matching the specified fingerprint.
	Fingerprint param.Field[string] `query:"fingerprint"`
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
	Sort param.Field[SshKeyListParamsSort] `query:"sort"`
}

// URLQuery serializes [SshKeyListParams]'s query parameters as `url.Values`.
func (r SshKeyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times.
type SshKeyListParamsSort string

const (
	SshKeyListParamsSortID       SshKeyListParamsSort = "id"
	SshKeyListParamsSortIDAsc    SshKeyListParamsSort = "id:asc"
	SshKeyListParamsSortIDDesc   SshKeyListParamsSort = "id:desc"
	SshKeyListParamsSortName     SshKeyListParamsSort = "name"
	SshKeyListParamsSortNameAsc  SshKeyListParamsSort = "name:asc"
	SshKeyListParamsSortNameDesc SshKeyListParamsSort = "name:desc"
)
