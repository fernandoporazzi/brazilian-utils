package states

import "testing"

func TestGetStates(t *testing.T) {
	t.Run("should return an array with 27 state", func(t *testing.T) {
		states := GetStates()
		l := len(states)

		if l != 27 {
			t.Errorf("Expected length of states to be 27, but got '%v'", l)
		}
	})
}
