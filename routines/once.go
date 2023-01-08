package main

import (

	"fmt"
	"sync"
)

var on sync.Once

func myinit(){

	fmt.Printf("In Init\n")
}

func goRoutine(id int, wg *sync.WaitGroup){

	on.Do(myinit)	
		
	fmt.Printf("In go routine id : %d\n", id)
	
	wg.Done()
}

func main(){

	var wg sync.WaitGroup
	
	wg.Add(3)
	
	go goRoutine(1, &wg)
	go goRoutine(2, &wg)
	go goRoutine(3, &wg)
	
	wg.Wait()

}
