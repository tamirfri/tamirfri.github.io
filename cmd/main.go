package main

import (
	"fmt"
	"math"
	"math/rand"
	"syscall/js"
	"time"
)

func monteCarloPi(n int) (s float64) {
	for i := 0; i < n; i++ {
		s += math.Abs(rand.NormFloat64())
	}
	s /= float64(n)
	return 2 / (s * s)
}

func main() {
	t := time.Now()
	out := fmt.Sprintf("Pi ~ %.16f (took %v)\n", monteCarloPi(80_000_000), time.Since(t))
	fmt.Print("log:", out)
	js.Global().Get("document").Get("body").Set("textContent", js.ValueOf(out))
}
