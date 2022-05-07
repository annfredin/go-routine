package main

import (
	"fmt"
	"runtime"
	"time"
)

func main(){
	fmt.Println("GO-ROUTINE DEMO:")

	go start()
    fmt.Println("Started")
	//if the below line commented, then start fn will never be executed, bec before that main fn will exit
    time.Sleep(1 * time.Second)

	noParentChild()
    time.Sleep(1 * time.Second)

    fmt.Println("Finished")
	fmt.Println("NumCPU: ", runtime.NumCPU())

	// chanTest()
	wgTest()
    time.Sleep(1 * time.Second)

}

func start() {
	fmt.Println("Inside Goroutine")
}

//go routine has no parent and child, all are executed in independent manner
func noParentChild(){
	fmt.Println("Proving Goroutine has no-parent/child")

	child1()
}

func child1(){
	fmt.Println("Goroutine Child-1")
	go child2()
}

func child2(){
	fmt.Println("Goroutine Child-2 (Still Executing, even child1 fn exits -> no-parent dependency)")
}