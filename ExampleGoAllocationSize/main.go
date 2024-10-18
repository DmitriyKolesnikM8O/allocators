package main

import (
	"testing"
	"fmt"
)

var (
	result string
	buffer []byte = make([]byte, 33)
)

func Function(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = string(buffer) + string(buffer)
}
}

/*
Explanation:
alocated 176: 48 bytes (not 33) + 48 bytes (not 33) + 80 bytes (not 66)
waste: 15 + 15 + 14 = 44 (25% waste)
*/
func main() {
	b := testing.Benchmark(Function)
	fmt.Println(b.AllocsPerOp()) //3
	fmt.Println(b.AllocedBytesPerOp()) //176
}
