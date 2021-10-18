package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Channels
	pass dta between different go routines
	designed to synchronize data transmission between multiple g routines



create
	ps1.
		//create channel
		ch:=make(chan int)//chan key words(we want to create a channel)
		//we can only send Integers though this channel(use string key words send strings
		// use *int key words only send integer pointers)
		//create a channel to accept messages of a certain type
		wg3.Add(2)//设置两个go routines
		go func(){
			i:= <-ch//receive data from the channel
			fmt.Println(i)
			wg3.Done()
		}()
		go func(){
			i:=42
			ch<-i//send message into a channel,箭头的方向是data want to flow
			i=27//when we passing data into a channel,we actually pass a copy of the data
			wg3.Done()
		}()
		wg3.Wait()


	ps2.
		for j :=0;j<5;j++{//create 5 sets of go routines
			wg3.Add(2)
			go func(){
				i:=<-ch
				fmt.Println(i)
				wg3.Done()
			}()
			go func() {
				ch<-42//this line will going to pause the execution of this go routine
				//-until these's a space available in the channel
				//by default,we're working with unbuffered channels,means only one message
				//-can be in the channel at one time
				wg3.Done()
			}()
		}//create 10 go routines,all of them use the single channel ch to communicate message
		//当把接受go routines 放在循环外时（5个发送1个接受）
		//不会工作：只接受一个，但是发送了5个，出现deadlock condition
		//向channel中发送int，但是没有go rutine可以进行处理
		wg3.Wait()



restrict data flow(限制数据流)
	basic channel is a two way street(send data in,send data out)
	sometime we want to send only channel or receive oly channel
	ps3.
		ch:=make(chan int)
		wg3.Add(2)//设置两个go routines
		go func(ch<-chan int){//set receive only channel
			i:= <-ch//receive data from the channel
			fmt.Println(i)
			//ch<-27//send to receive-only type <-chan int
			wg3.Done()
		}(ch)//将ch设置为形参
		go func(ch chan<-int){//set send only channel
			i:=42
			ch<-i
			//fmt.Println(<-ch)
			wg3.Done()
		}(ch)
		wg3.Wait()

buffered channels
	design channels to have an internal data store,they can store several message
	-at once in case the sender and receiver aren't processing data at the same rate

	ps4.
		ch:=make(chan int,50)//set buffer,this channel can store 50 integers
		//程序不会出现deadlock，但是会出现数据丢失，接受两次即可
		wg3.Add(2)//设置两个go routines
		go func(ch<-chan int){//set receive only channel
			i:= <-ch//receive data from the channel
			fmt.Println(i)
			i= <-ch//receive data from the channel
			fmt.Println(i)
			wg3.Done()
		}(ch)//将ch设置为形参
		go func(ch chan<-int){//set send only channel
			ch<-42
			ch<-27//all goroutines are asleep - deadlock!
			wg3.Done()
			//使用buffer解决多个发送一个接受的问题
		}(ch)
		wg3.Wait()

	sender and receiver needs a little more time to process ,and we don't
	-want to block the other side
close
	ps5.
		ch:=make(chan int,50)
		wg3.Add(2)
		go func(ch<-chan int) {
			for i :=range ch{//without close():dead lock here:continuing monitor,but we stop input
				//-for loop never end,this go routine never end,cause deadlock
				//对channel不需要使用两个值range
				fmt.Println(i)
			}
			wg3.Done()
		}(ch)
		go func(ch chan<-int) {
			ch<-27
			ch<-42
			close(ch)//for loop will know the length of channel
			ch<-33//panic: send on closed channel
			//cannot recover from a closed channel,even cannot detect if a channel
			//-is closed(the limition of go)
			wg3.Done()
		}(ch)
		wg3.Wait()

	ps6.
		ch:=make(chan int,50)
		wg3.Add(2)
		go func(ch<-chan int) {
			for {
				if i,ok:=<-ch;ok{//ok means channel is open or not
					fmt.Println(i)
				}else {
					break
				}
			}
			wg3.Done()
		}(ch)
		go func(ch chan<-int) {
			ch<-27
			ch<-42
			close(ch)//for loop will know the length of channel
			wg3.Done()
		}(ch)
		wg3.Wait()

for range loops with channels

select statement


*/
var wg3 = sync.WaitGroup{} //use for synchronize(we use channel and waitgroup for synchronize)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logERROR   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) //add an addition channel 使用struct with no field
//struct with no field in go is unque that it requires zero memory allocations
//a channel setup like this,the intention is it can send any data though
//-except for the face that a message was sent or received
//-called signal channel.
//There 's zero memory allocations required in sending the message.
//-but we do have the ability to just let the receiving side know that a message was sent

/*
没有使用done go routine时候结束
	an application is shut down as soon as the last statement of main function
	-finishes execution.
	When we finish ehd sleep call,the application terminates and everything is turn
	-down and all resources ar clean include go routines

	Our go routines are torn down forcibly,beacuse main fuction has done

	Many situations where we want to have much more control over a go routine
	-we should have a strategy for how our go rouutine is going to shut down when we create
	-our go routines

*/
func logger() {
	for {
		select {
		//整个语句将会block until a message is received on one of the channels that it's
		//-listening for
		//in this case:we are listening case for messages from logCh and
		//-listening for message from the doneCh.
		//-if we get messages from logCh we print.
		//-if we get messages from doneCh we break from the loop
		//This allow us:at the end of the application we can go ahead and pass in a
		//-message into our doneCh
		case entry := <-logCh:
			fmt.Printf("%v-[%v]%v\n", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		case <-doneCh:
			break
		}

	}
}

func main16() {
	go logger()
	/*defer func() {//当主函数结束时可以关闭channel
		close(logCh)
	}()*/

	logCh <- logEntry{time.Now(), logInfo, "APP is starting"}

	logCh <- logEntry{time.Now(), logInfo, "APP is shutting down"}
	time.Sleep(100 * time.Millisecond)
	doneCh <- struct{}{} //This allow us:at the end of the application we can go ahead and pass in a
	//-message into our doneCh
}
