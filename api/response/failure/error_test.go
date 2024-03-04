package failure

import "testing"

var testReasonCase = "Error reason"

func TestError(t *testing.T) {
	t.Run("Client error case", func(t *testing.T) {
		err := ClientError{Reason: testReasonCase}

		if err.Error() != testReasonCase {
			t.Errorf("Error resaon should be %v not %v", testReasonCase, err.Error())
		}
	})
	t.Run("Server error case", func(t *testing.T) {
		err := ServerError{Reason: testReasonCase}

		if err.Error() != testReasonCase {
			t.Errorf("Error resaon should be %v not %v", testReasonCase, err.Error())
		}
	})
	t.Run("Day limit error case", func(t *testing.T) {
		err := DayLimitError{Reason: testReasonCase}

		if err.Error() != testReasonCase {
			t.Errorf("Error resaon should be %v not %v", testReasonCase, err.Error())
		}
	})
	t.Run("Duration limit error case", func(t *testing.T) {
		err := DurationLimitError{Reason: testReasonCase}

		if err.Error() != testReasonCase {
			t.Errorf("Error resaon should be %v not %v", testReasonCase, err.Error())
		}
	})
}
