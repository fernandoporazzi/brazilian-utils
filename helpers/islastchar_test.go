package helpers

import "testing"

func TestIsLastChar(t *testing.T) {
	t.Run("Should return true when index is the same as last index of the string", func(t *testing.T) {
		entries := []struct {
			index int
			input string
			want  bool
		}{
			{0, "a", true},
			{1, "ab", true},
			{2, "abc", true},
			{3, "abcd", true},
			{4, "abcd ", true},
			{5, "abcd e", true},
		}

		for _, entry := range entries {
			got := IsLastChar(entry.index, entry.input)

			if got != entry.want {
				t.Errorf("Expected output for %v: '%v', but got '%v'", entry.input, entry.want, got)
			}
		}
	})

	t.Run("Should return false when index is not the same as last index of the string", func(t *testing.T) {
		entries := []struct {
			index int
			input string
			want  bool
		}{
			{1, "a", false},
			{2, "ab", false},
			{3, "abc", false},
		}

		for _, entry := range entries {
			got := IsLastChar(entry.index, entry.input)

			if got != entry.want {
				t.Errorf("Expected output for %v: '%v', but got '%v'", entry.input, entry.want, got)
			}
		}
	})
}
