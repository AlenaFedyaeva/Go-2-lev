package main

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"time"
)

var(
	ErrTask2 = errors.New("err task2")
	ErrTask3 = errors.New("err create file task3")
)

type MyEmplicitErr struct{
	Time time.Time
}

func (e *MyEmplicitErr) getTime() time.Time{
	return e.Time
} 

func (e *MyEmplicitErr) Error() string {
	return fmt.Sprintf("hw1: error time:  %s",e.Time.String())
}


func main(){
	fmt.Println("HW1")
	//1
	task1()
	//2
	err:=task2()
	fmt.Println("t2 -- ",err)
	
	//3
	err=task3()
	fmt.Println("t3 -- ",err)

	fmt.Println("after panic")
}

func task1(){
	defer panicHandler()

	implicitPanic()
}

func task2() (err error){
	defer func ()  {
		if v := recover(); v != nil {
			buff := make([]byte, 1024)
			runtime.Stack(buff, false)
			err = &MyEmplicitErr{Time: time.Now(),}
			
		}
		
	}()

	panic("task_2")
	
}


func task3() error{

	file, err := os.Create("3.txt")

	if err != nil{                          // если возникла ошибка
        fmt.Println("Unable to create file:", err) 
        return fmt.Errorf("%w: %s", ErrTask3, err )// 
    }
    defer file.Close()                      // закрываем файл
    
	fmt.Println(file.Name()) 

	return nil             

}

func implicitPanic(){
	arr :=[]int{1,2,3,4,5,6,7,8,9,10}
	for i := 0; i <= len(arr); i++ {
		fmt.Println(arr[i])
	}
} 

func panicHandler(){
	if v := recover(); v != nil {
		buff := make([]byte, 1024)
		runtime.Stack(buff, false)
		fmt.Printf("panic wiht value: %v, %s\n", v,buff)
	}
}