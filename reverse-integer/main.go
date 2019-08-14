package problem

func main(x int) (out int) {
	if x-int(int32(x)) > 0 {
		return 0
	}

	for ; x != 0; x /= 10 {
		out = out*10 + x%10
	}
	return out
}
