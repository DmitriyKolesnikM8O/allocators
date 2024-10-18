package main

import (
	"fmt"
	"errors"
	"unsafe"
)

type PoolAllocator struct {
	objectPool []byte
	freeObjects map[unsafe.Pointer]struct{}
	objectSize int
}

func New(capacity int, objectSize int) (PoolAllocator, error) {
	if (capacity <= 0 || objectSize <= 0 || capacity%objectSize != 0) {
		return PoolAllocator{}, errors.New("Incorrect capacity or objectSize")
	}

	allocator := PoolAllocator{
		objectPool: make([]byte, capacity),
		objectSize: objectSize,
		freeObjects: make(map[unsafe.Pointer]struct{}, capacity/objectSize),
	}
	
	allocator.resetMemoryState()
	
	return allocator, nil
}

func (a *PoolAllocator) Allocate() (unsafe.Pointer, error) {
	if len(a.freeObjects) == 0 {
		return nil, errors.New("Failed memory allocate")
	}

	var pointer unsafe.Pointer
	for freePointer := range a.freeObjects{
		pointer = freePointer
		break
	}

	return pointer, nil
}

func (a *PoolAllocator) Deallocate(pointer unsafe.Pointer) error {
	if pointer == nil {
		return errors.New("Failed deallocate memory")
	}

	a.freeObjects[pointer] = struct{}{}
	return nil
}

func (a *PoolAllocator) Free() {
	a.resetMemoryState()
}

func (a *PoolAllocator) resetMemoryState() {
	for offset := 0; offset < len(a.objectPool); offset += a.objectSize {
		pointer := unsafe.Pointer(&a.objectPool[offset])
		a.freeObjects[pointer] = struct{}{}
	}
}

func store[T any](pointer unsafe.Pointer, value T) {
	*(*T)(pointer) = value
}

func load[T any](pointer unsafe.Pointer) T {
	return *(*T)(pointer)
}

func main() {
	KB := 1 << 10
	allocator, err := New(KB, 4)
	if err != nil {
		fmt.Printf("Can`t create allocator: %s\n", err)
	}

	defer allocator.Free()

	pointer1, _ := allocator.Allocate()
	pointer2, _ := allocator.Allocate()
	store[int16](pointer1, 100)
	store[int32](pointer2, 200)

	value1 := load[int16](pointer1)
	value2 := load[int32](pointer2)

	fmt.Println("Value 1:", value1)
	fmt.Println("Value 2:", value2)

	fmt.Println("Pointer 1:", pointer1)
	fmt.Println("Pointer 2:", pointer2)

	allocator.Deallocate(pointer1)
	allocator.Deallocate(pointer2)
}
