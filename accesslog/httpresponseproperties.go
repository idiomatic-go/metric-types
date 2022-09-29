package accesslog

// [#next-free-field: 7]
type HTTPResponseProperties struct {
	// The HTTP response code returned by Envoy.
	ResponseCode uint32
	// Size of the HTTP response headers in bytes.
	//
	// This value is captured from the OSI layer 7 perspective, i.e. it does not
	// include overhead from framing or encoding at other networking layers.
	ResponseHeadersBytes uint64
	// Size of the HTTP response body in bytes.
	//
	// This value is captured from the OSI layer 7 perspective, i.e. it does not
	// include overhead from framing or encoding at other networking layers.
	ResponseBodyBytes uint64
	// Map of additional headers configured to be logged.
	ResponseHeaders map[string]string
	// Map of trailers configured to be logged.
	ResponseTrailers map[string]string
	// The HTTP response code details.
	ResponseCodeDetails string
}
