package problem

// Divide & Merge
func main(o []int) []int {
	return divide(o)
}

func divide(o []int) []int {
	if len(o) < 2 {
		return o
	}

	mid := len(o) / 2
	return merge(divide(o[:mid]), divide(o[mid:]))
}

func merge(l, r []int) []int {
	size, i, j := len(l)+len(r), 0, 0
	ret := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(l)-1 && j <= len(r)-1 {
			ret[k] = r[j]
			j++
		} else if j > len(r)-1 && i <= len(l)-1 {
			ret[k] = l[i]
			i++
		} else if l[i] < r[j] {
			ret[k] = l[i]
			i++
		} else {
			ret[k] = r[j]
			j++
		}
	}

	return ret
}
