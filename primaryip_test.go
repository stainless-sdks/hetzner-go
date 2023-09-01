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

func TestPrimaryIpNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.PrimaryIps.New(context.TODO(), hetzner.PrimaryIpNewParams{
		AssigneeType: hetzner.F(hetzner.PrimaryIpNewParamsAssigneeTypeServer),
		Name:         hetzner.F("my-ip"),
		Type:         hetzner.F(hetzner.PrimaryIpNewParamsTypeIpv4),
		AssigneeID:   hetzner.F(int64(17)),
		AutoDelete:   hetzner.F(false),
		Datacenter:   hetzner.F("fsn1-dc8"),
		Labels: hetzner.F(map[string]string{
			"foo": "string",
		}),
	})
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrimaryIpGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.PrimaryIps.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrimaryIpUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.PrimaryIps.Update(
		context.TODO(),
		int64(0),
		hetzner.PrimaryIpUpdateParams{
			AutoDelete: hetzner.F(true),
			Labels: hetzner.F(map[string]string{
				"foo": "string",
			}),
			Name: hetzner.F("my-ip"),
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

func TestPrimaryIpListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.PrimaryIps.List(context.TODO(), hetzner.PrimaryIpListParams{
		Ip:            hetzner.F("string"),
		LabelSelector: hetzner.F("string"),
		Name:          hetzner.F("string"),
		Page:          hetzner.F(int64(1)),
		PerPage:       hetzner.F(int64(1)),
		Sort:          hetzner.F(hetzner.PrimaryIpListParamsSortID),
	})
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPrimaryIpDelete(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	err := client.PrimaryIps.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
