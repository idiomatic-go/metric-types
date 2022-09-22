package accesslog

// Changes : only allow SocketAddress

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

type SocketAddress_PortValue struct {
	PortValue uint32
}

type SocketAddress struct {
	Protocol SocketAddress_Protocol
	Address  string
	// Types that are assignable to PortSpecifier:
	//	*SocketAddress_PortValue
	//	*SocketAddress_NamedPort
	PortSpecifier SocketAddress_PortValue
	ResolverName  string
	Ipv4Compat    bool
}

func GetProtocolName(s *SocketAddress) string {
	if s != nil {
		return SocketAddress_Protocol_name[int32(s.Protocol)]
	}
	return ""
}

// Only accept SocketAddress
type Address SocketAddress
