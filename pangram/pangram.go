package pangram

//IsPangram ...
func IsPangram(s string) bool {

	seen := make(map[int]bool)

	for i := 0; i < len(s); i++ {

		// if upper case
		if s[i] >= 65 && s[i] <= 90 {
			seen[int(s[i]+32)] = true

			// if lower case
		} else if s[i] >= 97 && s[i] <= 122 {
			seen[int(s[i])] = true

			// anything else
		} else {
			continue
		}
	}

	for _, v := range "abcdefghijklmnopqrstuvwxyz" {

		_, ok := seen[int(v)]
		if !ok {
			return false
		}
	}

	return true
}
