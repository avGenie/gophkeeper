package grpc

// GRPC processing errors
//
// FailedUserCredentials - returned if user credentials are not valid
// InternalServerError - returned in internal server error
const (
	FailedUserCredentials = "user credentials are not valid"
	InternalServerError   = "internal server error"
)
