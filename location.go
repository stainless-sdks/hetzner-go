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

// LocationService contains methods and other services that help with interacting
// with the hetzner API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewLocationService] method instead.
type LocationService struct {
	Options []option.RequestOption
}

// NewLocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLocationService(opts ...option.RequestOption) (r *LocationService) {
	r = &LocationService{}
	r.Options = opts
	return
}

// Returns a specific Location object.
func (r *LocationService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *LocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("locations/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all Location objects.
func (r *LocationService) List(ctx context.Context, query LocationListParams, opts ...option.RequestOption) (res *LocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/locations/{id}
type LocationGetResponse struct {
	// Location the Floating IP was created in. Routing is optimized for this Location.
	// | Location of the Volume. Volume can only be attached to Servers in the same
	// Location.
	Location LocationGetResponseLocation `json:"location,required"`
	JSON     locationGetResponseJSON
}

// locationGetResponseJSON contains the JSON metadata for the struct
// [LocationGetResponse]
type locationGetResponseJSON struct {
	Location    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type LocationGetResponseLocation struct {
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
	JSON        locationGetResponseLocationJSON
}

// locationGetResponseLocationJSON contains the JSON metadata for the struct
// [LocationGetResponseLocation]
type locationGetResponseLocationJSON struct {
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

func (r *LocationGetResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/locations
type LocationListResponse struct {
	Locations []LocationListResponseLocation `json:"locations,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta,required"`
	JSON locationListResponseJSON
}

// locationListResponseJSON contains the JSON metadata for the struct
// [LocationListResponse]
type locationListResponseJSON struct {
	Locations   apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Location the Floating IP was created in. Routing is optimized for this Location.
// | Location of the Volume. Volume can only be attached to Servers in the same
// Location.
type LocationListResponseLocation struct {
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
	JSON        locationListResponseLocationJSON
}

// locationListResponseLocationJSON contains the JSON metadata for the struct
// [LocationListResponseLocation]
type locationListResponseLocationJSON struct {
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

func (r *LocationListResponseLocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LocationListParams struct {
	// Can be used to filter Locations by their name. The response will only contain
	// the Location matching the specified name.
	Name param.Field[string] `query:"name"`
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
	// Can be used multiple times.
	Sort param.Field[LocationListParamsSort] `query:"sort"`
}

// URLQuery serializes [LocationListParams]'s query parameters as `url.Values`.
func (r LocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times.
type LocationListParamsSort string

const (
	LocationListParamsSortID       LocationListParamsSort = "id"
	LocationListParamsSortIDAsc    LocationListParamsSort = "id:asc"
	LocationListParamsSortIDDesc   LocationListParamsSort = "id:desc"
	LocationListParamsSortName     LocationListParamsSort = "name"
	LocationListParamsSortNameAsc  LocationListParamsSort = "name:asc"
	LocationListParamsSortNameDesc LocationListParamsSort = "name:desc"
)
