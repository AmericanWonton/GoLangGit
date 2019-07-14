package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	/*
		Different Go routines are accessing variables at the same time.
		This can be a problem called  a "Race". They 'write' over each other, depending on
		which routine is fastest.
	*/
	fmt.Println("CPU's: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex //Locks down code to prevent race conditions.

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()               //Ensures that no other function can access the counter variable
			v := counter            // Counter is 0
			time.Sleep(time.Second) //wait 1 second
			runtime.Gosched()       //Allows other Goroutines to run as well
			v++
			counter = v
			mu.Unlock() //Lets other functions use this variable when done.
			wg.Done()   //Finishes this Go Routine
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait() //This says, don't exit the program until everything is done.
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Count: ", counter)

}
