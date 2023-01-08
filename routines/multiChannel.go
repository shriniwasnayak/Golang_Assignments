package main

import (

	"fmt"
	"math/rand"
	"time"
)

func channelOne(c1 chan string){

	// put random delay

	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	c1 <- "\n In Channel One\n"
}

func channelTwo(c2 chan string){

	// put random delay

	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	
	c2 <- "\n In Channel Two\n"
}

func main(){

	var c1, c2 chan string
	
	c1 = make(chan string)

	c2 = make(chan string)

	go channelOne(c1)
	
	go channelTwo(c2)

	select{
	
		case a := <- c1:
		
			fmt.Printf(a)
			
		case b := <- c2:
		
			fmt.Printf(b)
	
	}

}
