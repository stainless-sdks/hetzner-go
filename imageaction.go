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

// ImageActionService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewImageActionService] method
// instead.
type ImageActionService struct {
	Options []option.RequestOption
}

// NewImageActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewImageActionService(opts ...option.RequestOption) (r *ImageActionService) {
	r = &ImageActionService{}
	r.Options = opts
	return
}

// Returns a specific Action for an Image.
func (r *ImageActionService) Get(ctx context.Context, id int64, actionID int64, opts ...option.RequestOption) (res *ImageActionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("images/%v/actions/%v", id, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all Action objects for an Image. You can sort the results by using the
// `sort` URI parameter, and filter them with the `status` parameter.
func (r *ImageActionService) List(ctx context.Context, id int64, query ImageActionListParams, opts ...option.RequestOption) (res *ImageActionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("images/%v/actions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Changes the protection configuration of the Image. Can only be used on
// snapshots.
func (r *ImageActionService) ChangeProtection(ctx context.Context, id int64, body ImageActionChangeProtectionParams, opts ...option.RequestOption) (res *ImageActionChangeProtectionResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("images/%v/actions/change_protection", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/images/{id}/actions/{action_id}
type ImageActionGetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   imageActionGetResponseJSON
}

// imageActionGetResponseJSON contains the JSON metadata for the struct
// [ImageActionGetResponse]
type imageActionGetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/images/{id}/actions
type ImageActionListResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON imageActionListResponseJSON
}

// imageActionListResponseJSON contains the JSON metadata for the struct
// [ImageActionListResponse]
type imageActionListResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST
// https://api.hetzner.cloud/v1/images/{id}/actions/change_protection
type ImageActionChangeProtectionResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   imageActionChangeProtectionResponseJSON
}

// imageActionChangeProtectionResponseJSON contains the JSON metadata for the
// struct [ImageActionChangeProtectionResponse]
type imageActionChangeProtectionResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ImageActionChangeProtectionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ImageActionListParams struct {
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

// URLQuery serializes [ImageActionListParams]'s query parameters as `url.Values`.
func (r ImageActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ImageActionChangeProtectionParams struct {
	// If true, prevents the snapshot from being deleted
	Delete param.Field[bool] `json:"delete"`
}

func (r ImageActionChangeProtectionParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
