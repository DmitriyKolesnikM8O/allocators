package main

import "testing"

// go test -bench=. benchmem

const bufferSize = 10

type ReaderWithSliceArguments struct {}

func (r ReaderWithSliceArguments) Read(p []byte) (int, error) {
	for i := 0; i < bufferSize; i++ {
		p[i] = byte(i)
	}

	return bufferSize, nil
}

type ReaderWithSliceReturn struct {}

func (r ReaderWithSliceReturn) Read(n int) (int, error) {
	p := make([]byte, n) // this allocation goes by heap
	for i := 0; i < bufferSize; i++ {
		p[i] = byte(i)
	}

	return bufferSize, nil
}

func BenchmarkExampleTwo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		p := make([]byte, bufferSize)
		reader := ReaderWithSliceArguments{}
		_, _ = reader.Read(p)
	} 
}

func BenchmarkExampleTwo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reader := ReaderWithSliceReturn{}
		_, _ = reader.Read(bufferSize)
	} 
}










