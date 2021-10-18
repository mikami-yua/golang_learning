package main

import (
	"fmt"
)
/*
Creatin pointers

Dereferencing pointers
	is Basically using a pointer to get at some underlying data

the New function

Working with nil

Types with internal pointers
 */

func main10() {
	a:=42
	b:=a//copy
	a=27
	fmt.Println(a,b)
	var aa int=52
	var bb *int //a pointer to an integer
	bb=&aa //address of operator,bb point to aa
	//bb is a pointer to an integer and we pointed to aa
	fmt.Println(aa,bb)//52 0xc00000a0d0
	//bb holding the memory location that's holding aa's data
	fmt.Println(&aa,bb)//0xc00000a0d0 0xc00000a0d0
	//& address of operator will give us the address of memory

	//derefernecing operator:figure out what calue is actually being stroed at a memry location
	fmt.Println(aa,*bb)
	// * before the type:declaring a pointer to data of that type
	// * brefor a pointer:dereference,ask go runtime to through the pointer find the memory location taht the pointer is pointing to,
	// -and the pull the value back out
	*bb=33
	fmt.Println(aa,*bb)//33

	//pointer arthmetic---not allowed! cannot jump around memory
	//unsafe package 提供指针运算的功能
	a1:=[3]int{1,2,3}
	a2:=&a1[0]
	a3:=&a1[1]
	fmt.Println("%v %p %p \n",a1,a2,a3)//[1 2 3] 0xc000012150 0xc000012158

	//when we want t work with pointers ,we donnot really care where the underlying
	//-data is stroed,we just need the ability to point to it wherever it's at
	var ms *myStruct
	ms=&myStruct{foo: 42}
	fmt.Println(ms)//&{42}
	var ms2 *myStruct
	fmt.Println(ms2)//<nil>
	ms2=new(myStruct)
	fmt.Println(ms2)//&{0}
	(*ms2).foo=25
	fmt.Println(*ms2)//{25}
	fmt.Println(ms2.foo)//25---指针并没有这个字段，是指针指向的对象有这个字段,go的语法糖
	//complex types are automatically dereferenced

	////////////
	b1:=a1//copy,
	// slice---actually a projection of the underlying array,slice doesn't contain the data it self
	//-slice contains the pointer to the first element that slice is pointing to on the underlying array
	//-slice is still copying,but it copy the pointer,not the underlying data it self
	//-sharing slices is actually always pointing at the same underlying data
	//map is also in this way
	fmt.Println(a1,b1)
	a1[1]=99
	fmt.Println(a1,b1)//[1 99 3] [1 2 3]
}

type myStruct struct {
	foo int
}

