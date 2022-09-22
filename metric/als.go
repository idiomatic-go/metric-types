package metric

import (
	v3 "github.com/idiomatic-go/metric-data/v3"
	v31 "github.com/idiomatic-go/metric-data/v31"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type StreamAccessLogsMessage_Identifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The node sending the access log messages over the stream.
	Node *v3.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// The friendly name of the log configured in :ref:`CommonGrpcAccessLogConfig
	// <envoy_v3_api_msg_extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig>`.
	LogName string `protobuf:"bytes,2,opt,name=log_name,json=logName,proto3" json:"log_name,omitempty"`
}

type isStreamAccessLogsMessage_LogEntries interface {
	isStreamAccessLogsMessage_LogEntries()
}

// Wrapper for batches of HTTP access log entries.
type StreamAccessLogsMessage_HTTPAccessLogEntries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogEntry []*v31.HTTPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
}

type StreamAccessLogsMessage_HttpLogs struct {
	HttpLogs *StreamAccessLogsMessage_HTTPAccessLogEntries `protobuf:"bytes,2,opt,name=http_logs,json=httpLogs,proto3,oneof"`
}

// Wrapper for batches of TCP access log entries.
type StreamAccessLogsMessage_TCPAccessLogEntries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogEntry []*v31.TCPAccessLogEntry `protobuf:"bytes,1,rep,name=log_entry,json=logEntry,proto3" json:"log_entry,omitempty"`
}

type StreamAccessLogsMessage_TcpLogs struct {
	TcpLogs *StreamAccessLogsMessage_TCPAccessLogEntries `protobuf:"bytes,3,opt,name=tcp_logs,json=tcpLogs,proto3,oneof"`
}

// Stream message for the StreamAccessLogs API. Envoy will open a stream to the server and stream
// access logs without ever expecting a response.
type StreamAccessLogsMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Identifier data that will only be sent in the first message on the stream. This is effectively
	// structured metadata and is a performance optimization.
	Identifier *StreamAccessLogsMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// Batches of log entries of a single type. Generally speaking, a given stream should only
	// ever include one type of log entry.
	//
	// Types that are assignable to LogEntries:
	//	*StreamAccessLogsMessage_HttpLogs
	//	*StreamAccessLogsMessage_TcpLogs
	LogEntries isStreamAccessLogsMessage_LogEntries `protobuf_oneof:"log_entries"`
}
