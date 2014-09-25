package raindrops

import "strconv"

// Convert gives us the raindrops for a given number as per the rules
func Convert(number int) string {
	var returnString string

	if number%3 == 0 {
		returnString += "Pling"
	}

	if number%5 == 0 {
		returnString += "Plang"
	}

	if number%7 == 0 {
		returnString += "Plong"
	}

	if returnString == "" {
		return strconv.Itoa(number)
	}
	return returnString
}
