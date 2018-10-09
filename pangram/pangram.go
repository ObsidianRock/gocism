package pangram

//IsPangram ...
func IsPangram(s string) bool {

	seen := make(map[int]bool)

	for i := 97; i < 122; i++ {
		seen[i] = false
	}

	for i := 0; i < len(s); i++ {

		if s[i] >= 65 && s[i] <= 90 {
			seen[int(s[i]+32)] = true
		} else if s[i] >= 97 && s[i] <= 122 {
			seen[int(s[i])] = true
		} else {
			continue
		}
	}

	for _, ok := range seen {
		if !ok {
			return false
		}
	}

	return true
}
