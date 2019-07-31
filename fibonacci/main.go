package problem

func fibonacci1(n int) int {
	if n <= 2 {
		return n
	}

	return fibonacci1(n-1) + fibonacci1(n-2)
}

func fibonacci2(n int) int {
	if n < 4 {
		return n
	}

	n1 := 2
	n2 := 3

	for i := 4; i <= n; i++ {
		n1, n2 = n2, n1+n2
	}

	return n2
}
