package metric

import (
	"github.com/idiomatic-go/metric-data/any"
	"github.com/idiomatic-go/metric-data/timestamp"
	v3 "github.com/idiomatic-go/metric-data/v3"
	duration "time"
)

type AccessLogCommon struct {
	SampleRate float64 `protobuf:"fixed64,1,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	// This field is the remote/origin address on which the request from the user was received.
	// Note: This may not be the physical peer. E.g, if the remote address is inferred from for
	// example the x-forwarder-for header, proxy protocol, etc.
	DownstreamRemoteAddress *v3.Address `protobuf:"bytes,2,opt,name=downstream_remote_address,json=downstreamRemoteAddress,proto3" json:"downstream_remote_address,omitempty"`
	// This field is the local/destination address on which the request from the user was received.
	DownstreamLocalAddress *v3.Address `protobuf:"bytes,3,opt,name=downstream_local_address,json=downstreamLocalAddress,proto3" json:"downstream_local_address,omitempty"`
	// If the connection is secure,S this field will contain TLS properties.
	TlsProperties *TLSProperties `protobuf:"bytes,4,opt,name=tls_properties,json=tlsProperties,proto3" json:"tls_properties,omitempty"`
	// The time that Envoy started servicing this request. This is effectively the time that the first
	// downstream byte is received.
	StartTime *timestamp.Timestamp `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Interval between the first downstream byte received and the last
	// downstream byte received (i.e. time it takes to receive a request).
	TimeToLastRxByte *duration.Duration `protobuf:"bytes,6,opt,name=time_to_last_rx_byte,json=timeToLastRxByte,proto3" json:"time_to_last_rx_byte,omitempty"`
	// Interval between the first downstream byte received and the first upstream byte sent. There may
	// by considerable delta between ``time_to_last_rx_byte`` and this value due to filters.
	// Additionally, the same caveats apply as documented in ``time_to_last_downstream_tx_byte`` about
	// not accounting for kernel socket buffer time, etc.
	TimeToFirstUpstreamTxByte *duration.Duration `protobuf:"bytes,7,opt,name=time_to_first_upstream_tx_byte,json=timeToFirstUpstreamTxByte,proto3" json:"time_to_first_upstream_tx_byte,omitempty"`
	// Interval between the first downstream byte received and the last upstream byte sent. There may
	// by considerable delta between ``time_to_last_rx_byte`` and this value due to filters.
	// Additionally, the same caveats apply as documented in ``time_to_last_downstream_tx_byte`` about
	// not accounting for kernel socket buffer time, etc.
	TimeToLastUpstreamTxByte *duration.Duration `protobuf:"bytes,8,opt,name=time_to_last_upstream_tx_byte,json=timeToLastUpstreamTxByte,proto3" json:"time_to_last_upstream_tx_byte,omitempty"`
	// Interval between the first downstream byte received and the first upstream
	// byte received (i.e. time it takes to start receiving a response).
	TimeToFirstUpstreamRxByte *duration.Duration `protobuf:"bytes,9,opt,name=time_to_first_upstream_rx_byte,json=timeToFirstUpstreamRxByte,proto3" json:"time_to_first_upstream_rx_byte,omitempty"`
	// Interval between the first downstream byte received and the last upstream
	// byte received (i.e. time it takes to receive a complete response).
	TimeToLastUpstreamRxByte *duration.Duration `protobuf:"bytes,10,opt,name=time_to_last_upstream_rx_byte,json=timeToLastUpstreamRxByte,proto3" json:"time_to_last_upstream_rx_byte,omitempty"`
	// Interval between the first downstream byte received and the first downstream byte sent.
	// There may be a considerable delta between the ``time_to_first_upstream_rx_byte`` and this field
	// due to filters. Additionally, the same caveats apply as documented in
	// ``time_to_last_downstream_tx_byte`` about not accounting for kernel socket buffer time, etc.
	TimeToFirstDownstreamTxByte *duration.Duration `protobuf:"bytes,11,opt,name=time_to_first_downstream_tx_byte,json=timeToFirstDownstreamTxByte,proto3" json:"time_to_first_downstream_tx_byte,omitempty"`
	// Interval between the first downstream byte received and the last downstream byte sent.
	// Depending on protocol, buffering, windowing, filters, etc. there may be a considerable delta
	// between ``time_to_last_upstream_rx_byte`` and this field. Note also that this is an approximate
	// time. In the current implementation it does not include kernel socket buffer time. In the
	// current implementation it also does not include send window buffering inside the HTTP/2 codec.
	// In the future it is likely that work will be done to make this duration more accurate.
	TimeToLastDownstreamTxByte *duration.Duration `protobuf:"bytes,12,opt,name=time_to_last_downstream_tx_byte,json=timeToLastDownstreamTxByte,proto3" json:"time_to_last_downstream_tx_byte,omitempty"`
	// The upstream remote/destination address that handles this exchange. This does not include
	// retries.
	UpstreamRemoteAddress *v3.Address `protobuf:"bytes,13,opt,name=upstream_remote_address,json=upstreamRemoteAddress,proto3" json:"upstream_remote_address,omitempty"`
	// The upstream local/origin address that handles this exchange. This does not include retries.
	UpstreamLocalAddress *v3.Address `protobuf:"bytes,14,opt,name=upstream_local_address,json=upstreamLocalAddress,proto3" json:"upstream_local_address,omitempty"`
	// The upstream cluster that ``upstream_remote_address`` belongs to.
	UpstreamCluster string `protobuf:"bytes,15,opt,name=upstream_cluster,json=upstreamCluster,proto3" json:"upstream_cluster,omitempty"`
	// Flags indicating occurrences during request/response processing.
	ResponseFlags *ResponseFlags `protobuf:"bytes,16,opt,name=response_flags,json=responseFlags,proto3" json:"response_flags,omitempty"`
	// All metadata encountered during request processing, including endpoint
	// selection.
	//
	// This can be used to associate IDs attached to the various configurations
	// used to process this request with the access log entry. For example, a
	// route created from a higher level forwarding rule with some ID can place
	// that ID in this field and cross reference later. It can also be used to
	// determine if a canary endpoint was used or not.
	Metadata *v3.Metadata `protobuf:"bytes,17,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// If upstream connection failed due to transport socket (e.g. TLS handshake), provides the
	// failure reason from the transport socket. The format of this field depends on the configured
	// upstream transport socket. Common TLS failures are in
	// :ref:`TLS trouble shooting <arch_overview_ssl_trouble_shooting>`.
	UpstreamTransportFailureReason string `protobuf:"bytes,18,opt,name=upstream_transport_failure_reason,json=upstreamTransportFailureReason,proto3" json:"upstream_transport_failure_reason,omitempty"`
	// The name of the route
	RouteName string `protobuf:"bytes,19,opt,name=route_name,json=routeName,proto3" json:"route_name,omitempty"`
	// This field is the downstream direct remote address on which the request from the user was
	// received. Note: This is always the physical peer, even if the remote address is inferred from
	// for example the x-forwarder-for header, proxy protocol, etc.
	DownstreamDirectRemoteAddress *v3.Address `protobuf:"bytes,20,opt,name=downstream_direct_remote_address,json=downstreamDirectRemoteAddress,proto3" json:"downstream_direct_remote_address,omitempty"`
	// Map of filter state in stream info that have been configured to be logged. If the filter
	// state serialized to any message other than ``google.protobuf.Any`` it will be packed into
	// ``google.protobuf.Any``.
	FilterStateObjects map[string]*any.Any `protobuf:"bytes,21,rep,name=filter_state_objects,json=filterStateObjects,proto3" json:"filter_state_objects,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of custom tags, which annotate logs with additional information.
	// To configure this value, users should configure
	// :ref:`custom_tags <envoy_v3_api_field_extensions.access_loggers.grpc.v3.CommonGrpcAccessLogConfig.custom_tags>`.
	CustomTags map[string]string `protobuf:"bytes,22,rep,name=custom_tags,json=customTags,proto3" json:"custom_tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// For HTTP: Total duration in milliseconds of the request from the start time to the last byte out.
	// For TCP: Total duration in milliseconds of the downstream connection.
	// This is the total duration of the request (i.e., when the request's ActiveStream is destroyed)
	// and may be longer than ``time_to_last_downstream_tx_byte``.
	Duration *duration.Duration `protobuf:"bytes,23,opt,name=duration,proto3" json:"duration,omitempty"`
	// For HTTP: Number of times the request is attempted upstream. Note that the field is omitted when the request was never attempted upstream.
	// For TCP: Number of times the connection request is attempted upstream. Note that the field is omitted when the connect request was never attempted upstream.
	UpstreamRequestAttemptCount uint32 `protobuf:"varint,24,opt,name=upstream_request_attempt_count,json=upstreamRequestAttemptCount,proto3" json:"upstream_request_attempt_count,omitempty"`
	// Connection termination details may provide additional information about why the connection was terminated by Envoy for L4 reasons.
	ConnectionTerminationDetails string `protobuf:"bytes,25,opt,name=connection_termination_details,json=connectionTerminationDetails,proto3" json:"connection_termination_details,omitempty"`
}
