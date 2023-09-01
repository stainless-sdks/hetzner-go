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

func TestLoadBalancerNewWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.New(context.TODO(), hetzner.LoadBalancerNewParams{
		Algorithm: hetzner.F(hetzner.LoadBalancerNewParamsAlgorithm{
			Type: hetzner.F(hetzner.LoadBalancerNewParamsAlgorithmTypeLeastConnections),
		}),
		LoadBalancerType: hetzner.F("lb11"),
		Name:             hetzner.F("Web Frontend"),
		Labels: hetzner.F(map[string]string{
			"foo": "string",
		}),
		Location:        hetzner.F("string"),
		Network:         hetzner.F(int64(123)),
		NetworkZone:     hetzner.F("eu-central"),
		PublicInterface: hetzner.F(true),
		Services: hetzner.F([]hetzner.LoadBalancerServiceModelParam{{
			DestinationPort: hetzner.F(int64(80)),
			HealthCheck: hetzner.F(hetzner.LoadBalancerServiceModelHealthCheckParam{
				HTTP: hetzner.F(hetzner.LoadBalancerServiceModelHealthCheckHTTPParam{
					Domain:      hetzner.F("example.com"),
					Path:        hetzner.F("/"),
					Response:    hetzner.F("{\"status\": \"ok\"}"),
					StatusCodes: hetzner.F([]string{"string", "string", "string"}),
					Tls:         hetzner.F(false),
				}),
				Interval: hetzner.F(int64(15)),
				Port:     hetzner.F(int64(4711)),
				Protocol: hetzner.F(hetzner.LoadBalancerServiceModelHealthCheckProtocolHTTP),
				Retries:  hetzner.F(int64(3)),
				Timeout:  hetzner.F(int64(10)),
			}),
			HTTP: hetzner.F(hetzner.LoadBalancerServiceModelHTTPParam{
				Certificates:   hetzner.F([]int64{int64(0), int64(0), int64(0)}),
				CookieLifetime: hetzner.F(int64(300)),
				CookieName:     hetzner.F("HCLBSTICKY"),
				RedirectHTTP:   hetzner.F(true),
				StickySessions: hetzner.F(true),
			}),
			ListenPort:    hetzner.F(int64(443)),
			Protocol:      hetzner.F(hetzner.LoadBalancerServiceModelProtocolHTTP),
			Proxyprotocol: hetzner.F(false),
		}, {
			DestinationPort: hetzner.F(int64(80)),
			HealthCheck: hetzner.F(hetzner.LoadBalancerServiceModelHealthCheckParam{
				HTTP: hetzner.F(hetzner.LoadBalancerServiceModelHealthCheckHTTPParam{
					Domain:      hetzner.F("example.com"),
					Path:        hetzner.F("/"),
					Response:    hetzner.F("{\"status\": \"ok\"}"),
					StatusCodes: hetzner.F([]string{"string", "string", "string"}),
					Tls:         hetzner.F(false),
				}),
				Interval: hetzner.F(int64(15)),
				Port:     hetzner.F(int64(4711)),
				Protocol: hetzner.F(hetzner.LoadBalancerServiceModelHealthCheckProtocolHTTP),
				Retries:  hetzner.F(int64(3)),
				Timeout:  hetzner.F(int64(10)),
			}),
			HTTP: hetzner.F(hetzner.LoadBalancerServiceModelHTTPParam{
				Certificates:   hetzner.F([]int64{int64(0), int64(0), int64(0)}),
				CookieLifetime: hetzner.F(int64(300)),
				CookieName:     hetzner.F("HCLBSTICKY"),
				RedirectHTTP:   hetzner.F(true),
				StickySessions: hetzner.F(true),
			}),
			ListenPort:    hetzner.F(int64(443)),
			Protocol:      hetzner.F(hetzner.LoadBalancerServiceModelProtocolHTTP),
			Proxyprotocol: hetzner.F(false),
		}, {
			DestinationPort: hetzner.F(int64(80)),
			HealthCheck: hetzner.F(hetzner.LoadBalancerServiceModelHealthCheckParam{
				HTTP: hetzner.F(hetzner.LoadBalancerServiceModelHealthCheckHTTPParam{
					Domain:      hetzner.F("example.com"),
					Path:        hetzner.F("/"),
					Response:    hetzner.F("{\"status\": \"ok\"}"),
					StatusCodes: hetzner.F([]string{"string", "string", "string"}),
					Tls:         hetzner.F(false),
				}),
				Interval: hetzner.F(int64(15)),
				Port:     hetzner.F(int64(4711)),
				Protocol: hetzner.F(hetzner.LoadBalancerServiceModelHealthCheckProtocolHTTP),
				Retries:  hetzner.F(int64(3)),
				Timeout:  hetzner.F(int64(10)),
			}),
			HTTP: hetzner.F(hetzner.LoadBalancerServiceModelHTTPParam{
				Certificates:   hetzner.F([]int64{int64(0), int64(0), int64(0)}),
				CookieLifetime: hetzner.F(int64(300)),
				CookieName:     hetzner.F("HCLBSTICKY"),
				RedirectHTTP:   hetzner.F(true),
				StickySessions: hetzner.F(true),
			}),
			ListenPort:    hetzner.F(int64(443)),
			Protocol:      hetzner.F(hetzner.LoadBalancerServiceModelProtocolHTTP),
			Proxyprotocol: hetzner.F(false),
		}}),
		Targets: hetzner.F([]hetzner.LoadBalancerNewParamsTarget{{
			HealthStatus: hetzner.F([]shared.HealthStatusParam{{
				ListenPort: hetzner.F(int64(443)),
				Status:     hetzner.F(shared.HealthStatusStatusHealthy),
			}, {
				ListenPort: hetzner.F(int64(443)),
				Status:     hetzner.F(shared.HealthStatusStatusHealthy),
			}, {
				ListenPort: hetzner.F(int64(443)),
				Status:     hetzner.F(shared.HealthStatusStatusHealthy),
			}}),
			IP: hetzner.F(hetzner.LoadBalancerTargetIPParam{
				IP: hetzner.F("203.0.113.1"),
			}),
			LabelSelector: hetzner.F(hetzner.LoadBalancerNewParamsTargetsLabelSelector{
				Selector: hetzner.F("env=prod"),
			}),
			Server: hetzner.F(hetzner.LoadBalancerNewParamsTargetsServer{
				ID: hetzner.F(int64(42)),
			}),
			Targets: hetzner.F([]hetzner.LoadBalancerNewParamsTargetsTarget{{
				HealthStatus: hetzner.F([]shared.HealthStatusParam{{
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}}),
				Server: hetzner.F(hetzner.LoadBalancerNewParamsTargetsTargetsServer{
					ID: hetzner.F(int64(42)),
				}),
				Type:         hetzner.F("server"),
				UsePrivateIP: hetzner.F(true),
			}, {
				HealthStatus: hetzner.F([]shared.HealthStatusParam{{
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}}),
				Server: hetzner.F(hetzner.LoadBalancerNewParamsTargetsTargetsServer{
					ID: hetzner.F(int64(42)),
				}),
				Type:         hetzner.F("server"),
				UsePrivateIP: hetzner.F(true),
			}, {
				HealthStatus: hetzner.F([]shared.HealthStatusParam{{
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}}),
				Server: hetzner.F(hetzner.LoadBalancerNewParamsTargetsTargetsServer{
					ID: hetzner.F(int64(42)),
				}),
				Type:         hetzner.F("server"),
				UsePrivateIP: hetzner.F(true),
			}}),
			Type:         hetzner.F(hetzner.LoadBalancerNewParamsTargetsTypeIP),
			UsePrivateIP: hetzner.F(true),
		}, {
			HealthStatus: hetzner.F([]shared.HealthStatusParam{{
				ListenPort: hetzner.F(int64(443)),
				Status:     hetzner.F(shared.HealthStatusStatusHealthy),
			}, {
				ListenPort: hetzner.F(int64(443)),
				Status:     hetzner.F(shared.HealthStatusStatusHealthy),
			}, {
				ListenPort: hetzner.F(int64(443)),
				Status:     hetzner.F(shared.HealthStatusStatusHealthy),
			}}),
			IP: hetzner.F(hetzner.LoadBalancerTargetIPParam{
				IP: hetzner.F("203.0.113.1"),
			}),
			LabelSelector: hetzner.F(hetzner.LoadBalancerNewParamsTargetsLabelSelector{
				Selector: hetzner.F("env=prod"),
			}),
			Server: hetzner.F(hetzner.LoadBalancerNewParamsTargetsServer{
				ID: hetzner.F(int64(42)),
			}),
			Targets: hetzner.F([]hetzner.LoadBalancerNewParamsTargetsTarget{{
				HealthStatus: hetzner.F([]shared.HealthStatusParam{{
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}}),
				Server: hetzner.F(hetzner.LoadBalancerNewParamsTargetsTargetsServer{
					ID: hetzner.F(int64(42)),
				}),
				Type:         hetzner.F("server"),
				UsePrivateIP: hetzner.F(true),
			}, {
				HealthStatus: hetzner.F([]shared.HealthStatusParam{{
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}}),
				Server: hetzner.F(hetzner.LoadBalancerNewParamsTargetsTargetsServer{
					ID: hetzner.F(int64(42)),
				}),
				Type:         hetzner.F("server"),
				UsePrivateIP: hetzner.F(true),
			}, {
				HealthStatus: hetzner.F([]shared.HealthStatusParam{{
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}}),
				Server: hetzner.F(hetzner.LoadBalancerNewParamsTargetsTargetsServer{
					ID: hetzner.F(int64(42)),
				}),
				Type:         hetzner.F("server"),
				UsePrivateIP: hetzner.F(true),
			}}),
			Type:         hetzner.F(hetzner.LoadBalancerNewParamsTargetsTypeIP),
			UsePrivateIP: hetzner.F(true),
		}, {
			HealthStatus: hetzner.F([]shared.HealthStatusParam{{
				ListenPort: hetzner.F(int64(443)),
				Status:     hetzner.F(shared.HealthStatusStatusHealthy),
			}, {
				ListenPort: hetzner.F(int64(443)),
				Status:     hetzner.F(shared.HealthStatusStatusHealthy),
			}, {
				ListenPort: hetzner.F(int64(443)),
				Status:     hetzner.F(shared.HealthStatusStatusHealthy),
			}}),
			IP: hetzner.F(hetzner.LoadBalancerTargetIPParam{
				IP: hetzner.F("203.0.113.1"),
			}),
			LabelSelector: hetzner.F(hetzner.LoadBalancerNewParamsTargetsLabelSelector{
				Selector: hetzner.F("env=prod"),
			}),
			Server: hetzner.F(hetzner.LoadBalancerNewParamsTargetsServer{
				ID: hetzner.F(int64(42)),
			}),
			Targets: hetzner.F([]hetzner.LoadBalancerNewParamsTargetsTarget{{
				HealthStatus: hetzner.F([]shared.HealthStatusParam{{
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}}),
				Server: hetzner.F(hetzner.LoadBalancerNewParamsTargetsTargetsServer{
					ID: hetzner.F(int64(42)),
				}),
				Type:         hetzner.F("server"),
				UsePrivateIP: hetzner.F(true),
			}, {
				HealthStatus: hetzner.F([]shared.HealthStatusParam{{
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}}),
				Server: hetzner.F(hetzner.LoadBalancerNewParamsTargetsTargetsServer{
					ID: hetzner.F(int64(42)),
				}),
				Type:         hetzner.F("server"),
				UsePrivateIP: hetzner.F(true),
			}, {
				HealthStatus: hetzner.F([]shared.HealthStatusParam{{
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}, {
					ListenPort: hetzner.F(int64(443)),
					Status:     hetzner.F(shared.HealthStatusStatusHealthy),
				}}),
				Server: hetzner.F(hetzner.LoadBalancerNewParamsTargetsTargetsServer{
					ID: hetzner.F(int64(42)),
				}),
				Type:         hetzner.F("server"),
				UsePrivateIP: hetzner.F(true),
			}}),
			Type:         hetzner.F(hetzner.LoadBalancerNewParamsTargetsTypeIP),
			UsePrivateIP: hetzner.F(true),
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

func TestLoadBalancerGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Get(context.TODO(), int64(0))
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerUpdateWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Update(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerUpdateParams{
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

func TestLoadBalancerListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.List(context.TODO(), hetzner.LoadBalancerListParams{
		LabelSelector: hetzner.F("string"),
		Name:          hetzner.F("string"),
		Page:          hetzner.F(int64(1)),
		PerPage:       hetzner.F(int64(1)),
		Sort:          hetzner.F(hetzner.LoadBalancerListParamsSortID),
	})
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerDelete(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	err := client.LoadBalancers.Delete(context.TODO(), int64(0))
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
