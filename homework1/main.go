package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

var (
	ErrTask2 = errors.New("err task2")
	ErrTask3 = errors.New("err create file task3")
)

type MyImplicitErr struct {
	time  time.Time
	trace string
	msg   string
}

func (e *MyImplicitErr) getTime() time.Time {
	return e.time
}
func (e *MyImplicitErr) getMsg() string {
	return e.trace
}

func New(msg string) error {
	return &MyImplicitErr{
		time:  time.Now(),
		trace: string(debug.Stack()),
		msg:   msg,
	}
}

func (e *MyImplicitErr) Error() string {
	return fmt.Sprintf("hw1: error time:  %s", e.time.String())
}

func main() {
	fmt.Println("HW1")
	//1
	task1()
	//2
	err := task2()
	fmt.Println("t2 -- ", err)

	//3
	err = task3()
	fmt.Println("t3 -- ", err)

	fmt.Println("after panic")
}

func task1() {
	//defer panicHandler()

	panicingFunc()
}

func task2() (err error) {
	defer func() {
		if v := recover(); v != nil {
			err = New("text err task2")

		}

	}()

	panic("task_2")

}

func task3() error {

	file, err := os.Create("3.txt")

	if err != nil { // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		return fmt.Errorf("%w: %s", ErrTask3, err) //
	}
	defer file.Close() // закрываем файл

	fmt.Println(file.Name())

	return nil

}

func panicingFunc() (err error) {
	defer panicHandler(func(e error) { err = e })

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i <= len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func panicHandler(setErr func(error)) {
	if v := recover(); v != nil {
		buff := make([]byte, 1024)
		runtime.Stack(buff, false)
		setErr(fmt.Errorf("panic wiht value: %v, %s\n", v, buff))
	}
}
