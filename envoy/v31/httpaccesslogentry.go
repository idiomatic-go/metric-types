package v31

import (
	"github.com/idiomatic-go/metric-data/envoy/accesslogv3"
)

type HTTPAccessLogEntry_HTTPVersion int32

type HTTPAccessLogEntry struct {

	// Common properties shared by all Envoy access logs.
	CommonProperties *accesslogv3.AccessLogCommon   `protobuf:"bytes,1,opt,name=common_properties,json=commonProperties,proto3" json:"common_properties,omitempty"`
	ProtocolVersion  HTTPAccessLogEntry_HTTPVersion `protobuf:"varint,2,opt,name=protocol_version,json=protocolVersion,proto3,enum=envoy.data.accesslogv3.v3.HTTPAccessLogEntry_HTTPVersion" json:"protocol_version,omitempty"`
	// Description of the incoming HTTP request.
	Request *accesslogv3.HTTPRequestProperties `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	// Description of the outgoing HTTP response.
	Response *accesslogv3.HTTPResponseProperties `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
}
