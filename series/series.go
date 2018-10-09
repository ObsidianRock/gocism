package series

//All returns a list of all substrings of s with length n.
func All(n int, s string) []string {
	var series []string

	end, front := 0, n

	for {
		if front > len(s) {
			break
		}
		series = append(series, string(s[end:front]))

		end++
		front++

	}

	return series
}
