package accesslogv3

import (
	v31 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
)

type TCPAccessLogEntry struct {
	// Common properties shared by all Envoy access logs.
	//CommonProperties *Common
	// Properties of the TCP connection.
	ConnectionProperties *v31.ConnectionProperties
}
