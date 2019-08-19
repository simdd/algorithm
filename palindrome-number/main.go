package problem

func main(x int) bool {
	if x < 0 {
		return false
	}

	in := x
	out := 0

	for ; x != 0; x /= 10 {
		out = out*10 + x%10
	}

	return out == in
}
