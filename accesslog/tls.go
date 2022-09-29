package accesslog

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
	TlsVersion string
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
