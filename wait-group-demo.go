package main

import (
	"fmt"
	"sync"
)

func wgTest(){
	fmt.Println("Wait Group Demo:")

	//====== %%% ALL the types in Sync package needs to be passed by ref (not by value)
	var wg sync.WaitGroup
	wg.Add(1)
	go hello(&wg)
	wg.Wait()

	world()
}

//bydefault fo params are passed by value, so in this case it will inititlize counter with zero value (not as 1), so when w.Done() will decrement value, becomes negative value, expecption occurs
//MUST -> always pass WaitGroup as ref in fn's
func hello(w *sync.WaitGroup) {
	fmt.Print("Hello")
	w.Done()
}

func world(){
	fmt.Print(" World\n")
}