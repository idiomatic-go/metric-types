package accesslogv3

type Address_Type int32

const (
	Address_TYPE_UNSPECIFIED Address_Type = 0
	Address_Socket           Address_Type = 1
	Address_Pipe             Address_Type = 2
	Address_EnvoyInternal    Address_Type = 3
)

// Enum value maps for Address_Type.
var (
	Address_Type_name = map[int32]string{
		0: "VERSION_UNSPECIFIED",
		1: "socket",
		2: "pipe",
		3: "envoy-internal",
	}
	Address_Type_value = map[string]int32{
		"VERSION_UNSPECIFIED": 0,
		"socket":              1,
		"pipe":                2,
		"envoy-internal":      3,
	}
)

type Address struct {
	Type                 Address_Type
	Pipe                 *Pipe
	SocketAddress        *SocketAddress
	EnvoyInternalAddress *EnvoyInternalAddress
}
