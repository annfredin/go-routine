package main

import (
	"fmt"
)

func chanTest(){
	fmt.Println("\n Chan Test")

	ch := make(chan int)
	go send(ch)
	go receive(ch)

	go lenCapChan()
}

func send(c chan int){
	//time.Sleep(1 * time.Second)
	fmt.Println("Sending")

	c <- 10 // will be blocked untill receive, chan init with unbuffered one.
	fmt.Println("Sended")
}

func receive(c chan int){

	fmt.Println("Receiving")
	a:= <-c //release blocking [send/receive both are block the execution, channel will take care of blocking/unblocking]
	fmt.Println("Received!")

	fmt.Println("Val: ", a)
}

func lenCapChan(){
	ch := make(chan int, 3)
	ch <- 5
	fmt.Printf("Len: %d\n", len(ch))

	ch <- 6
	fmt.Printf("Len: %d\n", len(ch))
	ch <- 7
	fmt.Printf("Len: %d\n", len(ch))
	fmt.Printf("Cap: %d\n", cap(ch))
}