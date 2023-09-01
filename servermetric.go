// File generated from our OpenAPI spec by Stainless.

package hetzner

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/hetzner/hetzner-go/internal/apijson"
	"github.com/hetzner/hetzner-go/internal/apiquery"
	"github.com/hetzner/hetzner-go/internal/param"
	"github.com/hetzner/hetzner-go/internal/requestconfig"
	"github.com/hetzner/hetzner-go/internal/shared"
	"github.com/hetzner/hetzner-go/option"
	"github.com/tidwall/gjson"
)

// ServerMetricService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewServerMetricService] method
// instead.
type ServerMetricService struct {
	Options []option.RequestOption
}

// NewServerMetricService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewServerMetricService(opts ...option.RequestOption) (r *ServerMetricService) {
	r = &ServerMetricService{}
	r.Options = opts
	return
}

// Get Metrics for specified Server.
//
// You must specify the type of metric to get: cpu, disk or network. You can also
// specify more than one type by comma separation, e.g. cpu,disk.
//
// # Depending on the type you will get different time series data
//
// | Type    | Timeseries              | Unit      | Description                                          |
// | ------- | ----------------------- | --------- | ---------------------------------------------------- |
// | cpu     | cpu                     | percent   | Percent CPU usage                                    |
// | disk    | disk.0.iops.read        | iop/s     | Number of read IO operations per second              |
// |         | disk.0.iops.write       | iop/s     | Number of write IO operations per second             |
// |         | disk.0.bandwidth.read   | bytes/s   | Bytes read per second                                |
// |         | disk.0.bandwidth.write  | bytes/s   | Bytes written per second                             |
// | network | network.0.pps.in        | packets/s | Public Network interface packets per second received |
// |         | network.0.pps.out       | packets/s | Public Network interface packets per second sent     |
// |         | network.0.bandwidth.in  | bytes/s   | Public Network interface bytes/s received            |
// |         | network.0.bandwidth.out | bytes/s   | Public Network interface bytes/s sent                |
//
// Metrics are available for the last 30 days only.
//
// If you do not provide the step argument we will automatically adjust it so that
// a maximum of 200 samples are returned.
//
// We limit the number of samples returned to a maximum of 500 and will adjust the
// step parameter accordingly.
func (r *ServerMetricService) List(ctx context.Context, id int64, query ServerMetricListParams, opts ...option.RequestOption) (res *ServerMetricListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("servers/%v/metrics", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/servers/{id}/metrics
type ServerMetricListResponse struct {
	// You must specify the type of metric to get: open_connections,
	// requests_per_second or bandwidth. You can also specify more than one type by
	// comma separation, e.g. requests_per_second,bandwidth. Depending on the type you
	// will get different time series data.
	Metrics ServerMetricListResponseMetrics `json:"metrics,required"`
	JSON    serverMetricListResponseJSON
}

// serverMetricListResponseJSON contains the JSON metadata for the struct
// [ServerMetricListResponse]
type serverMetricListResponseJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerMetricListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// You must specify the type of metric to get: open_connections,
// requests_per_second or bandwidth. You can also specify more than one type by
// comma separation, e.g. requests_per_second,bandwidth. Depending on the type you
// will get different time series data.
type ServerMetricListResponseMetrics struct {
	// End of period of metrics reported (in ISO-8601 format)
	End string `json:"end,required"`
	// Start of period of metrics reported (in ISO-8601 format)
	Start string `json:"start,required"`
	// Resolution of results in seconds.
	Step int64 `json:"step,required"`
	// Hash with timeseries information, containing the name of timeseries as key
	TimeSeries map[string]ServerMetricListResponseMetricsTimeSeries `json:"time_series,required"`
	JSON       serverMetricListResponseMetricsJSON
}

// serverMetricListResponseMetricsJSON contains the JSON metadata for the struct
// [ServerMetricListResponseMetrics]
type serverMetricListResponseMetricsJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Step        apijson.Field
	TimeSeries  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerMetricListResponseMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ServerMetricListResponseMetricsTimeSeries struct {
	// Metrics Timestamps with values
	Values [][]ServerMetricListResponseMetricsTimeSeriesValue `json:"values,required"`
	JSON   serverMetricListResponseMetricsTimeSeriesJSON
}

// serverMetricListResponseMetricsTimeSeriesJSON contains the JSON metadata for the
// struct [ServerMetricListResponseMetricsTimeSeries]
type serverMetricListResponseMetricsTimeSeriesJSON struct {
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServerMetricListResponseMetricsTimeSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type ServerMetricListResponseMetricsTimeSeriesValue interface {
	ImplementsServerMetricListResponseMetricsTimeSeriesValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ServerMetricListResponseMetricsTimeSeriesValue)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ServerMetricListParams struct {
	// End of period to get Metrics for (in ISO-8601 format)
	End param.Field[string] `query:"end,required"`
	// Start of period to get Metrics for (in ISO-8601 format)
	Start param.Field[string] `query:"start,required"`
	// Type of metrics to get
	Type param.Field[ServerMetricListParamsType] `query:"type,required"`
	// Resolution of results in seconds
	Step param.Field[string] `query:"step"`
}

// URLQuery serializes [ServerMetricListParams]'s query parameters as `url.Values`.
func (r ServerMetricListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Type of metrics to get
type ServerMetricListParamsType string

const (
	ServerMetricListParamsTypeCpu     ServerMetricListParamsType = "cpu"
	ServerMetricListParamsTypeDisk    ServerMetricListParamsType = "disk"
	ServerMetricListParamsTypeNetwork ServerMetricListParamsType = "network"
)
