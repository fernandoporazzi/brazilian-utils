package helpers

import (
	"testing"
)

func TestGenerateChecksum(t *testing.T) {
	entries := []struct {
		base   string
		weight interface{}
		want   int
	}{
		{"987654", []int{10, 8, 5, 6, 3, 7}, 268},
		{"987654", 10, 310},
		{"12", 10, 28},
		{"12", []int{10, 9}, 28},
	}

	for _, entry := range entries {
		got := GenerateChecksum(entry.base, entry.weight)

		if got != entry.want {
			t.Errorf("Expected Checksum to be %v, but got %v", entry.want, got)
		}
	}
}
