package main

import (
	"fmt"
	"strconv"
)


/*
Variable declaration
	var i int //declared variable i and the type is int
	var i int=27
	i:=27 //go know what type i is
	k:=99.  float64 as not declared  99 , float64

	When declared at package-level,have to use the full declaration
		var j float32=28

	At package-level could create a block of variable

Redeclaration and shadowing
	shadowing:作用域的就近原则

	all variable have to be used!


Visibility
	3-level
		package-level and lower spelling first letter any file in the package can read this var
		package-level and up spelling first letter any file out of the package can read this var
		block scope

Naming conventions

Type conversions
	var i int=42
	fmt.Printf("%v,%T\n",i,i)
	var j float32
	j=float32(i)//a conversion function
	fmt.Printf("%v,%T\n",j,j)

	float to int may case lose number 42.5->42

	var k string
		k=string(i)
		*,string //查找42的Unicode值 是一个*

	use strconv package for strings
 */


//使用格式化的字符串包

var(
	l float32=55 //go knows how to change int to float32
 	actorName string="Elisabeth Sladen"
 	companion string="Sarah Jane Smith"
 	doctorNumber int = 3
 	season int = 11
)

var (
	counter int = 0
)

func main2() {// main package's main function 程序的入口
	var i int=42
	fmt.Printf("%v,%T\n",i,i)
	var j float32
	j=float32(i)
	fmt.Printf("%v,%T\n",j,j)
	var k string
	k=strconv.Itoa(i)
	fmt.Printf("%v,%T\n",k,k)

}