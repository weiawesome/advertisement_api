package failure

type ClientError struct {
	Reason string `json:"reason"`
}

func (err ClientError) Error() string {
	return err.Reason
}

type ServerError struct {
	Reason string `json:"reason"`
}

func (err ServerError) Error() string {
	return err.Reason
}

type DayLimitError struct {
	Reason string `json:"reason"`
}

func (err DayLimitError) Error() string {
	return err.Reason
}

type TimeLimitError struct {
	Reason string `json:"reason"`
}

func (err TimeLimitError) Error() string {
	return err.Reason
}
