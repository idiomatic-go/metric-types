package accesslog

type HTTPAccessLogEntry_HTTPVersion int32

func ConvertHttpVersion(vers int) HTTPAccessLogEntry_HTTPVersion {
	return HTTPAccessLogEntry_HTTPVersion(vers)
}

type HTTPAccessLogEntry struct {
	// Common properties shared by all Envoy access logs.
	//CommonProperties *Common
	ProtocolVersion HTTPAccessLogEntry_HTTPVersion
	// Description of the incoming HTTP request.
	Request *HTTPRequestProperties
	// Description of the outgoing HTTP response.
	Response *HTTPResponseProperties
}
