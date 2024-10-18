package main

import "testing"

type DataForTestThree struct {
	pointer *int
}

func BenchmarkExampleThree1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var number int
		data := &DataForTestThree{
			pointer: &number,
		}
		_ = data
	}
}

func BenchmarkExampleThree2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var number int
		data := &DataForTestThree{}
		data.pointer = &number
		_ = data
	}
}
