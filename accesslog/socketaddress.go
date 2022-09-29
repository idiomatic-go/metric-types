package accesslog

type SocketAddress struct {
	Protocol      string // TCP,UDP
	Address       string // IP address
	ResolverName  string // Envoy configuration
	PortSpecifier string // uint32 PortValue, or string NamedPort
	Ipv4Compat    bool
}
