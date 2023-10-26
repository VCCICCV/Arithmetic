# Go语言算法库
# 本库用为VCCICCV学习记录,题库及解法来自力扣及网络
- 作者：力扣官方题解+网络
-  链接：https://leetcode.cn/problems/two-sum/solutions/434597/liang-shu-zhi-he-by-leetcode-solution/
-  来源：力扣（LeetCode）
-  著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
## 测试
```go
// 压测函数：测试函数运行平均耗时
func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	s := "abcabcbb" // 测试字符串
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(s) // 测试函数
	}
}
// 统计内存占用
func memoryUsage(t *testing.T) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	memUsage := float64(mem.Alloc) / (1024 * 1024)
	t.Logf("Memory usage for lengthOfLongestSubstring: %.4f MB", memUsage)
}
```