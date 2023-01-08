package main

import (

	"fmt"
	"time"
)

func printString(data string){

	var i int

	for i = 0; i<30; i++{
	
		fmt.Printf("%s : %d\n", data, i)
	}
}

func main(){

	fmt.Printf("hello\n")
	
	go printString("A")
	
	go printString("B")

	time.Sleep(time.Second)

	fmt.Printf("hello\n")
}

