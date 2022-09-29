package accesslogv3

import (
	v31 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
)

type ResponseFlags_Encoded uint32

// Flags indicating occurrences during request/response processing.
type ResponseFlags struct {
	Encoded             ResponseFlags_Encoded
	UnauthorizedDetails *v31.ResponseFlags_Unauthorized
}
