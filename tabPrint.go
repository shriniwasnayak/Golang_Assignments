package main


import (
	"fmt"
)

func main(){

	var num, range_start, range_end int
	
	fmt.Printf("\n Enter number for which table is required : ")
	fmt.Scanf("%d",&num)
	
	fmt.Printf("\n Enter range start : ")
	fmt.Scanf("%d",&range_start)

	fmt.Printf("\n Enter range end : ")
	fmt.Scanf("%d",&range_end)
	
	fmt.Printf("\n::: Table :::\n")
	
	for i:=range_start; i<=range_end; i++{
	
		fmt.Printf("%d * %d = %d\n",num, i, num*i)
	
	}
}
