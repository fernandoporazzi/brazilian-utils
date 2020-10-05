package phone

import (
	"log"
	"strconv"

	"github.com/fernandoporazzi/brazilian-utils/data"
	"github.com/fernandoporazzi/brazilian-utils/helpers"
)

var phoneMinLength = 10
var phoneMaxLength = 11
var landlineValidFirstNumbers = []int{2, 3, 4, 5}
var mobileValidFirstNumbers = []int{6, 7, 8, 9}

func parsePhoneDigits(phone string) (isValidDigits bool, digits string) {
	isValidDigits = true

	if phone == "" {
		isValidDigits = false
	}

	return isValidDigits, helpers.OnlyNumbers(phone)
}

func isValidMobilePhoneLength(phone string) bool {
	return len(phone) == phoneMaxLength
}

func isValidLandlinePhoneLength(phone string) bool {
	return len(phone) >= phoneMinLength && len(phone) < phoneMaxLength
}

func isValidLength(phone string) bool {
	return isValidLandlinePhoneLength(phone) || isValidMobilePhoneLength(phone)
}

func isValidLandlinePhoneFirstNumber(phone string) bool {
	runes := []rune(phone)

	i, err := strconv.Atoi(string(runes[2]))

	if err != nil {
		log.Fatal(err)
	}

	return helpers.Contains(landlineValidFirstNumbers, i)
}

func isValidMobilePhoneFirstNumber(phone string) bool {
	runes := []rune(phone)

	i, err := strconv.Atoi(string(runes[2]))

	if err != nil {
		log.Fatal(err)
	}

	return helpers.Contains(mobileValidFirstNumbers, i)
}

func isValidFirstNumber(phone string) bool {
	if len(phone) == phoneMinLength {
		return isValidLandlinePhoneFirstNumber(phone)
	}

	return isValidMobilePhoneFirstNumber(phone)
}

func isValidDDD(phone string) bool {
	runes := []rune(phone)

	i, err := strconv.Atoi(string(runes[:2]))

	if err != nil {
		log.Fatal(err)
	}

	return helpers.Contains(data.AreaCodes, i)
}

// IsValid checks if phone number (mobile or landline) is valid.
func IsValid(phone string) bool {
	isValidDigits, digits := parsePhoneDigits(phone)

	if !isValidDigits {
		return false
	}

	return isValidLength(digits) && isValidFirstNumber(digits) && isValidDDD(digits)
}

// IsValidMobilePhone checks if mobile phone number is valid.
func IsValidMobilePhone(phone string) bool {
	isValidDigits, digits := parsePhoneDigits(phone)

	if !isValidDigits {
		return false
	}

	return isValidMobilePhoneLength(digits) && isValidMobilePhoneFirstNumber(digits) && isValidDDD(digits)
}

// IsValidLandlinePhone checks if landline phone number is valid.
func IsValidLandlinePhone(phone string) bool {
	isValidDigits, digits := parsePhoneDigits(phone)

	if !isValidDigits {
		return false
	}

	return isValidLandlinePhoneLength(digits) && isValidLandlinePhoneFirstNumber(digits) && isValidDDD(digits)
}
