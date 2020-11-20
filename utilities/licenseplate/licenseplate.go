package licenseplate

import "regexp"

// IsValid checks if a Brazilian or Mercosul car plate is valid
func IsValid(plate string) bool {
	brazilianLicensePlate, _ := regexp.Compile(`(?m)^[a-zA-Z]{3}-?[0-9]{4}$`)
	mercosulLicensePlate, _ := regexp.Compile(`(?m)^[a-zA-Z]{3}[0-9]{1}[a-zA-Z]{1}[0-9]{2}$`)

	return brazilianLicensePlate.Match([]byte(plate)) || mercosulLicensePlate.Match([]byte(plate))
}
