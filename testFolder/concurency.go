//Concurrency Lesson

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup //This is a variable for waitgroup

fmt.Println("OS\t", runtime.GOOS)
fmt.Println("ARCH\t", runtime.GOARCH)
fmt.Println("CPUS\t", runtime.NumCPU)
fmt.Println("Goroutines\t", runtime.NumGoroutine)

wg.Add(1)
go foo() //to launch a go routine, just add go to a function
bar()

fmt.Println("CPUS\t", runtime.NumCPU)
fmt.Println("Goroutines\t", runtime.NumGoroutine)
wg.Wait()

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar: ", i)
	}
}

//Using Channels and Concurrency
/*
		How to make a concurrent, self closing function.
		Step 1, make a channel,(The example is an 'int channel')
		Step 2, write a function like below and feed that
		channel the result of a function.
		Step 3, make sure to add a closing parenthseis, '()',
		so it closes when done.
	*/
	ch := make(chan int)
	go func() {
		ch <- doSomething(10)
		/*
			This symbol '<-' sends values and shit to this channel
		*/
	}()

	fmt.Println(<-ch)

	func doSomething(x int) int {
		//Does something, already built
		return x * 5
	}

//Race Condition
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

	for i := 0; i < gs; i++ {
		go func() {
			v := counter            // Counter is 0
			time.Sleep(time.Second) //wait 1 second
			runtime.Gosched()       //Allows other Goroutines to run as well
			v++
			counter = v
			wg.Done() //Finishes this Go Routine
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait() //This says, don't exit the program until everything is done.
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Count: ", counter)

	//Using Mutex
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

	//Atomic Lesson (It's a package, not sure what for)
	/*
		Different Go routines are accessing variables at the same time.
		This can be a problem called  a "Race". They 'write' over each other, depending on
		which routine is fastest.
	*/
	fmt.Println("CPU's: ", runtime.NumCPU())
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	var counter int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter \t", atomic.LoadInt64(&counter))
			wg.Done() //Finishes this Go Routine
		}()
		fmt.Println("Goroutines: ", runtime.NumGoroutine())
	}
	wg.Wait() //This says, don't exit the program until everything is done.
	fmt.Println("Goroutines: ", runtime.NumGoroutine())
	fmt.Println("Count: ", counter)