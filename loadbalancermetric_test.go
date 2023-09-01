// File generated from our OpenAPI spec by Stainless.

package hetzner_test

import (
	"context"
	"errors"
	"testing"

	"github.com/hetzner/hetzner-go"
	"github.com/hetzner/hetzner-go/internal/testutil"
	"github.com/hetzner/hetzner-go/option"
)

func TestLoadBalancerMetricListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Metrics.List(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerMetricListParams{
			End:   hetzner.F("string"),
			Start: hetzner.F("string"),
			Type:  hetzner.F(hetzner.LoadBalancerMetricListParamsTypeOpenConnections),
			Step:  hetzner.F("string"),
		},
	)
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
