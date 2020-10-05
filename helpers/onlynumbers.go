package helpers

import (
	"fmt"
	"log"
	"regexp"
)

// OnlyNumbers remove all non numeric characters
func OnlyNumbers(input interface{}) string {
	i := fmt.Sprintf("%v", input)

	reg, err := regexp.Compile("[^0-9]+")

	if err != nil {
		log.Fatal(err)
	}

	return reg.ReplaceAllString(i, "")
}
