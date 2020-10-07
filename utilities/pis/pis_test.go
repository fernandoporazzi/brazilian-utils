package pis

import "testing"

func TestIsValid(t *testing.T) {
	entries := []struct {
		pis  string
		want bool
	}{
		{"", false},
		{"00000000000", false},
		{"11111111111", false},
		{"22222222222", false},
		{"33333333333", false},
		{"44444444444", false},
		{"55555555555", false},
		{"66666666666", false},
		{"77777777777", false},
		{"88888888888", false},
		{"99999999999", false},
		{"123456", false},
		{"12056Aabb412847", false},
		{"abcabcabcde", false},
		{"12056412547", false},
		{"12081636639", false},
		{"12056412847", true},
		{"120.5641.284-7", true},
		{"120.1213.266-0", true},
		{"120.7041.469-0", true},
		{"558.89873.84-1", true},
		{"236.50833.45-6", true},
		{"520.85778.06-3", true},
		{"454.03417.88-0", true},
		{"476.60681.13-3", true},
		{"44087491810", true},
		{"86836499503", true},
	}

	for _, entry := range entries {
		got := IsValid(entry.pis)

		if got != entry.want {
			t.Errorf("Expected PIS %v to be %v, but got %v", entry.pis, entry.want, got)
		}
	}
}
