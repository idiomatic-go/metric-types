package accesslogv3

import (
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type TLSProperties_TLSVersion int32

const (
	TLSProperties_VERSION_UNSPECIFIED TLSProperties_TLSVersion = 0
	TLSProperties_TLSv1               TLSProperties_TLSVersion = 1
	TLSProperties_TLSv1_1             TLSProperties_TLSVersion = 2
	TLSProperties_TLSv1_2             TLSProperties_TLSVersion = 3
	TLSProperties_TLSv1_3             TLSProperties_TLSVersion = 4
)

// Enum value maps for TLSProperties_TLSVersion.
var (
	TLSProperties_TLSVersion_name = map[int32]string{
		0: "VERSION_UNSPECIFIED",
		1: "TLSv1",
		2: "TLSv1_1",
		3: "TLSv1_2",
		4: "TLSv1_3",
	}
	TLSProperties_TLSVersion_value = map[string]int32{
		"VERSION_UNSPECIFIED": 0,
		"TLSv1":               1,
		"TLSv1_1":             2,
		"TLSv1_2":             3,
		"TLSv1_3":             4,
	}
)

type isTLSProperties_CertificateProperties_SubjectAltName_San interface {
	isTLSProperties_CertificateProperties_SubjectAltName_San()
}

type TLSProperties_CertificateProperties_SubjectAltName_Uri struct {
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3,oneof"`
}

type TLSProperties_CertificateProperties_SubjectAltName_Dns struct {
	// [#not-implemented-hide:]
	Dns string `protobuf:"bytes,2,opt,name=dns,proto3,oneof"`
}

type TLSProperties_CertificateProperties_SubjectAltName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to San:
	//	*TLSProperties_CertificateProperties_SubjectAltName_Uri
	//	*TLSProperties_CertificateProperties_SubjectAltName_Dns
	San isTLSProperties_CertificateProperties_SubjectAltName_San `protobuf_oneof:"san"`
}

type TLSProperties_CertificateProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SANs present in the certificate.
	SubjectAltName []*TLSProperties_CertificateProperties_SubjectAltName `protobuf:"bytes,1,rep,name=subject_alt_name,json=subjectAltName,proto3" json:"subject_alt_name,omitempty"`
	// The subject field of the certificate.
	Subject string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
}

type TLSProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Version of TLS that was negotiated.
	TlsVersion TLSProperties_TLSVersion `protobuf:"varint,1,opt,name=tls_version,json=tlsVersion,proto3,enum=envoy.data.accesslog.v3.TLSProperties_TLSVersion" json:"tls_version,omitempty"`
	// TLS cipher suite negotiated during handshake. The value is a
	// four-digit hex code defined by the IANA TLS Cipher Suite Registry
	// (e.g. ``009C`` for ``TLS_RSA_WITH_AES_128_GCM_SHA256``).
	//
	// Here it is expressed as an integer.
	TlsCipherSuite *wrappers.UInt32Value `protobuf:"bytes,2,opt,name=tls_cipher_suite,json=tlsCipherSuite,proto3" json:"tls_cipher_suite,omitempty"`
	// SNI hostname from handshake.
	TlsSniHostname string `protobuf:"bytes,3,opt,name=tls_sni_hostname,json=tlsSniHostname,proto3" json:"tls_sni_hostname,omitempty"`
	// Properties of the local certificate used to negotiate TLS.
	LocalCertificateProperties *TLSProperties_CertificateProperties `protobuf:"bytes,4,opt,name=local_certificate_properties,json=localCertificateProperties,proto3" json:"local_certificate_properties,omitempty"`
	// Properties of the peer certificate used to negotiate TLS.
	PeerCertificateProperties *TLSProperties_CertificateProperties `protobuf:"bytes,5,opt,name=peer_certificate_properties,json=peerCertificateProperties,proto3" json:"peer_certificate_properties,omitempty"`
	// The TLS session ID.
	TlsSessionId string `protobuf:"bytes,6,opt,name=tls_session_id,json=tlsSessionId,proto3" json:"tls_session_id,omitempty"`
	// The ``JA3`` fingerprint when ``JA3`` fingerprinting is enabled.
	Ja3Fingerprint string `protobuf:"bytes,7,opt,name=ja3_fingerprint,json=ja3Fingerprint,proto3" json:"ja3_fingerprint,omitempty"`
}
