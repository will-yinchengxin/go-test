package bench

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var StrData = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "a", "b", "c", "d", "e", "f", "g", "h", "i", "g", "k", "l", "m", "n"}

// 直接使用“+”号拼接
func BenchmarkStringsAdd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var s string
		for _, v := range StrData {
			s += v
		}
	}
	b.StopTimer()
	/*
	   GOMAXPROCS: 8
	   goos: darwin
	   goarch: arm64
	   pkg: owntest/bench
	   BenchmarkStringsAdd
	   BenchmarkStringsAdd-8   	 2198286	       543.3 ns/op
	   PASS
	*/
}

// fmt.Sprint拼接
func BenchmarkStringsFmt(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var _ string = fmt.Sprint(StrData)
	}
	b.StopTimer()
	/*
	   GOMAXPROCS: 8
	   goos: darwin
	   goarch: arm64
	   pkg: owntest/bench
	   BenchmarkStringsFmt
	   BenchmarkStringsFmt-8   	  830073	      1363 ns/op
	   PASS
	*/
}

// strings.Join拼接
func BenchmarkStringsJoin(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = strings.Join(StrData, "")
	}
	b.StopTimer()
	/*
		GOMAXPROCS: 8
		goos: darwin
		goarch: arm64
		pkg: owntest/bench
		BenchmarkStringsJoin
		BenchmarkStringsJoin-8   	 6169448	       191.7 ns/op
		PASS
	*/
}

// StringsBuffer拼接
func BenchmarkStringsBuffer(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n := len("") * (len(StrData) - 1)
		for i := 0; i < len(StrData); i++ {
			n += len(StrData[i])
		}
		var s bytes.Buffer
		s.Grow(n)
		for _, v := range StrData {
			s.WriteString(v)
		}
		_ = s.String()
	}

	b.StopTimer()
	/*
		GOMAXPROCS: 8
		goos: darwin
		goarch: arm64
		pkg: owntest/bench
		BenchmarkStringsBuffer
		BenchmarkStringsBuffer-8   	 8048534	       140.9 ns/op
		PASS
	*/
}

// 使用strings包里的builder类型拼接
func BenchmarkStringsBuilder(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n := len("") * (len(StrData) - 1)
		for i := 0; i < len(StrData); i++ {
			n += len(StrData[i])
		}

		var b strings.Builder
		b.Grow(n)
		b.WriteString(StrData[0])
		for _, s := range StrData[1:] {
			b.WriteString("")
			b.WriteString(s)
		}
		_ = b.String()
	}
	b.StopTimer()
	/*
		GOMAXPROCS: 8
		goos: darwin
		goarch: arm64
		pkg: owntest/bench
		BenchmarkStringsBuilder
		BenchmarkStringsBuilder-8   	 8317896	       144.6 ns/op
		PASS
	*/
}
