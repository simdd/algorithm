package problem

/**
* 遍历-选择最小的，放前面
**/
func main(o []int) []int {
	position := 0
	min := position

	for position < len(o)-1 {
		min = position

		for i := position + 1; i < len(o); i++ {
			if o[i] < o[min] {
				min = i
			}
		}

		o[position], o[min] = o[min], o[position]
		position++
	}

	return o
}
