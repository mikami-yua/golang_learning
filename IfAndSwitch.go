package main

import (
	"fmt"
	"math"
)

/*
if
	must use {} even only one line

switch
	不存在case穿透

 */
func main7() {
	statePopulations:=make(map[string]int)
	statePopulations=map[string]int{//key is string value is int
		"California": 392,
		"Texas":26,
		"Florida":206,
		"New York":197,
		"Ohio":116,
	}
	if pop,ok:=statePopulations["Florida"];ok{
		fmt.Println(pop)//206
	}

	number:=50
	guess:=30
	if guess <1 {
		fmt.Println("guess must between 1 ")
	}else if guess >100{
		fmt.Println("guess must between 100")
	} else {
		if guess < number {
			fmt.Println("too low")
		}
		if guess >number {
			fmt.Println("too high")
		}
		if guess==number {
			fmt.Println("right")
		}
	}
	/*if guess >=1 && guess <=100 {
		if guess < number {
			fmt.Println("too low")
		}
		if guess >number {
			fmt.Println("too high")
		}
		if guess==number {
			fmt.Println("right")
		}
	}*/
	fmt.Println(number<=guess)
	myNum:=0.123
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum),2)-1)<0.001 {
		fmt.Println("these are the same")//go中的浮点是近似的
	}else {
		fmt.Println("these are different")
	}

	//switch
	switch i:=2+2;i {
	case 1:
		fmt.Println("one")
	case 2,4,6:
		fmt.Println("2 4 6")
	default:
		fmt.Println("not one or two")
	}

	a:=10
	switch  {
	case a<=10:
		fmt.Println("minner than 10")
		fallthrough//提供case穿透功能,且不需要判断下一条语句是什么
	case a>=20:
		fmt.Println("minner than 20")
	case a<=50:
		fmt.Println("minner than 20")
	default:
		fmt.Println("greater than 20")
	}

	//type switch
	var kk interface{}=1.0//分配给interface任何东西都是允许的
	switch kk.(type) {//kk.(type) 告诉go找到接口的底层类型，并且为之后的使用提供方便
	case int:
		fmt.Println("Integer")
	case float64:
		fmt.Println("float64")
		break
		fmt.Println("this will not exe after break keyword")
	case string:
		fmt.Println("string")
	case [3]int:
		fmt.Println("this is a [3] array")
	default:
		fmt.Println("not in this")
	}
}

