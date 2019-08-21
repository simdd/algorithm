package problem

func main(s string) int {
	ret := 0
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	l := len(s)
	i := l - 1

	for i >= 0 {
		c := string(s[i])
		cnum := m[c]
		ret += cnum
		i--

		if i >= 0 {
			p := string(s[i])
			pnum := m[p]

			if pnum < cnum {
				ret -= pnum
				i--
			}
		}
	}

	return ret
}
