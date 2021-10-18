package main

import (
	"fmt"
	"log"
)

/*
Defer
	delay a function execution to some future point
	当函数退出时，会查看是否有需要defer调用的语句
	defer并不是移到main函数的结尾执行，实际是在main函数执行之后执行（main函数返回之前）
	多个defer是顺序是栈的顺序（先进后出）

	一般使用defer关闭资源
	defer 使用时，defer后的函数已经接受的参数，只是延迟调用
	open 和 close可以一起使用

Panic
	go app can enter a state where it can no longer continue to run
	go donnot have exceptions
	panic在defer之后执行
	Donot use when a file cannt be open
	use for unrecoverable events---cannt btain TCp port for web server

	function will stop
		defer func will still fire
	if nothing handle panic ,program will exit


recover
	go app get painc,and save the rest of program
	used to recover from panics
	only useful in deferred func(一旦panic 只有defer部分会继续执行)
	current function will not continue ，but the higher functions in call stack will
 */

func main9() {
	/*fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")*/

	/*a,b:=1,0
	ans:=a/b
	fmt.Println(ans)//panic: runtime error: integer divide by zero*/

	/*fmt.Println("start")
	defer fmt.Println("middle")
	panic("something bad happened")//panic: something bad happened
	fmt.Println("end")*/

	fmt.Println("start")
	defer func() {
		if err:=recover(); err != nil {
			log.Println("Error",err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}
