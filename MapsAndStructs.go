package main

import "fmt"
import "reflect"

/*
Map-----order of map is not guaranteed
	creation
		1.
		statePopulations :=map[string]int{//key is string value is int
			"California": 392,
			"Texas":26,
			"Florida":206,
			"New York":197,
			"Ohio":116,
		}
		fmt.Println(statePopulations)
		//map[California:392 Florida:206 New York:197 Ohio:116 Texas:26]
		2.
		statePopulations:=make(map[string]int)
		statePopulations=map[string]int{//key is string value is int
			"California": 392,
			"Texas":26,
			"Florida":206,
			"New York":197,
			"Ohio":116,
		}
		fmt.Println(statePopulations)
		//map[California:392 Florida:206 New York:197 Ohio:116 Texas:26]

		make function can have 2 parameter
		make(map[string]int,10)
			an effective hashing function requires the size of the hash table to be at least
			twice the number of elements. Reserving the size beforehand preempts the program
			initialising a larger map to ensure hashing efficiency.

		m:= map[[3]int]string{}//arrays can be the key ,slice cannot

	manipulate
		map:= is not copy,is the pointer

		fmt.Println(statePopulations["Ohio"])//116
		statePopulations["Georgia"]=113
		fmt.Println(statePopulations)
		delete(statePopulations,"Georgia")//with the not exited key will return 0
		//the not exited key will return false
		fmt.Println(statePopulations)
		fmt.Println(len(statePopulations))//5
		sp:=statePopulations
		delete(sp,"Ohio")
		fmt.Println(statePopulations)

Structs
	1.
	type Doctor struct {
		number int
		actorName string
		companions []string
	}
	aDoctor:=Doctor{
		number: 3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	2.anonymous struct
	aDoctor:= struct {
		name string
	}{name:"jone"}

	unlike map and slice ,use "&" to pointer at
	anotherDoctor:=aDoctor
	anotherDoctor.name="lili"
	fmt.Println(aDoctor)//jone

	tag
		type Animal struct {
			Name string `required max:"100"`//tag
			Origin string
		}
		t:=reflect.TypeOf(Animal{})
		field,_ := t.FieldByName("Name")
		fmt.Println(field.Tag)//required max:"100"

go dosenot support traditional object oriented principles
	-donnot have inheritance(is a)
	-use a modle similar to inheritance called composition(has a)
		(embedding)嵌入式组合
 */

type Animal struct {
	Name string `required max:"100"`//tag
	Origin string
}
type Bird struct {
	Animal//embed
	SpeedKPH float32
	CanFly bool
}

func main6() {
	b:=Bird{}
	b.Name="Emu"
	b.Origin="Austr"
	b.SpeedKPH=48
	b.CanFly=false
	fmt.Println(b.Name)

	t:=reflect.TypeOf(Animal{})
	field,_ := t.FieldByName("Name")
	fmt.Println(field.Tag)//required max:"100"
}

