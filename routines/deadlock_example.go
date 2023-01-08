package main

import (

	"fmt"
	"sync"
)

func worker(reader, writer chan int, wg *sync.WaitGroup, id int){

	fmt.Printf("\nIn Worker Id : %d\n", id)

	<- reader
	
	writer <- 1
	
	wg.Done()

}

func main(){

	var wg sync.WaitGroup
	
	var c1, c2 chan int
	
	wg.Add(2)
	
	worker(c1, c2, &wg, 1)
	
	worker(c2, c1, &wg, 2)
	
	wg.Wait()

}
