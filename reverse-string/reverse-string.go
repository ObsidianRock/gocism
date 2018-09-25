package reverse

func String(s string) string {

	if len(s) == 0 {
		return ""
	}

	v := []rune(s)

	for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {

		v[i], v[j] = v[j], v[i]

	}

	return string(v)
}
