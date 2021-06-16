package main

import (
	"fmt"
	"sync"
)


func main2() {
   var m1 sync.Mutex
	criticalSectionWithPanic(&m1)
	criticalSection(&m1)
}

func criticalSection(mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
}

func criticalSectionWithPanic(mutex *sync.Mutex) {
	defer func() {
		fmt.Println("recovered", recover())
      mutex.Unlock()
	}()
	mutex.Lock()
	panic("AAA!")
}
