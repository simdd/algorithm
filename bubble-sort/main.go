package problem

func main(o []int) []int {
	position := len(o) - 1

	for position > 0 {
		for i := 1; i <= position; i++ {
			if o[i] < o[i-1] {
				tmp := o[i-1]
				o[i-1] = o[i]
				o[i] = tmp
			}
		}

		position--
	}

	return o
}

func main2(o []int) []int {
	position := len(o) - 1

	for position > 0 {
		change := false

		for i := 1; i <= position; i++ {
			if o[i] < o[i-1] {
				tmp := o[i-1]
				o[i-1] = o[i]
				o[i] = tmp
				change = true
			}
		}

		if change == false {
			return o
		}

		position--
	}

	return o
}
