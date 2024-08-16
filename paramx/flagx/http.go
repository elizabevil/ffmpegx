package flagx

type AuthType = Flag
type HttpMethod = Flag

const (
	AuthType_none AuthType = "none"
	//Choose the HTTP authentication type automatically. This is the default.

	AuthType_basic AuthType = "basic"
	//Choose the HTTP basic authentication.
)

const (
	MethodGet     HttpMethod = "GET"
	MethodHead    HttpMethod = "HEAD"
	MethodPost    HttpMethod = "POST"
	MethodPut     HttpMethod = "PUT"
	MethodPatch   HttpMethod = "PATCH" // RFC 5789
	MethodDelete  HttpMethod = "DELETE"
	MethodConnect HttpMethod = "CONNECT"
	MethodOptions HttpMethod = "OPTIONS"
	MethodTrace   HttpMethod = "TRACE"
)
