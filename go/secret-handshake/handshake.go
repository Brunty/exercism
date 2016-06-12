package secret

import "sort"

var sequence sort.StringSlice = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(input int) []string {
    handshake := []string{}

    if input <= 0 {
        return handshake
    }

    for i, step := range sequence {
        if 1 << uint(i) & input > 0 {
            handshake = append(handshake, step)
        }
    }

    if 1 << uint(len(sequence)) & input > 0 {
        reverse(handshake)
    }

    return handshake
}

/* wondering if there's not a better way to do this with sort.* functions? */
func reverse(strings []string) {
    for i, j := 0, len(strings) - 1; i < j; i, j = i + 1, j - 1 {
        strings[i], strings[j] = strings[j], strings[i]
    }
}