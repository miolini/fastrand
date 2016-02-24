package fastrand

import (
	"testing"
)

func BenchmarkMod(b *testing.B) {
	n1 := 1234231
	n2 := 35
	n3 := 0
	for i := 0; i < b.N; i++ {
		n3 = n1 % n2
	}
	n1 = n3
}

func BenchmarkFastRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastRand(64)
	}
}

func TestFastRand(t *testing.T) {
	var iters int = 1e9
	var shards = 4
	data := map[int]int{}
	for i := 0; i < iters; i++ {
		data[FastRand(shards)]++
	}
	var min int = 1e15
	var max = 0
	for _, counter := range data {
		if counter < min {
			min = counter
		}
		if counter > max {
			max = counter
		}
	}
	jitter := (max + min) / 2
	t.Logf("n %d, iters %d %.5f %d (%d-%d)", shards, iters, float64(jitter)/float64(iters/shards), jitter, min, max)
	t.Logf("%v", data)
}
