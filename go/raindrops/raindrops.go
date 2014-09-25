package raindrops

import "strconv"

// primeFactors for us to use to determine what to output
var primeFactors = []struct {
	prime   int
	message string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

// Convert gives us the raindrops for a given number as per the rules
func Convert(input int) string {
	var output string

	for _, factor := range primeFactors {
		if input%factor.prime == 0 {
			output += factor.message
		}
	}

	if output == "" {
		return strconv.Itoa(input)
	}
	return output
}
