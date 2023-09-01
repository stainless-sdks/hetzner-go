// File generated from our OpenAPI spec by Stainless.

package hetzner

import (
	"github.com/hetzner/hetzner-go/internal/apierror"
	"github.com/hetzner/hetzner-go/internal/shared"
)

type Error = apierror.Error

// Actions show the results and progress of asynchronous requests to the API.
//
// This is an alias to an internal type.
type Action = shared.Action

// Error message for the Action if error occurred, otherwise null
//
// This is an alias to an internal type.
type ActionError = shared.ActionError

// This is an alias to an internal type.
type ActionResource = shared.ActionResource

// Status of the Action
//
// This is an alias to an internal type.
type ActionStatus = shared.ActionStatus

// This is an alias to an internal type.
type HealthStatus = shared.HealthStatus

// This is an alias to an internal type.
type HealthStatusStatus = shared.HealthStatusStatus

// This is an alias to an internal type.
type HealthStatusParam = shared.HealthStatusParam

// Metadata contained in the response
//
// This is an alias to an internal type.
type ResponseMeta = shared.ResponseMeta

// Information about the current pagination. The keys previous_page, next_page,
// last_page, and total_entries may be null when on the first page, last page, or
// when the total number of entries is unknown
//
// This is an alias to an internal type.
type ResponseMetaPagination = shared.ResponseMetaPagination

// This is an alias to an internal type.
type SortParam = shared.SortParam

// This is an alias to an internal type.
type StatusParam = shared.StatusParam
