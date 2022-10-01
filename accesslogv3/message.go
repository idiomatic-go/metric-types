package accesslogv3

import (
	"github.com/idiomatic-go/common-lib/util"
)

// Provenance contains the attributes needed to identify a log entry's place of origin, down to the instance
type Provenance struct {
	Locality    *Locality // Region (data center), Zone, SubZone (cluster)
	Application string    // (service name), (container name), (uri host)
	InstanceId  string    // (pod), (node), (physical host) identifier
}

type CombinedEntry struct {
	Common *CommonProperties
	Http   *HTTPAccessLogEntry
	Tcp    *TCPAccessLogEntry
}

type Message struct {
	Identifier *Provenance
	Dictionary *util.InvertedDictionary
	LogEntries []*CombinedEntry
}
