package series

//All returns a list of all substrings of s with length n
func All(n int, s string) (result []string) {
	for i := 0; i <= len(s)-n; i++ {
		result = append(result, s[i:i+n])
	}

	return
}

//UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

//First returns the first substring of swith length n and an identifier that it worked
func First(n int, s string) (string, bool) {
	if len(s) < n {
		return "", false
	}
	return s[0:n], true
}
