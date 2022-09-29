package accesslog

//https://github.com/envoyproxy/envoy/blob/main/envoy/stream_info/stream_info.h
//https://fossies.org/linux/envoy/source/common/stream_info/utility.h

// Reasons why the request was unauthorized
type ResponseFlags_Unauthorized_Reason int32

func ConvertResponseUnauthorizedReason(reason int32) ResponseFlags_Unauthorized_Reason {
	return ResponseFlags_Unauthorized_Reason(reason)
}

const (
	ResponseFlags_Unauthorized_REASON_UNSPECIFIED ResponseFlags_Unauthorized_Reason = 0
	// The request was denied by the external authorization service.
	ResponseFlags_Unauthorized_EXTERNAL_SERVICE ResponseFlags_Unauthorized_Reason = 1
)

// Enum value maps for ResponseFlags_Unauthorized_Reason.
var (
	ResponseFlags_Unauthorized_Reason_name = map[int32]string{
		0: "REASON_UNSPECIFIED",
		1: "EXTERNAL_SERVICE",
	}
	ResponseFlags_Unauthorized_Reason_value = map[string]int32{
		"REASON_UNSPECIFIED": 0,
		"EXTERNAL_SERVICE":   1,
	}
)

type ResponseFlags_Unauthorized struct {
	Reason ResponseFlags_Unauthorized_Reason `protobuf:"varint,1,opt,name=reason,proto3,enum=envoy.data.accesslog.v3.ResponseFlags_Unauthorized_Reason" json:"reason,omitempty"`
}

// Flags indicating occurrences during request/response processing.
// [#next-free-field: 28]
type ResponseFlags struct {
	// Indicates local server healthcheck failed.
	FailedLocalHealthcheck bool `protobuf:"varint,1,opt,name=failed_local_healthcheck,json=failedLocalHealthcheck,proto3" json:"failed_local_healthcheck,omitempty"`
	// Indicates there was no healthy upstream.
	NoHealthyUpstream bool `protobuf:"varint,2,opt,name=no_healthy_upstream,json=noHealthyUpstream,proto3" json:"no_healthy_upstream,omitempty"`
	// Indicates an there was an upstream request timeout.
	UpstreamRequestTimeout bool `protobuf:"varint,3,opt,name=upstream_request_timeout,json=upstreamRequestTimeout,proto3" json:"upstream_request_timeout,omitempty"`
	// Indicates local codec level reset was sent on the stream.
	LocalReset bool `protobuf:"varint,4,opt,name=local_reset,json=localReset,proto3" json:"local_reset,omitempty"`
	// Indicates remote codec level reset was received on the stream.
	UpstreamRemoteReset bool `protobuf:"varint,5,opt,name=upstream_remote_reset,json=upstreamRemoteReset,proto3" json:"upstream_remote_reset,omitempty"`
	// Indicates there was a local reset by a connection pool due to an initial connection failure.
	UpstreamConnectionFailure bool `protobuf:"varint,6,opt,name=upstream_connection_failure,json=upstreamConnectionFailure,proto3" json:"upstream_connection_failure,omitempty"`
	// Indicates the stream was reset due to an upstream connection termination.
	UpstreamConnectionTermination bool `protobuf:"varint,7,opt,name=upstream_connection_termination,json=upstreamConnectionTermination,proto3" json:"upstream_connection_termination,omitempty"`
	// Indicates the stream was reset because of a resource overflow.
	UpstreamOverflow bool `protobuf:"varint,8,opt,name=upstream_overflow,json=upstreamOverflow,proto3" json:"upstream_overflow,omitempty"`
	// Indicates no route was found for the request.
	NoRouteFound bool `protobuf:"varint,9,opt,name=no_route_found,json=noRouteFound,proto3" json:"no_route_found,omitempty"`
	// Indicates that the request was delayed before proxying.
	DelayInjected bool `protobuf:"varint,10,opt,name=delay_injected,json=delayInjected,proto3" json:"delay_injected,omitempty"`
	// Indicates that the request was aborted with an injected error code.
	FaultInjected bool `protobuf:"varint,11,opt,name=fault_injected,json=faultInjected,proto3" json:"fault_injected,omitempty"`
	// Indicates that the request was rate-limited locally.
	RateLimited bool `protobuf:"varint,12,opt,name=rate_limited,json=rateLimited,proto3" json:"rate_limited,omitempty"`
	// Indicates if the request was deemed unauthorized and the reason for it.
	UnauthorizedDetails *ResponseFlags_Unauthorized `protobuf:"bytes,13,opt,name=unauthorized_details,json=unauthorizedDetails,proto3" json:"unauthorized_details,omitempty"`
	// Indicates that the request was rejected because there was an error in rate limit service.
	RateLimitServiceError bool `protobuf:"varint,14,opt,name=rate_limit_service_error,json=rateLimitServiceError,proto3" json:"rate_limit_service_error,omitempty"`
	// Indicates the stream was reset due to a downstream connection termination.
	DownstreamConnectionTermination bool `protobuf:"varint,15,opt,name=downstream_connection_termination,json=downstreamConnectionTermination,proto3" json:"downstream_connection_termination,omitempty"`
	// Indicates that the upstream retry limit was exceeded, resulting in a downstream error.
	UpstreamRetryLimitExceeded bool `protobuf:"varint,16,opt,name=upstream_retry_limit_exceeded,json=upstreamRetryLimitExceeded,proto3" json:"upstream_retry_limit_exceeded,omitempty"`
	// Indicates that the stream idle timeout was hit, resulting in a downstream 408.
	StreamIdleTimeout bool `protobuf:"varint,17,opt,name=stream_idle_timeout,json=streamIdleTimeout,proto3" json:"stream_idle_timeout,omitempty"`
	// Indicates that the request was rejected because an envoy request header failed strict
	// validation.
	InvalidEnvoyRequestHeaders bool `protobuf:"varint,18,opt,name=invalid_envoy_request_headers,json=invalidEnvoyRequestHeaders,proto3" json:"invalid_envoy_request_headers,omitempty"`
	// Indicates there was an HTTP protocol error on the downstream request.
	DownstreamProtocolError bool `protobuf:"varint,19,opt,name=downstream_protocol_error,json=downstreamProtocolError,proto3" json:"downstream_protocol_error,omitempty"`
	// Indicates there was a max stream duration reached on the upstream request.
	UpstreamMaxStreamDurationReached bool `protobuf:"varint,20,opt,name=upstream_max_stream_duration_reached,json=upstreamMaxStreamDurationReached,proto3" json:"upstream_max_stream_duration_reached,omitempty"`
	// Indicates the response was served from a cache filter.
	ResponseFromCacheFilter bool `protobuf:"varint,21,opt,name=response_from_cache_filter,json=responseFromCacheFilter,proto3" json:"response_from_cache_filter,omitempty"`
	// Indicates that a filter configuration is not available.
	NoFilterConfigFound bool `protobuf:"varint,22,opt,name=no_filter_config_found,json=noFilterConfigFound,proto3" json:"no_filter_config_found,omitempty"`
	// Indicates that request or connection exceeded the downstream connection duration.
	DurationTimeout bool `protobuf:"varint,23,opt,name=duration_timeout,json=durationTimeout,proto3" json:"duration_timeout,omitempty"`
	// Indicates there was an HTTP protocol error in the upstream response.
	UpstreamProtocolError bool `protobuf:"varint,24,opt,name=upstream_protocol_error,json=upstreamProtocolError,proto3" json:"upstream_protocol_error,omitempty"`
	// Indicates no cluster was found for the request.
	NoClusterFound bool `protobuf:"varint,25,opt,name=no_cluster_found,json=noClusterFound,proto3" json:"no_cluster_found,omitempty"`
	// Indicates overload manager terminated the request.
	OverloadManager bool `protobuf:"varint,26,opt,name=overload_manager,json=overloadManager,proto3" json:"overload_manager,omitempty"`
	// Indicates a DNS resolution failed.
	DnsResolutionFailure bool `protobuf:"varint,27,opt,name=dns_resolution_failure,json=dnsResolutionFailure,proto3" json:"dns_resolution_failure,omitempty"`
}
