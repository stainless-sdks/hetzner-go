// File generated from our OpenAPI spec by Stainless.

package shared

type UnionString string

func (UnionString) ImplementsLoadBalancerMetricListResponseMetricsTimeSeriesValue() {}
func (UnionString) ImplementsServerMetricListResponseMetricsTimeSeriesValue()       {}

type UnionFloat float64

func (UnionFloat) ImplementsLoadBalancerMetricListResponseMetricsTimeSeriesValue() {}
func (UnionFloat) ImplementsServerMetricListResponseMetricsTimeSeriesValue()       {}
