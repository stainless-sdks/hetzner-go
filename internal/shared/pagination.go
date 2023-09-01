// File generated from our OpenAPI spec by Stainless.

package shared

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hetzner/hetzner-go/internal/apijson"
	"github.com/hetzner/hetzner-go/internal/requestconfig"
)

// Response to GET https://api.hetzner.cloud/v1/floating_ips
type FloatingIPsPage[T any] struct {
	FloatingIPs []T `json:"floating_ips,required"`
	// Metadata contained in the response
	Meta ResponseMeta `json:"meta,required"`
	JSON floatingIPsPageJSON
	cfg  *requestconfig.RequestConfig
	res  *http.Response
}

// floatingIPsPageJSON contains the JSON metadata for the struct
// [FloatingIPsPage[T]]
type floatingIPsPageJSON struct {
	FloatingIPs apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FloatingIPsPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *FloatingIPsPage[T]) GetNextPage() (res *FloatingIPsPage[T], err error) {
	currentPage := r.Meta.Pagination.Page
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *FloatingIPsPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type FloatingIPsPageAutoPager[T any] struct {
	page *FloatingIPsPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewFloatingIPsPageAutoPager[T any](page *FloatingIPsPage[T], err error) *FloatingIPsPageAutoPager[T] {
	return &FloatingIPsPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *FloatingIPsPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.FloatingIPs) == 0 {
		return false
	}
	if r.idx >= len(r.page.FloatingIPs) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.FloatingIPs[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *FloatingIPsPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *FloatingIPsPageAutoPager[T]) Err() error {
	return r.err
}

func (r *FloatingIPsPageAutoPager[T]) Index() int {
	return r.run
}

// Response to GET https://api.hetzner.cloud/v1/servers
type ServersPage[T any] struct {
	// Metadata contained in the response
	Meta    ResponseMeta `json:"meta,required"`
	Servers []T          `json:"servers,required"`
	JSON    serversPageJSON
	cfg     *requestconfig.RequestConfig
	res     *http.Response
}

// serversPageJSON contains the JSON metadata for the struct [ServersPage[T]]
type serversPageJSON struct {
	Meta        apijson.Field
	Servers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServersPage[T]) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// NextPage returns the next page as defined by this pagination style. When there
// is no next page, this function will return a 'nil' for the page value, but will
// not return an error
func (r *ServersPage[T]) GetNextPage() (res *ServersPage[T], err error) {
	currentPage := r.Meta.Pagination.Page
	cfg := r.cfg.Clone(context.Background())
	query := cfg.Request.URL.Query()
	query.Set("page", fmt.Sprintf("%d", currentPage+1))
	cfg.Request.URL.RawQuery = query.Encode()
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *ServersPage[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	r.cfg = cfg
	r.res = res
}

type ServersPageAutoPager[T any] struct {
	page *ServersPage[T]
	cur  T
	idx  int
	run  int
	err  error
}

func NewServersPageAutoPager[T any](page *ServersPage[T], err error) *ServersPageAutoPager[T] {
	return &ServersPageAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *ServersPageAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Servers) == 0 {
		return false
	}
	if r.idx >= len(r.page.Servers) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil {
			return false
		}
	}
	r.cur = r.page.Servers[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *ServersPageAutoPager[T]) Current() T {
	return r.cur
}

func (r *ServersPageAutoPager[T]) Err() error {
	return r.err
}

func (r *ServersPageAutoPager[T]) Index() int {
	return r.run
}
