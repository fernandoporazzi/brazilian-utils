package boleto

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/fernandoporazzi/brazilian-utils/helpers"
)

type partial struct {
	start, end, index int
}

var partials = []partial{
	{0, 9, 9},
	{10, 20, 20},
	{21, 31, 31},
}

var mod10Weights = []int{2, 1}
var mod11Weights = []int{2, 9}

var digitableLineToBoletoConvertPositions = []struct {
	start, end int
}{
	{0, 4},
	{32, 47},
	{4, 9},
	{10, 20},
	{21, 31},
}

var dotIndexes = []int{4, 14, 25}
var spaceIndexes = []int{9, 20, 31, 32}

const (
	length             = 47
	checkDigitPosition = 4
)

func mod10(partial string) int {
	splitted := strings.Split(partial, "")
	reversed := helpers.Reverse(splitted)

	acc := 0

	for index, digit := range reversed {
		i, err := strconv.Atoi(digit)
		if err != nil {
			log.Fatal(err)
		}

		result := i * mod10Weights[index%2]

		if result > 9 {
			r := 1 + (result % 10)
			acc = acc + r
		} else {
			acc = acc + result
		}
	}

	mod := acc % 10

	if mod > 0 {
		return 10 - mod
	}

	return 0
}

func mod11(value string) int {
	initial := mod11Weights[0]
	end := mod11Weights[1]
	weight := initial
	acc := 0
	splitted := strings.Split(value, "")
	reversed := helpers.Reverse(splitted)

	for _, digit := range reversed {
		d, err := strconv.Atoi(digit)
		if err != nil {
			log.Fatal(err)
		}

		result := d * weight

		if weight < end {
			weight++
		} else {
			weight = initial
		}

		acc = acc + result
	}

	mod := acc % 11

	if mod == 0 || mod == 1 {
		return 1
	}

	return 11 - mod
}

func isValidLength(digitableLine string) bool {
	return len(digitableLine) == length
}

func isValidPartials(digitableLine string) bool {
	every := true

	for _, partial := range partials {
		mod := mod10(digitableLine[partial.start:partial.end])

		i, err := strconv.Atoi(digitableLine[partial.index : partial.index+1])
		if err != nil {
			log.Fatal(err)
		}

		if i != mod {
			every = false
		}

	}

	return every
}

func parse(digitableLine string) string {
	var acc string
	for _, p := range digitableLineToBoletoConvertPositions {
		acc += digitableLine[p.start:p.end]
	}
	return acc
}

func isValidCheckDigit(parsedDigitableLine string) bool {
	mod := mod11(parsedDigitableLine[0:checkDigitPosition] + parsedDigitableLine[checkDigitPosition+1:])

	digit := parsedDigitableLine[checkDigitPosition : checkDigitPosition+1]

	i, err := strconv.Atoi(digit)
	if err != nil {
		log.Fatal(err)
	}

	return i == mod
}

// IsValid checks if boleto (brazilian payment method) is valid.
func IsValid(digitableLine string) bool {
	digits := helpers.OnlyNumbers(digitableLine)

	if !isValidLength(digits) {
		return false
	}

	if !isValidPartials(digits) {
		return false
	}

	parsedDigits := parse(digits)

	return isValidCheckDigit(parsedDigits)
}

// Format boleto (brazilian payment method)
func Format(boleto string) string {
	sliceAt := length

	if len(boleto) < length {
		sliceAt = len(boleto)
	}

	digits := helpers.OnlyNumbers(boleto)[0:sliceAt]

	acc := ""

	for i, digit := range digits {
		acc = acc + fmt.Sprintf("%c", digit)

		if helpers.IsLastChar(i, digits) == false {
			if helpers.Contains(dotIndexes, i) {
				acc = acc + "."
			}

			if helpers.Contains(spaceIndexes, i) {
				acc = acc + " "
			}
		}
	}

	return acc
}

func GetValueInCents(boleto string) int {
	if !IsValid(boleto) {
		return 0
	}

	digits := helpers.OnlyNumbers(boleto)

	value, err := strconv.Atoi(digits[len(digits)-10:])

	if err != nil {
		log.Fatal(err)
	}

	return value
}

func GetBankCode(boleto string) string {
	if !IsValid(boleto) {
		return ""
	}

	return boleto[0:3]
}

func GetExpirationDate(boleto string) time.Time {
	if !IsValid(boleto) {
		return time.Time{}
	}

	daysSinceBaseDayStartIndex := 14
	daysSinceBaseDayEndIndex := 10

	daysSinceBaseDay, err := strconv.Atoi(boleto[len(boleto)-daysSinceBaseDayStartIndex : len(boleto)-daysSinceBaseDayEndIndex])

	if err != nil {
		log.Fatal(err)
	}

	oneDayMilliseconds := 24 * 60 * 60 * 1000
	millisecondsSinceBaseDay := daysSinceBaseDay * oneDayMilliseconds

	dateSinceBaseDay := time.Unix(0, int64(millisecondsSinceBaseDay)*int64(time.Millisecond))

	bancoCentralBaseDate, err := time.Parse("2006-01-02", "1997-09-07")
	if err != nil {
		log.Fatal(err)
	}

	bancoCentralBaseDateMilliseconds := bancoCentralBaseDate.UnixNano() / int64(time.Millisecond)

	return dateSinceBaseDay.Add(time.Millisecond * time.Duration(bancoCentralBaseDateMilliseconds)).UTC()
}
