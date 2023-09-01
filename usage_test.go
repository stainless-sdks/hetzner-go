// File generated from our OpenAPI spec by Stainless.

package hetzner_test

import (
	"context"
	"testing"

	"github.com/hetzner/hetzner-go"
	"github.com/hetzner/hetzner-go/internal/testutil"
	"github.com/hetzner/hetzner-go/option"
)

func TestUsage(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	serverNewResponse, err := client.Servers.New(context.TODO(), hetzner.ServerNewParams{
		Image:      hetzner.F("ubuntu-20.04"),
		Name:       hetzner.F("my-server"),
		ServerType: hetzner.F("cx11"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", serverNewResponse.Server.ID)
}
