package helpers

import "testing"

func TestOnlyNumbers(t *testing.T) {
	t.Run("should remove all non numeric characters", func(t *testing.T) {
		entries := []struct {
			input interface{}
			want  string
		}{
			{"abcd", ""},
			{"12345", "12345"},
			{"12345 ", "12345"},
			{" 12345", "12345"},
			{"   12345  ", "12345"},
			{"abcd12345", "12345"},
			{"123abc456?.#789xyz 0", "1234567890"},
			{123, "123"},
			{-123, "123"},
			{0, "0"},
			{[]float32{123, 456}, "123456"},
			{true, ""},
		}

		for _, entry := range entries {
			got := OnlyNumbers(entry.input)

			if got != entry.want {
				t.Errorf("Expected %v to be equal %v", got, entry.want)
			}
		}
	})
}
