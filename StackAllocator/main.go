package main

import (
	"fmt"
	"errors"
	"unsafe"
	"math"
)

const headerSize = 2

type Allocator struct {
	data []byte
}

func New(capacity int) (Allocator, error) {
	if capacity <= 0 {
		return Allocator{}, errors.New("Incorrect capacity")
	} 

	return Allocator{
		data: make([]byte, 0, capacity),
	}, nil
}

func(a *Allocator) Allocate(size int) (unsafe.Pointer, error) {
	if size > math.MaxInt16 {
		return nil, errors.New("Incorrect size")
	}

	previousLength := len(a.data)
	newLength := previousLength + size + headerSize

	if newLength > cap(a.data) {
		return nil, errors.New("Not enough memory")
	}

	a.data = a.data[:newLength]
	header := unsafe.Pointer(&a.data[previousLength])
	pointer := unsafe.Pointer(&a.data[previousLength + headerSize])

	*(*int16)(header) = int16(size)
	return pointer, nil
}

func (a *Allocator) Deallocate(pointer unsafe.Pointer) error {
	if pointer == nil {
		return errors.New("Incorrect pointer")
	}

	header := unsafe.Add(pointer, -headerSize)
	size := *(*int16)(header)

	previousLength := len(a.data)
	newLength := previousLength - headerSize - int(size)

	a.data = a.data[:newLength]
	return nil
}

func (a *Allocator) Free() {
	a.data = a.data[:0]
}



func store[T any](pointer unsafe.Pointer, value T) {
	*(*T)(pointer) = value
}

func load[T any](pointer unsafe.Pointer) T {
	return *(*T)(pointer)
}


func main() {
	KB := 1 << 10
	allocator, err := New(KB)
	if err != nil {
		fmt.Println("Errors when allocate memory", err)
	}

	defer allocator.Free()


	pointer1, _ := allocator.Allocate(2)
	defer allocator.Deallocate(pointer1)
	pointer2, _ := allocator.Allocate(3)
	defer allocator.Deallocate(pointer2)

	store[int16](pointer1, 100)
	store[int32](pointer2, 200)

	value1 := load[int16](pointer1)
	value2 := load[int32](pointer2)

	fmt.Println("value1:", value1)
	fmt.Println("value2:", value2)
	fmt.Println("pointer1:", pointer1)
	fmt.Println("pointer2:", pointer2)
}
