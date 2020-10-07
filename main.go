package main

import (
	"fmt"

	"github.com/fernandoporazzi/brazilian-utils/helpers"
)

func main() {
	// fmt.Println(phone.IsValid("11900000000"))
	// fmt.Println(cep.Format("95690000"))

	fmt.Println(helpers.GenerateChecksum("987654", []int{10, 8, 5, 6, 3, 7}))
}
