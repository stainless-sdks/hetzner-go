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

// DatacenterService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewDatacenterService] method instead.
type DatacenterService struct {
	Options []option.RequestOption
}

// NewDatacenterService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDatacenterService(opts ...option.RequestOption) (r *DatacenterService) {
	r = &DatacenterService{}
	r.Options = opts
	return
}

// Returns a specific Datacenter object.
func (r *DatacenterService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *DatacenterGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("datacenters/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all Datacenter objects.
func (r *DatacenterService) List(ctx context.Context, query DatacenterListParams, opts ...option.RequestOption) (res *DatacenterListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "datacenters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/datacenters/{id}
type DatacenterGetResponse struct {
	// Datacenter this Primary IP is located at | Datacenter this Resource is located
	// at
	Datacenter DatacenterGetResponseDatacenter `json:"datacenter,required"`
	JSON       datacenterGetResponseJSON
}

// datacenterGetResponseJSON contains the JSON metadata for the struct
// [DatacenterGetResponse]
type datacenterGetResponseJSON struct {
	Datacenter  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatacenterGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Datacenter this Primary IP is located at | Datacenter this Resource is located
// at
type DatacenterGetResponseDatacenter struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Description of the Datacenter
	Description string `json:"description,required"`
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	Location DatacenterGetResponseDatacenterLocation `json:"location,required"`
	// Unique identifier of the Datacenter
	Name string `json:"name,required"`
	// The Server types the Datacenter can handle
	ServerTypes DatacenterGetResponseDatacenterServerTypes `json:"server_types,required"`
	JSON        datacenterGetResponseDatacenterJSON
}

// datacenterGetResponseDatacenterJSON contains the JSON metadata for the struct
// [DatacenterGetResponseDatacenter]
type datacenterGetResponseDatacenterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	ServerTypes apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatacenterGetResponseDatacenter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type DatacenterGetResponseDatacenterLocation struct {
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
	JSON        datacenterGetResponseDatacenterLocationJSON
}

// datacenterGetResponseDatacenterLocationJSON contains the JSON metadata for the
// struct [DatacenterGetResponseDatacenterLocation]
type datacenterGetResponseDatacenterLocationJSON struct {
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

func (r *DatacenterGetResponseDatacenterLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Server types the Datacenter can handle
type DatacenterGetResponseDatacenterServerTypes struct {
	// IDs of Server types that are supported and for which the Datacenter has enough
	// resources left
	Available []int64 `json:"available,required"`
	// IDs of Server types that are supported and for which the Datacenter has enough
	// resources left
	AvailableForMigration []int64 `json:"available_for_migration,required"`
	// IDs of Server types that are supported in the Datacenter
	Supported []int64 `json:"supported,required"`
	JSON      datacenterGetResponseDatacenterServerTypesJSON
}

// datacenterGetResponseDatacenterServerTypesJSON contains the JSON metadata for
// the struct [DatacenterGetResponseDatacenterServerTypes]
type datacenterGetResponseDatacenterServerTypesJSON struct {
	Available             apijson.Field
	AvailableForMigration apijson.Field
	Supported             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DatacenterGetResponseDatacenterServerTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/datacenters
type DatacenterListResponse struct {
	Datacenters []DatacenterListResponseDatacenter `json:"datacenters,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta,required"`
	// The Datacenter which is recommended to be used to create new Servers.
	Recommendation int64 `json:"recommendation,required"`
	JSON           datacenterListResponseJSON
}

// datacenterListResponseJSON contains the JSON metadata for the struct
// [DatacenterListResponse]
type datacenterListResponseJSON struct {
	Datacenters    apijson.Field
	Meta           apijson.Field
	Recommendation apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *DatacenterListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Datacenter this Primary IP is located at | Datacenter this Resource is located
// at
type DatacenterListResponseDatacenter struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Description of the Datacenter
	Description string `json:"description,required"`
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	Location DatacenterListResponseDatacentersLocation `json:"location,required"`
	// Unique identifier of the Datacenter
	Name string `json:"name,required"`
	// The Server types the Datacenter can handle
	ServerTypes DatacenterListResponseDatacentersServerTypes `json:"server_types,required"`
	JSON        datacenterListResponseDatacenterJSON
}

// datacenterListResponseDatacenterJSON contains the JSON metadata for the struct
// [DatacenterListResponseDatacenter]
type datacenterListResponseDatacenterJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Location    apijson.Field
	Name        apijson.Field
	ServerTypes apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatacenterListResponseDatacenter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type DatacenterListResponseDatacentersLocation struct {
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
	JSON        datacenterListResponseDatacentersLocationJSON
}

// datacenterListResponseDatacentersLocationJSON contains the JSON metadata for the
// struct [DatacenterListResponseDatacentersLocation]
type datacenterListResponseDatacentersLocationJSON struct {
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

func (r *DatacenterListResponseDatacentersLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The Server types the Datacenter can handle
type DatacenterListResponseDatacentersServerTypes struct {
	// IDs of Server types that are supported and for which the Datacenter has enough
	// resources left
	Available []int64 `json:"available,required"`
	// IDs of Server types that are supported and for which the Datacenter has enough
	// resources left
	AvailableForMigration []int64 `json:"available_for_migration,required"`
	// IDs of Server types that are supported in the Datacenter
	Supported []int64 `json:"supported,required"`
	JSON      datacenterListResponseDatacentersServerTypesJSON
}

// datacenterListResponseDatacentersServerTypesJSON contains the JSON metadata for
// the struct [DatacenterListResponseDatacentersServerTypes]
type datacenterListResponseDatacentersServerTypesJSON struct {
	Available             apijson.Field
	AvailableForMigration apijson.Field
	Supported             apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *DatacenterListResponseDatacentersServerTypes) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DatacenterListParams struct {
	// Can be used to filter Datacenters by their name. The response will only contain
	// the Datacenter matching the specified name. When the name does not match the
	// Datacenter name format, an `invalid_input` error is returned.
	Name param.Field[string] `query:"name"`
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
	// Can be used multiple times.
	Sort param.Field[DatacenterListParamsSort] `query:"sort"`
}

// URLQuery serializes [DatacenterListParams]'s query parameters as `url.Values`.
func (r DatacenterListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times.
type DatacenterListParamsSort string

const (
	DatacenterListParamsSortID       DatacenterListParamsSort = "id"
	DatacenterListParamsSortIDAsc    DatacenterListParamsSort = "id:asc"
	DatacenterListParamsSortIDDesc   DatacenterListParamsSort = "id:desc"
	DatacenterListParamsSortName     DatacenterListParamsSort = "name"
	DatacenterListParamsSortNameAsc  DatacenterListParamsSort = "name:asc"
	DatacenterListParamsSortNameDesc DatacenterListParamsSort = "name:desc"
)
