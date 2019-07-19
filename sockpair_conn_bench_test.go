package sockpairconn

import (
	"net"
	"testing"
)

func BenchmarkScoketPairConn128InConcurrent(b *testing.B) {
	benchmarkScoketPairConnInConcurrent(128, b)
}

func BenchmarkScoketPairConn1024InConcurrent(b *testing.B) {
	benchmarkScoketPairConnInConcurrent(1024, b)
}

func BenchmarkScoketPairConn2048InConcurrent(b *testing.B) {
	benchmarkScoketPairConnInConcurrent(2048, b)
}

func BenchmarkScoketPairConn4096InConcurrent(b *testing.B) {
	benchmarkScoketPairConnInConcurrent(4096, b)
}

func BenchmarkScoketPairConn10240InConcurrent(b *testing.B) {
	benchmarkScoketPairConnInConcurrent(10240, b)
}

func BenchmarkScoketPairConn128(b *testing.B) {
	benchmarkScoketPairConn(128, b)
}

func BenchmarkScoketPairConn1024(b *testing.B) {
	benchmarkScoketPairConn(1024, b)
}

func BenchmarkScoketPairConn2048(b *testing.B) {
	benchmarkScoketPairConn(2048, b)
}

func BenchmarkScoketPairConn4096(b *testing.B) {
	benchmarkScoketPairConn(4096, b)
}

func BenchmarkScoketPairConn10240(b *testing.B) {
	benchmarkScoketPairConn(10240, b)
}

func BenchmarkNetPipeConn128InConcurrent(b *testing.B) {
	benchmarkNetPipeConn(128, b)
}

func BenchmarkNetPipeConn1024InConcurrent(b *testing.B) {
	benchmarkNetPipeConn(1024, b)
}

func BenchmarkNetPipeConn2048InConcurrent(b *testing.B) {
	benchmarkNetPipeConn(2048, b)
}

func BenchmarkNetPipeConn4096InConcurrent(b *testing.B) {
	benchmarkNetPipeConn(4096, b)
}

func BenchmarkNetPipeConn128(b *testing.B) {
	benchmarkNetPipeConn(128, b)
}

func BenchmarkNetPipeConn1024(b *testing.B) {
	benchmarkNetPipeConn(1024, b)
}

func BenchmarkNetPipeConn2048(b *testing.B) {
	benchmarkNetPipeConn(2048, b)
}

func BenchmarkNetPipeConn4096(b *testing.B) {
	benchmarkNetPipeConn(4096, b)
}

func benchmarkScoketPairConn(size int, b *testing.B) {
	indata := make([]byte, 1024)
	outdata := make([]byte, 1024)

	sp0, sp1, err := NewSocketPairConn()
	if err != nil {
		b.Fatal(err)
	}

	go func() {
		for {
			_, err := sp1.Read(outdata)
			if err != nil {
				break
			}
		}
	}()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n, _ := sp0.Write(indata)
		b.SetBytes(int64(n))
	}
}

func benchmarkScoketPairConnInConcurrent(size int, b *testing.B) {
	indata := make([]byte, 1024)
	outdata := make([]byte, 1024)

	sp0, sp1, err := NewSocketPairConn()
	if err != nil {
		b.Fatal(err)
	}

	go func() {
		for {
			_, err := sp1.Read(outdata)
			if err != nil {
				break
			}
		}
	}()

	b.ReportAllocs()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			n, _ := sp0.Write(indata)
			b.SetBytes(int64(n))
		}
	})
}

func benchmarkNetPipeConn(size int, b *testing.B) {
	indata := make([]byte, 1024)
	outdata := make([]byte, 1024)

	sp0, sp1 := net.Pipe()

	go func() {
		for {
			_, err := sp1.Read(outdata)
			if err != nil {
				break
			}
		}
	}()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n, _ := sp0.Write(indata)
		b.SetBytes(int64(n))
	}
}
