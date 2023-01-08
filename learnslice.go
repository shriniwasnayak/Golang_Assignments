package main

import (

	"fmt"
)

func main(){

	sli := make([]int, 2)
	
	fmt.Printf("\n%v : %d : %d\n", sli, len(sli), cap(sli))
	
	sli[0] = 9
	
	fmt.Printf("\n%v : %d : %d\n", sli, len(sli), cap(sli))
	
	sli = append(sli, 7)
	
	fmt.Printf("\n%v : %d : %d\n", sli, len(sli), cap(sli))
	
	sli = append(sli, 11)
	
	fmt.Printf("\n%v : %d : %d\n", sli, len(sli), cap(sli))
	
	sli = append(sli, 13)
	
	fmt.Printf("\n%v : %d : %d\n", sli, len(sli), cap(sli))
}
