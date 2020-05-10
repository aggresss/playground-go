package factorial

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	var dt = []struct {
		name       string
		in         int
		out        uint64
		shouldFail bool
	}{
		{"-1", -1, 0, true},
		{"0", 0, 1, false},
		{"1", 1, 1, false},
		{"2", 2, 2, false},
		{"3", 3, 6, false},
		{"10", 10, 3628800, false},
	}
	for _, dtt := range dt {
		t.Run(dtt.name, func(t *testing.T) {
			var (
				v = dtt.in
			)
			out, err := Factorial(v)
			if dtt.shouldFail {
				if err == nil {
					t.Fatal("should fail")
				}
			} else {
				if err != nil {
					t.Errorf("err: %s", err)
				}
				if out != dtt.out {
					t.Errorf("wrong output: %d", out)
				}
			}
		})
	}
}

func BenchmarkAppendFloat(b *testing.B) {
	benchmarks := []struct {
		name string
		in   int
	}{
		{"10", 10},
		{"100", 100},
		{"1000", 1000},
		{"10000", 10000},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				v := bm.in
				Factorial(v)
			}
		})
	}
}
