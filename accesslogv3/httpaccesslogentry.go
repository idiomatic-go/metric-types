package accesslogv3

type HTTPAccessLogEntry_HTTPVersion int32

const (
	HTTPAccessLogEntry_PROTOCOL_UNSPECIFIED HTTPAccessLogEntry_HTTPVersion = 0
	HTTPAccessLogEntry_HTTP10               HTTPAccessLogEntry_HTTPVersion = 1
	HTTPAccessLogEntry_HTTP11               HTTPAccessLogEntry_HTTPVersion = 2
	HTTPAccessLogEntry_HTTP2                HTTPAccessLogEntry_HTTPVersion = 3
	HTTPAccessLogEntry_HTTP3                HTTPAccessLogEntry_HTTPVersion = 4
)

// Enum value maps for HTTPAccessLogEntry_HTTPVersion.
var (
	HTTPAccessLogEntry_HTTPVersion_name = map[int32]string{
		0: "PROTOCOL_UNSPECIFIED",
		1: "HTTP10",
		2: "HTTP11",
		3: "HTTP2",
		4: "HTTP3",
	}
	HTTPAccessLogEntry_HTTPVersion_value = map[string]int32{
		"PROTOCOL_UNSPECIFIED": 0,
		"HTTP10":               1,
		"HTTP11":               2,
		"HTTP2":                3,
		"HTTP3":                4,
	}
)

type HTTPAccessLogEntry struct {
	// Common properties shared by all Envoy access logs.
	//CommonProperties *Common
	ProtocolVersion HTTPAccessLogEntry_HTTPVersion
	// Description of the incoming HTTP request.
	Request *HTTPRequestProperties
	// Description of the outgoing HTTP response.
	Response *HTTPResponseProperties
}
