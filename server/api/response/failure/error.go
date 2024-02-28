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

type DurationLimitError struct {
	Reason string `json:"reason"`
}

func (err DurationLimitError) Error() string {
	return err.Reason
}
