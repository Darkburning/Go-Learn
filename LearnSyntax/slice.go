package main

import "fmt"

// slice的append的原理,其实同C++的vector很像
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

// 移除下标为i的元素并保持顺序
// 若无需保持顺序则只需用最后一个元素覆盖对应下标为i的元素即可
func removeInt(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:]) // copy函数改变了内存空间，实际上此时slice长度不变
	return slice[:len(slice)-1]
}

// 传入的是指针,赋值需要解引用
func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
func reverseInt(slice []int) []int {
	for i := 0; i < len(slice)/2; i++ {
		swap(&slice[i], &slice[len(slice)-i-1])
	}
	return slice
}
func main() {
	// testAppendInt
	x := make([]int, 0, 2)
	x = appendInt(x, 1)
	fmt.Println(len(x), cap(x))
	x = appendInt(x, 2)
	fmt.Println(len(x), cap(x))
	x = appendInt(x, 3)
	fmt.Println(len(x), cap(x))
	x = appendInt(x, 4)
	fmt.Println(len(x), cap(x))
	x = appendInt(x, 5)
	fmt.Println(len(x), cap(x))
	x = appendInt(x, 6)
	//testRemoveInt
	fmt.Println(x)
	fmt.Println(removeInt(x, 2))
	fmt.Println(x)
	fmt.Println(len(x), cap(x))
	fmt.Println(reverseInt(x))

}
