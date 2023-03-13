package LeetCode

func containsDuplicate(nums []int) bool {
	myMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		myMap[nums[i]] = i
	}
	return len(nums) != len(myMap)
}
