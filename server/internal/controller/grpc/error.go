package grpc

// GRPC processing errors
//
// FailedUserCredentials - returned if user credentials are not valid
// InternalServerError - returned in internal server error
// FailedData - returned if data is not valid
const (
	FailedUserCredentials = "user credentials are not valid"
	InternalServerError   = "internal server error"
	FailedData            = "data is not valid"
)
