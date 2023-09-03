package iteration

// Repeat a string for any number of times and return the concatenated string.
func Repeat(char string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += char
	}
	return repeated
}
