package accesslog

const (
	SocketAddressType   = "socket"
	PipeAddressType     = "pipe"
	InternalAddressType = "envoy-internal"
)

type Address struct {
	Type                 string
	Pipe                 *Pipe
	SocketAddress        *SocketAddress
	EnvoyInternalAddress *EnvoyInternalAddress
}
