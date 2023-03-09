package main

// No.26 easy
func removeDuplicates(nums []int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}

// No.216 easy
func containsDuplicate(nums []int) bool {
	myMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		myMap[nums[i]] = i
	}
	return len(nums) != len(myMap)
}

// No.997
func findJudge(n int, trust [][]int) int {
	if n == len(trust) {
		return -1
	}

	type degree struct {
		inDegree  int
		outDegree int
	}
	var count [1001]degree
	// count
	for i := 0; i < len(trust); i++ {
		count[trust[i][0]].outDegree++
		count[trust[i][1]].inDegree++
	}
	for i := 1; i < n+1; i++ {
		if count[i].inDegree == n-1 && count[i].outDegree == 0 {
			return i
		}
	}
	return -1
}
