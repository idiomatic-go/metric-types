package accesslogv3

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/idiomatic-go/metric-data/envoy/v3"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type HTTPRequestProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The request method (RFC 7231/2616).
	RequestMethod v3.RequestMethod `protobuf:"varint,1,opt,name=request_method,json=requestMethod,proto3,enum=envoy.config.core.v3.RequestMethod" json:"request_method,omitempty"`
	// The scheme portion of the incoming request URI.
	Scheme string `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	// HTTP/2 ``:authority`` or HTTP/1.1 ``Host`` header value.
	Authority string `protobuf:"bytes,3,opt,name=authority,proto3" json:"authority,omitempty"`
	// The port of the incoming request URI
	// (unused currently, as port is composed onto authority).
	Port *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	// The path portion from the incoming request URI.
	Path string `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	// Value of the ``User-Agent`` request header.
	UserAgent string `protobuf:"bytes,6,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	// Value of the ``Referer`` request header.
	Referer string `protobuf:"bytes,7,opt,name=referer,proto3" json:"referer,omitempty"`
	// Value of the ``X-Forwarded-For`` request header.
	ForwardedFor string `protobuf:"bytes,8,opt,name=forwarded_for,json=forwardedFor,proto3" json:"forwarded_for,omitempty"`
	// Value of the ``X-Request-Id`` request header
	//
	// This header is used by Envoy to uniquely identify a request.
	// It will be generated for all external requests and internal requests that
	// do not already have a request ID.
	RequestId string `protobuf:"bytes,9,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Value of the ``X-Envoy-Original-Path`` request header.
	OriginalPath string `protobuf:"bytes,10,opt,name=original_path,json=originalPath,proto3" json:"original_path,omitempty"`
	// Size of the HTTP request headers in bytes.
	//
	// This value is captured from the OSI layer 7 perspective, i.e. it does not
	// include overhead from framing or encoding at other networking layers.
	RequestHeadersBytes uint64 `protobuf:"varint,11,opt,name=request_headers_bytes,json=requestHeadersBytes,proto3" json:"request_headers_bytes,omitempty"`
	// Size of the HTTP request body in bytes.
	//
	// This value is captured from the OSI layer 7 perspective, i.e. it does not
	// include overhead from framing or encoding at other networking layers.
	RequestBodyBytes uint64 `protobuf:"varint,12,opt,name=request_body_bytes,json=requestBodyBytes,proto3" json:"request_body_bytes,omitempty"`
	// Map of additional headers that have been configured to be logged.
	RequestHeaders map[string]string `protobuf:"bytes,13,rep,name=request_headers,json=requestHeaders,proto3" json:"request_headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}
