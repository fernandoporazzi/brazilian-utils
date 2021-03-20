package cpf

import (
	"testing"
)

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
			{"192.837.465-19", false},
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

func TestFormat(t *testing.T) {
	t.Run("When CPF length does not match required length", func(t *testing.T) {
		entries := []struct {
			cpf string
		}{
			{""},
			{"403644788"},
			{"123435"},
			{"12343567890123456789"},
		}

		for _, entry := range entries {
			_, err := Format(entry.cpf)

			if err.Error() != "CPF length is not 11" {
				t.Errorf("Expected error to be '%v', instead got '%v'", "CPF length is not 11", err.Error())
			}
		}
	})

	t.Run("When CPF length matches required length", func(t *testing.T) {
		entries := []struct {
			cpf  string
			want string
		}{
			{"03136942017", "031.369.420-17"},
			{"11133344455", "111.333.444-55"},
			{"111a333b444c55", "111.333.444-55"},
		}

		for _, entry := range entries {
			got, _ := Format(entry.cpf)

			if got != entry.want {
				t.Errorf("Expected formatted CPF to be '%v', but got '%v' instead", entry.want, got)
			}
		}
	})
}

func TestGenerate(t *testing.T) {
	t.Run("Should generate valid CPF when params are not provided", func(t *testing.T) {
		generatedCpf, _ := Generate()

		if len(generatedCpf) != 11 {
			t.Errorf("Expected generated CPF to have length 11, but got '%v' instead", len(generatedCpf))
		}

		if !IsValid(generatedCpf) {
			t.Errorf("Expected generated CPF(%v)to be valid", generatedCpf)
		}
	})

	t.Run("Should generate valid CPF when state name or state code are provided", func(t *testing.T) {
		entries := []struct {
			input string
		}{
			{"RS"},
			{"Rio Grande do Sul"},
			{"RJ"},
			{"Rio de Janeiro"},
			{"ABC"}, // non-existent State Code or State Name
		}

		for _, entry := range entries {
			got, _ := Generate(entry.input)

			if len(got) != 11 {
				t.Errorf("Expected generated CPF to have length 11, but got '%v' instead", len(got))
			}

			if !IsValid(got) {
				t.Errorf("Expected generated CPF(%v)to be valid", got)
			}
		}
	})

	t.Run("Should return error when 2 or more parameters are provided", func(t *testing.T) {
		_, err := Generate("RS", "SC")

		if err == nil {
			t.Errorf("Expected error to be '%v'", err.Error())
		}
	})
}
