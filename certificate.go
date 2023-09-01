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

// CertificateService contains methods and other services that help with
// interacting with the hetzner API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewCertificateService] method
// instead.
type CertificateService struct {
	Options []option.RequestOption
	Actions *CertificateActionService
}

// NewCertificateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCertificateService(opts ...option.RequestOption) (r *CertificateService) {
	r = &CertificateService{}
	r.Options = opts
	r.Actions = NewCertificateActionService(opts...)
	return
}

// Creates a new Certificate.
//
// The default type **uploaded** allows for uploading your existing `certificate`
// and `private_key` in PEM format. You have to monitor its expiration date and
// handle renewal yourself.
//
// In contrast, type **managed** requests a new Certificate from _Let's Encrypt_
// for the specified `domain_names`. Only domains managed by _Hetzner DNS_ are
// supported. We handle renewal and timely alert the project owner via email if
// problems occur.
//
// For type `managed` Certificates the `action` key of the response contains the
// Action that allows for tracking the issuance process. For type `uploaded`
// Certificates the `action` is always null.
func (r *CertificateService) New(ctx context.Context, body CertificateNewParams, opts ...option.RequestOption) (res *CertificateNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "certificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets a specific Certificate object.
func (r *CertificateService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *CertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("certificates/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Certificate properties.
//
// Note that when updating labels, the Certificate’s current set of labels will be
// replaced with the labels provided in the request body. So, for example, if you
// want to add a new label, you have to provide all existing labels plus the new
// label in the request body.
//
// Note: if the Certificate object changes during the request, the response will be
// a “conflict” error.
func (r *CertificateService) Update(ctx context.Context, id int64, body CertificateUpdateParams, opts ...option.RequestOption) (res *CertificateUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("certificates/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all Certificate objects.
func (r *CertificateService) List(ctx context.Context, query CertificateListParams, opts ...option.RequestOption) (res *CertificateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "certificates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Deletes a Certificate.
func (r *CertificateService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("certificates/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// TLS/SSL Certificates prove the identity of a Server and are used to encrypt
// client traffic.
type Certificate struct {
	// ID of the Resource
	ID int64 `json:"id,required"`
	// Certificate and chain in PEM format, in order so that each record directly
	// certifies the one preceding
	Certificate string `json:"certificate,required,nullable"`
	// Point in time when the Resource was created (in ISO-8601 format)
	Created string `json:"created,required"`
	// Domains and subdomains covered by the Certificate
	DomainNames []string `json:"domain_names,required"`
	// SHA256 fingerprint of the Certificate
	Fingerprint string `json:"fingerprint,required,nullable"`
	// User-defined labels (key-value pairs)
	Labels map[string]string `json:"labels,required"`
	// Name of the Resource. Must be unique per Project.
	Name string `json:"name,required"`
	// Point in time when the Certificate stops being valid (in ISO-8601 format)
	NotValidAfter string `json:"not_valid_after,required,nullable"`
	// Point in time when the Certificate becomes valid (in ISO-8601 format)
	NotValidBefore string `json:"not_valid_before,required,nullable"`
	// Resources currently using the Certificate
	UsedBy []CertificateUsedBy `json:"used_by,required"`
	// Current status of a type `managed` Certificate, always _null_ for type
	// `uploaded` Certificates
	Status CertificateStatus `json:"status,nullable"`
	// Type of the Certificate
	Type CertificateType `json:"type"`
	JSON certificateJSON
}

// certificateJSON contains the JSON metadata for the struct [Certificate]
type certificateJSON struct {
	ID             apijson.Field
	Certificate    apijson.Field
	Created        apijson.Field
	DomainNames    apijson.Field
	Fingerprint    apijson.Field
	Labels         apijson.Field
	Name           apijson.Field
	NotValidAfter  apijson.Field
	NotValidBefore apijson.Field
	UsedBy         apijson.Field
	Status         apijson.Field
	Type           apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *Certificate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateUsedBy struct {
	// ID of the Resource | ID of resource referenced
	ID int64 `json:"id,required"`
	// Type of resource referenced
	Type string `json:"type,required"`
	JSON certificateUsedByJSON
}

// certificateUsedByJSON contains the JSON metadata for the struct
// [CertificateUsedBy]
type certificateUsedByJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateUsedBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of a type `managed` Certificate, always _null_ for type
// `uploaded` Certificates
type CertificateStatus struct {
	// If issuance or renewal reports `failed`, this property contains information
	// about what happened
	Error CertificateStatusError `json:"error,nullable"`
	// Status of the issuance process of the Certificate
	Issuance CertificateStatusIssuance `json:"issuance"`
	// Status of the renewal process of the Certificate.
	Renewal CertificateStatusRenewal `json:"renewal"`
	JSON    certificateStatusJSON
}

// certificateStatusJSON contains the JSON metadata for the struct
// [CertificateStatus]
type certificateStatusJSON struct {
	Error       apijson.Field
	Issuance    apijson.Field
	Renewal     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// If issuance or renewal reports `failed`, this property contains information
// about what happened
type CertificateStatusError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	JSON    certificateStatusErrorJSON
}

// certificateStatusErrorJSON contains the JSON metadata for the struct
// [CertificateStatusError]
type certificateStatusErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateStatusError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the issuance process of the Certificate
type CertificateStatusIssuance string

const (
	CertificateStatusIssuanceCompleted CertificateStatusIssuance = "completed"
	CertificateStatusIssuanceFailed    CertificateStatusIssuance = "failed"
	CertificateStatusIssuancePending   CertificateStatusIssuance = "pending"
)

// Status of the renewal process of the Certificate.
type CertificateStatusRenewal string

const (
	CertificateStatusRenewalFailed      CertificateStatusRenewal = "failed"
	CertificateStatusRenewalPending     CertificateStatusRenewal = "pending"
	CertificateStatusRenewalScheduled   CertificateStatusRenewal = "scheduled"
	CertificateStatusRenewalUnavailable CertificateStatusRenewal = "unavailable"
)

// Type of the Certificate
type CertificateType string

const (
	CertificateTypeManaged  CertificateType = "managed"
	CertificateTypeUploaded CertificateType = "uploaded"
)

// Response to POST https://api.hetzner.cloud/v1/certificates
type CertificateNewResponse struct {
	// TLS/SSL Certificates prove the identity of a Server and are used to encrypt
	// client traffic.
	Certificate Certificate `json:"certificate,required"`
	// Actions show the results and progress of asynchronous requests to the API.
	Action CertificateNewResponseAction `json:"action,nullable"`
	JSON   certificateNewResponseJSON
}

// certificateNewResponseJSON contains the JSON metadata for the struct
// [CertificateNewResponse]
type certificateNewResponseJSON struct {
	Certificate apijson.Field
	Action      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Actions show the results and progress of asynchronous requests to the API.
type CertificateNewResponseAction struct {
	// ID of the Action
	ID int64 `json:"id,required"`
	// Command executed in the Action
	Command string `json:"command,required"`
	// Error message for the Action if error occurred, otherwise null
	Error CertificateNewResponseActionError `json:"error,required,nullable"`
	// Point in time when the Action was finished (in ISO-8601 format). Only set if the
	// Action is finished otherwise null.
	Finished string `json:"finished,required,nullable"`
	// Progress of Action in percent
	Progress float64 `json:"progress,required"`
	// Resources the Action relates to
	Resources []CertificateNewResponseActionResource `json:"resources,required"`
	// Point in time when the Action was started (in ISO-8601 format)
	Started string `json:"started,required"`
	// Status of the Action
	Status CertificateNewResponseActionStatus `json:"status,required"`
	JSON   certificateNewResponseActionJSON
}

// certificateNewResponseActionJSON contains the JSON metadata for the struct
// [CertificateNewResponseAction]
type certificateNewResponseActionJSON struct {
	ID          apijson.Field
	Command     apijson.Field
	Error       apijson.Field
	Finished    apijson.Field
	Progress    apijson.Field
	Resources   apijson.Field
	Started     apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateNewResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Error message for the Action if error occurred, otherwise null
type CertificateNewResponseActionError struct {
	// Fixed machine readable code
	Code string `json:"code,required"`
	// Humanized error message
	Message string `json:"message,required"`
	JSON    certificateNewResponseActionErrorJSON
}

// certificateNewResponseActionErrorJSON contains the JSON metadata for the struct
// [CertificateNewResponseActionError]
type certificateNewResponseActionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateNewResponseActionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateNewResponseActionResource struct {
	// ID of the Resource | ID of resource referenced
	ID int64 `json:"id,required"`
	// Type of resource referenced
	Type string `json:"type,required"`
	JSON certificateNewResponseActionResourceJSON
}

// certificateNewResponseActionResourceJSON contains the JSON metadata for the
// struct [CertificateNewResponseActionResource]
type certificateNewResponseActionResourceJSON struct {
	ID          apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateNewResponseActionResource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Action
type CertificateNewResponseActionStatus string

const (
	CertificateNewResponseActionStatusError   CertificateNewResponseActionStatus = "error"
	CertificateNewResponseActionStatusRunning CertificateNewResponseActionStatus = "running"
	CertificateNewResponseActionStatusSuccess CertificateNewResponseActionStatus = "success"
)

// Response to GET https://api.hetzner.cloud/v1/certificates/{id}
type CertificateGetResponse struct {
	// TLS/SSL Certificates prove the identity of a Server and are used to encrypt
	// client traffic.
	Certificate Certificate `json:"certificate,required"`
	JSON        certificateGetResponseJSON
}

// certificateGetResponseJSON contains the JSON metadata for the struct
// [CertificateGetResponse]
type certificateGetResponseJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to PUT https://api.hetzner.cloud/v1/certificates/{id}
type CertificateUpdateResponse struct {
	// TLS/SSL Certificates prove the identity of a Server and are used to encrypt
	// client traffic.
	Certificate Certificate `json:"certificate,required"`
	JSON        certificateUpdateResponseJSON
}

// certificateUpdateResponseJSON contains the JSON metadata for the struct
// [CertificateUpdateResponse]
type certificateUpdateResponseJSON struct {
	Certificate apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CertificateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Response to GET https://api.hetzner.cloud/v1/certificates
type CertificateListResponse struct {
	Certificates []Certificate `json:"certificates,required"`
	// Metadata contained in the response
	Meta shared.ResponseMeta `json:"meta"`
	JSON certificateListResponseJSON
}

// certificateListResponseJSON contains the JSON metadata for the struct
// [CertificateListResponse]
type certificateListResponseJSON struct {
	Certificates apijson.Field
	Meta         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *CertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CertificateNewParams struct {
	// Name of the Certificate
	Name param.Field[string] `json:"name,required"`
	// Certificate and chain in PEM format, in order so that each record directly
	// certifies the one preceding. Required for type `uploaded` Certificates.
	Certificate param.Field[string] `json:"certificate"`
	// Domains and subdomains that should be contained in the Certificate issued by
	// _Let's Encrypt_. Required for type `managed` Certificates.
	DomainNames param.Field[[]string] `json:"domain_names"`
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// Certificate key in PEM format. Required for type `uploaded` Certificates.
	PrivateKey param.Field[string] `json:"private_key"`
	// Choose between uploading a Certificate in PEM format or requesting a managed
	// _Let's Encrypt_ Certificate. If omitted defaults to `uploaded`.
	Type param.Field[CertificateNewParamsType] `json:"type"`
}

func (r CertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Choose between uploading a Certificate in PEM format or requesting a managed
// _Let's Encrypt_ Certificate. If omitted defaults to `uploaded`.
type CertificateNewParamsType string

const (
	CertificateNewParamsTypeManaged  CertificateNewParamsType = "managed"
	CertificateNewParamsTypeUploaded CertificateNewParamsType = "uploaded"
)

type CertificateUpdateParams struct {
	// User-defined labels (key-value pairs)
	Labels param.Field[map[string]string] `json:"labels"`
	// New Certificate name
	Name param.Field[string] `json:"name"`
}

func (r CertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type CertificateListParams struct {
	// Can be used to filter resources by labels. The response will only contain
	// resources matching the label selector.
	LabelSelector param.Field[string] `query:"label_selector"`
	// Can be used to filter resources by their name. The response will only contain
	// the resources matching the specified name
	Name param.Field[string] `query:"name"`
	// Specifies the page to fetch. The number of the first page is 1
	Page param.Field[int64] `query:"page"`
	// Specifies the number of items returned per page. The default value is 25, the
	// maximum value is 50 except otherwise specified in the documentation.
	PerPage param.Field[int64] `query:"per_page"`
	// Can be used multiple times.
	Sort param.Field[CertificateListParamsSort] `query:"sort"`
	// Can be used multiple times. The response will only contain Certificates matching
	// the type.
	Type param.Field[CertificateListParamsType] `query:"type"`
}

// URLQuery serializes [CertificateListParams]'s query parameters as `url.Values`.
func (r CertificateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Can be used multiple times.
type CertificateListParamsSort string

const (
	CertificateListParamsSortID          CertificateListParamsSort = "id"
	CertificateListParamsSortIDAsc       CertificateListParamsSort = "id:asc"
	CertificateListParamsSortIDDesc      CertificateListParamsSort = "id:desc"
	CertificateListParamsSortName        CertificateListParamsSort = "name"
	CertificateListParamsSortNameAsc     CertificateListParamsSort = "name:asc"
	CertificateListParamsSortNameDesc    CertificateListParamsSort = "name:desc"
	CertificateListParamsSortCreated     CertificateListParamsSort = "created"
	CertificateListParamsSortCreatedAsc  CertificateListParamsSort = "created:asc"
	CertificateListParamsSortCreatedDesc CertificateListParamsSort = "created:desc"
)

// Can be used multiple times. The response will only contain Certificates matching
// the type.
type CertificateListParamsType string

const (
	CertificateListParamsTypeUploaded CertificateListParamsType = "uploaded"
	CertificateListParamsTypeManaged  CertificateListParamsType = "managed"
)
