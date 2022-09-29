package accesslogv3

//https://github.com/envoyproxy/envoy/blob/main/envoy/stream_info/stream_info.h
//https://fossies.org/linux/envoy/source/common/stream_info/utility.h

import (
	v31 "github.com/envoyproxy/go-control-plane/envoy/data/accesslog/v3"
	"strings"
)

const (
	NONE                                 = "-"
	DOWNSTREAM_CONNECTION_TERMINATION    = "DC"
	FAILED_LOCAL_HEALTH_CHECK            = "LH"
	NO_HEALTHY_UPSTREAM                  = "UH"
	UPSTREAM_REQUEST_TIMEOUT             = "UT"
	LOCAL_RESET                          = "LR"
	UPSTREAM_REMOTE_RESET                = "UR"
	UPSTREAM_CONNECTION_FAILURE          = "UF"
	UPSTREAM_CONNECTION_TERMINATION      = "UC"
	UPSTREAM_OVERFLOW                    = "UO"
	UPSTREAM_RETRY_LIMIT_EXCEEDED        = "URX"
	NO_ROUTE_FOUND                       = "NR"
	DELAY_INJECTED                       = "DI"
	FAULT_INJECTED                       = "FI"
	RATE_LIMITED                         = "RL"
	UNAUTHORIZED_EXTERNAL_SERVICE        = "UAEX"
	RATELIMIT_SERVICE_ERROR              = "RLSE"
	STREAM_IDLE_TIMEOUT                  = "SI"
	INVALID_ENVOY_REQUEST_HEADERS        = "IH"
	DOWNSTREAM_PROTOCOL_ERROR            = "DPE"
	UPSTREAM_MAX_STREAM_DURATION_REACHED = "UMSDR"
	RESPONSE_FROM_CACHE_FILTER           = "RFCF"
	NO_FILTER_CONFIG_FOUND               = "NFCF"
	DURATION_TIMEOUT                     = "DT"
	UPSTREAM_PROTOCOL_ERROR              = "UPE"
	NO_CLUSTER_FOUND                     = "NC"
	OVERLOAD_MANAGER                     = "OM"
	DNS_FAIL                             = "DF"
)

/*
52   static constexpr std::array ALL_RESPONSE_STRING_FLAGS{
53       FlagStringAndEnum{FAILED_LOCAL_HEALTH_CHECK, ResponseFlag::FailedLocalHealthCheck},
54       FlagStringAndEnum{NO_HEALTHY_UPSTREAM, ResponseFlag::NoHealthyUpstream},
55       FlagStringAndEnum{UPSTREAM_REQUEST_TIMEOUT, ResponseFlag::UpstreamRequestTimeout},
56       FlagStringAndEnum{LOCAL_RESET, ResponseFlag::LocalReset},
57       FlagStringAndEnum{UPSTREAM_REMOTE_RESET, ResponseFlag::UpstreamRemoteReset},
58       FlagStringAndEnum{UPSTREAM_CONNECTION_FAILURE, ResponseFlag::UpstreamConnectionFailure},
59       FlagStringAndEnum{UPSTREAM_CONNECTION_TERMINATION,
60                         ResponseFlag::UpstreamConnectionTermination},
61       FlagStringAndEnum{UPSTREAM_OVERFLOW, ResponseFlag::UpstreamOverflow},
62       FlagStringAndEnum{NO_ROUTE_FOUND, ResponseFlag::NoRouteFound},
63       FlagStringAndEnum{DELAY_INJECTED, ResponseFlag::DelayInjected},
64       FlagStringAndEnum{FAULT_INJECTED, ResponseFlag::FaultInjected},
65       FlagStringAndEnum{RATE_LIMITED, ResponseFlag::RateLimited},
66       FlagStringAndEnum{UNAUTHORIZED_EXTERNAL_SERVICE, ResponseFlag::UnauthorizedExternalService},
67       FlagStringAndEnum{RATELIMIT_SERVICE_ERROR, ResponseFlag::RateLimitServiceError},
68       FlagStringAndEnum{DOWNSTREAM_CONNECTION_TERMINATION,
69                         ResponseFlag::DownstreamConnectionTermination},
70       FlagStringAndEnum{UPSTREAM_RETRY_LIMIT_EXCEEDED, ResponseFlag::UpstreamRetryLimitExceeded},
71       FlagStringAndEnum{STREAM_IDLE_TIMEOUT, ResponseFlag::StreamIdleTimeout},
72       FlagStringAndEnum{INVALID_ENVOY_REQUEST_HEADERS, ResponseFlag::InvalidEnvoyRequestHeaders},
73       FlagStringAndEnum{DOWNSTREAM_PROTOCOL_ERROR, ResponseFlag::DownstreamProtocolError},
74       FlagStringAndEnum{UPSTREAM_MAX_STREAM_DURATION_REACHED,
75                         ResponseFlag::UpstreamMaxStreamDurationReached},
76       FlagStringAndEnum{RESPONSE_FROM_CACHE_FILTER, ResponseFlag::ResponseFromCacheFilter},
77       FlagStringAndEnum{NO_FILTER_CONFIG_FOUND, ResponseFlag::NoFilterConfigFound},
78       FlagStringAndEnum{DURATION_TIMEOUT, ResponseFlag::DurationTimeout},
79       FlagStringAndEnum{UPSTREAM_PROTOCOL_ERROR, ResponseFlag::UpstreamProtocolError},
80       FlagStringAndEnum{NO_CLUSTER_FOUND, ResponseFlag::NoClusterFound},
81       FlagStringAndEnum{OVERLOAD_MANAGER, ResponseFlag::OverloadManager},
82   };
83

*/

const (
	// Local server healthcheck failed.
	FailedLocalHealthCheck = ResponseFlags_Encoded(0x1)
	// No healthy upstream.
	NoHealthyUpstream = ResponseFlags_Encoded(0x2)
	// Request timeout on upstream.
	UpstreamRequestTimeout = ResponseFlags_Encoded(0x4)
	// Local codec level reset was sent on the stream.
	LocalReset = ResponseFlags_Encoded(0x8)
	// Remote codec level reset was received on the stream.
	UpstreamRemoteReset = ResponseFlags_Encoded(0x10)
	// Local reset by a connection pool due to an initial connection failure.
	UpstreamConnectionFailure = ResponseFlags_Encoded(0x20)
	// If the stream was locally reset due to connection termination.
	UpstreamConnectionTermination = ResponseFlags_Encoded(0x40)
	// The stream was reset because of a resource overflow.
	UpstreamOverflow = ResponseFlags_Encoded(0x80)
	// No route found for a given request.
	NoRouteFound = ResponseFlags_Encoded(0x100)
	// Request was delayed before proxying.
	DelayInjected = ResponseFlags_Encoded(0x200)
	// Abort with error code was injected.
	FaultInjected = ResponseFlags_Encoded(0x400)
	// Request was ratelimited locally by rate limit filter.
	RateLimited = ResponseFlags_Encoded(0x800)
	// Request was unauthorized by external authorization service.
	UnauthorizedExternalService = ResponseFlags_Encoded(0x1000)
	// Unable to call Ratelimit service.
	RateLimitServiceError = ResponseFlags_Encoded(0x2000)
	// If the stream was reset due to a downstream connection termination.
	DownstreamConnectionTermination = ResponseFlags_Encoded(0x4000)
	// Exceeded upstream retry limit.
	UpstreamRetryLimitExceeded = ResponseFlags_Encoded(0x8000)
	// Request hit the stream idle timeout, triggering a 408.
	StreamIdleTimeout = ResponseFlags_Encoded(0x10000)
	// Request specified x-envoy-* header values that failed strict header checks.
	InvalidEnvoyRequestHeaders = ResponseFlags_Encoded(0x20000)
	// Downstream request had an HTTP protocol error
	DownstreamProtocolError = ResponseFlags_Encoded(0x40000)
	// Upstream request reached to user defined max stream duration.
	UpstreamMaxStreamDurationReached = ResponseFlags_Encoded(0x80000)
	// True if the response was served from an Envoy cache filter.
	ResponseFromCacheFilter = ResponseFlags_Encoded(0x100000)
	// Filter config was not received within the permitted warming deadline.
	NoFilterConfigFound = ResponseFlags_Encoded(0x200000)
	// Request or connection exceeded the downstream connection duration.
	DurationTimeout = ResponseFlags_Encoded(0x400000)
	// Upstream response had an HTTP protocol error
	UpstreamProtocolError = ResponseFlags_Encoded(0x800000)
	// No cluster found for a given request.
	NoClusterFound = ResponseFlags_Encoded(0x1000000)
	// Overload Manager terminated the stream.
	OverloadManager = ResponseFlags_Encoded(0x2000000)
	// DNS resolution failed.
	DnsResolutionFailed = ResponseFlags_Encoded(0x4000000)
	// ATTENTION: MAKE SURE THIS REMAINS EQUAL TO THE LAST FLAG.
	LastFlag = DnsResolutionFailed
)

func GetResponseFlags(resp *ResponseFlags) (primary string, secondary string) {
	var flags string

	if resp == nil {
		return primary, secondary
	}

	// Start with response code generating flags, 404.429.503.504

	// Indicates there was an upstream request timeout.
	if resp.Encoded&UpstreamRequestTimeout == 1 {
		flags = addFlag(flags, UPSTREAM_REQUEST_TIMEOUT)
	}

	// Indicates there was no healthy upstream.
	if resp.Encoded&NoHealthyUpstream == 1 {
		flags = addFlag(flags, NO_HEALTHY_UPSTREAM)
	}
	// Indicates there was a local reset by a connection pool due to an initial connection failure.
	if resp.Encoded&UpstreamConnectionFailure == 1 {
		flags = addFlag(flags, UPSTREAM_CONNECTION_FAILURE)
	}
	// Indicates remote codec level reset was received on the stream.
	if resp.Encoded&UpstreamRemoteReset == 1 {
		flags = addFlag(flags, UPSTREAM_REMOTE_RESET)
	}
	// Indicates the stream was reset due to an upstream connection termination.
	if resp.Encoded&UpstreamConnectionTermination == 1 {
		flags = addFlag(flags, UPSTREAM_CONNECTION_TERMINATION)
	}
	// Indicates the stream was reset because of a resource overflow.
	if resp.Encoded&UpstreamOverflow == 1 {
		flags = addFlag(flags, UPSTREAM_OVERFLOW)
	}
	// Indicates no route was found for the request.
	if resp.Encoded&NoRouteFound == 1 {
		flags = addFlag(flags, NO_ROUTE_FOUND)
	}
	// Indicates that the request was rate-limited locally.
	if resp.Encoded&RateLimited == 1 {
		flags = addFlag(flags, RATE_LIMITED)
	}

	// Next, finish upstream errors
	// Indicates that the upstream retry limit was exceeded, resulting in a downstream error.
	if resp.Encoded&UpstreamRetryLimitExceeded == 1 {
		flags = addFlag(flags, UPSTREAM_RETRY_LIMIT_EXCEEDED)
	}
	// Indicates no cluster was found for the request.
	if resp.Encoded&NoClusterFound == 1 {
		flags = addFlag(flags, NO_CLUSTER_FOUND)
	}

	// Next, local/downstream connection

	// Indicates local codec level reset was sent on the stream.
	if resp.Encoded&LocalReset == 1 {
		flags = addFlag(flags, LOCAL_RESET)
	}
	// Indicates the stream was reset due to a downstream connection termination.
	if resp.Encoded&DownstreamConnectionTermination == 1 {
		flags = addFlag(flags, DOWNSTREAM_CONNECTION_TERMINATION)
	}
	// Indicates local server healthcheck failed.
	if resp.Encoded&FailedLocalHealthCheck == 1 {
		flags = addFlag(flags, FAILED_LOCAL_HEALTH_CHECK)
	}
	// Indicates that the stream idle timeout was hit, resulting in a downstream 408.
	if resp.Encoded&StreamIdleTimeout == 1 {
		flags = addFlag(flags, STREAM_IDLE_TIMEOUT)
	}

	// Add UAEX
	if resp.UnauthorizedDetails != nil {
		if resp.UnauthorizedDetails.GetReason() == v31.ResponseFlags_Unauthorized_EXTERNAL_SERVICE {
			flags = addFlag(flags, UNAUTHORIZED_EXTERNAL_SERVICE)
		}
	}

	// Indicates that the request was delayed before proxying.
	if resp.Encoded&DelayInjected == 1 {
		flags = addFlag(flags, DELAY_INJECTED)
	}
	// Indicates that the request was aborted with an injected error code.
	if resp.Encoded&FaultInjected == 1 {
		flags = addFlag(flags, FAULT_INJECTED)
	}
	// Indicates if the request was deemed unauthorized and the reason for it.
	//UnauthorizedDetails *ResponseFlags_Unauthorized
	// Indicates that the request was rejected because there was an error in rate limit service.
	if resp.Encoded&RateLimitServiceError == 1 {
		flags = addFlag(flags, RATELIMIT_SERVICE_ERROR)
	}
	// Indicates that the request was rejected because an envoy request header failed strict
	// validation.
	if resp.Encoded&InvalidEnvoyRequestHeaders == 1 {
		flags = addFlag(flags, INVALID_ENVOY_REQUEST_HEADERS)
	}
	// Indicates there was an HTTP protocol error on the downstream request.
	if resp.Encoded&DownstreamProtocolError == 1 {
		flags = addFlag(flags, DOWNSTREAM_PROTOCOL_ERROR)
	}
	// Indicates there was a max stream duration reached on the upstream request.
	if resp.Encoded&UpstreamMaxStreamDurationReached == 1 {
		flags = addFlag(flags, UPSTREAM_MAX_STREAM_DURATION_REACHED)
	}
	// Indicates the response was served from a cache filter.
	if resp.Encoded&ResponseFromCacheFilter == 1 {
		flags = addFlag(flags, RESPONSE_FROM_CACHE_FILTER)
	}
	// Indicates that a filter configuration is not available.
	if resp.Encoded&NoFilterConfigFound == 1 {
		flags = addFlag(flags, NO_FILTER_CONFIG_FOUND)
	}
	// Indicates that request or connection exceeded the downstream connection duration.
	if resp.Encoded&DurationTimeout == 1 {
		flags = addFlag(flags, DURATION_TIMEOUT)
	}
	// Indicates there was an HTTP protocol error in the upstream response.
	if resp.Encoded&UpstreamProtocolError == 1 {
		flags = addFlag(flags, UPSTREAM_PROTOCOL_ERROR)
	}
	// Indicates overload manager terminated the request.
	if resp.Encoded&OverloadManager == 1 {
		flags = addFlag(flags, OVERLOAD_MANAGER)
	}
	// Indicates a DNS resolution failed.
	if resp.Encoded&DnsResolutionFailed == 1 {
		flags = addFlag(flags, DNS_FAIL)
	}

	if flags == "" {
		return primary, secondary
	}
	// Put first flag in ResponseFlag, the rest of the flags will be a list on the ext string
	ls := strings.Split(flags, ",")
	if len(ls) > 0 {
		for i, s := range ls {
			if i == 0 {
				primary = s
			} else {
				secondary = addFlag(secondary, s)
			}
		}
	}

	return primary, secondary
}

func addFlag(flags string, token string) string {
	if flags != "" {
		flags += ","
	}
	flags += token
	return flags
}

func EncodeResponseFlags(resp *v31.ResponseFlags) ResponseFlags_Encoded {
	var flags ResponseFlags_Encoded

	// Start with response code generating flags, 404.429.503.504

	// Indicates there was an upstream request timeout.
	if resp.GetUpstreamRequestTimeout() {
		flags |= UpstreamRequestTimeout
	}
	// Indicates there was no healthy upstream.
	if resp.GetNoHealthyUpstream() {
		flags |= NoHealthyUpstream
	}
	// Indicates there was a local reset by a connection pool due to an initial connection failure.
	if resp.GetUpstreamConnectionFailure() {
		flags |= UpstreamConnectionFailure
	}
	// Indicates remote codec level reset was received on the stream.
	if resp.GetUpstreamRemoteReset() {
		flags |= UpstreamRemoteReset
	}
	// Indicates the stream was reset due to an upstream connection termination.
	if resp.GetUpstreamConnectionTermination() {
		flags |= UpstreamConnectionTermination
	}
	// Indicates the stream was reset because of a resource overflow.
	if resp.GetUpstreamOverflow() {
		flags |= UpstreamOverflow
	}
	// Indicates no route was found for the request.
	if resp.GetNoRouteFound() {
		flags |= NoRouteFound
	}
	// Indicates that the request was rate-limited locally.
	if resp.GetRateLimited() {
		flags |= RateLimited
	}

	// Next, finish upstream errors
	// Indicates that the upstream retry limit was exceeded, resulting in a downstream error.
	if resp.GetUpstreamRetryLimitExceeded() {
		flags |= UpstreamRetryLimitExceeded
	}
	// Indicates no cluster was found for the request.
	if resp.GetNoClusterFound() {
		flags |= NoClusterFound
	}
	// Next, local/downstream connection

	// Indicates local codec level reset was sent on the stream.
	if resp.GetLocalReset() {
		flags |= LocalReset
	}
	// Indicates the stream was reset due to a downstream connection termination.
	if resp.GetDownstreamConnectionTermination() {
		flags |= DownstreamConnectionTermination
	}
	// Indicates local server healthcheck failed.
	if resp.GetFailedLocalHealthcheck() {
		flags |= FailedLocalHealthCheck
	}
	// Indicates that the stream idle timeout was hit, resulting in a downstream 408.
	if resp.GetStreamIdleTimeout() {
		flags |= StreamIdleTimeout
	}

	// Add UAEX

	// Indicates that the request was delayed before proxying.
	if resp.GetDelayInjected() {
		flags |= DelayInjected
	}
	// Indicates that the request was aborted with an injected error code.
	if resp.GetFaultInjected() {
		flags |= FaultInjected
	}
	// Indicates if the request was deemed unauthorized and the reason for it.
	//UnauthorizedDetails *ResponseFlags_Unauthorized
	// Indicates that the request was rejected because there was an error in rate limit service.
	if resp.GetRateLimitServiceError() {
		flags |= RateLimitServiceError
	}
	// Indicates that the request was rejected because an envoy request header failed strict
	// validation.
	if resp.GetInvalidEnvoyRequestHeaders() {
		flags |= InvalidEnvoyRequestHeaders
	}
	// Indicates there was an HTTP protocol error on the downstream request.
	if resp.GetDownstreamProtocolError() {
		flags |= DownstreamProtocolError
	}
	// Indicates there was a max stream duration reached on the upstream request.
	if resp.GetUpstreamMaxStreamDurationReached() {
		flags |= UpstreamMaxStreamDurationReached
	}
	// Indicates the response was served from a cache filter.
	if resp.GetResponseFromCacheFilter() {
		flags |= ResponseFromCacheFilter
	}
	// Indicates that a filter configuration is not available.
	if resp.GetNoFilterConfigFound() {
		flags |= NoFilterConfigFound
	}
	// Indicates that request or connection exceeded the downstream connection duration.
	if resp.GetDurationTimeout() {
		flags |= DurationTimeout
	}
	// Indicates there was an HTTP protocol error in the upstream response.
	if resp.GetUpstreamProtocolError() {
		flags |= UpstreamProtocolError
	}
	// Indicates overload manager terminated the request.
	if resp.GetOverloadManager() {
		flags |= OverloadManager
	}
	// Indicates a DNS resolution failed.
	if resp.GetDnsResolutionFailure() {
		flags |= DnsResolutionFailed
	}
	return flags
}
