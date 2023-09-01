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

func TestFirewallNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Firewalls.New(context.TODO(), hetzner.FirewallNewParams{
		Name: hetzner.F("Corporate Intranet Protection"),
		ApplyTo: hetzner.F([]hetzner.FirewallNewParamsApplyTo{{
			LabelSelector: hetzner.F(hetzner.FirewallNewParamsApplyToLabelSelector{
				Selector: hetzner.F("env=prod"),
			}),
			Server: hetzner.F(hetzner.FirewallNewParamsApplyToServer{
				ID: hetzner.F(int64(42)),
			}),
			Type: hetzner.F(hetzner.FirewallNewParamsApplyToTypeLabelSelector),
		}, {
			LabelSelector: hetzner.F(hetzner.FirewallNewParamsApplyToLabelSelector{
				Selector: hetzner.F("env=prod"),
			}),
			Server: hetzner.F(hetzner.FirewallNewParamsApplyToServer{
				ID: hetzner.F(int64(42)),
			}),
			Type: hetzner.F(hetzner.FirewallNewParamsApplyToTypeLabelSelector),
		}, {
			LabelSelector: hetzner.F(hetzner.FirewallNewParamsApplyToLabelSelector{
				Selector: hetzner.F("env=prod"),
			}),
			Server: hetzner.F(hetzner.FirewallNewParamsApplyToServer{
				ID: hetzner.F(int64(42)),
			}),
			Type: hetzner.F(hetzner.FirewallNewParamsApplyToTypeLabelSelector),
		}}),
		Labels: hetzner.F(map[string]string{
			"foo": "string",
		}),
		Rules: hetzner.F([]hetzner.RuleParam{{
			Description:    hetzner.F("string"),
			DestinationIps: hetzner.F([]string{"string", "string", "string"}),
			Direction:      hetzner.F(hetzner.RuleDirectionIn),
			Port:           hetzner.F("80"),
			Protocol:       hetzner.F(hetzner.RuleProtocolEsp),
			SourceIps:      hetzner.F([]string{"string", "string", "string"}),
		}, {
			Description:    hetzner.F("string"),
			DestinationIps: hetzner.F([]string{"string", "string", "string"}),
			Direction:      hetzner.F(hetzner.RuleDirectionIn),
			Port:           hetzner.F("80"),
			Protocol:       hetzner.F(hetzner.RuleProtocolEsp),
			SourceIps:      hetzner.F([]string{"string", "string", "string"}),
		}, {
			Description:    hetzner.F("string"),
			DestinationIps: hetzner.F([]string{"string", "string", "string"}),
			Direction:      hetzner.F(hetzner.RuleDirectionIn),
			Port:           hetzner.F("80"),
			Protocol:       hetzner.F(hetzner.RuleProtocolEsp),
			SourceIps:      hetzner.F([]string{"string", "string", "string"}),
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

func TestFirewallGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Firewalls.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFirewallUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Firewalls.Update(
		context.TODO(),
		int64(0),
		hetzner.FirewallUpdateParams{
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

func TestFirewallListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Firewalls.List(context.TODO(), hetzner.FirewallListParams{
		LabelSelector: hetzner.F("string"),
		Name:          hetzner.F("string"),
		Page:          hetzner.F(int64(1)),
		PerPage:       hetzner.F(int64(1)),
		Sort:          hetzner.F(hetzner.FirewallListParamsSortID),
	})
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFirewallDelete(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	err := client.Firewalls.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
