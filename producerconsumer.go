package main

import (
	"fmt"
	"time"
	"math/rand"
)

func producer(ch chan <- int){
	fmt.Print("producer: ")
	ch <- rand.Intn(1000)
	time.Sleep(10 * time.Millisecond)
}

func consumer(ch <- chan int){
	fmt.Println("consumer: ",<-ch)
	time.Sleep(10 * time.Millisecond)
}

func main() {

	ch := make(chan int)
	go func() {for{producer(ch)}}()
	for{consumer(ch)}



}