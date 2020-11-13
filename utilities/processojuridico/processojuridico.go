package processojuridico

import (
	"fmt"

	"github.com/fernandoporazzi/brazilian-utils/helpers"
)

const length = 20

var dotIndexes = []int{8, 12, 15}
var hyphenIndexes = []int{6}

// Format adds mask to a processo jurídico
func Format(processoJuridico string) string {
	sliceAt := length

	if len(processoJuridico) < length {
		sliceAt = len(processoJuridico)
	}

	digits := helpers.OnlyNumbers(processoJuridico)[0:sliceAt]

	acc := ""

	for i, digit := range digits {
		acc = acc + fmt.Sprintf("%c", digit)

		if helpers.IsLastChar(i, digits) == false {
			if helpers.Contains(dotIndexes, i) {
				acc = acc + "."
			}

			if helpers.Contains(hyphenIndexes, i) {
				acc = acc + "-"
			}
		}
	}

	return acc
}
