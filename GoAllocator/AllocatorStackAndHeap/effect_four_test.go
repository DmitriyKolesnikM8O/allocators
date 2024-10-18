package main

import "testing"

//go:noinline
func Initialize(value *int) {
	*value = 1000
}

func BenchmarkExampleFour1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var number int
		Initialize(&number)
	}
}

func BenchmarkExampleFour2(b *testing.B) {
	var number int
	for i := 0; i < b.N; i++ {
		Initialize(&number)
	}
}
