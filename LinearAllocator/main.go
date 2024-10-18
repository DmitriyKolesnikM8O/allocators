package main

import (
	"errors"
	"unsafe"
	"fmt"
)

type LinearAllocator struct {
	data []byte
}

func New(capacity int) (LinearAllocator, error) {
	if capacity <= 0 {
		return LinearAllocator{}, errors.New("Incorrect capacity")
	}

	return LinearAllocator{
		data: make([]byte, 0, capacity),
	}, nil
}

func (a *LinearAllocator) Allocate(size int) (unsafe.Pointer, error) {
	previousLength := len(a.data)
	newLength := previousLength + size

	if newLength > cap(a.data) {
		return nil, errors.New("Can`t allocate memory")
	}

	a.data = a.data[:newLength]
	pointer := unsafe.Pointer(&a.data[previousLength])

	return pointer, nil
}

func (a *LinearAllocator) Free() {
	a.data = a.data[:0]
}

func store[T any](pointer unsafe.Pointer, value T) {
	*(*T)(pointer) = value
}

func load[T any](pointer unsafe.Pointer) T {
	return *(*T)(pointer)
}

func main() {
	MB := 1 << 20
	allocator, err := New(MB)
	if err != nil {
		fmt.Printf("Can`t create allocator: %s\n", err)
	}

	defer allocator.Free()

	pointer1, _ := allocator.Allocate(15)
	pointer2, _ := allocator.Allocate(4)
	store[int16](pointer1, 100)
	store[int32](pointer2, 200)

	value1 := load[int16](pointer1)
	value2 := load[int32](pointer2)

	fmt.Println("Value 1:", value1)
	fmt.Println("Value 2:", value2)

	fmt.Println("Pointer 1:", pointer1)
	fmt.Println("Pointer 2:", pointer2)

}
