package main

import (

	"fmt"

)

func data(s1 string, ss ...string){

	fmt.Printf("\n1st Arg : %s\n", s1)
	fmt.Printf("2nd Arg : %v\n", ss)

}

func main(){

	data("A1")

}
