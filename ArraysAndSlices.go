package main

import "fmt"

//1:45:58
/*
Arrays
	arrays are actually considered as values(not pointings),when copy an array
	we are actually creating a literal copy,not pointing to the same underlying.
		when the size of array is huge,the copy could slow the program
		use pointer to solute this problem

	Creation
		has a fixed size,has to be known at compile time
		same type
	Built-in functions
	Working with arrays


	grades := [...]int{97,85,93}
	fmt.Printf("%v\n",grades)//[97 85 93]
	var student [3]string
	fmt.Printf("%v\n",student)
	student[0]="Lisa"
	fmt.Printf("%v\n",student)
	fmt.Printf("%v\n",len(student))
	var Matrix [3][3]int =[3][3]int{[3]int{1,0,0},[3]int{0,1,0},[3]int{0,0,1}}
	fmt.Println(Matrix)
	b:=grades
	b[1]=5
	fmt.Println(grades)
	c:=&grades //c is going to point the same data that grades has
	c[1]=100
	fmt.Println(grades)

Slices------remember there is only one underlying array
	grades := []int{97,85,93}
	a[1]=5
	fmt.Printf("%v\n",grades)//[97 5 93]
	a and grades are actually pointing to the same underlying

	Creation
		a:=[]int{1,2,3,4,5,6,7,8,9,10}
		b:=a[:]//slice of all elements------what ever a is an array or slice
		c:=a[3:]//slice form 4th element to end
		d:=a[:6]//slice first 6 elements
		e:=a[3:6]//slice the 4th,5th,and 6th elements
		//first num is include and second num is exclude

		a:=make([]int,3,100)
			100 is the cap,underlying is array copy,if we know the size the copy will not happen
		fmt.Println(a)
		fmt.Println(len(a))
		fmt.Println(cap(a))

	Built-in functions
		a:=[]int{}
		fmt.Println(a)
		fmt.Println(len(a))
		fmt.Println(cap(a))
		a=append(a,1)
		fmt.Println(a)
		fmt.Println(len(a))
		fmt.Println(cap(a))
		a=append(a,2,3,4,5,6,7)
		fmt.Println(a)
		fmt.Println(len(a))
		fmt.Println(cap(a))//8
		a=append(a,[]int{8,9,10,11,12}...)//add 2 slices
		fmt.Println(a)
		fmt.Println(len(a))
		fmt.Println(cap(a))//16


	Working with slice
		remove element from the inside of a slice,make sure we dont have
		any other references to that underlying array

[...] is an array
[] is a slice
 */

func main5() {
	a:=[]int{2,3,4,5,6}
	b:=a[1:]// as the stack pop
	fmt.Print(b)
	c:=a[:len(a)-1]// as the queue pop
	fmt.Println(c)

	//remove mid element
	d:=append(a[:2],a[3:]...)
	fmt.Println(d)//[2 3 5 6]
	fmt.Println(a)//[2 3 5 6 6]
}
