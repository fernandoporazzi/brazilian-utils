package cpf

import "testing"

func TestIsValid(t *testing.T) {
	t.Run("When CPF length does not match required length", func(t *testing.T) {
		entries := []struct {
			cpf  string
			want bool
		}{
			{"", false},
			{"403644788", false},
		}

		for _, entry := range entries {
			got := IsValid(entry.cpf)

			if got != entry.want {
				t.Errorf("Expected CPF '%v' to be invalid", entry.cpf)
			}
		}
	})

	t.Run("When CPF belong to the reserved words list", func(t *testing.T) {
		entries := []struct {
			cpf  string
			want bool
		}{
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
		}

		for _, entry := range entries {
			got := IsValid(entry.cpf)

			if got != entry.want {
				t.Errorf("Expected CPF '%v' to be invalid", entry.cpf)
			}
		}
	})

	t.Run("When CPF contains only letter", func(t *testing.T) {
		entries := []struct {
			cpf  string
			want bool
		}{
			{"abc", false},
			{"abcdefghijk", false},
		}

		for _, entry := range entries {
			got := IsValid(entry.cpf)

			if got != entry.want {
				t.Errorf("Expected CPF '%v' to be invalid", entry.cpf)
			}
		}
	})

	t.Run("When CPF is valid", func(t *testing.T) {
		entries := []struct {
			cpf  string
			want bool
		}{
			{"40364478829", true},
			{"962.718.458-60", true},
			{"905.886.311-59", true},
			{"069.106.596-94", true},
			{"31110186959", true},
			{"295.326.33227", true},
			{"280.521037-97", true},
			{"07542897705", true},
			{"82765875855", true},
			{"311.344.135-80", true},
			{"02805510615", true},
			{"840.881.935-63", true},
			{"432.25217549", true},
			{"685.615.387-24", true},
		}

		for _, entry := range entries {
			got := IsValid(entry.cpf)

			if got != entry.want {
				t.Errorf("Expected CPF '%v' to be valid", entry.cpf)
			}
		}
	})

	t.Run("When CPF is invalid", func(t *testing.T) {
		entries := []struct {
			cpf  string
			want bool
		}{
			{"123.456.789-10", false},
			{"73.500.457-35", false},
			{"987.654.321-98", false},
			{"11122233344", false},
			{"192.837.465-19", true},
			{"546.372.819-73", false},
			{"02036942010", false},
			{"871.298.056-12", false},
		}

		for _, entry := range entries {
			got := IsValid(entry.cpf)

			if got != entry.want {
				t.Errorf("Expected CPF '%v' to be invalid", entry.cpf)
			}
		}
	})

}
