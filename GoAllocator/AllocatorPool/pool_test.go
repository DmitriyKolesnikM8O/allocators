package main

// go test -bench=. -benchmem

import (
	"testing"
	"sync"
)

type Person struct {
	name string
}

type PersonPool struct {
	pool sync.Pool
}

func NewPersonPool() *PersonPool {
	return &PersonPool{
		pool: sync.Pool{
			New: func() interface{} { return new(Person) },
		},
	}
}

func (p *PersonPool) Get() *Person {
	return p.pool.Get().(*Person)
}

func (p *PersonPool) Put(person *Person) {
	p.pool.Put(person)
}

var gPerson *Person

func BenchmarkWithPool(b *testing.B) {
	pool := NewPersonPool()
	for i := 0; i < b.N; i++ {
		person := pool.Get()
		person.name = "Ivan"
		gPerson = person
		pool.Put(person)
	}
}

func BenchmarkWithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		person := &Person{name: "Ivan"}
		gPerson = person
	}
}

