/*
All response for the failure request.
Including client's error, server's error, day-limit and duration-limit error.
Furthermore, every error has its reason help client know what's happened.
*/

package failure

// ClientError is the error from client's reason
// The reason is why the error happen
type ClientError struct {
	Reason string `json:"reason"`
}

// Error is a function to show the reason
func (err ClientError) Error() string {
	return err.Reason
}

// ServerError is the error from server's reason
// The reason is why the error happen
type ServerError struct {
	Reason string `json:"reason"`
}

// Error is a function to show the reason
func (err ServerError) Error() string {
	return err.Reason
}

// DayLimitError is the error from system's limit
// For example, day's advertisement addition can't more than 3000.
// If it exceeds the limitation(3000), it will return this type of error.
// Furthermore, the reason is why the error happen
type DayLimitError struct {
	Reason string `json:"reason"`
}

// Error is a function to show the reason
func (err DayLimitError) Error() string {
	return err.Reason
}

// DurationLimitError is the error from system's limit
// For example, advertisement in a period can't more than 1000.
// If the period in advertisement post exceeds the limitation(1000), it will return this type of error.
// Furthermore, the reason is why the error happen
type DurationLimitError struct {
	Reason string `json:"reason"`
}

// Error is a function to show the reason
func (err DurationLimitError) Error() string {
	return err.Reason
}
