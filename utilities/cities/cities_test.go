package cities

import "testing"

func TestGetCities(t *testing.T) {
	t.Run("When no State is provided", func(t *testing.T) {
		wantedLength := 5570
		got := len(GetCities())

		if got != wantedLength {
			t.Errorf("GetCities should have '%v' cities, but got '%v'", wantedLength, got)
		}
	})

	t.Run("When State Code is provided", func(t *testing.T) {
		entries := []struct {
			code         string
			wantedLength int
		}{
			{"GO", 246},
			{"MG", 853},
			{"PA", 144},
			{"CE", 184},
			{"BA", 417},
			{"PR", 399},
			{"SC", 295},
			{"PE", 185},
			{"TO", 139},
			{"RN", 167},
			{"PI", 224},
			{"RS", 497},
			{"MT", 141},
			{"AC", 22},
			{"SP", 645},
			{"ES", 78},
			{"MA", 217},
			{"PB", 223},
			{"MS", 79},
			{"RO", 52},
			{"RR", 15},
			{"AM", 62},
			{"AP", 16},
			{"SE", 75},
			{"AL", 102},
			{"RJ", 92},
			{"DF", 1},
		}

		for _, entry := range entries {
			got := len(GetCities(entry.code))

			if got != entry.wantedLength {
				t.Errorf("Expected GetCities(%v) to have length '%v', but got '%v'", entry.code, entry.wantedLength, got)
			}
		}
	})

	t.Run("When State Name is provided", func(t *testing.T) {
		entries := []struct {
			name         string
			wantedLength int
		}{
			{"Goiás", 246},
			{"Minas Gerais", 853},
			{"Pará", 144},
			{"Ceará", 184},
			{"Bahia", 417},
			{"Paraná", 399},
			{"Santa Catarina", 295},
			{"Pernambuco", 185},
			{"Tocantins", 139},
			{"Rio Grande do Norte", 167},
			{"Piauí", 224},
			{"Rio Grande do Sul", 497},
			{"Mato Grosso", 141},
			{"Acre", 22},
			{"São Paulo", 645},
			{"Espírito Santo", 78},
			{"Maranhão", 217},
			{"Paraíba", 223},
			{"Mato Grosso do Sul", 79},
			{"Rondônia", 52},
			{"Roraima", 15},
			{"Amazonas", 62},
			{"Amapá", 16},
			{"Sergipe", 75},
			{"Alagoas", 102},
			{"Rio de Janeiro", 92},
			{"Distrito Federal", 1},
		}

		for _, entry := range entries {
			got := len(GetCities(entry.name))

			if got != entry.wantedLength {
				t.Errorf("Expected GetCities(%v) to have length '%v', but got '%v'", entry.name, entry.wantedLength, got)
			}
		}
	})

	t.Run("When non-existing State Name or Code is provided", func(t *testing.T) {
		state := "GM"
		got := len(GetCities(state))
		expectedLength := 0

		if got != expectedLength {
			t.Errorf("Expected GetCities(%v) to have length '%v', but got '%v'", state, expectedLength, got)
		}
	})
}
