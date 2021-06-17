package main

// GODEBUG=schedtrace=500 GOMAXPROCS=2 go run task2.go

import (
	"log"
	"runtime"
)

func main() {

	for i := 0; i < 2; i += 1 {
		go func(num int) {
			
			for i := 0; ; i += 1 {
				log.Println("I'm working! num ", num)
				if i%1e6 == 0 {
					runtime.Gosched()
				}
			}
		}(i)
	}

	for i := 0; ; i += 1 {
		log.Println("I'm working! main ")
		if i%1e6 == 0 {
			runtime.Gosched()
		}
	}
}
