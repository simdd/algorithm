package prime_number

func primeNumber(o int) []int {
	var arr []int

	if o > 1 {
		arr = append(arr, 2)
	}

	for i := 3; i < o; i++ {
		res := false

		for _, v := range arr {
			if i%v == 0 {
				res = true
				break
			}
		}

		if res == false {
			arr = append(arr, i)
		}
	}

	return arr
}
