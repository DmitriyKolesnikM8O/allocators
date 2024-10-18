package main

import "testing"

//go test -bench=. -benchmem

type Data struct {
	intValue    int
	stringValue string
	boolValue   bool
}

//go:noinline
func NewByValue() Data {
	return Data{intValue: 100, stringValue: "100", boolValue: false}
}

//go:noinline
func NewByPointer() *Data {
	return &Data{intValue: 100, stringValue: "100", boolValue: false}
}

func BenchmarkExampleOne1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewByValue()
	}
}

func BenchmarkExampleOne2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = NewByPointer()
	}
}
