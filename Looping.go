package main

import (
	"fmt"
)

/*
go中没有“，”运算符，不能在for中写i和j两个变量
 */
func main8() {
	for i:=0;i<5;i++ {
		fmt.Println(i)
	}
	i:=0
	for ;i<5;i++ {
		fmt.Println(i)
	}
	//go中同时使用i j的方法
	/*for i,j:=0,0;i<5;i,j=i+1,j+1 {
		fmt.Println(i)
	}*/
	fmt.Println("========================")
	for  {
		fmt.Println(i)
		i++
		if i==10 {
			break
		}
	}
	fmt.Println("========================")
	for i:=0;i<10;i++ {
		if i%2==0 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("========================")
	Loop:
	for i:=1;i<=3;i++ {
		for j:=1;j<=3;j++ {
			fmt.Println(i*j)
			if i*j>=3 {
				break Loop//指出想要跳出的地方
			}
		}
	}
	fmt.Println("========================")
	//s:=[]int{1,2,3}//arrays slice map string are same
	ss:="hello go"
	for k,v:=range ss{
		//look at collection s,and take each value one at a time,give us
		//key and value,then we rae going to be able to work with those values
		fmt.Println(k,string(v))
	}
	for _,v:=range ss{//只需要value的情况
		fmt.Println(string(v))
	}

}