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

func TestServerNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Servers.New(context.TODO(), hetzner.ServerNewParams{
		Image:      hetzner.F("ubuntu-20.04"),
		Name:       hetzner.F("my-server"),
		ServerType: hetzner.F("cx11"),
		Automount:  hetzner.F(false),
		Datacenter: hetzner.F("nbg1-dc3"),
		Firewalls: hetzner.F([]hetzner.ServerNewParamsFirewall{{
			Firewall: hetzner.F(int64(0)),
		}, {
			Firewall: hetzner.F(int64(0)),
		}, {
			Firewall: hetzner.F(int64(0)),
		}}),
		Labels: hetzner.F(map[string]string{
			"foo": "string",
		}),
		Location:       hetzner.F("nbg1"),
		Networks:       hetzner.F([]int64{int64(0), int64(0), int64(0)}),
		PlacementGroup: hetzner.F(int64(1)),
		PublicNet: hetzner.F(hetzner.ServerNewParamsPublicNet{
			EnableIpv4: hetzner.F(true),
			EnableIpv6: hetzner.F(true),
			Ipv4:       hetzner.F(int64(0)),
			Ipv6:       hetzner.F(int64(0)),
		}),
		SshKeys:          hetzner.F([]string{"string", "string", "string"}),
		StartAfterCreate: hetzner.F(true),
		UserData:         hetzner.F("#cloud-config\nruncmd:\n- [touch, /root/cloud-init-worked]\n"),
		Volumes:          hetzner.F([]int64{int64(0), int64(0), int64(0)}),
	})
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestServerGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Servers.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestServerUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Servers.Update(
		context.TODO(),
		int64(0),
		hetzner.ServerUpdateParams{
			Labels: hetzner.F(map[string]string{
				"foo": "string",
			}),
			Name: hetzner.F("my-server"),
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

func TestServerListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Servers.List(context.TODO(), hetzner.ServerListParams{
		LabelSelector: hetzner.F("string"),
		Name:          hetzner.F("string"),
		Page:          hetzner.F(int64(1)),
		PerPage:       hetzner.F(int64(1)),
		Sort:          hetzner.F(hetzner.ServerListParamsSortID),
		Status:        hetzner.F(hetzner.ServerListParamsStatusInitializing),
	})
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestServerDelete(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Servers.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
