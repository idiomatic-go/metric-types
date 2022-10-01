package accesslogv3

import "github.com/idiomatic-go/common-lib/util"

type HTTPRequestProperties struct {
	// The request method (RFC 7231/2616).
	RequestMethod RequestMethod
	// The scheme portion of the incoming request URI.
	//Scheme string
	Scheme util.DictionaryValue
	// HTTP/2 ``:authority`` or HTTP/1.1 ``Host`` header value.
	//Authority string
	Authority util.DictionaryValue
	// The port of the incoming request URI
	// (unused currently, as port is composed onto authority).
	Port uint32
	// The path portion from the incoming request URI.
	//Path string
	Path util.DictionaryValue
	// Value of the ``User-Agent`` request header.
	//UserAgent string
	UserAgent util.DictionaryValue
	// Value of the ``Referer`` request header.
	//Referer string
	Referer util.DictionaryValue
	// Value of the ``X-Forwarded-For`` request header.
	//ForwardedFor string
	ForwardedFor util.DictionaryValue
	// Value of the ``X-Request-Id`` request header
	//
	// This header is used by Envoy to uniquely identify a request.
	// It will be generated for all external requests and internal requests that
	// do not already have a request ID.
	RequestId string
	// Value of the ``X-Envoy-Original-Path`` request header.
	//OriginalPath string
	OriginalPath util.DictionaryValue
	// Size of the HTTP request headers in bytes.
	//
	// This value is captured from the OSI layer 7 perspective, i.e. it does not
	// include overhead from framing or encoding at other networking layers.
	RequestHeadersBytes uint64
	// Size of the HTTP request body in bytes.
	//
	// This value is captured from the OSI layer 7 perspective, i.e. it does not
	// include overhead from framing or encoding at other networking layers.
	RequestBodyBytes uint64
	// Map of additional headers that have been configured to be logged.
	RequestHeaders map[string]string
}
