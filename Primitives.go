package main

import "fmt"

/*
Boolean type
	var n bool=true
	fmt.Println("helloWorld")
	fmt.Printf("%v,%T\n",n,n)
	a:=1==1
	b:=2==1
	fmt.Printf("%v,%T\n",a,a)
	fmt.Printf("%v,%T\n",b,b)

	every time initialize a var it has a zero var
	initialize boolean value is false

Numeric types
	Integer---varying size,but min 32 bits
		int8 -128~127
		int16 -32768~32767
		int32
		int64

		cannot add an int var with an int8 var
		go very confervison with data conversion

	Floating point
		float32 cannot add with float64

		n:=3.14
		n=13.7e72
		n=2.1E14
		fmt.Printf("%v,%T\n",n,n)//2.1e+14,float64
	Complex number
		var n complex64=1+2i
		fmt.Printf("%v,%T\n",n,n)//(1+2i),complex64

		var n complex64=1+2i
		fmt.Printf("%v,%T\n",real(n),real(n))//1,float32
		fmt.Printf("%v,%T\n",imag(n),imag(n))//2,float32
		var m complex128=1+2i
		fmt.Printf("%v,%T\n",real(m),real(m))//1,float64
		fmt.Printf("%v,%T\n",imag(m),imag(m))//2,float64
		var j complex64=complex(5,12)
		fmt.Printf("%v,%T\n",j,j)

Text types
	string---UTF8 arrays
		cannot assign a string to a byte,hava to do a conversion
		Cannot manipulate the value of the string
		string can add together
		convert string to a collcetions of bytes,which in go is called a slice of bytes
			send to other place use this
		s:="this is a string"
		fmt.Printf("%v,%T\n",s,s)//this is a string,string
		fmt.Printf("%v,%T\n",string(s[2]),s[2])//i,uint8
		s2:="this is also a string"
		fmt.Printf("%v,%T\n",s+s2,s+s2)//this is a stringthis is also a string,string
		b:=[]byte(s2)
		fmt.Printf("%v,%T\n",b,b)
		//[116 104 105 115 32 105 115 32 97 108 115 111 32 97 32 115 116 114 105 110 103],[]uint8

	rune---UTF32

		rune is an alias for integer32
 */

func main3() {
	var r rune='a'
	fmt.Printf("%v,%T\n",r,r)//97,int32

}
