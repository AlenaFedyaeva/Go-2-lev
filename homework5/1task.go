package main

import (
	"flag"
	"fmt"
	"sync"
	"sync/atomic"
)

// const count = 1000

func main() {

	var n int
	flag.IntVar(&n, "n", 1234, "flag n ")
	fmt.Println("f", n)
	var (
		counter int32

		// Создаем экземпляр
		wg = sync.WaitGroup{}
	)

	for i := 0; i < n; i += 1 {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&counter, 1)

			// Выполняем декремент семафора
			wg.Done()
		}()
	}
	// Ждем обнуления семафора
	wg.Wait()

	// Выводим показание общего счетчика
	fmt.Println(counter)
}
