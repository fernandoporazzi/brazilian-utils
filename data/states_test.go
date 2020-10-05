package data

import (
	"testing"
)

func TestAreaCodes(t *testing.T) {
	t.Run("It has area codes", func(t *testing.T) {
		want := 67
		got := len(AreaCodes)

		if want != got {
			t.Errorf("Expected length of AreaCodes to be equal %v, got %v", want, got)
		}
	})
}
