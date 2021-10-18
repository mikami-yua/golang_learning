package main

import (
	"bytes"
	"fmt"
)

/*
Basics
	structs are data containers
	interface describe behaviors,we storing method definitions
Composing
	如果需要包装一个具体类型，而没有人发布过一个接口，可以创建一个他们类型的实现接口
	单方法的接口分程强大，他们定义了非常具体的行为。可以被以各种方式进行实现

Type conversion
	empty interface
	type switch

Implementing with value vs. pointers
 */
func main12() {
	var w Writer = ConsoleWriter{}
	//w holding a Writer,which is something that implements the Writer Interface
	//-donnot know the concrete type
	w.Write([]byte("hello Go!"))//know how to call the method because,taht's
	//-defined by interface,we can replace ConsoleWriter to TCPWriter or FileWriter
	//实现多态
	myInt:=IntCounter(0)
	var  inc Incrementer = &myInt
	for i:=0;i<10;i++ {
		fmt.Println(inc.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("hello youtube listeners,this is test"))
	wc.Close()

	bwc:=wc.(*BufferedWriterCloser)
	fmt.Println(bwc)//&{0xc00007a3c0}


}

type Writer interface {
	//define
	Write([]byte) (int,error)//int the number of bytes written
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

type BufferedWriterCloser struct {//embedding
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser ) Write(data []byte) (int,error)  {
	n,err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v:=make([]byte,8)
	for bwc.buffer.Len()>8 {
		_,err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_,err=fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len()>0 {
		data:=bwc.buffer.Next(8)
		_,err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

//implement 隐式实现
//implicitly implement the interface,by create a method on struct taht has
//-the signature of a Writer interface
type ConsoleWriter struct {}

func (cw ConsoleWriter ) Write(data []byte) (int,error) {
	n,err := fmt.Println(string(data))
	return n, err
}

type Incrementer interface {
	Increment() int
}

type IntCounter int//use any type to implement interface
//we can add method to any type ,and we can implement interface with it

func (ic *IntCounter) Increment() int  {
	*ic++
	return int(*ic)
}