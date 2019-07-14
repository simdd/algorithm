package problem

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		rest := target - nums[i]
		if _, ok := m[rest]; ok {
			return []int{m[rest], i}
		}
		m[nums[i]] = i
	}
	return nil
}
