package LeetCode

func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}
	KMPTable := getKMPTable(needle)
	i, j := 0, 0
	for i < len(haystack) {
		if haystack[i] == needle[j] {
			i++
			j++
		} else {
			if j > 0 {
				j = KMPTable[j-1]
			} else {
				i++
			}
		}
		if j == len(needle) {
			return i - j
		}
	}
	return -1
}

func getKMPTable(needle string) []int {
	i, j := 1, 0
	KMPTable := make([]int, len(needle))
	KMPTable[0] = 0
	for i < len(needle) {
		if needle[i] == needle[j] {
			//j++自增符号是语句而不是表达式
			j++
			KMPTable[i] = j
			i++
		} else {
			if j > 0 {
				j = KMPTable[j-1]
			} else {
				KMPTable[i] = 0
				i++
			}
		}
	}
	return KMPTable
}
