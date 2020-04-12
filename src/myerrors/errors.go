package myerrors

// Exception is generic error
type Exception struct {
	Message string
	Code    int
	Reason  string
}
