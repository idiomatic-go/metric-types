package accesslogv3

type ResponseFlags_Encoded uint32

// Reasons why the request was unauthorized
type ResponseFlags_Unauthorized_Reason int32

const (
	ResponseFlags_Unauthorized_REASON_UNSPECIFIED ResponseFlags_Unauthorized_Reason = 0
	// The request was denied by the external authorization service.
	ResponseFlags_Unauthorized_EXTERNAL_SERVICE ResponseFlags_Unauthorized_Reason = 1
)

// Enum value maps for ResponseFlags_Unauthorized_Reason.
var (
	ResponseFlags_Unauthorized_Reason_name = map[int32]string{
		0: "REASON_UNSPECIFIED",
		1: "EXTERNAL_SERVICE",
	}
	ResponseFlags_Unauthorized_Reason_value = map[string]int32{
		"REASON_UNSPECIFIED": 0,
		"EXTERNAL_SERVICE":   1,
	}
)

type ResponseFlags_Unauthorized struct {
	Reason ResponseFlags_Unauthorized_Reason
}

// Flags indicating occurrences during request/response processing.
type ResponseFlags struct {
	Encoded             ResponseFlags_Encoded
	UnauthorizedDetails *ResponseFlags_Unauthorized
}
