package main

import (

	"fmt"
	"strconv"
	"strings"

)

func Swap(arr []int, id int){

	var temp int

	temp = arr[id]
	arr[id] = arr[id+1]
	arr[id+1] = temp	

}

func BubbleSort(arr []int){

	var size int = len(arr)
	
	var i,j int
	
	for i = 0; i<size ; i++{
	
		for j = 0; j<size-1; j++{
		
			if(arr[j] > arr[j+1]){
			
				Swap(arr, j)			
				
			}
		
		}
	
	}

}

func readData(arr []int) ([]int) {

	var data string = "-1"	

	for ;;{
	
		if(len(arr) == 10){
		
			fmt.Printf("\n ::: Maximum Capacity Reached :::\n")
			break
		}
	
		fmt.Printf("\nEnter Number OR 'X' to exit : ")
		
		fmt.Scanf("%s", &data)
		
		data = strings.ToUpper(data)
		
		if(data == "X"){
		
			break
		}
		
		val,err := strconv.Atoi(data)
		
		if( err != nil){
			fmt.Print(err)
		} else{
		
			arr = append(arr, val)
		}
	}

	return arr

}

func main(){

	arr := make([]int, 0, 0)

	arr = readData(arr)
	
	fmt.Printf("\nEntered Array : %v\n", arr)

	BubbleSort(arr)

	fmt.Printf("\nSorted Array : %v\n", arr)

}

