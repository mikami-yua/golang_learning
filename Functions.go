package main
//4:21:33

import (
	"fmt"
)

/*
函数第一个字母的大小写决定了函数的可见性

go具有将局部变量作为指针返回的能力
	go执行return语句时：actually copied that result to another variable
	go的函数执行结束之后弹出栈与别的语言不同之处在于
		when go reconizes that we have returning a value taht's generated
		-on the local stack,it's automatically going to promote this variadble
		-for you to the shared memory in the computer,what's also called the heap
		-memory. we donnot have to be worry about this value being cleared

函数本身可以被是为types，passed around as variables
	将函数作为变量时需要注意函数是否已经被声明
*/


func main11() {
	sayMessage("hello go")
	for i:=0;i<5;i++ {
		sayIndex("hello",i)
	}
	greeting:="liil"
	name:="ppp"
	sayGreeting(greeting,name)
	sayGreeting2(& greeting,&name)
	fmt.Println(name)
	fmt.Println("===================")
	sum("hill",1,2,3,4,5,6,7,8)
	fmt.Println("===================")
	fmt.Println(sum2(1,2,3,4,5,6,7,8))
	fmt.Println("===================")
	fmt.Println(*sum3(1,2,3,4,5,6,7,8))
	fmt.Println("===================")
	fmt.Println(sum4(1,2,3,4,5,6,7,8))
	fmt.Println("===================")
	d,err:=divide(5.0,3.0)
	if err!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
	fmt.Println("===================")
	//anonymous functions
	func(){
		msg:="anonymous functions"
		fmt.Println(msg)
	}()

	var f func()= func() {
		fmt.Println("heool go")
	}
	f()

	var divide func(float64, float64) (float64,error)
	divide= func(a ,b float64) (float64, error) {
		if b==0.0 {
			return 0.0,fmt.Errorf("cannot divide by zero")
		}else {
			return a/b,nil
		}
	}
	d1,err1:=divide(5.0,3.0)
	if err1!=nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d1)

	//method
	g:=greeter{
		greeting: "hello",
		name: "zhang",
	}
	g.greet()
}

func sayMessage(msg string)  {
	fmt.Println(msg)
}
func sayIndex(msg string,idx int)  {
	fmt.Println(msg,idx)
}

func sayGreeting(greeting,name string)  {
	//copy value,data is not going to be changed
	fmt.Println(greeting,name)
	name="ted"
	fmt.Println(name)
}

func sayGreeting2(greeting,name *string)  {
	//passing pointers to func,data will change
	fmt.Println(*greeting,*name)
	*name="ted"
	fmt.Println(*name)
}

//可变参数 variadic parameter 只能有一个，且位置需要在最后
func sum(msg string,values ...int)  {
	//tell go runtime to take in all of the last arguments that are passed in
	//-and wrap them up into a slice,that has a name of "values"
	fmt.Println(values)
	result:=0
	for _,v:=range values{
		result+=v
	}
	fmt.Println(result)
}
func sum2(values ...int) int {
	fmt.Println(values)
	result:=0
	for _,v:=range values{
		result+=v
	}
	return result
}
//go 返回指针
func sum3(values ...int) *int {
	fmt.Println(values)
	result:=0
	for _,v:=range values{
		result+=v
	}
	return &result
}

//name return value
func sum4(values ...int) (result int) {
	fmt.Println(values)
	for _,v:=range values{
		result+=v
	}
	return
}
//multiple return value
func divide(a,b float64) (float64,error)  {
	if b==0.0 {
		return  0.0,fmt.Errorf("cannot divide by zero")
	}
	return a/b,nil
}

//method---a functin that executes in context of a type
type greeter struct {
	greeting string
	name string
}

func (g greeter) greet()  {//传递拷贝的对象
	//调用greet方法时，greet method is oing to get a copy of greeter object
	//-and that's going to be given the name g in the context of method
	fmt.Println(g.name,g.greeting)
}

func (g *greeter) greet2()  {
	//传递真是的对象
	fmt.Println(g.name,g.greeting)
}
