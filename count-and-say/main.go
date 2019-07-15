package problem

// 1.     1
// 2.     11
// 3.     21
// 4.     1211
// 5.     111221

// read([2,1]) [1,2,1,1]
func read(n []int) []int {
	val := n[0]
	count := 1
	ret := []int{}

	for i:=1; i<len(n); i++ {
		if n[i] == val {
			count++
		}else {
			ret = append(ret, count)
			ret = append(ret, val)
			val = n[i]
			count = 1
		}
	}

	ret = append(ret, count)
	ret = append(ret, val)

	return ret
}

// func countAndSay(n int) string {}