package v3

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	"github.com/idiomatic-go/metric-data/envoy/v31"
	"google.golang.org/protobuf/runtime/protoimpl"
)

// Version and identification for an Envoy extension.
// [#next-free-field: 7]
type Extension struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is the name of the Envoy filter as specified in the Envoy
	// configuration, e.g. envoy.filters.http.router, com.acme.widget.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Category of the extension.
	// Extension category names use reverse DNS notation. For instance "envoy.filters.listener"
	// for Envoy's built-in listener filters or "com.acme.filters.http" for HTTP filters from
	// acme.com vendor.
	// [#comment:TODO(yanavlasov): Link to the doc with existing envoy category names.]
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	// [#not-implemented-hide:] Type descriptor of extension configuration proto.
	// [#comment:TODO(yanavlasov): Link to the doc with existing configuration protos.]
	// [#comment:TODO(yanavlasov): Add tests when PR #9391 lands.]
	//
	// Deprecated: Do not use.
	TypeDescriptor string `protobuf:"bytes,3,opt,name=type_descriptor,json=typeDescriptor,proto3" json:"type_descriptor,omitempty"`
	// The version is a property of the extension and maintained independently
	// of other extensions and the Envoy API.
	// This field is not set when extension did not provide version information.
	Version *BuildVersion `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// Indicates that the extension is present but was disabled via dynamic configuration.
	Disabled bool `protobuf:"varint,5,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Type URLs of extension configuration protos.
	TypeUrls []string `protobuf:"bytes,6,rep,name=type_urls,json=typeUrls,proto3" json:"type_urls,omitempty"`
}

// BuildVersion combines SemVer version of extension with free-form build information
// (i.e. 'alpha', 'private-build') as a set of strings.
type BuildVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SemVer version of extension.
	// Note..: Had to remove the v3. prefix
	Version *SemanticVersion `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Free-form build information.
	// Envoy defines several well known keys in the source/common/version/version.h file
	Metadata *_struct.Struct `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

// Identifies location of where either Envoy runs or where upstream hosts run.
type Locality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Region this :ref:`zone <envoy_v3_api_field_config.core.v3.Locality.zone>` belongs to.
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// Defines the local service zone where Envoy is running. Though optional, it
	// should be set if discovery service routing is used and the discovery
	// service exposes :ref:`zone data <envoy_v3_api_field_config.endpoint.v3.LocalityLbEndpoints.locality>`,
	// either in this message or via :option:`--service-zone`. The meaning of zone
	// is context dependent, e.g. `Availability Zone (AZ)
	// <https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html>`_
	// on AWS, `Zone <https://cloud.google.com/compute/docs/regions-zones/>`_ on
	// GCP, etc.
	Zone string `protobuf:"bytes,2,opt,name=zone,proto3" json:"zone,omitempty"`
	// When used for locality of upstream hosts, this field further splits zone
	// into smaller chunks of sub-zones so they can be load balanced
	// independently.
	SubZone string `protobuf:"bytes,3,opt,name=sub_zone,json=subZone,proto3" json:"sub_zone,omitempty"`
}

type isNode_UserAgentVersionType interface {
	isNode_UserAgentVersionType()
}

type Node_UserAgentVersion struct {
	// Free-form string that identifies the version of the entity requesting config.
	// E.g. "1.12.2" or "abcd1234", or "SpecialEnvoyBuild"
	UserAgentVersion string `protobuf:"bytes,7,opt,name=user_agent_version,json=userAgentVersion,proto3,oneof"`
}

type Node_UserAgentBuildVersion struct {
	// Structured version of the entity requesting config.
	UserAgentBuildVersion *BuildVersion `protobuf:"bytes,8,opt,name=user_agent_build_version,json=userAgentBuildVersion,proto3,oneof"`
}

// Identifies a specific Envoy instance. The node identifier is presented to the
// management server, which may use this identifier to distinguish per Envoy
// configuration for serving.
// [#next-free-field: 13]
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An opaque node identifier for the Envoy node. This also provides the local
	// service node name. It should be set if any of the following features are
	// used: :ref:`statsd <arch_overview_statistics>`, :ref:`CDS
	// <config_cluster_manager_cds>`, and :ref:`HTTP tracing
	// <arch_overview_tracing>`, either in this message or via
	// :option:`--service-node`.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Defines the local service cluster name where Envoy is running. Though
	// optional, it should be set if any of the following features are used:
	// :ref:`statsd <arch_overview_statistics>`, :ref:`health check cluster
	// verification
	// <envoy_v3_api_field_config.core.v3.HealthCheck.HttpHealthCheck.service_name_matcher>`,
	// :ref:`runtime override directory <envoy_v3_api_msg_config.bootstrap.v3.Runtime>`,
	// :ref:`user agent addition
	// <envoy_v3_api_field_extensions.filters.network.http_connection_manager.v3.HttpConnectionManager.add_user_agent>`,
	// :ref:`HTTP global rate limiting <config_http_filters_rate_limit>`,
	// :ref:`CDS <config_cluster_manager_cds>`, and :ref:`HTTP tracing
	// <arch_overview_tracing>`, either in this message or via
	// :option:`--service-cluster`.
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Opaque metadata extending the node identifier. Envoy will pass this
	// directly to the management server.
	Metadata *_struct.Struct `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Map from xDS resource type URL to dynamic context parameters. These may vary at runtime (unlike
	// other fields in this message). For example, the xDS client may have a shard identifier that
	// changes during the lifetime of the xDS client. In Envoy, this would be achieved by updating the
	// dynamic context on the Server::Instance's LocalInfo context provider. The shard ID dynamic
	// parameter then appears in this field during future discovery requests.
	DynamicParameters map[string]*v31.ContextParams `protobuf:"bytes,12,rep,name=dynamic_parameters,json=dynamicParameters,proto3" json:"dynamic_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Locality specifying where the Envoy instance is running.
	Locality *Locality `protobuf:"bytes,4,opt,name=locality,proto3" json:"locality,omitempty"`
	// Free-form string that identifies the entity requesting config.
	// E.g. "envoy" or "grpc"
	UserAgentName string `protobuf:"bytes,6,opt,name=user_agent_name,json=userAgentName,proto3" json:"user_agent_name,omitempty"`
	// Types that are assignable to UserAgentVersionType:
	//	*Node_UserAgentVersion
	//	*Node_UserAgentBuildVersion
	UserAgentVersionType isNode_UserAgentVersionType `protobuf_oneof:"user_agent_version_type"`
	// List of extensions and their versions supported by the node.
	Extensions []*Extension `protobuf:"bytes,9,rep,name=extensions,proto3" json:"extensions,omitempty"`
	// Client feature support list. These are well known features described
	// in the Envoy API repository for a given major version of an API. Client features
	// use reverse DNS naming scheme, for example ``com.acme.feature``.
	// See :ref:`the list of features <client_features>` that xDS client may
	// support.
	ClientFeatures []string `protobuf:"bytes,10,rep,name=client_features,json=clientFeatures,proto3" json:"client_features,omitempty"`
	// Known listening ports on the node as a generic hint to the management server
	// for filtering :ref:`listeners <config_listeners>` to be returned. For example,
	// if there is a listener bound to port 80, the list can optionally contain the
	// SocketAddress ``(0.0.0.0,80)``. The field is optional and just a hint.
	//
	// Deprecated: Do not use.
	ListeningAddresses []*Address `protobuf:"bytes,11,rep,name=listening_addresses,json=listeningAddresses,proto3" json:"listening_addresses,omitempty"`
}
