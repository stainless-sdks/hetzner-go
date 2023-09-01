// File generated from our OpenAPI spec by Stainless.

package hetzner_test

import (
	"context"
	"errors"
	"testing"

	"github.com/hetzner/hetzner-go"
	"github.com/hetzner/hetzner-go/internal/shared"
	"github.com/hetzner/hetzner-go/internal/testutil"
	"github.com/hetzner/hetzner-go/option"
)

func TestFirewallActionGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Firewalls.Actions.Get(
		context.TODO(),
		int64(0),
		int64(0),
	)
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFirewallActionListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Firewalls.Actions.List(
		context.TODO(),
		int64(0),
		hetzner.FirewallActionListParams{
			Page:    hetzner.F(int64(1)),
			PerPage: hetzner.F(int64(1)),
			Sort:    hetzner.F(shared.SortParamID),
			Status:  hetzner.F(shared.StatusParamRunning),
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

func TestFirewallActionApplyToResources(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Firewalls.Actions.ApplyToResources(
		context.TODO(),
		int64(0),
		hetzner.FirewallActionApplyToResourcesParams{
			ApplyTo: hetzner.F([]hetzner.FirewallActionApplyToResourcesParamsApplyTo{{
				LabelSelector: hetzner.F(hetzner.FirewallActionApplyToResourcesParamsApplyToLabelSelector{
					Selector: hetzner.F("env=prod"),
				}),
				Server: hetzner.F(hetzner.FirewallActionApplyToResourcesParamsApplyToServer{
					ID: hetzner.F(int64(42)),
				}),
				Type: hetzner.F(hetzner.FirewallActionApplyToResourcesParamsApplyToTypeLabelSelector),
			}, {
				LabelSelector: hetzner.F(hetzner.FirewallActionApplyToResourcesParamsApplyToLabelSelector{
					Selector: hetzner.F("env=prod"),
				}),
				Server: hetzner.F(hetzner.FirewallActionApplyToResourcesParamsApplyToServer{
					ID: hetzner.F(int64(42)),
				}),
				Type: hetzner.F(hetzner.FirewallActionApplyToResourcesParamsApplyToTypeLabelSelector),
			}, {
				LabelSelector: hetzner.F(hetzner.FirewallActionApplyToResourcesParamsApplyToLabelSelector{
					Selector: hetzner.F("env=prod"),
				}),
				Server: hetzner.F(hetzner.FirewallActionApplyToResourcesParamsApplyToServer{
					ID: hetzner.F(int64(42)),
				}),
				Type: hetzner.F(hetzner.FirewallActionApplyToResourcesParamsApplyToTypeLabelSelector),
			}}),
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

func TestFirewallActionRemoveFromResources(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Firewalls.Actions.RemoveFromResources(
		context.TODO(),
		int64(0),
		hetzner.FirewallActionRemoveFromResourcesParams{
			RemoveFrom: hetzner.F([]hetzner.FirewallActionRemoveFromResourcesParamsRemoveFrom{{
				LabelSelector: hetzner.F(hetzner.FirewallActionRemoveFromResourcesParamsRemoveFromLabelSelector{
					Selector: hetzner.F("env=prod"),
				}),
				Server: hetzner.F(hetzner.FirewallActionRemoveFromResourcesParamsRemoveFromServer{
					ID: hetzner.F(int64(42)),
				}),
				Type: hetzner.F(hetzner.FirewallActionRemoveFromResourcesParamsRemoveFromTypeLabelSelector),
			}, {
				LabelSelector: hetzner.F(hetzner.FirewallActionRemoveFromResourcesParamsRemoveFromLabelSelector{
					Selector: hetzner.F("env=prod"),
				}),
				Server: hetzner.F(hetzner.FirewallActionRemoveFromResourcesParamsRemoveFromServer{
					ID: hetzner.F(int64(42)),
				}),
				Type: hetzner.F(hetzner.FirewallActionRemoveFromResourcesParamsRemoveFromTypeLabelSelector),
			}, {
				LabelSelector: hetzner.F(hetzner.FirewallActionRemoveFromResourcesParamsRemoveFromLabelSelector{
					Selector: hetzner.F("env=prod"),
				}),
				Server: hetzner.F(hetzner.FirewallActionRemoveFromResourcesParamsRemoveFromServer{
					ID: hetzner.F(int64(42)),
				}),
				Type: hetzner.F(hetzner.FirewallActionRemoveFromResourcesParamsRemoveFromTypeLabelSelector),
			}}),
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

func TestFirewallActionSetRules(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.Firewalls.Actions.SetRules(
		context.TODO(),
		int64(0),
		hetzner.FirewallActionSetRulesParams{
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
			}, {
				Description:    hetzner.F("string"),
				DestinationIps: hetzner.F([]string{"string", "string", "string"}),
				Direction:      hetzner.F(hetzner.RuleDirectionIn),
				Port:           hetzner.F("80"),
				Protocol:       hetzner.F(hetzner.RuleProtocolEsp),
				SourceIps:      hetzner.F([]string{"string", "string", "string"}),
			}}),
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
