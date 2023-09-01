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

func TestNetworkNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Networks.New(context.TODO(), hetzner.NetworkNewParams{
		IPRange:               hetzner.F("10.0.0.0/16"),
		Name:                  hetzner.F("mynet"),
		ExposeRoutesToVswitch: hetzner.F(false),
		Labels: hetzner.F(map[string]string{
			"foo": "string",
		}),
		Routes: hetzner.F([]hetzner.NetworkNewParamsRoute{{
			Destination: hetzner.F("10.100.1.0/24"),
			Gateway:     hetzner.F("10.0.1.1"),
		}, {
			Destination: hetzner.F("10.100.1.0/24"),
			Gateway:     hetzner.F("10.0.1.1"),
		}, {
			Destination: hetzner.F("10.100.1.0/24"),
			Gateway:     hetzner.F("10.0.1.1"),
		}}),
		Subnets: hetzner.F([]hetzner.NetworkNewParamsSubnet{{
			IPRange:     hetzner.F("10.0.1.0/24"),
			NetworkZone: hetzner.F("eu-central"),
			Type:        hetzner.F(hetzner.NetworkNewParamsSubnetsTypeCloud),
			VswitchID:   hetzner.F(int64(1000)),
		}, {
			IPRange:     hetzner.F("10.0.1.0/24"),
			NetworkZone: hetzner.F("eu-central"),
			Type:        hetzner.F(hetzner.NetworkNewParamsSubnetsTypeCloud),
			VswitchID:   hetzner.F(int64(1000)),
		}, {
			IPRange:     hetzner.F("10.0.1.0/24"),
			NetworkZone: hetzner.F("eu-central"),
			Type:        hetzner.F(hetzner.NetworkNewParamsSubnetsTypeCloud),
			VswitchID:   hetzner.F(int64(1000)),
		}}),
	})
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Networks.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Networks.Update(
		context.TODO(),
		int64(0),
		hetzner.NetworkUpdateParams{
			ExposeRoutesToVswitch: hetzner.F(false),
			Labels: hetzner.F(map[string]string{
				"foo": "string",
			}),
			Name: hetzner.F("new-name"),
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

func TestNetworkListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Networks.List(context.TODO(), hetzner.NetworkListParams{
		LabelSelector: hetzner.F("string"),
		Name:          hetzner.F("string"),
		Page:          hetzner.F(int64(1)),
		PerPage:       hetzner.F(int64(1)),
	})
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestNetworkDelete(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	err := client.Networks.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
