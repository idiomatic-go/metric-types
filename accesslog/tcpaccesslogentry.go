package accesslog

type TCPAccessLogEntry struct {
	Ingress bool
	// Common properties shared by all Envoy access logs.
	CommonProperties *AccessLogCommon `protobuf:"bytes,1,opt,name=common_properties,json=commonProperties,proto3" json:"common_properties,omitempty"`
	// Properties of the TCP connection.
	ConnectionProperties *ConnectionProperties `protobuf:"bytes,2,opt,name=connection_properties,json=connectionProperties,proto3" json:"connection_properties,omitempty"`
}
