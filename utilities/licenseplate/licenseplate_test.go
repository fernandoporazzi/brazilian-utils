package licenseplate

import "testing"

func TestIsValid(t *testing.T) {
	entries := []struct {
		plate string
		want  bool
	}{
		{"ABC1234", true},
		{"abc1234", true},
		{"ABC-1234", true},
		{"abc-1234", true},
		{"AbC1234", true},
		{"AbC-1234", true},
		{"abc1d23", true},
		{"ABC1D23", true},
		{"lorem ipsum", false},
		{"abc123", false},
		{"1abc1234", false},
		{"abc1234a", false},
		{"abc01234a", false},
		{"a0c01234a", false},
		{"abc-1234a", false},
		{"aabc1234", false},
	}

	for _, entry := range entries {
		got := IsValid(entry.plate)

		if got != entry.want {
			t.Errorf("Expected plate %v to be %v, but got %v", entry.plate, entry.want, got)
		}
	}

}
