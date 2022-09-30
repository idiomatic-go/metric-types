package accesslogv3

import "google.golang.org/protobuf/runtime/protoimpl"

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

type isSocketAddress_PortSpecifier interface {
	isSocketAddress_PortSpecifier()
}

type SocketAddress_PortValue struct {
	PortValue uint32 `protobuf:"varint,3,opt,name=port_value,json=portValue,proto3,oneof"`
}

type SocketAddress_NamedPort struct {
	// This is only valid if :ref:`resolver_name
	// <envoy_v3_api_field_config.core.v3.SocketAddress.resolver_name>` is specified below and the
	// named resolver is capable of named port resolution.
	NamedPort string `protobuf:"bytes,4,opt,name=named_port,json=namedPort,proto3,oneof"`
}

type SocketAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol SocketAddress_Protocol `protobuf:"varint,1,opt,name=protocol,proto3,enum=envoy.config.core.v3.SocketAddress_Protocol" json:"protocol,omitempty"`
	// The address for this socket. :ref:`Listeners <config_listeners>` will bind
	// to the address. An empty address is not allowed. Specify ``0.0.0.0`` or ``::``
	// to bind to any address. [#comment:TODO(zuercher) reinstate when implemented:
	// It is possible to distinguish a Listener address via the prefix/suffix matching
	// in :ref:`FilterChainMatch <envoy_v3_api_msg_config.listener.v3.FilterChainMatch>`.] When used
	// within an upstream :ref:`BindConfig <envoy_v3_api_msg_config.core.v3.BindConfig>`, the address
	// controls the source address of outbound connections. For :ref:`clusters
	// <envoy_v3_api_msg_config.cluster.v3.Cluster>`, the cluster type determines whether the
	// address must be an IP (``STATIC`` or ``EDS`` clusters) or a hostname resolved by DNS
	// (``STRICT_DNS`` or ``LOGICAL_DNS`` clusters). Address resolution can be customized
	// via :ref:`resolver_name <envoy_v3_api_field_config.core.v3.SocketAddress.resolver_name>`.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are assignable to PortSpecifier:
	//	*SocketAddress_PortValue
	//	*SocketAddress_NamedPort
	PortSpecifier isSocketAddress_PortSpecifier `protobuf_oneof:"port_specifier"`
	// The name of the custom resolver. This must have been registered with Envoy. If
	// this is empty, a context dependent default applies. If the address is a concrete
	// IP address, no resolution will occur. If address is a hostname this
	// should be set for resolution other than DNS. Specifying a custom resolver with
	// ``STRICT_DNS`` or ``LOGICAL_DNS`` will generate an error at runtime.
	ResolverName string `protobuf:"bytes,5,opt,name=resolver_name,json=resolverName,proto3" json:"resolver_name,omitempty"`
	// When binding to an IPv6 address above, this enables `IPv4 compatibility
	// <https://tools.ietf.org/html/rfc3493#page-11>`_. Binding to ``::`` will
	// allow both IPv4 and IPv6 connections, with peer IPv4 addresses mapped into
	// IPv6 space as ``::FFFF:<IPv4-address>``.
	Ipv4Compat bool `protobuf:"varint,6,opt,name=ipv4_compat,json=ipv4Compat,proto3" json:"ipv4_compat,omitempty"`
}

type Address_SocketAddress struct {
	SocketAddress *SocketAddress `protobuf:"bytes,1,opt,name=socket_address,json=socketAddress,proto3,oneof"`
}

type Pipe struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unix Domain Socket path. On Linux, paths starting with '@' will use the
	// abstract namespace. The starting '@' is replaced by a null byte by Envoy.
	// Paths starting with '@' will result in an error in environments other than
	// Linux.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// The mode for the Pipe. Not applicable for abstract sockets.
	Mode uint32 `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
}

type Address_Pipe struct {
	Pipe *Pipe `protobuf:"bytes,2,opt,name=pipe,proto3,oneof"`
}

type isEnvoyInternalAddress_AddressNameSpecifier interface {
	isEnvoyInternalAddress_AddressNameSpecifier()
}

type EnvoyInternalAddress_ServerListenerName struct {
	// Specifies the :ref:`name <envoy_v3_api_field_config.listener.v3.Listener.name>` of the
	// internal listener.
	ServerListenerName string `protobuf:"bytes,1,opt,name=server_listener_name,json=serverListenerName,proto3,oneof"`
}

// The address represents an envoy internal listener.
// [#comment: TODO(asraa): When address available, remove workaround from test/server/server_fuzz_test.cc:30.]
type EnvoyInternalAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AddressNameSpecifier:
	//	*EnvoyInternalAddress_ServerListenerName
	AddressNameSpecifier isEnvoyInternalAddress_AddressNameSpecifier `protobuf_oneof:"address_name_specifier"`
	// Specifies an endpoint identifier to distinguish between multiple endpoints for the same internal listener in a
	// single upstream pool. Only used in the upstream addresses for tracking changes to individual endpoints. This, for
	// example, may be set to the final destination IP for the target internal listener.
	EndpointId string `protobuf:"bytes,2,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
}

type Address_EnvoyInternalAddress struct {
	// Specifies a user-space address handled by :ref:`internal listeners
	// <envoy_v3_api_field_config.listener.v3.Listener.internal_listener>`.
	EnvoyInternalAddress *EnvoyInternalAddress `protobuf:"bytes,3,opt,name=envoy_internal_address,json=envoyInternalAddress,proto3,oneof"`
}

type isAddress_Address interface {
	isAddress_Address()
}

// Addresses specify either a logical or physical address and port, which are
// used to tell Envoy where to bind/listen, connect to upstream and find
// management servers.
type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Address:
	//	*Address_SocketAddress
	//	*Address_Pipe
	//	*Address_EnvoyInternalAddress
	Address isAddress_Address `protobuf_oneof:"address"`
}
