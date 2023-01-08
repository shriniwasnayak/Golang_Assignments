package main

import(

	"fmt"
	"sync"
)

func printHello(wg *sync.WaitGroup){

	fmt.Printf("%s","In Hello func\n")
	
	wg.Done()
}

func main(){

	var wg sync.WaitGroup
	
	wg.Add(2)
	
	go printHello(&wg)

	wg.Wait()
}
