package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cancelChan := make(chan os.Signal, 1)
    // catch SIGETRM or SIGINTERRUPT
    signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)
    go func() {
        // start your software here. Maybe your need to replace the for loop with other code
        for {
            // replace the time.Sleep with your code
            log.Println("Loop tick")
            time.Sleep(time.Second)
        }
    }()
    sig := <-cancelChan
    log.Printf("Caught SIGTERM %v", sig)
    // shutdown other goroutines gracefully
    // close other resources
}
