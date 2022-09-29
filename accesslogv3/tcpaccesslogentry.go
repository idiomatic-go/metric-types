package accesslogv3

type ConnectionProperties struct {
	// Number of bytes received from downstream.
	ReceivedBytes uint64
	// Number of bytes sent to downstream.
	SentBytes uint64
}

type TCPAccessLogEntry struct {
	// Common properties shared by all Envoy access logs.
	//CommonProperties *Common
	// Properties of the TCP connection.
	ConnectionProperties *ConnectionProperties
}
