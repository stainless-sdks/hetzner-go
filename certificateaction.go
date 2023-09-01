// File generated from our OpenAPI spec by Stainless.

package hetzner

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/hetzner/hetzner-go/internal/apijson"
	"github.com/hetzner/hetzner-go/internal/apiquery"
	"github.com/hetzner/hetzner-go/internal/param"
	"github.com/hetzner/hetzner-go/internal/requestconfig"
	"github.com/hetzner/hetzner-go/internal/shared"
	"github.com/hetzner/hetzner-go/option"
)

// CertificateActionService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCertificateActionService] method
// instead.
type CertificateActionService struct {
	Options []option.RequestOption
}

// NewCertificateActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCertificateActionService(opts ...option.RequestOption) (r *CertificateActionService) {
	r = &CertificateActionService{}
	r.Options = opts
	return
}

// Returns a specific Action for a Certificate. Only type `managed` Certificates
// have Actions.
func (r *CertificateActionService) Get(ctx context.Context, id int64, actionID int64, opts ...option.RequestOption) (res *CertificateActionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("certificates/%v/actions/%v", id, actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all Action objects for a Certificate. You can sort the results by using
// the `sort` URI parameter, and filter them with the `status` parameter.
//
// Only type `managed` Certificates can have Actions. For type `uploaded`
// Certificates the `actions` key will always contain an empty array.
func (r *CertificateActionService) List(ctx context.Context, id int64, query CertificateActionListParams, opts ...option.RequestOption) (res *CertificateActionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("certificates/%v/actions", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retry a failed Certificate issuance or renewal.
//
// Only applicable if the type of the Certificate is `managed` and the issuance or
// renewal status is `failed`.
//
// #### Call specific error codes
//
// | Code                                                    | Description                                                               |
// | ------------------------------------------------------- | ------------------------------------------------------------------------- |
// | `caa_record_does_not_allow_ca`                          | CAA record does not allow certificate authority                           |
// | `ca_dns_validation_failed`                              | Certificate Authority: DNS validation failed                              |
// | `ca_too_many_authorizations_failed_recently`            | Certificate Authority: Too many authorizations failed recently            |
// | `ca_too_many_certificates_issued_for_registered_domain` | Certificate Authority: Too many certificates issued for registered domain |
// | `ca_too_many_duplicate_certificates`                    | Certificate Authority: Too many duplicate certificates                    |
// | `could_not_verify_domain_delegated_to_zone`             | Could not verify domain delegated to zone                                 |
// | `dns_zone_not_found`                                    | DNS zone not found                                                        |
// | `dns_zone_is_secondary_zone`                            | DNS zone is a secondary zone                                              |
func (r *CertificateActionService) Retry(ctx context.Context, id int64, opts ...option.RequestOption) (res *CertificateActionRetryResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("certificates/%v/actions/retry", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Response to GET
// https://api.hetzner.cloud/v1/certificates/{id}/actions/{action_id}
type CertificateActionGetResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   certificateActionGetResponseJSON
}

// certificateActionGetResponseJSON contains the JSON metadata for the struct
// [CertificateActionGetResponse]
type certificateActionGetResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateActionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/certificates/{id}/actions
type CertificateActionListResponse struct {
	Actions []shared.Action `json:"actions,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON certificateActionListResponseJSON
}

// certificateActionListResponseJSON contains the JSON metadata for the struct
// [CertificateActionListResponse]
type certificateActionListResponseJSON struct {
	Actions     apijson.Field
	Meta        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateActionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to POST https://api.hetzner.cloud/v1/certificates/{id}/actions/retry
type CertificateActionRetryResponse struct {
	// Actions show the results and progress of asynchronous requests to the API.
	Action shared.Action `json:"action,required"`
	JSON   certificateActionRetryResponseJSON
}

// certificateActionRetryResponseJSON contains the JSON metadata for the struct
// [CertificateActionRetryResponse]
type certificateActionRetryResponseJSON struct {
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateActionRetryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateActionListParams struct {
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
	// Can be used multiple times.
	Sort param.Field[shared.SortParam] `query:"sort"`
	// Can be used multiple times, the response will contain only Actions with
	// specified statuses
	Status param.Field[shared.StatusParam] `query:"status"`
}

// URLQuery serializes [CertificateActionListParams]'s query parameters as
// `url.Values`.
func (r CertificateActionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
