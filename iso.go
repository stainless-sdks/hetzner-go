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

// ISOService contains methods and other services that help with interacting with
// the hetzner API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewISOService] method instead.
type ISOService struct {
	Options []option.RequestOption
}

// NewISOService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewISOService(opts ...option.RequestOption) (r *ISOService) {
	r = &ISOService{}
	r.Options = opts
	return
}

// Returns a specific ISO object.
func (r *ISOService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ISOGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("isos/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all available ISO objects.
func (r *ISOService) List(ctx context.Context, query ISOListParams, opts ...option.RequestOption) (res *ISOListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "isos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/isos/{id}
type ISOGetResponse struct {
	ISO  ISOGetResponseISO `json:"iso,required"`
	JSON isoGetResponseJSON
}

// isoGetResponseJSON contains the JSON metadata for the struct [ISOGetResponse]
type isoGetResponseJSON struct {
	ISO         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ISOGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ISOGetResponseISO struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Type of cpu architecture this iso is compatible with. Null indicates no
	// restriction on the architecture (wildcard).
	Architecture ISOGetResponseISOArchitecture `json:"architecture,required,nullable"`
	// ISO 8601 timestamp of deprecation, null if ISO is still available. After the
	// deprecation time it will no longer be possible to attach the ISO to Servers.
	Deprecated string `json:"deprecated,required,nullable"`
	// Description of the ISO
	Description string `json:"description,required"`
	// Unique identifier of the ISO. Only set for public ISOs
	Name string `json:"name,required,nullable"`
	// Type of the ISO
	Type ISOGetResponseISOType `json:"type,required"`
	JSON isoGetResponseISOJSON
}

// isoGetResponseISOJSON contains the JSON metadata for the struct
// [ISOGetResponseISO]
type isoGetResponseISOJSON struct {
	ID           apijson.Field
	Architecture apijson.Field
	Deprecated   apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ISOGetResponseISO) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cpu architecture this iso is compatible with. Null indicates no
// restriction on the architecture (wildcard).
type ISOGetResponseISOArchitecture string

const (
	ISOGetResponseISOArchitectureArm ISOGetResponseISOArchitecture = "arm"
	ISOGetResponseISOArchitectureX86 ISOGetResponseISOArchitecture = "x86"
)

// Type of the ISO
type ISOGetResponseISOType string

const (
	ISOGetResponseISOTypePrivate ISOGetResponseISOType = "private"
	ISOGetResponseISOTypePublic  ISOGetResponseISOType = "public"
)

// Response to GET https://api.hetzner.cloud/v1/isos
type ISOListResponse struct {
	ISOs []ISOListResponseISO `json:"isos,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON isoListResponseJSON
}

// isoListResponseJSON contains the JSON metadata for the struct [ISOListResponse]
type isoListResponseJSON struct {
	ISOs        apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ISOListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ISOListResponseISO struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Type of cpu architecture this iso is compatible with. Null indicates no
	// restriction on the architecture (wildcard).
	Architecture ISOListResponseISOsArchitecture `json:"architecture,required,nullable"`
	// ISO 8601 timestamp of deprecation, null if ISO is still available. After the
	// deprecation time it will no longer be possible to attach the ISO to Servers.
	Deprecated string `json:"deprecated,required,nullable"`
	// Description of the ISO
	Description string `json:"description,required"`
	// Unique identifier of the ISO. Only set for public ISOs
	Name string `json:"name,required,nullable"`
	// Type of the ISO
	Type ISOListResponseISOsType `json:"type,required"`
	JSON isoListResponseISOJSON
}

// isoListResponseISOJSON contains the JSON metadata for the struct
// [ISOListResponseISO]
type isoListResponseISOJSON struct {
	ID           apijson.Field
	Architecture apijson.Field
	Deprecated   apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ISOListResponseISO) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cpu architecture this iso is compatible with. Null indicates no
// restriction on the architecture (wildcard).
type ISOListResponseISOsArchitecture string

const (
	ISOListResponseISOsArchitectureArm ISOListResponseISOsArchitecture = "arm"
	ISOListResponseISOsArchitectureX86 ISOListResponseISOsArchitecture = "x86"
)

// Type of the ISO
type ISOListResponseISOsType string

const (
	ISOListResponseISOsTypePrivate ISOListResponseISOsType = "private"
	ISOListResponseISOsTypePublic  ISOListResponseISOsType = "public"
)

type ISOListParams struct {
	// Return only ISOs with the given architecture.
	Architecture param.Field[string] `query:"architecture"`
	// Include Images with wildcard architecture (architecture is null). Works only if
	// architecture filter is specified.
	IncludeArchitectureWildcard param.Field[bool] `query:"include_architecture_wildcard"`
	// Can be used to filter ISOs by their name. The response will only contain the ISO
	// matching the specified name.
	Name param.Field[string] `query:"name"`
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
}

// URLQuery serializes [ISOListParams]'s query parameters as `url.Values`.
func (r ISOListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
