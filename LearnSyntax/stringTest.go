package main

import (
	"fmt"
	"unicode/utf8"
)

// 以下函数strings包均有实现
// 由于UTF-8是无前缀编码，因此不需要解码，直接判断即可
func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
func contains(s, subString string) bool {
	for i := 0; i < len(s)-len(subString); i++ {
		if hasPrefix(s[i:], subString) {
			return true
		}
	}
	return false
}

func main() {
	prefix := "abc"
	suffix := "efg"
	text1 := "abcdefh"
	text2 := "abdefg"
	text3 := "efikabcgh"
	fmt.Println("text1 has prefix:", hasPrefix(text1, prefix))
	fmt.Println("text2 has prefix:", hasPrefix(text2, prefix))
	fmt.Println("text1 has suffix:", hasSuffix(text1, suffix))
	fmt.Println("text2 has suffix:", hasSuffix(text2, suffix))
	fmt.Println("text3 contains prefix:", contains(text3, prefix))
	fmt.Println("text3 contains suffix:", contains(text3, suffix))

	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13" 因为使用的是rune，返回的是字节数
	fmt.Println(utf8.RuneCountInString(s)) //"9"	用以统计字符串中字符数目
	// 解码字符串
	for i, r := range "Hello, 世界" {
		fmt.Printf("%d\t%q\t%d\n", i, r, r) // %q返回字面量
	}

	// "program" in Japanese katakana
	s1 := "プログラム"
	fmt.Printf("% x\n", s1) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s1)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"
}
