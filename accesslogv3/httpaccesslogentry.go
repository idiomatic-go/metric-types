package accesslogv3

import (
	v31 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
)

type HTTPAccessLogEntry struct {
	// Common properties shared by all Envoy access logs.
	//CommonProperties *Common
	ProtocolVersion v31.HTTPAccessLogEntry_HTTPVersion
	// Description of the incoming HTTP request.
	Request *v31.HTTPRequestProperties
	// Description of the outgoing HTTP response.
	Response *v31.HTTPResponseProperties
}
