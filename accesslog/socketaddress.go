package accesslog

type SocketAddress_Protocol int32

func ConvertSocketAddressProtocol(protocol int) SocketAddress_Protocol {
	return SocketAddress_Protocol(protocol)
}

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

func GetProtocolName(addr SocketAddress_Protocol) string {
	if addr != 0 || addr != 1 {
		return ""
	}
	return SocketAddress_Protocol_name[int32(addr)]
}

type SocketAddress struct {
	Protocol      string // TCP,UDP
	Address       string // IP address
	ResolverName  string // Envoy configuration
	PortSpecifier string // uint32 PortValue, or string NamedPort
	Ipv4Compat    bool
}
