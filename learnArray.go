package main

import (

	"fmt"

)

func printArr(arr [5]int, s1 []int, s2 []int){

	fmt.Printf("\nArr : %v\nS1 : %v\nS2 : %v\n", arr, s1, s2)
}

func main(){

	var arr [5]int = [...] int{1,2,3,4,5}

	s1 := arr[0:3]

	s2 := arr[1:3]

	s1[0] = 100

	s2[1] = 300

	printArr(arr, s1, s2)

}
