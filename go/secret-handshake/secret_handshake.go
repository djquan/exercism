package secret

//Handshake takes an int and returns a set of secret phrases
func Handshake(code uint) []string {
	result := make([]string, 0)

	if code&1 == 1 {
		result = append(result, "wink")
	}

	if code&2 == 2 {
		result = append(result, "double blink")
	}

	if code&4 == 4 {
		result = append(result, "close your eyes")
	}

	if code&8 == 8 {
		result = append(result, "jump")
	}

	if code&16 == 16 {
		for i := len(result)/2 - 1; i >= 0; i-- {
			opp := len(result) - 1 - i
			result[i], result[opp] = result[opp], result[i]
		}
	}

	return result
}
