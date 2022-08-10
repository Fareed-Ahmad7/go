package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main(){
	go sayHello()
	//  case 1
	// it should print farid but it prints ahmed
	var msg = "farid"
	go func(){
		fmt.Println(msg)
	}()
	msg = "ahmed"
	time.Sleep(100 * time.Millisecond)

	// case 2
	// to get away with this we can var as arguments
	var msg = "farid"
	go func(msg string){
		fmt.Println(msg)
	}(msg)
	msg = "ahmed"
	time.Sleep(100 * time.Millisecond)


	// case 3
	// to avoid using sleep we use waitGroup
	// wait group is designed to synchronize multiple go routines together
	var msg = "farid"
	wg.Add(1)

	go func(msg string){
		fmt.Println(msg)
		wg.Done()
	}(msg)

	msg = "ahmed"
	wg.Wait()


}

func sayHello(){
	fmt.Println("hello")
}