package cep

import "testing"

func TestIsValid(t *testing.T) {
	entries := []struct {
		cep  string
		want bool
	}{
		{"", false},
		{"12345", false},
		{"123456789", false},
		{"92500000", true},
		{"92500-000", true},
	}

	for _, entry := range entries {
		got := IsValid(entry.cep)

		if got != entry.want {
			t.Errorf("Expected CEP '%v' to be equal %v, but got %v", entry.cep, entry.want, got)
		}
	}
}

func TestFormat(t *testing.T) {
	entries := []struct {
		cep  string
		want string
	}{
		{"01001000", "01001-000"},
		{"956900001234", "95690-000"},
		{"95690000", "95690-000"},
	}

	for _, entry := range entries {
		got := Format(entry.cep)

		if got != entry.want {
			t.Errorf("Expected formatted cep to be %v, but got %v", entry.want, got)
		}
	}
}
