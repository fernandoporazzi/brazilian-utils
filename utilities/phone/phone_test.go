package phone

import "testing"

func TestIsValid(t *testing.T) {
	t.Run("Should return false", func(t *testing.T) {
		t.Run("When it is empty string", func(t *testing.T) {
			if IsValid("") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}
			if IsValidLandlinePhone("") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}
			if IsValidMobilePhone("") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}
		})

		t.Run("When it is a mobile phone with mask and code state invalid", func(t *testing.T) {
			if IsValid("(00) 3 0000-0000") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}

			if IsValidMobilePhone("(00) 3 0000-0000") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}
		})

		t.Run("When is a landline with mask and code state invalid", func(t *testing.T) {
			if IsValid("(11) 9000-0000") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}

			if IsValidLandlinePhone("(11) 9000-0000") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}
		})

		t.Run("When dont match with phone min length", func(t *testing.T) {
			if IsValid("11") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}
			if IsValidLandlinePhone("11") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}
			if IsValidMobilePhone("11") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}
		})

		t.Run("When dont match with phone max length", func(t *testing.T) {
			if IsValid("11300000001130000000") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}
			if IsValidLandlinePhone("11300000001130000000") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}
			if IsValidMobilePhone("11300000001130000000") != false {
				t.Errorf("Expected %v to be equal %v", false, false)
			}
		})
	})

	t.Run("Should return true", func(t *testing.T) {
		t.Run("When it is a mobile phone valid with mask", func(t *testing.T) {
			if IsValid("(11) 9 0000-0000") != true {
				t.Errorf("Expected %v to be equal %v", true, true)
			}

			if IsValidMobilePhone("(11) 9 0000-0000") != true {
				t.Errorf("Expected %v to be equal %v", true, true)
			}
		})

		t.Run("When it is a landline valid with mask", func(t *testing.T) {
			if IsValid("(11) 3000-0000") != true {
				t.Errorf("Expected %v to be equal %v", true, true)
			}

			if IsValidLandlinePhone("(11) 3000-0000") != true {
				t.Errorf("Expected %v to be equal %v", true, true)
			}
		})

		t.Run("When it is a mobile phone valid without mask", func(t *testing.T) {
			if IsValid("11900000000") != true {
				t.Errorf("Expected %v to be equal %v", true, true)
			}

			if IsValidMobilePhone("11900000000") != true {
				t.Errorf("Expected %v to be equal %v", true, true)
			}
		})

		t.Run("When it is a landline valid without mask", func(t *testing.T) {
			if IsValid("1130000000") != true {
				t.Errorf("Expected %v to be equal %v", true, true)
			}

			if IsValidLandlinePhone("1130000000") != true {
				t.Errorf("Expected %v to be equal %v", true, true)
			}
		})
	})
}
