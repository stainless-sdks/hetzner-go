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

// IsoService contains methods and other services that help with interacting with
// the hetzner API. Note, unlike clients, this service does not read variables from
// the environment automatically. You should not instantiate this service directly,
// and instead use the [NewIsoService] method instead.
type IsoService struct {
	Options []option.RequestOption
}

// NewIsoService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewIsoService(opts ...option.RequestOption) (r *IsoService) {
	r = &IsoService{}
	r.Options = opts
	return
}

// Returns a specific ISO object.
func (r *IsoService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *IsoGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("isos/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all available ISO objects.
func (r *IsoService) List(ctx context.Context, query IsoListParams, opts ...option.RequestOption) (res *IsoListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "isos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/isos/{id}
type IsoGetResponse struct {
	Iso  IsoGetResponseIso `json:"iso,required"`
	JSON isoGetResponseJSON
}

// isoGetResponseJSON contains the JSON metadata for the struct [IsoGetResponse]
type isoGetResponseJSON struct {
	Iso         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IsoGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IsoGetResponseIso struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Type of cpu architecture this iso is compatible with. Null indicates no
	// restriction on the architecture (wildcard).
	Architecture IsoGetResponseIsoArchitecture `json:"architecture,required,nullable"`
	// ISO 8601 timestamp of deprecation, null if ISO is still available. After the
	// deprecation time it will no longer be possible to attach the ISO to Servers.
	Deprecated string `json:"deprecated,required,nullable"`
	// Description of the ISO
	Description string `json:"description,required"`
	// Unique identifier of the ISO. Only set for public ISOs
	Name string `json:"name,required,nullable"`
	// Type of the ISO
	Type IsoGetResponseIsoType `json:"type,required"`
	JSON isoGetResponseIsoJSON
}

// isoGetResponseIsoJSON contains the JSON metadata for the struct
// [IsoGetResponseIso]
type isoGetResponseIsoJSON struct {
	ID           apijson.Field
	Architecture apijson.Field
	Deprecated   apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IsoGetResponseIso) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cpu architecture this iso is compatible with. Null indicates no
// restriction on the architecture (wildcard).
type IsoGetResponseIsoArchitecture string

const (
	IsoGetResponseIsoArchitectureArm IsoGetResponseIsoArchitecture = "arm"
	IsoGetResponseIsoArchitectureX86 IsoGetResponseIsoArchitecture = "x86"
)

// Type of the ISO
type IsoGetResponseIsoType string

const (
	IsoGetResponseIsoTypePrivate IsoGetResponseIsoType = "private"
	IsoGetResponseIsoTypePublic  IsoGetResponseIsoType = "public"
)

// Response to GET https://api.hetzner.cloud/v1/isos
type IsoListResponse struct {
	Isos []IsoListResponseIso `json:"isos,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON isoListResponseJSON
}

// isoListResponseJSON contains the JSON metadata for the struct [IsoListResponse]
type isoListResponseJSON struct {
	Isos        apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IsoListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IsoListResponseIso struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Type of cpu architecture this iso is compatible with. Null indicates no
	// restriction on the architecture (wildcard).
	Architecture IsoListResponseIsosArchitecture `json:"architecture,required,nullable"`
	// ISO 8601 timestamp of deprecation, null if ISO is still available. After the
	// deprecation time it will no longer be possible to attach the ISO to Servers.
	Deprecated string `json:"deprecated,required,nullable"`
	// Description of the ISO
	Description string `json:"description,required"`
	// Unique identifier of the ISO. Only set for public ISOs
	Name string `json:"name,required,nullable"`
	// Type of the ISO
	Type IsoListResponseIsosType `json:"type,required"`
	JSON isoListResponseIsoJSON
}

// isoListResponseIsoJSON contains the JSON metadata for the struct
// [IsoListResponseIso]
type isoListResponseIsoJSON struct {
	ID           apijson.Field
	Architecture apijson.Field
	Deprecated   apijson.Field
	Description  apijson.Field
	Name         apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *IsoListResponseIso) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Type of cpu architecture this iso is compatible with. Null indicates no
// restriction on the architecture (wildcard).
type IsoListResponseIsosArchitecture string

const (
	IsoListResponseIsosArchitectureArm IsoListResponseIsosArchitecture = "arm"
	IsoListResponseIsosArchitectureX86 IsoListResponseIsosArchitecture = "x86"
)

// Type of the ISO
type IsoListResponseIsosType string

const (
	IsoListResponseIsosTypePrivate IsoListResponseIsosType = "private"
	IsoListResponseIsosTypePublic  IsoListResponseIsosType = "public"
)

type IsoListParams struct {
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

// URLQuery serializes [IsoListParams]'s query parameters as `url.Values`.
func (r IsoListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
