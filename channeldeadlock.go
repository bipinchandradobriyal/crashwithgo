package main

import (
	"time"
	"fmt"
)

/*
	Chan send and receive should be in different routines

	Channel operations :
 		Sending to channel			(channel_name <- 12)
 		Receiving from channel 		(value := <- channel_name)

	Channel states:
		Empty
		Non-Empty

	If the channel is empty, the receiver blocks
	If the channel is non-empty, the send blocks

	In case of unbuffered channel, if any upper two blocks happen then the next statement
	does not execute

	Handling of "fatal error: all goroutines are asleep - deadlock!"

	Ordering of code is important while dealing with channel. The channel operations
	should be done in different go routines.

	Should use close channel for the deadlock handling

*/

func main() {

	// creates a unbuffered or synchronized channel
	ch := make(chan int,4)

	ch <- 5
	go func() {
		for val := range ch{
			fmt.Print(" VAL:: ",val)
		}
	}()
	fmt.Println("done")
	for{
		time.Sleep(5000 * time.Millisecond)
	}


	/*go func (){
		//defer close(ch)
		var i int
		for i=1; i<=10; i++{
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()*/

	/*for val := range ch{
		fmt.Print(" VAL:: ",val)
	}*/


/*
	Post closing a channel, the below points need to be taken care:
		Subsequent send operations will cause a program to panic
		Receive operations never block (regardless of whether buffered or unbuffered)
		All receive operations return the zero value of the channel's element type

	go func() {
		for i:=1; i<=10; i++{
			ch <- i
			time.Sleep(500 * time.Millisecond)
			if(i == 5){
				close(ch)
				return
			}
		}
	}()
	for i:=1; i<=10; i++{
		if val, opened := <-ch; opened {
			fmt.Println(val)
		} else {
			fmt.Println(val)
			fmt.Println("Channel closed!")
		}
	}
*/


}
