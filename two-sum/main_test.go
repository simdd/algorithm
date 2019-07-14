package problem

import "testing"

func TestTwoSum(t *testing.T) {
	origin := [][]int{
		[]int{1, 2, 3, 4},
		[]int{2, 3, 5},
	}

	target := []int{
		6,
		7,
	}

	result := [][]int{
		[]int{1, 3},
		[]int{0, 2},
	}

	for i := 0; i < len(origin); i++ {
		if ret := twoSum(origin[i], target[i]); ret[0] != result[i][0] && ret[1] != result[i][1] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}
