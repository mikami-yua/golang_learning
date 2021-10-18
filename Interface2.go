package main

import (
	"fmt"
	"math"
)

/*
go 中接口是一组方法的墙面，当类型为接口中的所有方法提供定义，成为实现接口
	接口中只写方法的定义，方法的实现又类型提供

	go中接口和实现类的关系是非嵌入式的
		不需要声明那个类型实现了接口，默默的把接口当成方法实现就是实现类

当使用接口类型对象的时候，可以使用热议实现类对象代替
接口对象不能访问实现类的属性

接口最大的意义是解耦合
	把功能的定义和功能的实现分离开

go 通过接口模拟多态
	一个接口的实现
		1.看成实现本身的类型，能够访问实现类中的属性和方法
		2.看成对应的接口类型，只能访问接口中的方法
	接口类型可以赋值任意实现类，接口类容器可以存放任意实现类

空接口中不包含任何方法，所有类型都实现了空接口，空接口可以存储任意类型的数值

如果想定义一个函数，可以接受任意类型的数据，就能定义为空接口

接口嵌套
	想要实现最终的接口，需要把继承的接口都实现

接口断言
	判断当前接口对象具体是那个类型的实现

	1.instance,ok := 接口对象.(实际类型) //安全
	2.switch语句

*/
type A interface {
}

type Cat struct {
	color string
}
type Person struct {
	name string
	age  int
}

//a是空接口，可以接受任意类型的数据
func anyTypeFunc(a A) {
	fmt.Println(a)
}

//使用匿名空接口
func anyTypeFunc2(a interface{}) {
	fmt.Println(a)
}

func main() {
	//创建mouse类型
	m1 := Mouse{"logic"}
	fmt.Println(m1.name)
	f1 := FlashDisk{"sandisk"}
	fmt.Println(f1.name)
	f1.deleteData()

	testInterface(m1) //m1是mouse类型，但是已经实现了接口
	testInterface(f1)

	var usb USB
	usb = f1
	usb.start()
	//fmt.Println(usb.name)报错，name字段是mouse特有的不是接口的
	//接口不能访问类新增的属性
	fmt.Println("=======================")
	var a1 A = Cat{"huamao"}
	var a2 A = Person{"wang", 30}
	var a3 = "haha"
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	anyTypeFunc(a1)
	anyTypeFunc(100)

	//map,  key为string value为任意类型数据
	map1 := make(map[string]interface{})
	map1["name"] = "lixiaohau"
	map1["age"] = 23
	map1["friend"] = Person{"jone", 18} //map[age:23 friend:{jone 18} name:lixiaohau]
	fmt.Println(map1)

	//存储任意类型的切片
	slice1 := make([]interface{}, 0, 10) //长度0，容量10
	slice1 = append(slice1, a1, a2, a3, 100, "hello")
	fmt.Println(slice1)
	//[{huamao} {wang 30} haha 100 hello]
	anySlice(slice1)

	fmt.Println("=======================")
	//miao 到底是那个实现，看具体是怎么定义的
	var mmm Miao = Miao{}
	mmm.test1()
	mmm.test2()
	mmm.test3()

	var aaaaa AAA = mmm
	aaaaa.test1() //只能调用test1 方法aaaaa是AAA类型

	var ccc CCC = mmm
	ccc.test1()
	ccc.test2()
	ccc.test3()

	fmt.Println("=======================")
	//断言
	var t1 Triang = Triang{3, 4, 5}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())

	var c1 Circle = Circle{4}
	fmt.Println(c1.peri())
	fmt.Println(c1.area())

	var s1, s2 Shape
	s1 = t1
	s2 = c1
	getType(s1) //sanjiaoxing 3 4 5
	getType(s2)
	var t2 *Triang = &Triang{5, 6, 7}
	getType(t2) //not sanjiaoxing
	getType2(s2)

}

//接受任意类型的切片的函数
func anySlice(slice2 []interface{}) {
	for i := 0; i < len(slice2); i++ {
		fmt.Println(slice2[i])
	}
}

//1.定义接口
type USB interface {
	start() //usb设备开始工作
	end()   //usb设备结束工作
}

//2.实现类
type Mouse struct {
	name string
}
type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println("mouse is ok")
}

func (m Mouse) end() {
	fmt.Println("mouse is end")
}

func (f FlashDisk) start() {
	fmt.Println("flashdisk is ok")
}

func (f FlashDisk) end() {
	fmt.Println("flashdisk is end")
}

//在类中新增非接口方法
func (f FlashDisk) deleteData() {
	fmt.Println("all data is cleaned")
}

//测试方法，测试接口
func testInterface(usb USB) {
	usb.start()
	usb.end()
}

//接口嵌套
type AAA interface {
	test1()
}

type BBB interface {
	test2()
}

type CCC interface {
	AAA
	BBB
	test3()
}

type Miao struct {
}

func (m Miao) test1() { //实现test1 说明miao是接口AAA的实现
	fmt.Println("test1")
}

func (m Miao) test2() { //实现test2 说明miao是接口BBB的实现
	fmt.Println("test2")
}

func (m Miao) test3() { //实现test3 说明miao是接口CCC的实现
	fmt.Println("test3")
}

//断言
type Shape interface {
	peri() float64 //形状周长
	area() float64 //面积
}

type Triang struct {
	a float64
	b float64
	c float64
}

func (t Triang) peri() float64 {
	return t.a + t.b + t.c
}

func (t Triang) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
	return s
}

type Circle struct {
	radius float64
}

func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}
func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

//断言判断类型
func getType(s Shape) {
	if ins, ok := s.(Triang); ok {
		fmt.Println("sanjiaoxing", ins.a, ins.b, ins.c) //这里接口可以按具体实现类型访问
	} else if ins2, ok2 := s.(*Triang); ok2 {
		fmt.Println("sanjiaoxing", ins2.a, ins2.b, ins2.c)
	} else {
		fmt.Println("not sanjiaoxing")
	}

}

func getType2(s Shape) {
	switch ins := s.(type) {
	case Triang:
		fmt.Println("sanjiaoxing", ins.b, ins.a, ins.c)
	case Circle:
		fmt.Println("yuanxing", ins.radius)
	}
}
