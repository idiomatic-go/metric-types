package accesslogv3

func ConvertTLSVersion(vers int32) TLSProperties_TLSVersion {
	if vers < int32(TLSProperties_VERSION_UNSPECIFIED) || vers > int32(TLSProperties_TLSv1_3) {
		return TLSProperties_VERSION_UNSPECIFIED
	}
	return TLSProperties_TLSVersion(vers)
}

func GetTLSVersionName(v TLSProperties_TLSVersion) string {
	if v <= TLSProperties_VERSION_UNSPECIFIED || v > TLSProperties_TLSv1_3 {
		v = TLSProperties_VERSION_UNSPECIFIED
	}
	return TLSProperties_TLSVersion_name[int32(v)]
}

func ConvertHTTPVersion(vers int32) HTTPAccessLogEntry_HTTPVersion {
	return HTTPAccessLogEntry_HTTPVersion(vers)
}
