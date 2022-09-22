package v31

import (
	"github.com/idiomatic-go/metric-data/metric"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type ConnectionProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Number of bytes received from downstream.
	ReceivedBytes uint64 `protobuf:"varint,1,opt,name=received_bytes,json=receivedBytes,proto3" json:"received_bytes,omitempty"`
	// Number of bytes sent to downstream.
	SentBytes uint64 `protobuf:"varint,2,opt,name=sent_bytes,json=sentBytes,proto3" json:"sent_bytes,omitempty"`
}

type TCPAccessLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Common properties shared by all Envoy access logs.
	CommonProperties *metric.AccessLogCommon `protobuf:"bytes,1,opt,name=common_properties,json=commonProperties,proto3" json:"common_properties,omitempty"`
	// Properties of the TCP connection.
	ConnectionProperties *ConnectionProperties `protobuf:"bytes,2,opt,name=connection_properties,json=connectionProperties,proto3" json:"connection_properties,omitempty"`
}
