package accesslogv3

type Pipe struct {
	// Unix Domain Socket path. On Linux, paths starting with '@' will use the
	// abstract namespace. The starting '@' is replaced by a null byte by Envoy.
	// Paths starting with '@' will result in an error in environments other than
	// Linux.
	Path string
	// The mode for the Pipe. Not applicable for abstract sockets.
	Mode uint32
}
