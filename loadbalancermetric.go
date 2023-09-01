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

// LoadBalancerMetricService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerMetricService] method
// instead.
type LoadBalancerMetricService struct {
	Options []option.RequestOption
}

// NewLoadBalancerMetricService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerMetricService(opts ...option.RequestOption) (r *LoadBalancerMetricService) {
	r = &LoadBalancerMetricService{}
	r.Options = opts
	return
}

// You must specify the type of metric to get: `open_connections`,
// `connections_per_second`, `requests_per_second` or `bandwidth`. You can also
// specify more than one type by comma separation, e.g.
// `requests_per_second,bandwidth`.
//
// Depending on the type you will get different time series data:
//
// | Type                   | Timeseries             | Unit          | Description            |
// | ---------------------- | ---------------------- | ------------- | ---------------------- |
// | open_connections       | open_connections       | number        | Open connections       |
// | connections_per_second | connections_per_second | connections/s | Connections per second |
// | requests_per_second    | requests_per_second    | requests/s    | Requests per second    |
// | bandwidth              | bandwidth.in           | bytes/s       | Ingress bandwidth      |
// |                        | bandwidth.out          | bytes/s       | Egress bandwidth       |
//
// Metrics are available for the last 30 days only.
//
// If you do not provide the step argument we will automatically adjust it so that
// 200 samples are returned.
//
// We limit the number of samples to a maximum of 500 and will adjust the step
// parameter accordingly.
func (r *LoadBalancerMetricService) List(ctx context.Context, id int64, query LoadBalancerMetricListParams, opts ...option.RequestOption) (res *LoadBalancerMetricListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("load_balancers/%v/metrics", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response to GET https://api.hetzner.cloud/v1/load_balancers/{id}/metrics
type LoadBalancerMetricListResponse struct {
	// You must specify the type of metric to get: open_connections,
	// requests_per_second or bandwidth. You can also specify more than one type by
	// comma separation, e.g. requests_per_second,bandwidth. Depending on the type you
	// will get different time series data.
	Metrics LoadBalancerMetricListResponseMetrics `json:"metrics,required"`
	JSON    loadBalancerMetricListResponseJSON
}

// loadBalancerMetricListResponseJSON contains the JSON metadata for the struct
// [LoadBalancerMetricListResponse]
type loadBalancerMetricListResponseJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMetricListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// You must specify the type of metric to get: open_connections,
// requests_per_second or bandwidth. You can also specify more than one type by
// comma separation, e.g. requests_per_second,bandwidth. Depending on the type you
// will get different time series data.
type LoadBalancerMetricListResponseMetrics struct {
	// End of period of metrics reported (in ISO-8601 format)
	End string `json:"end,required"`
	// Start of period of metrics reported (in ISO-8601 format)
	Start string `json:"start,required"`
	// Resolution of results in seconds.
	Step int64 `json:"step,required"`
	// Hash with timeseries information, containing the name of timeseries as key
	TimeSeries map[string]LoadBalancerMetricListResponseMetricsTimeSeries `json:"time_series,required"`
	JSON       loadBalancerMetricListResponseMetricsJSON
}

// loadBalancerMetricListResponseMetricsJSON contains the JSON metadata for the
// struct [LoadBalancerMetricListResponseMetrics]
type loadBalancerMetricListResponseMetricsJSON struct {
	End         apijson.Field
	Start       apijson.Field
	Step        apijson.Field
	TimeSeries  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMetricListResponseMetrics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerMetricListResponseMetricsTimeSeries struct {
	// Metrics Timestamps with values
	Values [][]LoadBalancerMetricListResponseMetricsTimeSeriesValue `json:"values,required"`
	JSON   loadBalancerMetricListResponseMetricsTimeSeriesJSON
}

// loadBalancerMetricListResponseMetricsTimeSeriesJSON contains the JSON metadata
// for the struct [LoadBalancerMetricListResponseMetricsTimeSeries]
type loadBalancerMetricListResponseMetricsTimeSeriesJSON struct {
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerMetricListResponseMetricsTimeSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [shared.UnionFloat] or [shared.UnionString].
type LoadBalancerMetricListResponseMetricsTimeSeriesValue interface {
	ImplementsLoadBalancerMetricListResponseMetricsTimeSeriesValue()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LoadBalancerMetricListResponseMetricsTimeSeriesValue)(nil)).Elem(),
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

type LoadBalancerMetricListParams struct {
	// End of period to get Metrics for (in ISO-8601 format)
	End param.Field[string] `query:"end,required"`
	// Start of period to get Metrics for (in ISO-8601 format)
	Start param.Field[string] `query:"start,required"`
	// Type of metrics to get
	Type param.Field[LoadBalancerMetricListParamsType] `query:"type,required"`
	// Resolution of results in seconds
	Step param.Field[string] `query:"step"`
}

// URLQuery serializes [LoadBalancerMetricListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerMetricListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Type of metrics to get
type LoadBalancerMetricListParamsType string

const (
	LoadBalancerMetricListParamsTypeOpenConnections      LoadBalancerMetricListParamsType = "open_connections"
	LoadBalancerMetricListParamsTypeConnectionsPerSecond LoadBalancerMetricListParamsType = "connections_per_second"
	LoadBalancerMetricListParamsTypeRequestsPerSecond    LoadBalancerMetricListParamsType = "requests_per_second"
	LoadBalancerMetricListParamsTypeBandwidth            LoadBalancerMetricListParamsType = "bandwidth"
)
