package helpers

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	t.Run("It reverses slice of string", func(t *testing.T) {
		entries := []struct {
			input []string
			want  []string
		}{
			{[]string{"a"}, []string{"a"}},
			{[]string{"a", "b"}, []string{"b", "a"}},
			{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
			{[]string{"1", "2", "3"}, []string{"3", "2", "1"}},
			{[]string{"a1", "b2", "c3"}, []string{"c3", "b2", "a1"}},
		}

		for _, entry := range entries {
			got := Reverse(entry.input)

			if !reflect.DeepEqual(got, entry.want) {
				t.Errorf("Expected reverse array to be '%v', but got '%v'", entry.want, got)
			}
		}
	})
}
