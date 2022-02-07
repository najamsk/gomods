package math

import (
	"fmt"
	"testing"

	"github.com/ericlagergren/decimal"
)

func TestIssue70(t *testing.T) {
	x, _ := new(decimal.Big).SetString("1E+41")
	x.Context.Precision = 1
	Log10(x, x)
	if x.Cmp(decimal.New(40, 0)) != 0 {
		t.Fatalf(`Log10(1e+41)
wanted: 40
got   : %s
`, x)
	}
}

/*
Benchmarks from "Handbook of Continued Fractions for Special Functions."

alg1 - 2.4.1
alg2 - 2.4.4
alg3 - 2.4.7

With a small value for x ("42.2") and a precision of 16, the number of iterations
was:

alg1 - 210
alg2 - 75
alg3 - 150

Code for the algorithm behind the benchmarks:
https://gist.github.com/ericlagergren/cc95be6530aec21e7f91e2204173fd4f

BenchmarkLog_alg1_9-4     	     500	   3209008 ns/op
BenchmarkLog_alg1_19-4    	     300	   5114247 ns/op
BenchmarkLog_alg1_38-4    	     200	  12034146 ns/op
BenchmarkLog_alg1_500-4   	       2	 535323033 ns/op
BenchmarkLog_alg2_9-4     	    3000	    478031 ns/op
BenchmarkLog_alg2_19-4    	    1000	   1954844 ns/op
BenchmarkLog_alg2_38-4    	     300	   4615867 ns/op
BenchmarkLog_alg2_500-4   	       5	 238076617 ns/op
BenchmarkLog_alg3_9-4     	    2000	   1043696 ns/op
BenchmarkLog_alg3_19-4    	     300	   4317666 ns/op
BenchmarkLog_alg3_38-4    	     200	   8040413 ns/op
BenchmarkLog_alg3_500-4   	       2	 550735383 ns/op
*/

var log_X, _ = new(decimal.Big).SetString("123456.789")

func BenchmarkLog(b *testing.B) {
	for _, prec := range benchPrecs {
		b.Run(fmt.Sprintf("%d", prec), func(b *testing.B) {
			b.ReportAllocs()
			z := decimal.WithPrecision(prec)
			for j := 0; j < b.N; j++ {
				Log(z, log_X)
			}
			gB = z
		})
	}
}

func BenchmarkLn10_CF(b *testing.B) {
	for _, prec := range benchPrecs {
		b.Run(fmt.Sprintf("%d", prec), func(b *testing.B) {
			b.ReportAllocs()
			z := decimal.WithPrecision(prec)
			for j := 0; j < b.N; j++ {
				ln10(z, prec)
			}
			gB = z
		})
	}
}

func BenchmarkLn10_Taylor(b *testing.B) {
	for _, prec := range benchPrecs {
		b.Run(fmt.Sprintf("%d", prec), func(b *testing.B) {
			b.ReportAllocs()
			z := decimal.WithPrecision(prec)
			for j := 0; j < b.N; j++ {
				ln10_t(z, prec)
			}
			gB = z
		})
	}
}
