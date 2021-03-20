package cpf

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/fernandoporazzi/brazilian-utils/data"
	"github.com/fernandoporazzi/brazilian-utils/helpers"
	"github.com/fernandoporazzi/brazilian-utils/utilities/states"
)

const (
	length = 11
)

var checkDigitsIndexes = []int{9, 10}
var dotIndexes = []int{2, 5}
var hyphenIndexes = []int{8}

// Format takes a string and returns a valid-formatted CPF
func Format(cpf string) (string, error) {
	digits := helpers.OnlyNumbers(cpf)

	if len(digits) != length {
		return "", errors.New("CPF length is not 11")
	}

	acc := ""

	for i, digit := range digits {
		acc = acc + fmt.Sprintf("%c", digit)

		if !helpers.IsLastChar(i, digits) {
			if helpers.Contains(dotIndexes, i) {
				acc = acc + "."
			}

			if helpers.Contains(hyphenIndexes, i) {
				acc = acc + "-"
			}
		}
	}

	return acc, nil
}

func Generate(params ...string) (string, error) {
	paramsLength := len(params)

	var stateCode string

	if paramsLength > 1 {
		return "", fmt.Errorf("Expected zero or one argument as State, received '%v'", paramsLength)
	}

	if paramsLength == 0 {
		stateCode = helpers.GenerateRandomNumber(1)
	} else {
		state := params[0]
		states := states.GetStates()

		found := ""

		for _, s := range states {
			if s.Code == state || s.Name == state {
				found = s.Code
				break
			}
		}

		if found == "" {
			stateCode = helpers.GenerateRandomNumber(1)
		} else {
			stateCode = data.States[found].Code
		}
	}

	baseCpf := helpers.GenerateRandomNumber(length-3) + stateCode

	firstCheckDigitMod := helpers.GenerateChecksum(baseCpf, 10) % 11
	var firstCheckDigit string

	if firstCheckDigitMod < 2 {
		firstCheckDigit = "0"
	} else {
		firstCheckDigit = strconv.Itoa(11 - firstCheckDigitMod)
	}

	secondCheckDigitMod := helpers.GenerateChecksum(baseCpf+firstCheckDigit, 11) % 11
	var secondCheckDigit string

	if secondCheckDigitMod < 2 {
		secondCheckDigit = "0"
	} else {
		secondCheckDigit = strconv.Itoa(11 - secondCheckDigitMod)
	}

	return baseCpf + firstCheckDigit + secondCheckDigit, nil
}

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
