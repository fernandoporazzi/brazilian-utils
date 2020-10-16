package boleto

import "testing"

func TestIsValid(t *testing.T) {
	entries := []struct {
		digitableLine string
		want          bool
	}{
		{"", false},
		{"12345", false},
		{"123456789", false},
		{"00190000090114971860168524522114675860000102656", true},
		{"0019000009 01149.718601 68524.522114 6 75860000102656", true},
		{"00190000020114971860168524522114675860000102656", false},
		{"0034191.79001 01043.510047 91020.150008 7 84100026000", false},
		{"001 9 05009 ( 5 ) 401448 1606 ( 9 ) 0680935031 ( 4 ) 337370000000100", true},
		{"00190500954014481606906809350314337370000000100", true},
	}

	for _, entry := range entries {
		got := IsValid(entry.digitableLine)

		if got != entry.want {
			t.Errorf("Expected Boleto '%v' to be equal %v, but got %v", entry.digitableLine, entry.want, got)
		}
	}
}
