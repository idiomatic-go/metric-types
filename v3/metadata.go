package v3

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	"github.com/idiomatic-go/metric-data/any"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// Metadata provides additional inputs to filters based on matched listeners,
// filter chains, routes and endpoints. It is structured as a map, usually from
// filter name (in reverse DNS format) to metadata specific to the filter. Metadata
// key-values for a filter are merged as connection and request handling occurs,
// with later values for the same key overriding earlier values.
//
// An example use of metadata is providing additional values to
// http_connection_manager in the envoy.http_connection_manager.access_log
// namespace.
//
// Another example use of metadata is to per service config info in cluster metadata, which may get
// consumed by multiple filters.
//
// For load balancing, Metadata provides a means to subset cluster endpoints.
// Endpoints have a Metadata object associated and routes contain a Metadata
// object to match against. There are some well defined metadata used today for
// this purpose:
//
// * ``{"envoy.lb": {"canary": <bool> }}`` This indicates the canary status of an
//   endpoint and is also used during header processing
//   (x-envoy-upstream-canary) and for stats purposes.
// [#next-major-version: move to type/metadata/v2]
type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Key is the reverse DNS filter name, e.g. com.acme.widget. The ``envoy.*``
	// namespace is reserved for Envoy's built-in filters.
	// If both ``filter_metadata`` and
	// :ref:`typed_filter_metadata <envoy_v3_api_field_config.core.v3.Metadata.typed_filter_metadata>`
	// fields are present in the metadata with same keys,
	// only ``typed_filter_metadata`` field will be parsed.
	FilterMetadata map[string]*_struct.Struct `protobuf:"bytes,1,rep,name=filter_metadata,json=filterMetadata,proto3" json:"filter_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Key is the reverse DNS filter name, e.g. com.acme.widget. The ``envoy.*``
	// namespace is reserved for Envoy's built-in filters.
	// The value is encoded as google.protobuf.Any.
	// If both :ref:`filter_metadata <envoy_v3_api_field_config.core.v3.Metadata.filter_metadata>`
	// and ``typed_filter_metadata`` fields are present in the metadata with same keys,
	// only ``typed_filter_metadata`` field will be parsed.
	TypedFilterMetadata map[string]*any.Any `protobuf:"bytes,2,rep,name=typed_filter_metadata,json=typedFilterMetadata,proto3" json:"typed_filter_metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}
