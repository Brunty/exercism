package secret

type SecretHandshake struct {
	action string
	binary int
}

const REVERSE = 16

var sequence = []SecretHandshake{{"wink", 1}, {"double blink", 2}, {"close your eyes", 4}, {"jump", 8}}

func Handshake(givenCode int) []string {
	handshake := []string{}

	if givenCode <= 0 {
		return handshake
	}

	for _, secret := range sequence {
		if secret.binary&givenCode == secret.binary {
			handshake = append(handshake, secret.action)
		}
	}

	if REVERSE&givenCode > 0 {
		reverse(handshake)
	}

	return handshake
}

/* wondering if there's not a better way to do this with sort.* functions? */
func reverse(strings []string) {
	for i, j := 0, len(strings)-1; i < j; i, j = i+1, j-1 {
		strings[i], strings[j] = strings[j], strings[i]
	}
}
