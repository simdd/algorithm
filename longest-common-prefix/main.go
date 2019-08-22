package problem

func main(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	p := ""
	i := 0
	o := strs[0]

	for i < len(o) {
		j := 1
		b := true

		for j < len(strs) {
			n := strs[j]
			if o[i] != n[i] {
				b = false
				break
			}

			j++
		}

		if b == false {
			break
		} else {
			p += string(o[i])
		}

		i++
	}

	return p
}
