package pis

import (
	"log"
	"regexp"
	"strconv"

	"github.com/fernandoporazzi/brazilian-utils/helpers"
)

func isValidLength(pis string) bool {
	return len(pis) == 11
}

func isReservedNumber(pis string) bool {
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

	return helpers.SliceContainsString(reservedNumbers, pis)
}

func hasLetter(pis string) bool {
	reg, _ := regexp.Compile("[^0-9]+")

	return reg.Match([]byte(pis))
}

func removeSeparators(pis string) string {
	reg, err := regexp.Compile("[ ().,*-]+")

	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(pis, "")
}

// IsValid checks if PIS is valid
func IsValid(pis string) bool {
	digits := removeSeparators(pis)

	if hasLetter(digits) || !isValidLength(digits) || isReservedNumber(digits) {
		return false
	}

	weights := []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	weightedChecksum := helpers.GenerateChecksum(digits[:len(digits)-1], weights)
	verifyingDigit, err := strconv.Atoi(digits[len(digits)-1:])

	if err != nil {
		log.Fatal(err)
	}

	calculatedDigit := 11 - (weightedChecksum % 11)

	return calculatedDigit == verifyingDigit ||
		(calculatedDigit == 10 && verifyingDigit == 0) ||
		(calculatedDigit == 11 && verifyingDigit == 0)
}
