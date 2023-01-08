package main

import (

	"fmt"
)

func one(){

	fmt.Printf("In one\n")
	two()
}

func two(){

	fmt.Printf("In two\n")
}

func main(){

	one()
	
}
