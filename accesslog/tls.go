package accesslog

// Changes : removed wrappers

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

type TLSProperties_CertificateProperties_SubjectAltName struct {
	// Types that are assignable to San:
	//	*TLSProperties_CertificateProperties_SubjectAltName_Uri
	//	*TLSProperties_CertificateProperties_SubjectAltName_Dns
	//San isTLSProperties_CertificateProperties_SubjectAltName_San `protobuf_oneof:"san"`
	Dns string
	Uri string
}

type TLSProperties_CertificateProperties struct {
	// SANs present in the certificate.
	SubjectAltName []*TLSProperties_CertificateProperties_SubjectAltName
	// The subject field of the certificate.
	Subject string
}

type TLSProperties struct {
	// Version of TLS that was negotiated.
	TlsVersion TLSProperties_TLSVersion
	// TLS cipher suite negotiated during handshake. The value is a
	// four-digit hex code defined by the IANA TLS Cipher Suite Registry
	// (e.g. ``009C`` for ``TLS_RSA_WITH_AES_128_GCM_SHA256``).
	//
	// Here it is expressed as an integer.
	TlsCipherSuite uint32
	// SNI hostname from handshake.
	TlsSniHostname string
	// Properties of the local certificate used to negotiate TLS.
	LocalCertificateProperties *TLSProperties_CertificateProperties
	// Properties of the peer certificate used to negotiate TLS.
	PeerCertificateProperties *TLSProperties_CertificateProperties
	// The TLS session ID.
	TlsSessionId string
	// The ``JA3`` fingerprint when ``JA3`` fingerprinting is enabled.
	Ja3Fingerprint string
}
