package LeetCode

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 获得入边表
	inDegree := make([]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		inDegree[prerequisites[i][1]] += 1
	}
	// 将入边为0的节点入队
	queue := make([]int, 0, numCourses) // 创建长度为0容量为n的int切片当作队列
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	// 遍历其所有出边指向的节点,并将对应的入度-1，若入度变为0则入队
	cnt := 0
	for len(queue) > 0 {
		cnt++
		v := queue[0]
		queue = queue[1:]

		for i1 := 0; i1 < len(prerequisites); i1++ {
			if prerequisites[i1][0] == v {
				inDegree[prerequisites[i1][1]] -= 1
				//若因为访问入度变为0则入队
				if inDegree[prerequisites[i1][1]] == 0 {
					queue = append(queue, prerequisites[i1][1])
				}
			}
		}
	}
	return cnt == numCourses
}
