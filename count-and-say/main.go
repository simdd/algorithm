package problem

import (
	"strconv"
	"strings"
)

// 1.     1
// 2.     11
// 3.     21
// 4.     1211
// 5.     111221

// read([2,1]) [1,2,1,1]
func read(n []string) []string {
	val := n[0]
	count := 1
	ret := []string{}

	for i := 1; i < len(n); i++ {
		if n[i] == val {
			count++
		} else {
			ret = append(ret, strconv.Itoa(count))
			ret = append(ret, val)
			val = n[i]
			count = 1
		}
	}

	ret = append(ret, strconv.Itoa(count))
	ret = append(ret, val)

	return ret
}

func countAndSay(n int) string {
	ret := []string{strconv.Itoa(1)}

	for i := 1; i < n; i++ {
		ret = read(ret)
	}

	return strings.Join(ret, "")
}
