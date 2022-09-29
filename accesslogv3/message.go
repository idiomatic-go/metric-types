package accesslogv3

// OriginIdentifier contains the data needed to uniquely identify a log entries
// Do we really need the node Id ?
type OriginIdentifier struct {
	DataCenter  string
	Zone        string
	Application string
	InstanceId  string
	NodeId      string
}

type CombinedEntry struct {
	Common *AccessLogCommon
	Http   *HTTPAccessLogEntry
	Tcp    *TCPAccessLogEntry
}

type Message struct {
	Origin OriginIdentifier
	// Batches of log entries of a single type. Generally speaking, a given stream should only
	// ever include one type of log entry.
	//
	// Types that are assignable to LogEntries:
	//	*StreamAccessLogsMessage_HttpLogs
	//	*StreamAccessLogsMessage_TcpLogs
	// HttpLogs
	LogEntries []*CombinedEntry
}
