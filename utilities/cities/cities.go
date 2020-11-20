package cities

import (
	"fmt"
	"log"

	"github.com/fernandoporazzi/brazilian-utils/data"
	"github.com/fernandoporazzi/brazilian-utils/utilities/states"
)

// GetCities returns all Brazilian cities if no state is passed as argument
func GetCities(params ...string) []string {
	length := len(params)

	if length > 1 {
		log.Fatal("Expected zero or one argument as State, received " + fmt.Sprintf("%v", len(params)))
	}

	if length == 1 {
		state := params[0]
		states := states.GetStates()

		found := ""

		for _, s := range states {
			if s.Code == state || s.Name == state {
				found = s.Code
				break
			}
		}

		return data.Cities[found]
	}

	allCities := []string{}

	for _, cities := range data.Cities {
		allCities = append(allCities, cities...)
	}

	return allCities
}
