package accesslogv3

type SocketAddress_Protocol int32

const (
	SocketAddress_TCP SocketAddress_Protocol = 0
	SocketAddress_UDP SocketAddress_Protocol = 1
)

// Enum value maps for SocketAddress_Protocol.
var (
	SocketAddress_Protocol_name = map[int32]string{
		0: "TCP",
		1: "UDP",
	}
	SocketAddress_Protocol_value = map[string]int32{
		"TCP": 0,
		"UDP": 1,
	}
)

type SocketAddress struct {
	Protocol      SocketAddress_Protocol // TCP,UDP
	Address       string                 // IP address
	ResolverName  string                 // Envoy configuration
	PortSpecifier string                 // uint32 PortValue, or string NamedPort
	Ipv4Compat    bool
}
