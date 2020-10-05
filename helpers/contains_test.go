package helpers

import "testing"

func TestContains(t *testing.T) {
	entries := []struct {
		arr     []int
		element int
		want    bool
	}{
		{[]int{9, 12, 34, 36, 111}, 36, true},
		{[]int{9, 12, 34, 36, 0, 111}, 0, true},
		{[]int{-9, 12, 34, 36, 0, 111}, -9, true},
		{[]int{9, 12, 34, 36, 0, 111}, -9, false},
		{[]int{9, 12, 34, 36, 111}, 0, false},
	}

	for _, entry := range entries {
		got := Contains(entry.arr, entry.element)

		if got != entry.want {
			t.Errorf("Expected %v to be equal %v", got, entry.want)
		}
	}
}
