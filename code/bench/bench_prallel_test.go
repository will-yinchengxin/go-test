package bench

//func init() {
//	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(-1))
//}
//func BenchmarkParallelHash(b *testing.B) {
//	testBytes := make([]byte, 1024)
//	// ResetTimer将基准时间和内存分配计数器归零 删除用户报告的指标
//	b.ResetTimer()
//	// SetParallelism 设置 RunParallel 使用的 goroutine 的数量为 p*GOMAXPROCS, cpu 绑定的基准测试通常不需要调用 SetParallelism。
//	// 如果 p 小于1，这个调用将没有效果
//	b.SetParallelism(20)
//	b.RunParallel(func(pb *testing.PB) {
//		x := sha256.New()
//		for pb.Next() {
//			x.Reset()
//			x.Write(testBytes)
//			_ = x.Sum(nil)
//		}
//	})
//	/*
//		GOMAXPROCS: 8
//		goos: darwin
//		goarch: arm64
//		pkg: owntest/bench
//		BenchmarkParallelHash
//		BenchmarkParallelHash-8   	11700856	       103.7 ns/op
//	*/
//}
//func BenchmarkNonParallelHash(b *testing.B) {
//	testBytes := make([]byte, 1024)
//	b.ResetTimer()
//	x := sha256.New()
//	for i := 0; i < b.N; i++ {
//		x.Reset()
//		x.Write(testBytes)
//		_ = x.Sum(nil)
//	}
//	/*
//		GOMAXPROCS: 8
//		goos: darwin
//		goarch: arm64
//		pkg: owntest/bench
//		BenchmarkNonParallelHash
//		BenchmarkNonParallelHash-8   	 1889756	       637.5 ns/op
//		PASS
//	*/
//}
