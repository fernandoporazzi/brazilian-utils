package helpers

// Contains returns whether or not an array of integers includes a given number
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
