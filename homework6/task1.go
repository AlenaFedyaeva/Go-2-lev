package main

import (
	"fmt"
	"sync"
	"time"
)
  
  const count = 1000
  
  func main() {
	var (
	   counter int
	   mutex sync.Mutex
  
	   // Вспомогательная часть нашего кода
	   ch = make(chan struct{}, count)
	)
	for i := 0; i < count; i += 1 {
	   go func() {
		  // Захват мьютекса
		  mutex.Lock()
		  counter += 1
		  // Освобождение мьютекса
		  mutex.Unlock()
  
		  // Фиксация факта запуска горутины в канале
		  ch <- struct{}{}
	   }()
	}
	time.Sleep(2*time.Second)
	close(ch)
  
	i := 0
	for range ch {
	   i += 1
	}
	// Выводим показание счетчика
	fmt.Println(counter)
	// Выводим показания канала
	fmt.Println(i)
  }
  