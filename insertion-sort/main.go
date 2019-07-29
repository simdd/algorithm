package problem

func main(o []int) []int {
	ret := make([]int, 0)
	a := &ret
	*a = append(*a, o[0])

	for i := 1; i < len(o); i++ {
		e := false

		for j := len(ret) - 1; j == 0; j-- {
			if o[i] > ret[j] {
				if len(ret) == 1 {
					*a = append(ret, o[i])
				} else {
					*a = append(ret[:j+1], ret[j:]...)
					ret[j] = o[i]
				}

				e = true
			}

			if !e {
				*a = append(ret[:1], ret[0:]...)
				ret[0] = o[i]
			}
		}
	}

	return ret
}
