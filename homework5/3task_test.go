package main

//go test -bench=. -benchmem 3task_test.go

// 3. Протестируйте производительность операций чтения
//  и записи на множестве
// действительных чисел, безопасность
//  которого обеспечивается sync.Mutex и sync.RWMutex для разных
//  вариантов использования: 10% запись,
//   90% чтение; 50% запись, 50% чтение; 90% запись, 10% чтение

import (
	"fmt"
	"sync"
	"testing"
)

type SetMutex struct {
	sync.Mutex
	mm map[int]struct{}
}

func NewSet() *SetMutex {
	return &SetMutex{
		mm: map[int]struct{}{},
	}
}

func (s *SetMutex) Add(i int) {
	s.Lock()
	defer s.Unlock()
	s.mm[i] = struct{}{}

}

//операция чтения
func (s *SetMutex) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.mm[i]
	return ok
}

func benchmarkReadWrite(input float32, b *testing.B) {
	b.StopTimer()
	num := float32(b.N) * input
	var set = NewSet()
	fmt.Println("N ", b.N, num)
	b.StartTimer()

	for n := 0; n < int(num); n++ {
		set.Add(int(n))
	}
	for n := num; n < float32(b.N); n++ {
		set.Has(int(n))
	}
}
func Benchmark90(b *testing.B) {
	benchmarkReadWrite(0.9, b)
}

func Benchmark50(b *testing.B) {
	benchmarkReadWrite(0.5, b)
}

func Benchmark10(b *testing.B) {
	benchmarkReadWrite(0.1, b)
}

// -------------------RWMUTEX-------------------------------------------------
  type SetRWMutex struct {
	sync.RWMutex
	mm map[int]struct{}
  }

  func(s *SetRWMutex) AddRWMutex(i int) {
	s.Lock()
	s.mm[i] = struct{}{}
	s.Unlock()
  }

  func(s *SetRWMutex) HasRWMutex(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.mm[i]
	return ok
  }

  func NewRWSet() *SetRWMutex {
	return &SetRWMutex{
	   mm: map[int]struct{}{},
	}
  }

  func benchmarkRWReadWrite(input float32, b *testing.B) {
	
	b.StopTimer()
	num := float32(b.N) * input
	var set = NewRWSet()
	fmt.Println("N ", b.N, num, b.N, b.N)
	b.StartTimer()

	for n := 0; n < int(num); n++ {
		set.AddRWMutex(int(n))
	}
	for n := num; n < float32(b.N); n++ {
		set.HasRWMutex(int(n))
	}
}
func BenchmarkRW90(b *testing.B) {
	benchmarkRWReadWrite(0.9, b)
}

func BenchmarkRW50(b *testing.B) {
	benchmarkRWReadWrite(0.5, b)
}

func BenchmarkRW10(b *testing.B) {
	benchmarkRWReadWrite(0.1, b)
}
