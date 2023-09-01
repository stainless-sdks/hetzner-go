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

// VolumeActionService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewVolumeActionService] method
// instead.
type VolumeActionService struct {
	Options []option.RequestOption
}

// NewVolumeActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVolumeActionService(opts ...option.RequestOption) (r *VolumeActionService) {
	r = &VolumeActionService{}
	r.Options = opts
	return
}

// Returns a specific Action object.
func (r *VolumeActionService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *VolumeActionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("volumes/actions/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all Action objects for a Volume. You can `sort` the results by using the
// sort URI parameter, and filter them with the `status` parameter.
func (r *VolumeActionService) List(ctx context.Context, id int64, query VolumeActionListParams, opts ...option.RequestOption) (res *VolumeActionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("volumes/%v/actions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Attaches a Volume to a Server. Works only if the Server is in the same Location
// as the Volume.
func (r *VolumeActionService) Attach(ctx context.Context, id int64, body VolumeActionAttachParams, opts ...option.RequestOption) (res *VolumeActionAttachResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("volumes/%v/actions/attach", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Changes the protection configuration of a Volume.
func (r *VolumeActionService) ChangeProtection(ctx context.Context, id int64, body VolumeActionChangeProtectionParams, opts ...option.RequestOption) (res *VolumeActionChangeProtectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("volumes/%v/actions/change_protection", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Detaches a Volume from the Server it’s attached to. You may attach it to a
// Server again at a later time.
func (r *VolumeActionService) Detach(ctx context.Context, id int64, opts ...option.RequestOption) (res *VolumeActionDetachResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("volumes/%v/actions/detach", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Changes the size of a Volume. Note that downsizing a Volume is not possible.
func (r *VolumeActionService) Resize(ctx context.Context, id int64, body VolumeActionResizeParams, opts ...option.RequestOption) (res *VolumeActionResizeResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("volumes/%v/actions/resize", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/{resource}/actions
type VolumeActionGetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   volumeActionGetResponseJSON
}

// volumeActionGetResponseJSON contains the JSON metadata for the struct
// [VolumeActionGetResponse]
type volumeActionGetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/volumes/{id}/actions
type VolumeActionListResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON volumeActionListResponseJSON
}

// volumeActionListResponseJSON contains the JSON metadata for the struct
// [VolumeActionListResponse]
type volumeActionListResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/volumes/{id}/actions/attach
type VolumeActionAttachResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   volumeActionAttachResponseJSON
}

// volumeActionAttachResponseJSON contains the JSON metadata for the struct
// [VolumeActionAttachResponse]
type volumeActionAttachResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeActionAttachResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/volumes/{id}/actions/change_protection
type VolumeActionChangeProtectionResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   volumeActionChangeProtectionResponseJSON
}

// volumeActionChangeProtectionResponseJSON contains the JSON metadata for the
// struct [VolumeActionChangeProtectionResponse]
type volumeActionChangeProtectionResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeActionChangeProtectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/volumes/{id}/actions/detach
type VolumeActionDetachResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   volumeActionDetachResponseJSON
}

// volumeActionDetachResponseJSON contains the JSON metadata for the struct
// [VolumeActionDetachResponse]
type volumeActionDetachResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeActionDetachResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/volumes/{id}/actions/resize
type VolumeActionResizeResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   volumeActionResizeResponseJSON
}

// volumeActionResizeResponseJSON contains the JSON metadata for the struct
// [VolumeActionResizeResponse]
type volumeActionResizeResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VolumeActionResizeResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeActionListParams struct {
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

// URLQuery serializes [VolumeActionListParams]'s query parameters as `url.Values`.
func (r VolumeActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VolumeActionAttachParams struct {
	// ID of the Server the Volume will be attached to
	Server param.Field[int64] `json:"server,required"`
	// Auto-mount the Volume after attaching it
	Automount param.Field[bool] `json:"automount"`
}

func (r VolumeActionAttachParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VolumeActionChangeProtectionParams struct {
	// If true, prevents the Volume from being deleted
	Delete param.Field[bool] `json:"delete"`
}

func (r VolumeActionChangeProtectionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type VolumeActionResizeParams struct {
	// New Volume size in GB (must be greater than current size)
	Size param.Field[float64] `json:"size,required"`
}

func (r VolumeActionResizeParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
