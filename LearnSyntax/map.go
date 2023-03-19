package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

// 按key值对map排序
func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
		"harry":   35,
		"bob":     40,
	}

	var names []string
	for name := range ages { // 默认去第一维遍历，可省略下划线
		fmt.Println(name)
		names = append(names, name)
	}
	// 将第一维读入names
	sort.Strings(names) // 对names按字母表顺序进行排序
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	// 如果key不存在会返回该类型对应的零值，但有时需要区分key存在返回的0和key不存在返回零值的0
	if age, ok := ages["apple"]; !ok {
		fmt.Println(age)
		fmt.Println("apple is not in ages!")
	}
	// 利用map实现一个简单的map
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
			fmt.Println(seen)
		}
		if line == "break" {
			os.Exit(1)
		}
	}
	// 错误处理
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}
