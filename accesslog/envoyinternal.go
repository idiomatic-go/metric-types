package accesslog

type EnvoyInternalAddress struct {
	ServerListenerName string
	// Specifies an endpoint identifier to distinguish between multiple endpoints for the same internal listener in a
	// single upstream pool. Only used in the upstream addresses for tracking changes to individual endpoints. This, for
	// example, may be set to the final destination IP for the target internal listener.
	EndpointId string
}
