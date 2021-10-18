package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
go routine

Create
	go sayHello()
	most language use OS thread,that means that they've got an individual
	-function call stack dedicated to the execution of whatever code is handed
	-to that thread.they tend to be very large.For example,about 1MB of RAM
	-they take quite a bit of time for the application to set up.That is why we get into
	-concepts of thread pooling,the creation and destruction of threads is very expensive

	GreenThread
		instead of creating these very massive heavy overhead threads,we're
		-create an abstraction of a thread that we're going to call a go routine

		inside of go runtime,we have got a scheduler that's going to map these go routines
		-onto these OS threads for periods of time,and the scheduler will then take
		-turns with every CPU thread that's available and assign the different go routines a
		-certain amount of processing time on those threads
		we don't have to interact with those low level thread directly
		we interacting with these high level go routines

		advantage:
			this abstraction go routines can start with very small stack spaces,
			-because they can reallocated(重新分配) very quickly.

	when create go routine must know how it will end!
	Check for race conditions at compile time

Synchronization(同步)
	WaitGroup
		synchronize multiple o routines together
		synchronize main function to this anonymous fo routine

		when the go rutine is done,it can tel the wait group taht it's actually
		-completed its execution
	Mutexes(互斥锁)

Parallelism(并行性)不同于concurrency（并发性）
	concurrency：the ability of the application to work on multiple things at
			-the same time.doesn't mean it can work on them at the same time.
			-just means has multiple things that it can be doing
	Parallelism: take go applications and enable them to work on those concurrent
			-work calculations in parallel,introduce parallelism into our applications

 */
var wg = sync.WaitGroup{}//使用{}初始化
//这个程序有两个go routine 一个main一个内置函数

func main13() {
	//create go routine
	//go sayHello()//tell go to spin off what's called a green thread,and run
	//-this function in that green thread
	//nothing print: main function is actually executing in a go routine itself
	//told main function to generate another go routine,but application exits
	time.Sleep(100)

	var msg = "miaomiao"
	wg.Add(1)
	go func(msg string){//解耦合main from go routine
		fmt.Println(msg)
		wg.Done()//tell wait group it has complete,会将前面的1 再-1 当为0时可以继续执行
	}(msg)
	msg="goodbye"//goodbye printed out in the go routine
	//beacuse most of time the ghost scheduler is not going to interrupt the main
	//-thread until it hit this sleep call.Which means even launch another go routine
	//-it dosen't actually give it any love yet,it still executing the main function
	//-it is still executing the main function,and it actually gets to
	//-reassigns(重新赋值) the value of msg before the go routine has a chance
	//-to print it out
	//this is actually creating what's called a race condition

	wg.Wait()//等待wg add的数 都通过done减完了就行

	/*go func(msg string){//解耦合main from go routine
		fmt.Println(msg)
	}(msg)
	msg="hav hav"//not change
	time.Sleep(100)*/
}

func sayHello() {
	fmt.Printf("hello #%v\n",counter2)
	wg.Done()
}

func increment()  {
	counter2++
	wg.Done()
}

var counter2 = 0

func main14() {
	for i:=0;i<10;i++ {
		wg.Add(2)
		go sayHello()
		go increment()
		/* go routine racing against
		have no synchronization between go routines
		hello #1
		hello #0
		hello #4
		hello #5
		hello #6
		hello #8
		hello #2
		hello #9
		hello #9
		hello #8
		 */
	}
	//启动20个go routine
	wg.Wait()
}

var m = sync.RWMutex{}//互斥锁（Read Write Mutex）
//many can read only one can write,cannot writing while reading

func main15() {
	runtime.GOMAXPROCS(100)
	/*for i:=0;i<10;i++{
		wg.Add(2)
		go satHello2()
		go increment2()
	}*/
	/* satHello2执行后increment2并没有紧接着执行
	hello #0
	hello #1
	hello #1
	hello #1
	hello #1
	hello #1
	hello #1
	hello #1
	hello #1
	hello #1

	需要在 go routine的上下问之外锁定互斥锁

	*/
	for i:=0;i<10;i++{
		wg.Add(2)
		m.RLock()
		go satHello2()
		m.Lock()
		go increment2()
	}
	/*
	this is working is actually locking the mutex in a single context
	-So,the mian function is actually executing the locks
	-and then asynchronously(异步的) will unlock them once done with the
	-asynchronously operation(一旦完成异步操作就解锁)

	destroyed concurrency


	hello #0
	hello #1
	hello #2
	hello #3
	hello #4
	hello #5
	hello #6
	hello #7
	hello #8
	hello #9
	 */

	wg.Wait()

	fmt.Printf("theads: %v\n",runtime.GOMAXPROCS(-1))
	//默认分配的最大线程数和机器和核心数相同
	//设置为Integer数就是最多几个线程
}

func satHello2()  {
	fmt.Printf("hello #%v\n",counter2)
	m.RUnlock()
	wg.Done()
}

func increment2()  {
	counter2++
	m.Unlock()
	wg.Done()
}