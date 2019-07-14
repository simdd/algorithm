package problem

import "testing"

func TestCountAndSay(t *testing.T) {
	for i := 0; i <= 30; i++ {
		println(countAndSay(i))
	}
}
