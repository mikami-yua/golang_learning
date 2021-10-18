package main

import "fmt"

/*
constants
	value
		cannot set constants equal to something that has to be determined
		at runtime.

		at the package const block can do some math

		constants values can be shadowed.
			inner constant shadows outer constant

		const myConst int=42
		fmt.Printf("%v,%T",myConst,myConst)
		var  b int16 = 27
		fmt.Printf("%v,%T", int(b)+myConst, int(b)+myConst) //(mismatched types int and int16)

	iota
		zero value is 0
		as a block counter

 */
const (
	a = iota
	b
	c
)

const (
	a2 = iota
)

const (
	_ = iota//ignore first calue by assigning to blank identifier
	KB=1<<(10*iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func main4() {
	fileSize:=4000000000.
	fmt.Printf("%.2fGB",fileSize/GB)

}
