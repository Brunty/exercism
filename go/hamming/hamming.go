package hamming

import (
	"errors"
	"fmt"
)

const testVersion = 4

func Distance(a, b string) (int, error) {
	var lengthOfA = len(a)
	var lengthOfB = len(b)

	if lengthOfA != lengthOfB {
		return -1, errors.New(fmt.Sprintf("The twi strings are not equal - a: %s b: %s", lengthOfA, lengthOfB))
	}

	var diff int = 0

	for i := range a {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil
}
