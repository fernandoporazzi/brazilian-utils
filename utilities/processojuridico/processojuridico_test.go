package processojuridico

import "testing"

func TestFormat(t *testing.T) {
	t.Run("Should format processo juridico with mask", func(t *testing.T) {
		entries := []struct {
			processoJuridico string
			want             string
		}{
			{"", ""},
			{"0", "0"},
			{"00", "00"},
			{"000", "000"},
			{"0002", "0002"},
			{"00020", "00020"},
			{"000208", "000208"},
			{"0002080", "0002080"},
			{"00020802", "0002080-2"},
			{"000208025", "0002080-25"},
			{"0002080252", "0002080-25.2"},
			{"00020802520", "0002080-25.20"},
			{"000208025201", "0002080-25.201"},
			{"0002080252012", "0002080-25.2012"},
			{"00020802520125", "0002080-25.2012.5"},
			{"000208025201251", "0002080-25.2012.51"},
			{"0002080252012515", "0002080-25.2012.515"},
			{"00020802520125150", "0002080-25.2012.515.0"},
			{"000208025201251500", "0002080-25.2012.515.00"},
			{"0002080252012515004", "0002080-25.2012.515.004"},
			{"00020802520125150049", "0002080-25.2012.515.0049"},
		}

		for _, entry := range entries {
			got := Format(entry.processoJuridico)

			if got != entry.want {
				t.Errorf("Expected formatted processo juridico to be '%v', but got '%v'", entry.want, got)
			}
		}
	})

	t.Run("Should NOT add digits after the processo juridico length 20", func(t *testing.T) {
		got := Format("00020802520125150049123123")
		want := "0002080-25.2012.515.0049"

		if got != want {
			t.Errorf("Expected formatted processo juridico to be '%v', but got '%v'", want, got)
		}
	})

	t.Run("Should remove all non numeric characters", func(t *testing.T) {
		got := Format("0002080@$25201%!@2515.%0049123123")
		want := "0002080-25.2012.515.0049"

		if got != want {
			t.Errorf("Expected formatted processo juridico to be '%v', but got '%v'", want, got)
		}
	})
}
