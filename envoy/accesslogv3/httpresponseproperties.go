package accesslogv3

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// [#next-free-field: 7]
type HTTPResponseProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The HTTP response code returned by Envoy.
	ResponseCode *wrappers.UInt32Value `protobuf:"bytes,1,opt,name=response_code,json=responseCode,proto3" json:"response_code,omitempty"`
	// Size of the HTTP response headers in bytes.
	//
	// This value is captured from the OSI layer 7 perspective, i.e. it does not
	// include overhead from framing or encoding at other networking layers.
	ResponseHeadersBytes uint64 `protobuf:"varint,2,opt,name=response_headers_bytes,json=responseHeadersBytes,proto3" json:"response_headers_bytes,omitempty"`
	// Size of the HTTP response body in bytes.
	//
	// This value is captured from the OSI layer 7 perspective, i.e. it does not
	// include overhead from framing or encoding at other networking layers.
	ResponseBodyBytes uint64 `protobuf:"varint,3,opt,name=response_body_bytes,json=responseBodyBytes,proto3" json:"response_body_bytes,omitempty"`
	// Map of additional headers configured to be logged.
	ResponseHeaders map[string]string `protobuf:"bytes,4,rep,name=response_headers,json=responseHeaders,proto3" json:"response_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Map of trailers configured to be logged.
	ResponseTrailers map[string]string `protobuf:"bytes,5,rep,name=response_trailers,json=responseTrailers,proto3" json:"response_trailers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The HTTP response code details.
	ResponseCodeDetails string `protobuf:"bytes,6,opt,name=response_code_details,json=responseCodeDetails,proto3" json:"response_code_details,omitempty"`
}
