package helpers

// IsLastChar returns whether or not a char of a given string is the last
func IsLastChar(index int, input string) bool {
	return index == len(input)-1
}
