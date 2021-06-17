package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

// GOMAXPROCS=1 go run task1.go 2>traceTask1.out
// go tool trace traceTask1.out
const count = 1000

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()
	
	var (
		counter int
		mutex   sync.Mutex

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
	time.Sleep(2 * time.Second)
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
