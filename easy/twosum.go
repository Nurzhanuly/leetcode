package easy

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i := range nums {
		diff := target - nums[i]
		if v, exist := m[diff]; exist {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return nil
}
