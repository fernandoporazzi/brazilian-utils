package main

import (
	"fmt"

	"github.com/fernandoporazzi/brazilian-utils/utilities/cep"
	"github.com/fernandoporazzi/brazilian-utils/utilities/phone"
)

func main() {
	fmt.Println(phone.IsValid("11900000000"))
	fmt.Println(cep.Format("95690000"))
}
