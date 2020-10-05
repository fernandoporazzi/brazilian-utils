package main

import (
	"fmt"

	"github.com/fernandoporazzi/brazilian-utils/utilities/phone"
)

func main() {
	// fmt.Println(helpers.OnlyNumbers("123abc"))
	fmt.Println(phone.IsValid("11900000000"))
}
