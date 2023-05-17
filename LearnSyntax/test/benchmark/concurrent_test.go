package main

import "testing"

// cd benchmark then run command: go test -v -bench=".*"
// Ps:".*"是正则表达式语法，"."表示匹配出换行符和回车符外的单个其他字符，"*"表示匹配前面的表达式零或多次
// 因此上述命令指明运行所有的Benchmark测试函数

// testing.B是测试框架定义好的一个struct
// b.N会帮我动态调整执行操作的次数从而得到一个较好的测试效果
func BenchmarkConcurrentAtomicAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcurrentAtomicAdd()
	}
}
func BenchmarkConcurrentMutexAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcurrentMutexAdd()
	}
}
