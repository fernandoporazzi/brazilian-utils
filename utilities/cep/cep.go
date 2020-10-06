package cep

import "github.com/fernandoporazzi/brazilian-utils/helpers"

const (
	length = 8
)

func isValidLength(cep string) bool {
	return len(cep) == length
}

// IsValid checks if a Brazilian Postal Code is valid
func IsValid(cep string) bool {
	digits := helpers.OnlyNumbers(cep)

	return isValidLength(digits)
}

// Format CEP
//
// 95690000 -> 95690-000
func Format(cep string) string {
	digits := helpers.OnlyNumbers(cep)[:length]

	return digits[:5] + "-" + digits[5:]
}
