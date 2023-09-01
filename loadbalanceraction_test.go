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

func TestLoadBalancerActionGet(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.Get(
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

func TestLoadBalancerActionListWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.List(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerActionListParams{
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

func TestLoadBalancerActionAddServiceWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.AddService(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerActionAddServiceParams{
			DestinationPort: hetzner.F(int64(80)),
			HealthCheck: hetzner.F(hetzner.LoadBalancerActionAddServiceParamsHealthCheck{
				HTTP: hetzner.F(hetzner.LoadBalancerActionAddServiceParamsHealthCheckHTTP{
					Domain:      hetzner.F("example.com"),
					Path:        hetzner.F("/"),
					Response:    hetzner.F("{\"status\": \"ok\"}"),
					StatusCodes: hetzner.F([]string{"string", "string", "string"}),
					Tls:         hetzner.F(false),
				}),
				Interval: hetzner.F(int64(15)),
				Port:     hetzner.F(int64(4711)),
				Protocol: hetzner.F(hetzner.LoadBalancerActionAddServiceParamsHealthCheckProtocolHTTP),
				Retries:  hetzner.F(int64(3)),
				Timeout:  hetzner.F(int64(10)),
			}),
			ListenPort:    hetzner.F(int64(443)),
			Protocol:      hetzner.F(hetzner.LoadBalancerActionAddServiceParamsProtocolHTTP),
			Proxyprotocol: hetzner.F(false),
			HTTP: hetzner.F(hetzner.LoadBalancerActionAddServiceParamsHTTP{
				Certificates:   hetzner.F([]int64{int64(0), int64(0), int64(0)}),
				CookieLifetime: hetzner.F(int64(300)),
				CookieName:     hetzner.F("HCLBSTICKY"),
				RedirectHTTP:   hetzner.F(true),
				StickySessions: hetzner.F(true),
			}),
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

func TestLoadBalancerActionAssTargetWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.AssTarget(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerActionAssTargetParams{
			Type: hetzner.F(hetzner.LoadBalancerActionAssTargetParamsTypeIp),
			Ip: hetzner.F(hetzner.LoadBalancerTargetIpParam{
				Ip: hetzner.F("203.0.113.1"),
			}),
			LabelSelector: hetzner.F(hetzner.LoadBalancerActionAssTargetParamsLabelSelector{
				Selector: hetzner.F("env=prod"),
			}),
			Server: hetzner.F(hetzner.LoadBalancerActionAssTargetParamsServer{
				ID: hetzner.F(int64(80)),
			}),
			UsePrivateIp: hetzner.F(true),
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

func TestLoadBalancerActionAttachToNetworkWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.AttachToNetwork(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerActionAttachToNetworkParams{
			Network: hetzner.F(int64(4711)),
			Ip:      hetzner.F("10.0.1.1"),
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

func TestLoadBalancerActionChangeAlgorithm(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.ChangeAlgorithm(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerActionChangeAlgorithmParams{
			Type: hetzner.F(hetzner.LoadBalancerActionChangeAlgorithmParamsTypeLeastConnections),
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

func TestLoadBalancerActionChangeDnsPtr(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.ChangeDnsPtr(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerActionChangeDnsPtrParams{
			DnsPtr: hetzner.F("lb1.example.com"),
			Ip:     hetzner.F("1.2.3.4"),
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

func TestLoadBalancerActionChangeProtectionWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.ChangeProtection(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerActionChangeProtectionParams{
			Delete: hetzner.F(true),
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

func TestLoadBalancerActionChangeType(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.ChangeType(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerActionChangeTypeParams{
			LoadBalancerType: hetzner.F("lb21"),
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

func TestLoadBalancerActionDeleteService(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.DeleteService(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerActionDeleteServiceParams{
			ListenPort: hetzner.F(int64(4711)),
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

func TestLoadBalancerActionDetachFromNetwork(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.DetachFromNetwork(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerActionDetachFromNetworkParams{
			Network: hetzner.F(int64(4711)),
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

func TestLoadBalancerActionDisablePublicInterface(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.DisablePublicInterface(context.TODO(), int64(0))
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerActionEnablePublicInterface(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.EnablePublicInterface(context.TODO(), int64(0))
	if err != nil {
		var apierr *hetzner.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestLoadBalancerActionRemoveTargetWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.RemoveTarget(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerActionRemoveTargetParams{
			Type: hetzner.F(hetzner.LoadBalancerActionRemoveTargetParamsTypeIp),
			Ip: hetzner.F(hetzner.LoadBalancerTargetIpParam{
				Ip: hetzner.F("203.0.113.1"),
			}),
			LabelSelector: hetzner.F(hetzner.LoadBalancerActionRemoveTargetParamsLabelSelector{
				Selector: hetzner.F("env=prod"),
			}),
			Server: hetzner.F(hetzner.LoadBalancerActionRemoveTargetParamsServer{
				ID: hetzner.F(int64(80)),
			}),
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

func TestLoadBalancerActionUpdateServiceWithOptionalParams(t *testing.T) {
	if !testutil.CheckTestServer(t) {
		return
	}
	client := hetzner.NewClient(
		option.WithBaseURL("http://127.0.0.1:4010"),
		option.WithAPIToken("APIToken"),
	)
	_, err := client.LoadBalancers.Actions.UpdateService(
		context.TODO(),
		int64(0),
		hetzner.LoadBalancerActionUpdateServiceParams{
			DestinationPort: hetzner.F(int64(80)),
			HealthCheck: hetzner.F(hetzner.LoadBalancerActionUpdateServiceParamsHealthCheck{
				HTTP: hetzner.F(hetzner.LoadBalancerActionUpdateServiceParamsHealthCheckHTTP{
					Domain:      hetzner.F("example.com"),
					Path:        hetzner.F("/"),
					Response:    hetzner.F("{\"status\": \"ok\"}"),
					StatusCodes: hetzner.F([]string{"string", "string", "string"}),
					Tls:         hetzner.F(false),
				}),
				Interval: hetzner.F(int64(15)),
				Port:     hetzner.F(int64(4711)),
				Protocol: hetzner.F(hetzner.LoadBalancerActionUpdateServiceParamsHealthCheckProtocolHTTP),
				Retries:  hetzner.F(int64(3)),
				Timeout:  hetzner.F(int64(10)),
			}),
			ListenPort:    hetzner.F(int64(443)),
			Protocol:      hetzner.F(hetzner.LoadBalancerActionUpdateServiceParamsProtocolHTTP),
			Proxyprotocol: hetzner.F(false),
			HTTP: hetzner.F(hetzner.LoadBalancerActionUpdateServiceParamsHTTP{
				Certificates:   hetzner.F([]int64{int64(0), int64(0), int64(0)}),
				CookieLifetime: hetzner.F(int64(300)),
				CookieName:     hetzner.F("HCLBSTICKY"),
				RedirectHTTP:   hetzner.F(true),
				StickySessions: hetzner.F(true),
			}),
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
