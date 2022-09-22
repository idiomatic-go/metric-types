package accesslog

type ConnectionProperties struct {
	// Number of bytes received from downstream.
	ReceivedBytes uint64 `protobuf:"varint,1,opt,name=received_bytes,json=receivedBytes,proto3" json:"received_bytes,omitempty"`
	// Number of bytes sent to downstream.
	SentBytes uint64 `protobuf:"varint,2,opt,name=sent_bytes,json=sentBytes,proto3" json:"sent_bytes,omitempty"`
}
