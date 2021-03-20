package helpers

import "testing"

func TestGenerateRandomNumber(t *testing.T) {
	t.Run("It creates a random number with the right amount of chars", func(t *testing.T) {
		entries := []struct {
			length int
		}{
			{0},
			{1},
			{2},
			{5},
			{10},
			{100},
			{1000},
		}

		for _, entry := range entries {
			got := GenerateRandomNumber(entry.length)

			if len(got) != entry.length {
				t.Errorf("Expected random number to have length '%v', but got '%v'", entry.length, len(got))
			}
		}
	})
}
