package accesslog

type HTTPAccessLogEntry struct {
	// Common properties shared by all Envoy access logs.
	//CommonProperties *Common
	ProtocolVersion string
	// Description of the incoming HTTP request.
	Request *HTTPRequestProperties
	// Description of the outgoing HTTP response.
	Response *HTTPResponseProperties
}
