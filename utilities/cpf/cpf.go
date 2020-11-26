package cpf

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/fernandoporazzi/brazilian-utils/helpers"
)

var checkDigitsIndexes = []int{9, 10}

func isValidFormat(cpf string) bool {
	r, _ := regexp.Compile(`(?m)^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$`)

	return r.Match([]byte(cpf))
}

func isReservedNumber(cpf string) bool {
	reservedNumbers := []string{
		"00000000000",
		"11111111111",
		"22222222222",
		"33333333333",
		"44444444444",
		"55555555555",
		"66666666666",
		"77777777777",
		"88888888888",
		"99999999999",
	}

	return helpers.SliceContainsString(reservedNumbers, cpf)
}

func isValidChecksum(cpf string) bool {
	allTrue := true

	for _, i := range checkDigitsIndexes {
		mod := helpers.GenerateChecksum(cpf[0:i], i+1) % 11

		var m string

		if mod < 2 {
			m = "0"
		} else {
			m = strconv.Itoa(11 - mod)
		}

		chartAtIndex := fmt.Sprintf("%c", cpf[i])

		if m != chartAtIndex {
			allTrue = false
			break
		}
	}

	return allTrue
}

// IsValid check if CPF is valid
func IsValid(cpf string) bool {
	digits := helpers.OnlyNumbers(cpf)

	return isValidFormat(cpf) && !isReservedNumber(digits) && isValidChecksum(digits)
}
