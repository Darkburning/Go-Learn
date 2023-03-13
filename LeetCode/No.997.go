package LeetCode

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
