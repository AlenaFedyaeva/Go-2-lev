package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func doWorkContext(ctx context.Context) {
	fmt.Println("ctxT")
	ctxT, cancelFunc := context.WithTimeout(ctx, 30*time.Second)
	defer cancelFunc()
	cancelCh := make(chan os.Signal, 1)
	// catch SIGETRM or SIGINTERRUPT
	signal.Notify(cancelCh, syscall.SIGTERM, syscall.SIGINT)

	// Используем select для выхода по истечении времени жизни контекста
	select {
	case <-ctxT.Done():
		fmt.Println("ctx.Done: Time to exit")

	case sig := <-cancelCh:
		fmt.Printf("Caught SIGTERM %v", sig)
	}
}

func main() {
	// Создаем контекст background
	ctx := context.Background()
	// Производим контекст с отменой
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)

	defer func() {
		fmt.Println("Main Defer: canceling context")
		cancelFunction()
	}()

	
	// Выполнение работы
	doWorkContext(ctxWithCancel)
	fmt.Println("Bye!")
}
